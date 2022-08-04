package admin_api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
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

	storageClient, storageEndpoint, err := getStorageClient(context)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to create storage client")}), err
	}

	// The Firebase Storage emulator has a non-standard format when referring to files
	// Normally, files should be referred to as "folder/filename"
	//  but the storage emulator expects them to be as "folder%2Ffilename" in Storage API paths
	// Therefore we need to activate override logic when using the emulator
	//  to access files directly (REST API, outside of SDK)

	urlEncodeRestAPIPath := (storageEndpoint != "")

	// Use signed URLs only when talking to the real Cloud Storage APIs
	// Otherwise, create public, unsigned URLs directly to the storage service
	//
	// The Cloud Storage SDK has support for working against local emulators,
	//  via the STORAGE_EMULATOR_HOST setting. However, this setting does not
	//  work properly for the SignedURL() functions when using local emulators:
	// The SignURL() function will always return URLs that point to the real
	//   Cloud Storage API, even when STORAGE_EMULATOR_HOST is set.
	// Because of this, when we use local emulators, we fall back to manually
	//  constructing download URLs.
	useSignedURLs := (storageEndpoint == "")

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

			if useSignedURLs {

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

				// The Firebase Storage emulator requires the path to be on the format "folder%2Ffilename"

				restAPIPath := path

				if urlEncodeRestAPIPath {
					restAPIPath = strings.ReplaceAll(path, "/", "%2F")
				}

				objectURL = fmt.Sprintf("%s/b/%s/o?uploadType=media&name=%s", storageEndpoint, symbolStoreBucketName, restAPIPath)

				log.Printf("Object %v has a non-signed URL %v", restAPIPath, objectURL)

			}

		} else {
			log.Printf("Object %v already exists in bucket %v", path, symbolStoreBucketName)
		}

		createStoreUploadResponse.Files = append(createStoreUploadResponse.Files, openapi.UploadFileResponse{
			FileName: uploadFileRequest.FileName,
			Hash:     uploadFileRequest.Hash,
			Url:      objectURL,
		})
	}

	// Log upload to Firestore DB

	files := make([]FileDBEntry, 0)

	for _, file := range createStoreUploadResponse.Files {
		status := FileDBEntry_Status_AlreadyPresent
		if file.Url != "" {
			status = FileDBEntry_Status_Pending
		}
		files = append(files, FileDBEntry{
			FileName: file.FileName,
			Hash:     file.Hash,
			Status:   status,
		})
	}

	uploadId, err := logUpload(context, storeId, createStoreUploadRequest.Description, createStoreUploadRequest.BuildId, files)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Internal error while logging upload to DB"}), errors.New("Internal error while logging upload to DB")
	}

	createStoreUploadResponse.Id = uploadId

	log.Printf("Response: %v", createStoreUploadResponse)

	return openapi.Response(http.StatusOK, createStoreUploadResponse), nil
}

func logUpload(ctx context.Context, storeId string, description string, buildId string, files []FileDBEntry) (string, error) {

	uploadContent := map[string]interface{}{
		"description": description,
		"buildId":     buildId,
		"files":       files,
		"timestamp":   time.Now().Format(time.RFC3339),
		"status":      StoreUploadEntry_Status_InProgress,
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
		storeDoc, err := tx.Get(storeDocRef)
		if err != nil {
			return err
		}
		storeEntry := &StoreEntry{}
		if err := storeDoc.DataTo(storeEntry); err != nil {
			return err
		}

		newUploadId = storeEntry.LatestUploadId + 1
		storeEntry.LatestUploadId = newUploadId

		err = tx.Set(storeDocRef, storeEntry)
		if err != nil {
			return err
		}

		uploadDocRef := storeDocRef.Collection(storeUploadsCollectionName).Doc(fmt.Sprint(newUploadId))

		err = tx.Create(uploadDocRef, uploadContent)

		return err
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
