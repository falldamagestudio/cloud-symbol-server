package migrate_firestore_to_postgres

import (
	"context"
	"errors"
	"log"

	"cloud.google.com/go/firestore"
)

func FirestoreClient(context context.Context, gcpProjectId string) (*firestore.Client, error) {

	firestoreClient, err := firestore.NewClient(context, gcpProjectId)
	if err != nil {
		log.Printf("Unable to create firestoreClient: %v", err)
		return nil, errors.New("unable to create firestoreClient")
	}

	return firestoreClient, nil
}
