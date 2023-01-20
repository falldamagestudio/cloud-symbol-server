package main

import (
	"context"
	"fmt"

	migrate_firestore_to_postgres "github.com/falldamagestudio/cloud-symbol-server/migrate-firestore-to-postgres"
	firestore2 "github.com/falldamagestudio/cloud-symbol-server/migrate-firestore-to-postgres/firestore2"
	postgres "github.com/falldamagestudio/cloud-symbol-server/migrate-firestore-to-postgres/postgres"
)

func main() {
	postgres.InitSQL()

	firestoreClient, err := firestore2.FirestoreClient(context.Background())

	if err != nil {
		panic(fmt.Sprintf("Unable to create firestore client: %v", err))
	}

	err = migrate_firestore_to_postgres.MigratePATs(context.Background(), firestoreClient)
	if err != nil {
		panic(fmt.Sprintf("Error while migrating PATs: %v", err))
	}

	err = migrate_firestore_to_postgres.MigrateStores(context.Background(), firestoreClient)
	if err != nil {
		panic(fmt.Sprintf("Error while migrating PATs: %v", err))
	}

}
