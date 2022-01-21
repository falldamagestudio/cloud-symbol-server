package download_api

import (
	"context"
	"log"
)

func getStoresConfig(context context.Context) ([]string, error) {

	firestoreClient, err := firestoreClient(context)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return nil, err
	}

	storesDocSnapshots, err := firestoreClient.Collection("stores").Documents(context).GetAll()
	if err != nil {
		log.Printf("Error when fetching stores, err = %v", err)
		return nil, err
	}

	stores := make([]string, len(storesDocSnapshots))

	for storeIndex, storeDocSnapshot := range storesDocSnapshots {
		stores[storeIndex] = storeDocSnapshot.Ref.ID
	}

	return stores, nil
}
