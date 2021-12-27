package hello

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type storageRequestHandler struct {
	StorageBucketHost string
}

func getStorageBucketURL(host string, path string) string {
	return fmt.Sprintf("%s%s", host, path)
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {

	handler := &storageRequestHandler{"localhost:9000"}

	log.Print("Creating storage client")

	_ = os.Setenv("STORAGE_EMULATOR_HOST", "localhost:9000")

	// Create client as usual.
	storageClient, err := storage.NewClient(r.Context())
	if err != nil {
		log.Print("Unable to create storageClient")
		http.Error(w, "Unable to create storageClient", http.StatusInternalServerError)
		return
	}

	log.Print("Querying for objects")

	query := &storage.Query{Prefix: ""}
	it := storageClient.Bucket("example-bucket").Objects(r.Context(), query)
	log.Print("Query completed, iteration time")
	for {
		log.Printf("before next")
		attrs, err := it.Next()
		log.Printf("after next")
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Encountered item %v", attrs.Name)
		fmt.Fprintf(w, "Object in bucket: %v\n", attrs.Name)
	}
	log.Print("iteration done")

	storageBucketURL := getStorageBucketURL(handler.StorageBucketHost, r.URL.Path)

	log.Printf("Path %v redirected to %v", r.URL.Path, storageBucketURL)
	http.Redirect(w, r, storageBucketURL, http.StatusTemporaryRedirect)
}
