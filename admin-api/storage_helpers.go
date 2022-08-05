package admin_api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
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

func deleteAllObjectsInStore(context context.Context, storageClient *storage.Client, bucketName string, storeId string) error {
	query := &storage.Query{
		Prefix: fmt.Sprintf("stores/%v/", storeId),
	}
	iter := storageClient.Bucket(bucketName).Objects(context, query)
	for {
		attrs, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		err = storageClient.Bucket(bucketName).Object(attrs.Name).Delete(context)
		if err != nil {
			return err
		}
	}

	return nil
}
