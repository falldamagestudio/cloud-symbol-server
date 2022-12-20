package helpers

import (
	"context"
	"errors"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

var localFirestoreClient *firestore.Client

func FirestoreClient(context context.Context) (*firestore.Client, error) {

	if localFirestoreClient == nil {

		gcpProjectId := os.Getenv("GCP_PROJECT_ID")
		if gcpProjectId == "" {
			log.Print("No GCP Project ID configured")
			return nil, errors.New("no GCP Project ID configured")
		}

		err := (error)(nil)
		localFirestoreClient, err = firestore.NewClient(context, gcpProjectId)
		if err != nil {
			log.Printf("Unable to create firestoreClient: %v", err)
			return nil, errors.New("unable to create firestoreClient")
		}
	}

	return localFirestoreClient, nil
}
