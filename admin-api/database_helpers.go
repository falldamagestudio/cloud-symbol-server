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

type ErrFirestore struct {
	Inner error
}

func (err ErrFirestore) Error() string {
	return fmt.Sprintf("Firestore error; err = %v", err.Inner)
}

func (err ErrFirestore) Unwrap() error {
	return err.Inner
}

func runDBTransaction(ctx context.Context, f func(context.Context, *firestore.Client, *firestore.Transaction) error) error {

	firestoreClient, err := firestoreClient(ctx)
	if err != nil {
		return &ErrFirestore{Inner: err}
	}

	wrappedF := func(ctx context.Context, tx *firestore.Transaction) error {
		return f(ctx, firestoreClient, tx)
	}

	err = firestoreClient.RunTransaction(ctx, wrappedF)
	if err != nil {
		return err
	}

	return nil
}

func runDBOperation(ctx context.Context, f func(context.Context, *firestore.Client) error) error {

	firestoreClient, err := firestoreClient(ctx)
	if err != nil {
		return &ErrFirestore{Inner: err}
	}

	wrappedF := func(ctx context.Context) error {
		return f(ctx, firestoreClient)
	}

	err = wrappedF(ctx)
	if err != nil {
		return err
	}

	return nil
}

func getStoreIds(ctx context.Context, client *firestore.Client) ([]string, error) {

	storesDocSnapshots, err := client.Collection(storesCollectionName).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	stores := make([]string, len(storesDocSnapshots))

	for storeIndex, storeDocSnapshot := range storesDocSnapshots {
		stores[storeIndex] = storeDocSnapshot.Ref.ID
	}

	return stores, nil
}

func getStoreFileIds(context context.Context, client *firestore.Client, storeId string) ([]string, error) {

	fileDocsIterator := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Documents(context)
	fileDocs, err := fileDocsIterator.GetAll()
	if err != nil {
		return nil, err
	}

	fileIds := make([]string, len(fileDocs))
	for fileDocIndex, fileDoc := range fileDocs {
		fileIds[fileDocIndex] = fileDoc.Ref.ID
	}
	return fileIds, nil
}

func getStoreUploadIds(context context.Context, client *firestore.Client, storeId string) ([]string, error) {

	uploadDocsIterator := client.Collection(storesCollectionName).Doc(storeId).Collection(storeUploadsCollectionName).Documents(context)
	uploadDocs, err := uploadDocsIterator.GetAll()
	if err != nil {
		return nil, err
	}

	uploadIds := make([]string, len(uploadDocs))
	for uploadDocIndex, uploadDoc := range uploadDocs {
		uploadIds[uploadDocIndex] = uploadDoc.Ref.ID
	}
	return uploadIds, nil
}

func getStoreEntry(client *firestore.Client, tx *firestore.Transaction, storeId string) (*StoreEntry, error) {
	storeDocRef := client.Collection(storesCollectionName).Doc(storeId)
	storeDoc, err := tx.Get(storeDocRef)
	if err != nil {
		return nil, err
	}

	var storeEntry StoreEntry
	if err = storeDoc.DataTo(&storeEntry); err != nil {
		return nil, err
	}

	return &storeEntry, nil
}

func createStoreEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, storeEntry *StoreEntry) error {
	storeUploadDocRef := client.Collection(storesCollectionName).Doc(storeId)
	err := tx.Create(storeUploadDocRef, storeEntry)
	return err
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
		return nil, err
	}

	var storeUploadEntry StoreUploadEntry
	if err = storeUploadDoc.DataTo(&storeUploadEntry); err != nil {
		return nil, err
	}

	return &storeUploadEntry, nil
}

func updateStoreUploadEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, uploadId string, storeUploadEntry *StoreUploadEntry) error {

	storeUploadRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeUploadsCollectionName).Doc(uploadId)
	err := tx.Set(storeUploadRef, storeUploadEntry)
	return err
}

func getStoreFileEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, fileId string) (*StoreFileEntry, error) {

	storeFileRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Doc(fileId)
	storeFileDoc, err := tx.Get(storeFileRef)
	if err != nil {
		return nil, err
	}

	var storeFileEntry StoreFileEntry
	if err = storeFileDoc.DataTo(&storeFileEntry); err != nil {
		return nil, err
	}

	return &storeFileEntry, nil
}

func updateStoreFileEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, fileId string, content *StoreFileEntry) error {
	storeFileDocRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Doc(fileId)
	err := tx.Set(storeFileDocRef, content)
	return err
}

func deleteStoreFileEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, fileId string) error {
	storeFileDocRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Doc(fileId)
	err := tx.Delete(storeFileDocRef)
	return err
}

func getStoreFileHashEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, fileId string, hashId string) (*StoreFileHashEntry, error) {

	storeFileHashRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Doc(fileId).Collection(storeFileHashesCollectionName).Doc(hashId)
	storeFileHashDoc, err := tx.Get(storeFileHashRef)
	if err != nil {
		return nil, err
	}

	var storeFileHashEntry StoreFileHashEntry
	if err = storeFileHashDoc.DataTo(&storeFileHashEntry); err != nil {
		return nil, err
	}

	return &storeFileHashEntry, nil
}

func updateStoreFileHashEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, fileId string, hashId string, content *StoreFileHashEntry) error {
	storeFileHashDocRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Doc(fileId).Collection(storeFileHashesCollectionName).Doc(hashId)
	err := tx.Set(storeFileHashDocRef, content)
	return err
}

func deleteStoreFileHashEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, fileId string, hashId string) error {
	storeFileHashDocRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Doc(fileId).Collection(storeFileHashesCollectionName).Doc(hashId)
	err := tx.Delete(storeFileHashDocRef)
	return err
}

func createStoreFileHashUploadEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, fileId string, hashId string, uploadId int64, content *StoreFileHashUploadEntry) error {
	storeFileHashUploadDocRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Doc(fileId).Collection(storeFileHashesCollectionName).Doc(hashId).Collection(storeFileHashUploadsCollectionName).Doc(fmt.Sprint(uploadId))
	err := tx.Create(storeFileHashUploadDocRef, content)
	return err
}

func deleteStoreFileHashUploadEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, fileId string, hashId string, uploadId string) error {
	storeFileHashUploadDocRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Doc(fileId).Collection(storeFileHashesCollectionName).Doc(hashId).Collection(storeFileHashUploadsCollectionName).Doc(uploadId)
	err := tx.Delete(storeFileHashUploadDocRef)
	return err
}

func getStoreFileHashUploadIds(ctx context.Context, client *firestore.Client, storeId string, fileId string, hashId string) ([]string, error) {

	storeFileHashUploadDocSnapshots, err := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Doc(fileId).Collection(storeFileHashesCollectionName).Doc(hashId).Collection(storeFileHashUploadsCollectionName).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	uploadIds := make([]string, len(storeFileHashUploadDocSnapshots))

	for uploadIndex, storeFileHashUploadDocSnapshot := range storeFileHashUploadDocSnapshots {
		uploadIds[uploadIndex] = storeFileHashUploadDocSnapshot.Ref.ID
	}

	return uploadIds, nil
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
