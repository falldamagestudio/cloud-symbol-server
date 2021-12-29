package hello

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {

	signedURLExpirationSeconds := 15 * 60

	// This is only set when the service is configured to run against a
	//  local emulator. When run against the real Cloud Storage APIs,
	//  the environment variable will be empty.
	storageEmulatorHost := os.Getenv("STORAGE_EMULATOR_HOST")

	// Use signed URLs only when talking to the real Cloud Storage APIs
	// Otherwise, create public, unsigned URLs directly to the storage service
	//
	// The Cloud Storage SDK has support for working against local emulators,
	//  via the STORAGE_EMULATOR_HOST setting. However, this setting does not
	//  work properly for the SignedURL() functions when using local emulators:
	// The SignURL() function will always return URLs that point to the real
	//   Cloud Storage API, even when STORAGE_EMULATOR_HOST is set.
	// Because of this, when we use local emulators, we fall back to manually
	//  constructing download URLs.
	useSignedURLs := (storageEmulatorHost == "")

	symbolStoreBucketName := os.Getenv("SYMBOL_STORE_BUCKET_NAME")
	if symbolStoreBucketName == "" {
		log.Print("No storage bucket configured")
		http.Error(w, "No storage bucket configured", http.StatusInternalServerError)
		return
	}

	gcpProjectId := os.Getenv("GCP_PROJECT_ID")
	if gcpProjectId == "" {
		log.Print("No GCP Project ID configured")
		http.Error(w, "No GCP Project ID configured", http.StatusInternalServerError)
		return
	}

	firestoreClient, err := firestore.NewClient(r.Context(), gcpProjectId)
	if err != nil {
		log.Printf("Unable to create firestoreClient: %v", err)
		http.Error(w, "Unable to create firestoreClient", http.StatusInternalServerError)
	}

	err = handlePATAuthentication(r, w, firestoreClient)

	if err != nil {
		return
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

	log.Printf("Validating whether object %v does exists in bucket %v", path, symbolStoreBucketName)

	_, err = storageClient.Bucket(symbolStoreBucketName).Object(path).Attrs(r.Context())
	if err != nil {
		log.Printf("Object %v does not exist in bucket %v", path, symbolStoreBucketName)
		http.Error(w, fmt.Sprintf("Object %v does not exist in bucket", path), http.StatusNotFound)
		return
	}

	// Object exists in bucket; respond with a redirect URL

	log.Printf("Object %v exists in bucket %v, preparing a redirect", path, symbolStoreBucketName)

	objectURL := ""

	if useSignedURLs {

		objectURL, err = storageClient.Bucket(symbolStoreBucketName).SignedURL(path, &storage.SignedURLOptions{
			Method:  "GET",
			Expires: time.Now().Add(time.Duration(signedURLExpirationSeconds) * time.Second),
		})

		if err != nil {
			log.Printf("Unable to create signed URL for %v: %v", path, err)
			http.Error(w, fmt.Sprintf("Unable to create signed URL for %v", path), http.StatusInternalServerError)
			return
		}

		log.Printf("Object %v has a signed URL %v, valid for %d seconds", path, objectURL, signedURLExpirationSeconds)

	} else {

		objectURL = fmt.Sprintf("http://%s/storage/v1/b/%s/o/%s?alt=media", storageEmulatorHost, symbolStoreBucketName, path)

		log.Printf("Object %v has a non-signed URL %v", path, objectURL)
	}

	log.Printf("Path %v redirected to %v", path, objectURL)
	http.Redirect(w, r, objectURL, http.StatusTemporaryRedirect)
}

func handlePATAuthentication(r *http.Request, w http.ResponseWriter, firestoreClient *firestore.Client) error {

	// Fetch email + PAT from Basic Authentication header of WWW request
	// The caller will have supplied this in its GET call
	//  by performaing GET to https://<email>:<pat>@<site>/<file path>
	// Also, the email/pat contents are expected to be URL encoded
	//  (so hello%40example.com will translate to hello@example.com, etc)

	email, pat, basicAuthPresent := r.BasicAuth()

	if !basicAuthPresent {

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		log.Print("Basic auth header (with email/token) not provided")
		http.Error(w, "Unauthorized; please provide email + personal access token using Basic Authentication", http.StatusUnauthorized)
		return errors.New("basic auth header (with email/token) not provided")

	}

	log.Printf("Basic Auth header present, email: %v, PAT: %v", email, pat)

	// Validate that email + PAT exists in database

	patDocRef := firestoreClient.Collection("users").Doc(email).Collection("pats").Doc(pat)

	if _, err := patDocRef.Get(r.Context()); err != nil {

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		log.Printf("Unable to find email %v, pat %v combination in database: %v", email, pat, err)
		http.Error(w, "Unauthorized; unable to find email / pat combination in database", http.StatusUnauthorized)
		return fmt.Errorf("unable to find email %v, pat %v combination in database: %v", email, pat, err)
	}

	log.Printf("Email/PAT pair %v, %v exist in database; authentication successful", email, pat)

	return nil
}
