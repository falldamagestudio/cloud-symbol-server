package download_api

import (
	"context"
	"errors"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func getStorageClient(context context.Context) (*storage.Client, error) {

	storageClientOpts := []option.ClientOption{}

	storageClient, err := storage.NewClient(context, storageClientOpts...)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return nil, err
	}

	return storageClient, nil
}

func getSymbolStoreBucketName() (string, error) {

	symbolStoreBucketName := os.Getenv("SYMBOL_STORE_BUCKET_NAME")
	if symbolStoreBucketName == "" {
		log.Print("No storage bucket configured")
		return "", errors.New("No storage bucket configured")
	}

	return symbolStoreBucketName, nil
}
