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
	env_SOURCE_GCP_PROJECT_ID = "SOURCE_GCP_PROJECT_ID"
)

func getSourceGCPProjectID() (string, error) {

	sourceGcpProjectId := os.Getenv(env_SOURCE_GCP_PROJECT_ID)
	if sourceGcpProjectId == "" {
		return "", &ErrSourceGCPProjectID{}
	}

	return sourceGcpProjectId, nil
}

type ErrSourceGCPProjectID struct {
}

func (err ErrSourceGCPProjectID) Error() string {
	return "No source GCP project ID configured"
}

func FirestoreClient(context context.Context) (*firestore.Client, error) {

	gcpProjectId, err := getSourceGCPProjectID()
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
