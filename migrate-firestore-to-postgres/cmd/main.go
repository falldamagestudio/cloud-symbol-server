package main

import (
	"context"
	"log"

	migrate_firestore_to_postgres "github.com/falldamagestudio/cloud-symbol-server/migrate-firestore-to-postgres"
)

func main() {
	const sourceGcpProjectId = "test-cloud-symbol-server"

	firestoreClient, err := migrate_firestore_to_postgres.FirestoreClient(context.Background(), sourceGcpProjectId)

	if err != nil {
		log.Printf("Unable to create firestore client: %v", err)
		panic("nope")
	}

	err = migrate_firestore_to_postgres.MigratePATs(context.Background(), firestoreClient)
	if err != nil {
		log.Printf("Error while migrating PATs: %v", err)
		panic("nope")
	}

}
