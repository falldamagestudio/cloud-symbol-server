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

	storeDoc, err := getStoreDoc(context, storeId)
	if err != nil {
		log.Printf("Unable to fetch store document for %v, err = %v", storeId, err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to fetch store document for %v", storeId)}), err
	}
	if storeDoc == nil {
		log.Printf("Store %v does not exist", storeId)
		return openapi.Response(http.StatusNotFound, &openapi.MessageResponse{Message: fmt.Sprintf("Store %v does not exist", storeId)}), err
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

	uploadId, err := logUpload(context, storeId, uploadStatus, createStoreUploadRequest.Description, createStoreUploadRequest.BuildId, files)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Internal error while logging upload to DB"}), errors.New("Internal error while logging upload to DB")
	}

	createStoreUploadResponse.Id = uploadId

	log.Printf("Response: %v", createStoreUploadResponse)

	return openapi.Response(http.StatusOK, createStoreUploadResponse), nil
}

func incrementStoreUploadId(tx *firestore.Transaction, storeDocRef *firestore.DocumentRef) (int64, error) {
	storeDoc, err := tx.Get(storeDocRef)
	if err != nil {
		return 0, err
	}
	storeEntry := &StoreEntry{}
	if err := storeDoc.DataTo(storeEntry); err != nil {
		return 0, err
	}

	newUploadId := storeEntry.LatestUploadId + 1
	storeEntry.LatestUploadId = newUploadId

	err = tx.Set(storeDocRef, storeEntry)
	if err != nil {
		return 0, err
	}

	return newUploadId, nil
}

func createUploadDoc(tx *firestore.Transaction, storeDocRef *firestore.DocumentRef, newUploadId int64, uploadContent StoreUploadEntry) error {
	uploadDocRef := storeDocRef.Collection(storeUploadsCollectionName).Doc(fmt.Sprint(newUploadId))

	err := tx.Create(uploadDocRef, uploadContent)
	return err
}

func logUpload(ctx context.Context, storeId string, uploadStatus string, description string, buildId string, files []FileDBEntry) (string, error) {

	uploadContent := StoreUploadEntry{
		Description: description,
		BuildId:     buildId,
		Files:       files,
		Timestamp:   time.Now().Format(time.RFC3339),
		Status:      uploadStatus,
	}

	log.Printf("Writing upload to database: %v", uploadContent)

	firestoreClient, err := firestoreClient(ctx)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return "", err
	}

	storeDocRef := firestoreClient.Collection(storesCollectionName).Doc(storeId)

	newUploadId := int64(0)

	err = firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {

		newUploadId, err := incrementStoreUploadId(tx, storeDocRef)
		if err != nil {
			return err
		}

		err = createUploadDoc(tx, storeDocRef, newUploadId, uploadContent)
		if err != nil {
			return err
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

	log.Printf("Upload is given ID %v", newUploadId)

	return fmt.Sprint(newUploadId), nil
}
