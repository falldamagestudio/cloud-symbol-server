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

func (s *ApiService) CreateTransaction(context context.Context, storeId string, uploadTransactionRequest openapi.UploadTransactionRequest) (openapi.ImplResponse, error) {

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

	uploadTransactionResponse := openapi.UploadTransactionResponse{}

	for _, uploadFileRequest := range uploadTransactionRequest.Files {

		// Validate whether object exists in bucket
		// This will talk to the Cloud Storage APIs

		path := uploadFileRequestToPath(storeId, uploadFileRequest)
		log.Printf("Validating whether object %v does exists in bucket %v", path, symbolStoreBucketName)
		_, err = storageClient.Bucket(symbolStoreBucketName).Object(path).Attrs(context)
		if err != nil {
			log.Printf("Object %v does not exist in bucket %v, preparing a redirect", path, symbolStoreBucketName)

			// Object does not exist in bucket; determine upload URL

			objectURL := ""

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

			uploadTransactionResponse.Files = append(uploadTransactionResponse.Files, openapi.UploadFileResponse{
				FileName: uploadFileRequest.FileName,
				Hash:     uploadFileRequest.Hash,
				Url:      objectURL,
			})

		} else {
			log.Printf("Object %v already exists in bucket %v", path, symbolStoreBucketName)
		}

	}

	// Log transaction to Firestore DB

	transactionId, err := logTransaction(context, storeId, uploadTransactionRequest, uploadTransactionResponse)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Internal error while logging transaction to DB"}), errors.New("Internal error while logging transaction to DB")
	}

	uploadTransactionResponse.Id = transactionId

	log.Printf("Response: %v", uploadTransactionResponse)

	return openapi.Response(http.StatusOK, uploadTransactionResponse), nil
}

func logTransaction(ctx context.Context, storeId string, uploadTransactionRequest openapi.UploadTransactionRequest, uploadTransactionResponse openapi.UploadTransactionResponse) (string, error) {

	transactionContent := map[string]interface{}{
		"description": uploadTransactionRequest.Description,
		"buildId":     uploadTransactionRequest.BuildId,
		"files":       uploadTransactionRequest.Files,
		"timestamp":   time.Now().Format(time.RFC3339),
	}

	log.Printf("Writing transaction to database: %v", transactionContent)

	firestoreClient, err := firestoreClient(ctx)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return "", err
	}

	storeDocRef := firestoreClient.Collection(storesCollectionName).Doc(storeId)

	newTransactionId := int64(0)

	err = firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		storeDoc, err := tx.Get(storeDocRef)
		if err != nil {
			return err
		}
		storeEntry := &StoreEntry{}
		if err := storeDoc.DataTo(storeEntry); err != nil {
			return err
		}

		newTransactionId = storeEntry.LatestTransactionId + 1
		storeEntry.LatestTransactionId = newTransactionId

		err = tx.Set(storeDocRef, storeEntry)
		if err != nil {
			return err
		}

		transactionDocRef := storeDocRef.Collection(storeUploadsCollectionName).Doc(fmt.Sprint(newTransactionId))

		err = tx.Create(transactionDocRef, transactionContent)

		return err
	})
	if err != nil {
		// Handle any errors appropriately in this section.
		log.Printf("An error has occurred: %s", err)
	}

	if err != nil {
		log.Printf("Error when logging transaction, err = %v", err)
		return "", err
	}

	log.Printf("Transaction is given ID %v", newTransactionId)

	return fmt.Sprint(newTransactionId), nil
}
