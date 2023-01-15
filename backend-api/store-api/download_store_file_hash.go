package store_api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/storage"

	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	cloud_storage "github.com/falldamagestudio/cloud-symbol-server/backend-api/cloud_storage"
)

func GetStoreFileHashDownloadUrl(ctx context.Context, storeId string, fileId string, hashId string) (openapi.ImplResponse, error) {

	log.Printf("Get store file hash download URL; store %v, file %v, hash %v", storeId, fileId, hashId)

	signedURLExpirationSeconds := 15 * 60

	storageClient, err := cloud_storage.GetStorageClient(ctx)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return openapi.Response(http.StatusInternalServerError, "Unable to create storageClient"), err
	}

	symbolStoreBucketName, err := cloud_storage.GetSymbolStoreBucketName()
	if err != nil {
		log.Print("No storage bucket configured")
		return openapi.Response(http.StatusInternalServerError, "No storage bucket configured"), err
	}

	fullPath := fmt.Sprintf("stores/%s/%s/%s/%s", storeId, fileId, hashId, fileId)

	// Validate that object exists in bucket

	_, err = storageClient.Bucket(symbolStoreBucketName).Object(fullPath).Attrs(ctx)
	if err != nil {
		log.Printf("Object %v does not exist in bucket %v", fullPath, symbolStoreBucketName)
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Object %v does not exist in bucket %v", fullPath, symbolStoreBucketName)}), nil
	}

	// Object exists in bucket; respond with download URL

	log.Printf("Preparing a download URL for %v", fullPath)

	const method = "GET"

	objectURL, err := storageClient.Bucket(symbolStoreBucketName).SignedURL(fullPath, &storage.SignedURLOptions{
		Method:  method,
		Expires: time.Now().Add(time.Duration(signedURLExpirationSeconds) * time.Second),
	})

	if err != nil {
		log.Printf("Unable to create signed URL for %v: %v", fullPath, err)
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Unable to create signed URL for %v", fullPath)}), err
	}

	log.Printf("Object %v has a signed URL %v, valid for %d seconds", fullPath, objectURL, signedURLExpirationSeconds)

	getStoreFileHashDownloadUrlResponse := &openapi.GetStoreFileHashDownloadUrlResponse{
		Method: method,
		Url: objectURL,
	}

	log.Printf("Response: %v", getStoreFileHashDownloadUrlResponse)
	return openapi.Response(http.StatusOK, getStoreFileHashDownloadUrlResponse), nil
}
