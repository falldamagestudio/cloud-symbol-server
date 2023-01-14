package http_symbol_store

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/storage"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	cloud_storage "github.com/falldamagestudio/cloud-symbol-server/backend-api/cloud_storage"
	models "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/sql-db-models"
	postgres "github.com/falldamagestudio/cloud-symbol-server/backend-api/postgres"
)

type HttpSymbolStoreHandler struct {
}

func CreateHttpSymbolStoreHandler() *HttpSymbolStoreHandler {
	return &HttpSymbolStoreHandler{}
}

func (*HttpSymbolStoreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	signedURLExpirationSeconds := 15 * 60

	storageClient, err := cloud_storage.GetStorageClient(r.Context())
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		http.Error(w, "Unable to create storageClient", http.StatusInternalServerError)
		return
	}

	symbolStoreBucketName, err := cloud_storage.GetSymbolStoreBucketName()
	if err != nil {
		log.Print("No storage bucket configured")
		http.Error(w, "No storage bucket configured", http.StatusInternalServerError)
		return
	}

	storeIds, err := GetStores(r.Context())
	if err != nil {
		log.Printf("Error while decoding local stores configuration: %v", err)
		http.Error(w, "Error while decoding local stores configuration", http.StatusInternalServerError)
		return
	}

	// Paths are on the format "/folder/filename"
	//  but the GCS API wants a file path on the format, "folder/filename"

	path := strings.TrimPrefix(r.URL.Path, "/")

	// Validate whether object exists in bucket
	// This will talk to the Cloud Storage APIs

	log.Printf("Validating whether object %v does exists in bucket %v", path, symbolStoreBucketName)

	fullPath := findObjectInStores(storageClient, symbolStoreBucketName, storeIds, path, r.Context())
	if fullPath == nil {
		log.Printf("Object %v does not exist in any store in bucket %v", path, symbolStoreBucketName)
		http.Error(w, fmt.Sprintf("Object %v does not exist in any store in bucket", path), http.StatusNotFound)
		return
	}

	// Object exists in bucket; respond with a redirect URL

	log.Printf("Preparing a redirect for %v", path)

	objectURL, err := storageClient.Bucket(symbolStoreBucketName).SignedURL(*fullPath, &storage.SignedURLOptions{
		Method:  "GET",
		Expires: time.Now().Add(time.Duration(signedURLExpirationSeconds) * time.Second),
	})

	if err != nil {
		log.Printf("Unable to create signed URL for %v: %v", *fullPath, err)
		http.Error(w, fmt.Sprintf("Unable to create signed URL for %v", *fullPath), http.StatusInternalServerError)
		return
	}

	log.Printf("Object %v has a signed URL %v, valid for %d seconds", *fullPath, objectURL, signedURLExpirationSeconds)

	log.Printf("Path %v redirected to %v", *fullPath, objectURL)
	http.Redirect(w, r, objectURL, http.StatusTemporaryRedirect)
}

func findObjectInStores(storageClient *storage.Client, symbolStoreBucketName string, stores []string, path string, context context.Context) *string {

	for _, storeId := range stores {

		fullPath := fmt.Sprintf("stores/%s/%s", storeId, path)

		_, err := storageClient.Bucket(symbolStoreBucketName).Object(fullPath).Attrs(context)
		if err != nil {
			log.Printf("Object %v does not exist in bucket %v store %v", path, symbolStoreBucketName, storeId)
		} else {
			log.Printf("Object %v exists in bucket %v store %v", path, symbolStoreBucketName, storeId)
			return &fullPath
		}
	}

	return nil
}

func GetStores(ctx context.Context) ([]string, error) {

	log.Printf("Getting store names")

	db := postgres.GetDB()
	if db == nil {
		return nil, errors.New("no DB")
	}

	// Fetch names of all stores
	stores, err := models.Stores(
		qm.Select(models.StoreColumns.Name), qm.OrderBy(models.StoreColumns.StoreID),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	// Convert DB query result to a plain array of strings
	var storeIds = make([]string, len(stores))

	for index, store := range stores {
		storeIds[index] = store.Name
	}

	return storeIds, nil
}
