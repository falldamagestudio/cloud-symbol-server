package cloud_storage

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type ErrStorageClient struct {
	Inner error
}

func (err ErrStorageClient) Error() string {
	return fmt.Sprintf("Unable to create storage client; err = %v", err.Inner)
}

func (err ErrStorageClient) Unwrap() error {
	return err.Inner
}

type ErrStoreBucketName struct {
}

func (err ErrStoreBucketName) Error() string {
	return "No store bucket name configured"
}

func GetStorageClient(context context.Context) (*storage.Client, error) {

	storageClientOpts := []option.ClientOption{}

	storageClient, err := storage.NewClient(context, storageClientOpts...)
	if err != nil {
		return nil, &ErrStorageClient{Inner: err}
	}

	return storageClient, nil
}

func GetSymbolStoreBucketName() (string, error) {

	symbolStoreBucketName := os.Getenv("SYMBOL_STORE_BUCKET_NAME")
	if symbolStoreBucketName == "" {
		return "", &ErrStoreBucketName{}
	}

	return symbolStoreBucketName, nil
}

func DeleteAllObjectsInStore(context context.Context, storageClient *storage.Client, storeId string) error {

	bucketName, err := GetSymbolStoreBucketName()
	if err != nil {
		return err
	}

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

func storefileNameHashToPath(storeId string, fileName string, hash string) string {
	return fmt.Sprintf("stores/%s/%s/%s/%s", storeId, fileName, hash, fileName)
}

func DeleteObjectInStore(context context.Context, storageClient *storage.Client, storeId string, fileName string, hash string) error {

	bucketName, err := GetSymbolStoreBucketName()
	if err != nil {
		return err
	}

	path := storefileNameHashToPath(storeId, fileName, hash)
	err = storageClient.Bucket(bucketName).Object(path).Delete(context)
	return err
}
