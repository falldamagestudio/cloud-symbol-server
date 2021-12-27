package hello

import (
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
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

	// Create client as usual.
	storageClient, err := storage.NewClient(r.Context())
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		http.Error(w, "Unable to create storageClient", http.StatusInternalServerError)
		return
	}

	log.Print("Querying for objects")

	query := &storage.Query{Prefix: ""}
	it := storageClient.Bucket(handler.BucketName).Objects(r.Context(), query)
	for {
		_, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Print("iteration done")

	storageBucketURL := getStorageBucketURL(handler.StorageBucketHost, r.URL.Path)

	log.Printf("Path %v redirected to %v", r.URL.Path, storageBucketURL)
	http.Redirect(w, r, storageBucketURL, http.StatusTemporaryRedirect)
}
