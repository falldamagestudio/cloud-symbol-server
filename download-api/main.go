package hello

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/storage"
)

type storageRequestHandler struct {
	StorageBucketHost string
	BucketName        string
}

func getStorageBucketURL(host string, path string) string {
	return fmt.Sprintf("%s%s", host, path)
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {

	symbolStoreBucketName := os.Getenv("SYMBOL_STORE_BUCKET_NAME")
	if symbolStoreBucketName == "" {
		log.Print("No storage bucket configured")
		http.Error(w, "No storage bucket configured", http.StatusInternalServerError)
		return
	}

	// Settings for live testing against GCS (test env)
	// handler := &storageRequestHandler{
	// 	StorageBucketHost: "https://storage.googleapis.com/test-cloud-symbol-store-symbols",
	// 	BucketName:        symbolStoreBucketName,
	// }

	// Settings for local testing
	handler := &storageRequestHandler{
		StorageBucketHost: "http://localhost:9000/",
		BucketName:        symbolStoreBucketName,
	}

	storageClient, err := storage.NewClient(r.Context())
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		http.Error(w, "Unable to create storageClient", http.StatusInternalServerError)
		return
	}

	// Paths are on the format "/folder/filename"
	//  but the GCS API wants a file path on the format, "folder/filename"

	path := strings.TrimPrefix(r.URL.Path, "/")

	// Validate whether object exists in bucket
	// This will talk to the Cloud Storage APIs

	log.Printf("Validating whether object %v does exists in bucket %v", path, handler.BucketName)

	_, err = storageClient.Bucket(handler.BucketName).Object(path).Attrs(r.Context())
	if err != nil {
		log.Printf("Object %v does not exist in bucket %v", path, handler.BucketName)
		http.Error(w, fmt.Sprintf("Object %v does not exist in bucket", path), http.StatusNotFound)
		return
	}

	// Object exists in bucket; respond with a redirect URL

	log.Printf("Object %v exists in bucket %v, preparing a redirect", path, handler.BucketName)

	storageBucketURL := getStorageBucketURL(handler.StorageBucketHost, path)

	log.Printf("Path %v redirected to %v", path, storageBucketURL)
	http.Redirect(w, r, storageBucketURL, http.StatusTemporaryRedirect)
}
