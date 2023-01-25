package upload_api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/storage"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	cloud_storage "github.com/falldamagestudio/cloud-symbol-server/backend-api/cloud_storage"
	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/sql-db-models"
	postgres "github.com/falldamagestudio/cloud-symbol-server/backend-api/postgres"
)

func uploadFileRequestToPath(storeId string, createStoreUploadFileRequest openapi.CreateStoreUploadFileRequest) string {
	// Legacy API users provide the blob identifier in 'Hash' rather than 'BlobIdentifier' member
	blobIdentifier := createStoreUploadFileRequest.Hash
	if blobIdentifier == "" {
		blobIdentifier = createStoreUploadFileRequest.BlobIdentifier
	}

	return fmt.Sprintf("stores/%s/%s/%s/%s", storeId, createStoreUploadFileRequest.FileName, blobIdentifier, createStoreUploadFileRequest.FileName)
}

type FileEntry struct {
	FileName       string
	BlobIdentifier string
	Type           string
	Size           int64
	ContentHash    string
	Status         string
}

type StoreUploadEntry struct {
	Description string
	BuildId     string
	Files       []FileEntry
	Status      string
}

func CreateStoreUpload(context context.Context, storeId string, createStoreUploadRequest openapi.CreateStoreUploadRequest) (openapi.ImplResponse, error) {

	signedURLExpirationSeconds := 15 * 60

	storageClient, err := cloud_storage.GetStorageClient(context)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to create storage client")}), err
	}

	symbolStoreBucketName, err := cloud_storage.GetSymbolStoreBucketName()
	if err != nil {
		log.Printf("Unable to determine symbol store bucket name: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to determine symbol store bucket name")}), err
	}

	// Legacy API users (those who do not use the progress API) expect the response to filter
	//  out any files that already are present; those that do use the progress API
	//  expect to have all files listed in the response, even if they should not be uploaded
	includeAlreadyPresentFiles := createStoreUploadRequest.UseProgressApi

	createStoreUploadResponse := openapi.CreateStoreUploadResponse{}

	logFiles := make([]FileEntry, 0)

	for _, createStoreUploadFileRequest := range createStoreUploadRequest.Files {

		objectURL := ""

		// Validate whether object exists in bucket
		// This will talk to the Cloud Storage APIs

		path := uploadFileRequestToPath(storeId, createStoreUploadFileRequest)
		log.Printf("Validating whether object %v does exists in bucket %v", path, symbolStoreBucketName)
		_, err = storageClient.Bucket(symbolStoreBucketName).Object(path).Attrs(context)
		if err != nil {
			log.Printf("Object %v does not exist in bucket %v, preparing a redirect", path, symbolStoreBucketName)

			// Object does not exist in bucket; determine upload URL

			objectURL, err = storageClient.Bucket(symbolStoreBucketName).SignedURL(path, &storage.SignedURLOptions{
				Method:  "PUT",
				Expires: time.Now().Add(time.Duration(signedURLExpirationSeconds) * time.Second),
			})

			if err != nil {
				log.Printf("Unable to create signed URL for %v: %v", path, err)
				return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to create signed URL for %v", path)}), errors.New(fmt.Sprintf("Unable to create signed URL for %v", path))
			}

			log.Printf("Object %v has a signed URL %v, valid for %d seconds", path, objectURL, signedURLExpirationSeconds)

		} else {
			log.Printf("Object %v already exists in bucket %v", path, symbolStoreBucketName)
		}

		if includeAlreadyPresentFiles || (objectURL != "") {
			// Legacy API users provide the blob identifier in 'Hash' rather than 'BlobIdentifier' member
			blobIdentifier := createStoreUploadFileRequest.Hash
			if blobIdentifier == "" {
				blobIdentifier = createStoreUploadFileRequest.BlobIdentifier
			}

			createStoreUploadResponse.Files = append(createStoreUploadResponse.Files, openapi.UploadFileResponse{
				FileName:       createStoreUploadFileRequest.FileName,
				BlobIdentifier: blobIdentifier,
				Url:            objectURL,
				// Legacy API users expect 'Hash' instead of 'BlobIdentifier' as name for this member
				Hash: 			blobIdentifier,
			})

			logStatus := models.StoreUploadFileStatusAlreadyPresent
			if objectURL != "" {
				// Legacy API users will not report completion of individual file upload; therefore the file's stats will be assumed to be completed immediately
				// For those that use the progress API, however, we can say for sure that the file is pending upload at this point
				if createStoreUploadRequest.UseProgressApi {
					logStatus = models.StoreUploadFileStatusPending
				} else {
					logStatus = models.StoreUploadFileStatusCompleted
				}
			}

			logType := models.StoreFileBlobTypeUnknown
			if createStoreUploadFileRequest.Type != "" {
				logType = string(createStoreUploadFileRequest.Type)
			}

			logFiles = append(logFiles, FileEntry{
				FileName:       createStoreUploadFileRequest.FileName,
				BlobIdentifier: blobIdentifier,
				Type:           logType,
				Size:           createStoreUploadFileRequest.Size,
				ContentHash:    createStoreUploadFileRequest.ContentHash,
				Status:         logStatus,
			})
		}
	}

	// Legacy API users will not report completion/abortion of the upload operation; therefore the upload's state will be assumed to be completed immediately
	// For those that use the progress API, however, we can say for sure that the upload is in progress at this point
	uploadStatus := models.StoreUploadStatusCompleted
	if createStoreUploadRequest.UseProgressApi {
		uploadStatus = models.StoreUploadStatusInProgress
	}

	uploadContent := StoreUploadEntry{
		Description: createStoreUploadRequest.Description,
		BuildId:     createStoreUploadRequest.BuildId,
		Files:       logFiles,
		Status:      uploadStatus,
	}

	uploadId, err := logUpload(context, storeId, uploadContent)
	if err != nil {
		if err == sql.ErrNoRows {
			return openapi.Response(http.StatusNotFound, &openapi.MessageResponse{Message: fmt.Sprintf("Store %v not found", storeId)}), nil
		} else {
			return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Internal error while logging upload to DB"}), nil
		}
	}

	createStoreUploadResponse.Id = fmt.Sprint(uploadId)
	createStoreUploadResponse.UploadId = int32(uploadId)

	log.Printf("Response: %v", createStoreUploadResponse)

	return openapi.Response(http.StatusOK, createStoreUploadResponse), nil
}

func logUpload(ctx context.Context, storeId string, storeUploadEntry StoreUploadEntry) (int, error) {

	log.Printf("Writing upload to database: %v", storeUploadEntry)

	tx, err := postgres.BeginDBTransaction(ctx)
	if err != nil {
		return 0, err
	}

	// Locate store in DB, and ensure store remains throughout entire txn
	store, err := models.Stores(
		qm.Where(models.StoreColumns.Name+" = ?", storeId),
		qm.For("update"),
	).One(ctx, tx)
	if err != nil {
		log.Printf("error while accessing store: %v", err)
		tx.Rollback()
		return 0, err
	}

	var storeUploadIndex = store.NextStoreUploadIndex

	store.NextStoreUploadIndex++

	// Increase next store upload index in store
	_, err = store.Update(ctx, tx, boil.Infer())
	if err != nil {
		log.Printf("error while updating store: %v", err)
		tx.Rollback()
		return 0, err
	}

	var timestamp = time.Now()

	// Add upload entry to DB
	var upload = models.StoreUpload{
		StoreID:          null.IntFrom(store.StoreID),
		StoreUploadIndex: storeUploadIndex,
		Description:      storeUploadEntry.Description,
		Build:            storeUploadEntry.BuildId,
		Timestamp:        timestamp,
		Status:           storeUploadEntry.Status,
	}
	err = upload.Insert(ctx, tx, boil.Infer())
	if err != nil {
		log.Printf("unable to insert upload: %v", err)
		tx.Rollback()
		return 0, err
	}

	var uploadFileIndex = 0

	// Add entries for each upload-file in DB
	for _, file := range storeUploadEntry.Files {

		// Create store-file, in case one doesn't exist yet
		storeFile, err := models.StoreFiles(
			qm.Where(models.StoreFileColumns.StoreID+" = ?", store.StoreID),
			qm.And(models.StoreFileColumns.FileName+" = ?", file.FileName),
		).One(ctx, tx)
		if err == sql.ErrNoRows {
			storeFile = &models.StoreFile{
				StoreID:  null.IntFrom(store.StoreID),
				FileName: file.FileName,
			}
			err = storeFile.Insert(ctx, tx, boil.Infer())
			if err != nil {
				log.Printf("unable to insert file: %v", err)
				tx.Rollback()
				return 0, err
			}
		} else if err != nil {
			log.Printf("error while locating store-file: %v", err)
			tx.Rollback()
			return 0, err
		}

		// Create store-file-blob, in case one doesn't exist yet
		storeFileBlob, err := models.StoreFileBlobs(
			qm.Where(models.StoreFileBlobColumns.FileID+" = ?", storeFile.FileID),
			qm.And(models.StoreFileBlobColumns.BlobIdentifier+" = ?", file.BlobIdentifier),
		).One(ctx, tx)
		if err == sql.ErrNoRows {

			// Translate all storefileblob statuses except "pending" to "present"
			// We assume that all uploads using non-progression API succeeds
			fileBlobStatus := models.StoreFileBlobStatusPresent
			if file.Status == models.StoreUploadFileStatusPending {
				fileBlobStatus = models.StoreFileBlobStatusPending
			}

			storeFileBlob = &models.StoreFileBlob{
				FileID:          null.IntFrom(storeFile.FileID),
				BlobIdentifier:  file.BlobIdentifier,
				UploadTimestamp: timestamp,
				Type:            file.Type,
				Size:            file.Size,
				ContentHash:     file.ContentHash,
				Status:          fileBlobStatus,
			}
			err = storeFileBlob.Insert(ctx, tx, boil.Infer())
			if err != nil {
				log.Printf("unable to insert file-blob: %v", err)
				tx.Rollback()
				return 0, err
			}
		} else if err != nil {
			log.Printf("error while locating file-blob: %v", err)
			tx.Rollback()
			return 0, err
		}

		// Create store-upload-file
		var uploadFile = models.StoreUploadFile{
			UploadID:           null.IntFrom(upload.UploadID),
			BlobID:             null.IntFrom(storeFileBlob.BlobID),
			UploadFileIndex:    uploadFileIndex,
			Status:             file.Status,
			FileName:           file.FileName,
			FileBlobIdentifier: file.BlobIdentifier,
		}
		err = uploadFile.Insert(ctx, tx, boil.Infer())
		if err != nil {
			log.Printf("unable to insert upload-file: %v", err)
			tx.Rollback()
			return 0, err
		}

		uploadFileIndex++
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("unable to commit transaction: %v", err)
		return 0, err
	}

	log.Printf("Upload is given ID %v", storeUploadIndex)

	return storeUploadIndex, nil
}
