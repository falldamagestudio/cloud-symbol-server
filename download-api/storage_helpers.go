package download_api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func getStorageEndpoint() string {
	// This is only set when the service is configured to run against a
	//  local emulator. When run against the real Cloud Storage APIs,
	//  the environment variable will be empty.
	storageEmulatorHost := os.Getenv("STORAGE_EMULATOR_HOST")

	// The Firebase Storage emulator has a non-standard endpoint
	// Normally, the API endpoint would be at http[s]://<site>:<port>/storage/v1
	//  but the storage emulator has it at http[s]://<site>:<port>
	// Therefore we need to specify an explicit endpoint when using the emulator

	storageEndpoint := ""
	if storageEmulatorHost != "" {
		storageEndpoint = fmt.Sprintf("http://%s", storageEmulatorHost)
	}

	return storageEndpoint
}

func getStorageClient(context context.Context) (*storage.Client, string, error) {

	storageEndpoint := getStorageEndpoint()

	storageClientOpts := []option.ClientOption{}

	if storageEndpoint != "" {
		storageClientOpts = append(storageClientOpts, option.WithEndpoint(storageEndpoint))
	}

	storageClient, err := storage.NewClient(context, storageClientOpts...)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return nil, "", err
	}

	return storageClient, storageEndpoint, nil
}

func getSymbolStoreBucketName() (string, error) {

	symbolStoreBucketName := os.Getenv("SYMBOL_STORE_BUCKET_NAME")
	if symbolStoreBucketName == "" {
		log.Print("No storage bucket configured")
		return "", errors.New("No storage bucket configured")
	}

	return symbolStoreBucketName, nil
}
