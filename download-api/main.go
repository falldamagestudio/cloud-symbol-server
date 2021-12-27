package hello

import (
	"fmt"
	"log"
	"net/http"

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

	handler := &storageRequestHandler{
		StorageBucketHost: "http://localhost:9000",
		BucketName:        "example-bucket",
	}

	storageClient, err := storage.NewClient(r.Context())
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		http.Error(w, "Unable to create storageClient", http.StatusInternalServerError)
		return
	}

	// Validate whether object exists in bucket
	// This will talk to the Cloud Storage APIs

	log.Printf("Validating whether object %v does exists in bucket %v", r.URL.Path, handler.BucketName)

	_, err = storageClient.Bucket(handler.BucketName).Object(r.URL.Path).Attrs(r.Context())
	if err != nil {
		log.Printf("Object %v does not exist in bucket %v", r.URL.Path, handler.BucketName)
		http.Error(w, fmt.Sprintf("Object %v does not exist in bucket", r.URL.Path), http.StatusNotFound)
		return
	}

	// Object exists in bucket; respond with a redirect URL

	log.Printf("Object %v exists in bucket %v, preparing a redirect", r.URL.Path, handler.BucketName)

	storageBucketURL := getStorageBucketURL(handler.StorageBucketHost, r.URL.Path)

	log.Printf("Path %v redirected to %v", r.URL.Path, storageBucketURL)
	http.Redirect(w, r, storageBucketURL, http.StatusTemporaryRedirect)
}
