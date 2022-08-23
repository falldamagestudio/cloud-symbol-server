package admin_api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func uploadFileRequestToPath(storeId string, uploadFileRequest openapi.UploadFileRequest) string {
	return fmt.Sprintf("stores/%s/%s/%s/%s", storeId, uploadFileRequest.FileName, uploadFileRequest.Hash, uploadFileRequest.FileName)
}

func (s *ApiService) CreateStoreUpload(context context.Context, storeId string, createStoreUploadRequest openapi.CreateStoreUploadRequest) (openapi.ImplResponse, error) {

	signedURLExpirationSeconds := 15 * 60

	storageClient, err := getStorageClient(context)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to create storage client")}), err
	}

	symbolStoreBucketName, err := getSymbolStoreBucketName()
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

	// Log upload to Firestore DB

	files := make([]FileDBEntry, 0)

	for _, file := range createStoreUploadResponse.Files {
		status := FileDBEntry_Status_AlreadyPresent
		if file.Url != "" {
			// Legacy API users will not report completion of individual file upload; therefore the file's stats will remain unknown
			// For those that use the progress API, however, we can say for sure that the file is pending upload at this point
			if createStoreUploadRequest.UseProgressApi {
				status = FileDBEntry_Status_Pending
			} else {
				status = FileDBEntry_Status_Unknown
			}
		}
		files = append(files, FileDBEntry{
			FileName: file.FileName,
			Hash:     file.Hash,
			Status:   status,
		})
	}

	// Legacy API users will not report completion/abortion of the upload operation; therefore the upload's state will remain unkonwn
	// For those that use the progress API, however, we can say for sure that the upload is in progress at this point
	uploadStatus := StoreUploadEntry_Status_Unknown
	if createStoreUploadRequest.UseProgressApi {
		uploadStatus = StoreUploadEntry_Status_InProgress
	}

	uploadContent := StoreUploadEntry{
		Description: createStoreUploadRequest.Description,
		BuildId:     createStoreUploadRequest.BuildId,
		Files:       files,
		Timestamp:   time.Now().Format(time.RFC3339),
		Status:      uploadStatus,
	}

	uploadId, err := logUpload(context, storeId, uploadContent)
	if err != nil {
		if status.Code(err) == codes.NotFound {
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

	uploadId := int64(0)

	err := runDBTransaction(ctx, func(ctx context.Context, client *firestore.Client, tx *firestore.Transaction) error {

		storeEntry, err := getStoreEntry(client, tx, storeId)
		if err != nil {
			return err
		}

		uploadId = storeEntry.LatestUploadId + 1

		newStoreEntry := StoreEntry{
			LatestUploadId: uploadId,
		}

		err = updateStoreEntry(client, tx, storeId, &newStoreEntry)
		if err != nil {
			return err
		}

		err = createStoreUploadEntry(client, tx, storeId, uploadId, &storeUploadEntry)
		if err != nil {
			return err
		}

		for _, file := range storeUploadEntry.Files {

			err = updateStoreFileHashEntry(client, tx, storeId, file.FileName, file.Hash, &StoreFileHashEntry{})
			if err != nil {
				return err
			}

			err = createStoreFileHashUploadEntry(client, tx, storeId, file.FileName, file.Hash, uploadId, &StoreFileHashUploadEntry{})
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		// Handle any errors appropriately in this section.
		log.Printf("An error has occurred: %s", err)
	}

	if err != nil {
		log.Printf("Error when logging upload, err = %v", err)
		return "", err
	}

	log.Printf("Upload is given ID %v", uploadId)

	return fmt.Sprint(uploadId), nil
}
