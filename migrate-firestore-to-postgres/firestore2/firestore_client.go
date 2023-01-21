package firestore2

import (
	"context"
	"errors"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

const (
	// GCP project ID containing Firestore DB
	// Example: test-cloud-symbol-server
	env_GCP_PROJECT_ID = "GCP_PROJECT_ID"
)

func getGCPProjectID() (string, error) {

	gcpProjectId := os.Getenv(env_GCP_PROJECT_ID)
	if gcpProjectId == "" {
		return "", &ErrGCPProjectID{}
	}

	return gcpProjectId, nil
}

type ErrGCPProjectID struct {
}

func (err ErrGCPProjectID) Error() string {
	return "No GCP project ID configured"
}

func FirestoreClient(context context.Context) (*firestore.Client, error) {

	gcpProjectId, err := getGCPProjectID()
	if err != nil {
		return nil, err
	}

	firestoreClient, err := firestore.NewClient(context, gcpProjectId)
	if err != nil {
		log.Printf("Unable to create firestoreClient: %v", err)
		return nil, errors.New("unable to create firestoreClient")
	}

	return firestoreClient, nil
}
