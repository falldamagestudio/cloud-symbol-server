package admin_api

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func getStoreDoc(context context.Context, storeId string) (*firestore.DocumentSnapshot, error) {

	log.Printf("Fetching store document for %v", storeId)

	firestoreClient, err := firestoreClient(context)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return nil, err
	}

	storeDoc, err := firestoreClient.Collection("stores").Doc(storeId).Get(context)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		} else {
			log.Printf("Unable to fetch store document for %v, err = %v", storeId, err)
			return nil, err
		}
	}

	return storeDoc, nil
}

func getTransactionDoc(context context.Context, storeId string, transactionId string) (*firestore.DocumentSnapshot, error) {

	log.Printf("Fetching transaction document for %v/%v", storeId, transactionId)

	firestoreClient, err := firestoreClient(context)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return nil, err
	}

	transactionDoc, err := firestoreClient.Collection("stores").Doc(storeId).Collection("transactions").Doc(transactionId).Get(context)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		} else {
			log.Printf("Unable to fetch transaction document for %v/%v, err = %v", storeId, transactionId, err)
			return nil, err
		}
	}

	return transactionDoc, nil
}
