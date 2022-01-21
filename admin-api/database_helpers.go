package admin_api

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"

	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	storesCollectionName       = "stores"
	storeUploadsCollectionName = "uploads"
)

func getStoresConfig(context context.Context) ([]string, error) {

	firestoreClient, err := firestoreClient(context)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return nil, err
	}

	storesDocSnapshots, err := firestoreClient.Collection(storesCollectionName).Documents(context).GetAll()
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

	storeDoc, err := firestoreClient.Collection(storesCollectionName).Doc(storeId).Get(context)

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

	transactionDoc, err := firestoreClient.Collection(storesCollectionName).Doc(storeId).Collection(storeUploadsCollectionName).Doc(transactionId).Get(context)

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

// Reference: https://firebase.google.com/docs/firestore/manage-data/delete-data#collections
func deleteAllDocumentsInCollection(ctx context.Context, client *firestore.Client, ref *firestore.CollectionRef, batchSize int) error {

	log.Printf("Deleting all documents in collection %v", ref.ID)

	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}
