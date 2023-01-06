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

func uploadFileRequestToPath(storeId string, uploadFileRequest openapi.UploadFileRequest) string {
	return fmt.Sprintf("stores/%s/%s/%s/%s", storeId, uploadFileRequest.FileName, uploadFileRequest.Hash, uploadFileRequest.FileName)
}

type FileEntry struct {
	FileName string
	Hash     string
	Status   string
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

	for _, uploadFileRequest := range createStoreUploadRequest.Files {

		objectURL := ""

		// Validate whether object exists in bucket
		// This will talk to the Cloud Storage APIs

		path := uploadFileRequestToPath(storeId, uploadFileRequest)
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
			createStoreUploadResponse.Files = append(createStoreUploadResponse.Files, openapi.UploadFileResponse{
				FileName: uploadFileRequest.FileName,
				Hash:     uploadFileRequest.Hash,
				Url:      objectURL,
			})
		}
	}

	// Log upload to database

	files := make([]FileEntry, 0)

	for _, file := range createStoreUploadResponse.Files {
		status := models.StoreUploadFileStatusAlreadyPresent
		if file.Url != "" {
			// Legacy API users will not report completion of individual file upload; therefore the file's stats will be assumed to be completed immediately
			// For those that use the progress API, however, we can say for sure that the file is pending upload at this point
			if createStoreUploadRequest.UseProgressApi {
				status = models.StoreUploadFileStatusPending
			} else {
				status = models.StoreUploadFileStatusCompleted
			}
		}
		files = append(files, FileEntry{
			FileName: file.FileName,
			Hash:     file.Hash,
			Status:   status,
		})
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
		Files:       files,
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

	createStoreUploadResponse.Id = uploadId

	log.Printf("Response: %v", createStoreUploadResponse)

	return openapi.Response(http.StatusOK, createStoreUploadResponse), nil
}

func logUpload(ctx context.Context, storeId string, storeUploadEntry StoreUploadEntry) (string, error) {

	log.Printf("Writing upload to database: %v", storeUploadEntry)

	tx, err := postgres.BeginDBTransaction(ctx)
	if err != nil {
		return "", err
	}

	// Locate store in DB, and ensure store remains throughout entire txn
	store, err := models.Stores(
		qm.Where(models.StoreColumns.Name+" = ?", storeId),
		qm.For("update"),
	).One(ctx, tx)
	if err != nil {
		log.Printf("error while accessing store: %v", err)
		tx.Rollback()
		return "", err
	}

	var storeUploadIndex = store.NextStoreUploadIndex

	store.NextStoreUploadIndex++

	// Increase next store upload index in store
	_, err = store.Update(ctx, tx, boil.Infer())
	if err != nil {
		log.Printf("error while updating store: %v", err)
		tx.Rollback()
		return "", err
	}

	// Add upload entry to DB
	var upload = models.StoreUpload{
		StoreID:          null.IntFrom(store.StoreID),
		StoreUploadIndex: storeUploadIndex,
		Description:      storeUploadEntry.Description,
		Build:            storeUploadEntry.BuildId,
		Timestamp:        time.Now(),
		Status:           storeUploadEntry.Status,
	}
	err = upload.Insert(ctx, tx, boil.Infer())
	if err != nil {
		log.Printf("unable to insert upload: %v", err)
		tx.Rollback()
		return "", err
	}

	var uploadFileIndex = 0

	// Add entries for each upload-file in DB
	for _, file := range storeUploadEntry.Files {

		// Create store-file, in case one doesn't exist yet
		storeFile, err := models.StoreFiles(
			qm.Where(models.StoreFileColumns.StoreID+" = ? AND "+models.StoreFileColumns.FileName+" = ?", store.StoreID, file.FileName),
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
				return "", err
			}
		} else if err != nil {
			log.Printf("error while locating store-file: %v", err)
			tx.Rollback()
			return "", err
		}

		// Create store-file-hash, in case one doesn't exist yet
		storeFileHash, err := models.StoreFileHashes(
			qm.Where(models.StoreFileHashColumns.FileID+" = ?", storeFile.FileID),
		).One(ctx, tx)
		if err == sql.ErrNoRows {

			// Translate all storefilehash statuses except "pending" to "present"
			// We assume that all uploads using non-progression API succeeds
			fileHashStatus := models.StoreFileHashStatusPresent
			if file.Status == models.StoreUploadFileStatusPending {
				fileHashStatus = models.StoreFileHashStatusPending
			}

			storeFileHash = &models.StoreFileHash{
				FileID: null.IntFrom(storeFile.FileID),
				Hash:   file.Hash,
				Status: fileHashStatus,
			}
			err = storeFileHash.Insert(ctx, tx, boil.Infer())
			if err != nil {
				log.Printf("unable to insert file-hash: %v", err)
				tx.Rollback()
				return "", err
			}
		} else if err != nil {
			log.Printf("error while locating file-hash: %v", err)
			tx.Rollback()
			return "", err
		}

		// Create store-upload-file
		var uploadFile = models.StoreUploadFile{
			UploadID:        null.IntFrom(upload.UploadID),
			HashID:          null.IntFrom(storeFileHash.HashID),
			UploadFileIndex: uploadFileIndex,
			Status:          file.Status,
			FileName:        file.FileName,
			FileHash:        file.Hash,
		}
		err = uploadFile.Insert(ctx, tx, boil.Infer())
		if err != nil {
			log.Printf("unable to insert upload-file: %v", err)
			tx.Rollback()
			return "", err
		}

		uploadFileIndex++
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("unable to commit transaction: %v", err)
		return "", err
	}

	log.Printf("Upload is given ID %v", storeUploadIndex)

	return fmt.Sprint(storeUploadIndex), nil
}