package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type storageRequestHandler struct {
	StorageBucketHost string
}

func getStorageBucketURL(host string, path string) string {
	return fmt.Sprintf("%s%s", host, path)
}

func (handler *storageRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	storageBucketURL := getStorageBucketURL(handler.StorageBucketHost, r.URL.Path)

	log.Printf("Path %v redirected to %v", r.URL.Path, storageBucketURL)
	http.Redirect(w, r, storageBucketURL, http.StatusTemporaryRedirect)
}

var storageBucketHost string = "http://localhost:9000"

func main() {

	// Initialize template parameters.
	service := os.Getenv("K_SERVICE")
	if service == "" {
		service = "???"
	}

	revision := os.Getenv("K_REVISION")
	if revision == "" {
		revision = "???"
	}

	// Define HTTP server.
	http.Handle("/", &storageRequestHandler{StorageBucketHost: storageBucketHost})

	// PORT environment variable is provided by Cloud Run.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Print("Hello from Cloud Run! The container started successfully and is listening for HTTP requests on $PORT")
	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
