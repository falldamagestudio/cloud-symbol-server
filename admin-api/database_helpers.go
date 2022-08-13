package admin_api

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"

	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	storesCollectionName               = "stores"
	storeUploadsCollectionName         = "uploads"
	storeFilesCollectionName           = "files"
	storeFileHashesCollectionName      = "hashes"
	storeFileHashUploadsCollectionName = "uploads"
)

func runDBTransaction(ctx context.Context, f func(context.Context, *firestore.Client, *firestore.Transaction) error) error {

	firestoreClient, err := firestoreClient(ctx)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return err
	}

	wrappedF := func(ctx context.Context, tx *firestore.Transaction) error {
		return f(ctx, firestoreClient, tx)
	}

	err = firestoreClient.RunTransaction(ctx, wrappedF)
	if err != nil {
		log.Printf("Transaction failed: %v", err)
		return err
	}

	return nil
}

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

func getStoreUploadIds(context context.Context, storeId string) ([]string, error) {

	log.Printf("Fetching all upload document IDs for %v", storeId)

	firestoreClient, err := firestoreClient(context)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return nil, err
	}

	uploadDocsIterator := firestoreClient.Collection(storesCollectionName).Doc(storeId).Collection(storeUploadsCollectionName).Documents(context)
	uploadDocs, err := uploadDocsIterator.GetAll()
	if err != nil {
		log.Printf("Error while fetching documents in %v: %v", storeId, err)
		return nil, err
	}

	uploadIds := make([]string, len(uploadDocs))
	for uploadDocIndex, uploadDoc := range uploadDocs {
		uploadIds[uploadDocIndex] = uploadDoc.Ref.ID
	}
	return uploadIds, nil
}

func getStoreUploadRef(context context.Context, storeId string, uploadId string) (*firestore.DocumentRef, error) {

	log.Printf("Creating upload ref for %v/%v", storeId, uploadId)

	firestoreClient, err := firestoreClient(context)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return nil, err
	}

	uploadRef := firestoreClient.Collection(storesCollectionName).Doc(storeId).Collection(storeUploadsCollectionName).Doc(uploadId)

	return uploadRef, nil
}

func getStoreUploadDoc(context context.Context, storeId string, uploadId string) (*firestore.DocumentSnapshot, error) {

	log.Printf("Fetching upload document for %v/%v", storeId, uploadId)

	firestoreClient, err := firestoreClient(context)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return nil, err
	}

	uploadDoc, err := firestoreClient.Collection(storesCollectionName).Doc(storeId).Collection(storeUploadsCollectionName).Doc(uploadId).Get(context)

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		} else {
			log.Printf("Unable to fetch upload document for %v/%v, err = %v", storeId, uploadId, err)
			return nil, err
		}
	}

	return uploadDoc, nil
}

func getStoreEntry(client *firestore.Client, tx *firestore.Transaction, storeId string) (*StoreEntry, error) {
	storeDocRef := client.Collection(storesCollectionName).Doc(storeId)
	storeDoc, err := tx.Get(storeDocRef)
	if err != nil {
		log.Printf("Unable to fetch store document for %v, err = %v", storeId, err)
		return nil, err
	}

	log.Printf("Extracting store doc data")
	var storeEntry StoreEntry
	if err = storeDoc.DataTo(&storeEntry); err != nil {
		log.Printf("Extracting store doc data failed")
		return nil, err
	}

	return &storeEntry, nil
}

func updateStoreEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, storeEntry *StoreEntry) error {
	storeUploadDocRef := client.Collection(storesCollectionName).Doc(storeId)
	err := tx.Set(storeUploadDocRef, storeEntry)
	return err
}

func createStoreUploadEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, uploadId int64, uploadContent *StoreUploadEntry) error {
	storeUploadDocRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeUploadsCollectionName).Doc(fmt.Sprint(uploadId))
	err := tx.Create(storeUploadDocRef, uploadContent)
	return err
}

func getStoreUploadEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, uploadId string) (*StoreUploadEntry, error) {

	storeUploadRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeUploadsCollectionName).Doc(uploadId)
	storeUploadDoc, err := tx.Get(storeUploadRef)
	if err != nil {
		log.Printf("Unable to fetch upload document for %v/%v, err = %v", storeId, uploadId, err)
		return nil, err
	}

	log.Printf("Extracting upload doc data")
	var storeUploadEntry StoreUploadEntry
	if err = storeUploadDoc.DataTo(&storeUploadEntry); err != nil {
		log.Printf("Extracting upload doc data failed")
		return nil, err
	}

	return &storeUploadEntry, nil
}

func updateStoreUploadEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, uploadId string, storeUploadEntry *StoreUploadEntry) error {

	storeUploadRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeUploadsCollectionName).Doc(uploadId)
	err := tx.Set(storeUploadRef, storeUploadEntry)
	if err != nil {
		log.Printf("Unable to update store upload entry for %v/%v, err = %v", storeId, uploadId, err)
		return err
	}

	return nil
}

func updateStoreFileHashEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, fileId string, hashId string, content *StoreFileHashEntry) error {
	storeFileHashDocRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Doc(fileId).Collection(storeFileHashesCollectionName).Doc(hashId)
	err := tx.Set(storeFileHashDocRef, content)
	return err
}

func createStoreFileHashUploadEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, fileId string, hashId string, uploadId int64, content *StoreFileHashUploadEntry) error {
	storeFileHashUploadDocRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Doc(fileId).Collection(storeFileHashesCollectionName).Doc(hashId).Collection(storeFileHashUploadsCollectionName).Doc(fmt.Sprint(uploadId))
	err := tx.Create(storeFileHashUploadDocRef, content)
	return err
}

type DeleteDocumentsRecursiveState struct {
	Ctx             *context.Context
	Client          *firestore.Client
	Recurse         bool
	Batch           *firestore.WriteBatch
	NumItemsInBatch int
}

const DeleteDocumentsRecursiveState_MaxItemsInBatch = 100

func commitStateBatch(state *DeleteDocumentsRecursiveState) error {

	if state.Batch != nil {
		_, err := state.Batch.Commit(*state.Ctx)
		state.Batch = nil
		state.NumItemsInBatch = 0
		return err
	} else {
		return nil
	}
}

func addItemToStateBatch(state *DeleteDocumentsRecursiveState, documentRef *firestore.DocumentRef) error {
	if state.Batch == nil {
		state.Batch = state.Client.Batch()
	}

	log.Printf("Adding item %v for batch deletion", documentRef.Path)
	state.Batch.Delete(documentRef)
	state.NumItemsInBatch++

	if state.NumItemsInBatch >= DeleteDocumentsRecursiveState_MaxItemsInBatch {
		err := commitStateBatch(state)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteAllDocumentsInCollectionRecurse(state *DeleteDocumentsRecursiveState, collectionRef *firestore.CollectionRef) error {

	// Get all documents in collection
	documentRefIter := collectionRef.DocumentRefs(*state.Ctx)

	for {
		documentRef, err := documentRefIter.Next()
		if err == iterator.Done {
			break
		}

		_, err = documentRef.Get(*state.Ctx)
		if err == nil {
			err = addItemToStateBatch(state, documentRef)
			if err != nil {
				return err
			}

		} else if status.Code(err) != codes.NotFound {
			return err
		}

		if state.Recurse {

			err = deleteAllCollectionsInDocumentRecurse(state, documentRef)
			if err != nil {
				return err
			}

			subCollectionRefIter := documentRef.Collections(*state.Ctx)

			for {
				subCollectionRef, err := subCollectionRefIter.Next()
				if err == iterator.Done {
					break
				} else if err != nil {
					return err
				}

				err = deleteAllDocumentsInCollectionRecurse(state, subCollectionRef)
				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}

func deleteAllCollectionsInDocumentRecurse(state *DeleteDocumentsRecursiveState, documentRef *firestore.DocumentRef) error {

	collectionRefIter := documentRef.Collections(*state.Ctx)

	for {
		collectionRef, err := collectionRefIter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			return err
		}

		err = deleteAllDocumentsInCollectionRecurse(state, collectionRef)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete all documents from a Firestore collection
// If recurse == true, also delete any documents within sub-collections to arbitrary depth
// This will find and delete orphaned subcollection documents as well
func deleteAllDocumentsInCollection(ctx context.Context, client *firestore.Client, ref *firestore.CollectionRef, recurse bool) error {

	state := &DeleteDocumentsRecursiveState{
		Ctx:             &ctx,
		Client:          client,
		Recurse:         recurse,
		Batch:           nil,
		NumItemsInBatch: 0,
	}

	err := deleteAllDocumentsInCollectionRecurse(state, ref)
	if err != nil {
		return err
	}
	err = commitStateBatch(state)
	return err
}

// Delete a document from a Firestore database
// If recurse == true, also delete any documents within sub-collections to arbitrary depth
// This will find and delete orphaned subcollection documents as well
func deleteDocument(ctx context.Context, client *firestore.Client, ref *firestore.DocumentRef, recurse bool) error {

	log.Printf("Deleting document %v; recurse: %v", ref.ID, recurse)

	state := &DeleteDocumentsRecursiveState{
		Ctx:             &ctx,
		Client:          client,
		Recurse:         recurse,
		Batch:           nil,
		NumItemsInBatch: 0,
	}

	err := addItemToStateBatch(state, ref)
	if err != nil {
		return err
	}

	if recurse {
		err := deleteAllCollectionsInDocumentRecurse(state, ref)
		if err != nil {
			return err
		}
	}

	err = commitStateBatch(state)
	return err
}
