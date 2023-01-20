package main

import (
	"context"
	"fmt"
	"os"

	migrate_firestore_to_postgres "github.com/falldamagestudio/cloud-symbol-server/migrate-firestore-to-postgres"
	postgres "github.com/falldamagestudio/cloud-symbol-server/migrate-firestore-to-postgres/postgres"
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

func main() {
	sourceGcpProjectId, err := getSourceGCPProjectID()
	if err != nil {
		panic(err)
	}

	postgres.InitSQL()

	firestoreClient, err := migrate_firestore_to_postgres.FirestoreClient(context.Background(), sourceGcpProjectId)

	if err != nil {
		panic(fmt.Sprintf("Unable to create firestore client: %v", err))
	}

	err = migrate_firestore_to_postgres.MigratePATs(context.Background(), firestoreClient)
	if err != nil {
		panic(fmt.Sprintf("Error while migrating PATs: %v", err))
	}

}
