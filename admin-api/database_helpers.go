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

type ErrEntryRef interface {
	Path() string
}

type ErrUnknown struct {
	EntryRef ErrEntryRef
	Inner    error
}

func (err ErrUnknown) Error() string {
	return fmt.Sprintf("Error while accessing %v; err = %v", err.EntryRef.Path(), err.Inner)
}

func (err ErrUnknown) Unwrap() error {
	return err.Inner
}

type ErrEntryNotFound struct {
	EntryRef ErrEntryRef
	Inner    error
}

func (err ErrEntryNotFound) Error() string {
	return fmt.Sprintf("%v not found; err = %v", err.EntryRef.Path(), err.Inner)
}

func (err ErrEntryNotFound) Unwrap() error {
	return err.Inner
}

type ErrDocToStructFailed struct {
	EntryRef ErrEntryRef
	Inner    error
}

func (err ErrDocToStructFailed) Error() string {
	return fmt.Sprintf("%v document to struct conversion failed; err = %v", err.EntryRef.Path(), err.Inner)
}

func (err ErrDocToStructFailed) Unwrap() error {
	return err.Inner
}

type ErrStoresRef struct {
}

func (ref ErrStoresRef) Path() string {
	return fmt.Sprintf("Stores")
}

type ErrStoreRef struct {
	StoreId string
}

func (ref ErrStoreRef) Path() string {
	return fmt.Sprintf("Store %v", ref.StoreId)
}

type ErrStoreUploadRef struct {
	StoreId  string
	UploadId string
}

func (ref ErrStoreUploadRef) Path() string {
	return fmt.Sprintf("Store %v / Upload %v", ref.StoreId, ref.UploadId)
}

type ErrStoreFileHashRef struct {
	StoreId string
	FileId  string
	HashId  string
}

func (ref ErrStoreFileHashRef) Path() string {
	return fmt.Sprintf("Store %v / Upload %v / File %v / Hash %v", ref.StoreId, ref.FileId, ref.HashId)
}

type ErrStoreFileHashUploadRef struct {
	StoreId  string
	FileId   string
	HashId   string
	UploadId string
}

func (ref ErrStoreFileHashUploadRef) Path() string {
	return fmt.Sprintf("Store %v / Upload %v / File %v / Hash %v / Upload %v", ref.StoreId, ref.FileId, ref.HashId, ref.UploadId)
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
		return nil, &ErrUnknown{EntryRef: &ErrStoresRef{}, Inner: err}
	}

	stores := make([]string, len(storesDocSnapshots))

	for storeIndex, storeDocSnapshot := range storesDocSnapshots {
		stores[storeIndex] = storeDocSnapshot.Ref.ID
	}

	return stores, nil
}

func getStoreUploadIds(context context.Context, storeId string) ([]string, error) {

	log.Printf("Fetching all upload document IDs for %v", storeId)

	firestoreClient, err := firestoreClient(context)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return nil, &ErrFirestore{Inner: err}
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

func getStoreEntry(client *firestore.Client, tx *firestore.Transaction, storeId string) (*StoreEntry, error) {
	storeDocRef := client.Collection(storesCollectionName).Doc(storeId)
	storeDoc, err := tx.Get(storeDocRef)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, &ErrEntryNotFound{EntryRef: &ErrStoreRef{StoreId: storeId}, Inner: err}
		} else {
			return nil, &ErrUnknown{EntryRef: &ErrStoreRef{StoreId: storeId}, Inner: err}
		}
	}

	var storeEntry StoreEntry
	if err = storeDoc.DataTo(&storeEntry); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, &ErrDocToStructFailed{EntryRef: &ErrStoreRef{StoreId: storeId}, Inner: err}
		} else {
			return nil, &ErrUnknown{EntryRef: &ErrStoreRef{StoreId: storeId}, Inner: err}
		}
	}

	return &storeEntry, nil
}

func updateStoreEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, storeEntry *StoreEntry) error {
	storeUploadDocRef := client.Collection(storesCollectionName).Doc(storeId)
	err := tx.Set(storeUploadDocRef, storeEntry)
	if err != nil {
		return &ErrUnknown{EntryRef: &ErrStoreRef{StoreId: storeId}, Inner: err}
	}
	return nil
}

func createStoreUploadEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, uploadId int64, uploadContent *StoreUploadEntry) error {
	storeUploadDocRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeUploadsCollectionName).Doc(fmt.Sprint(uploadId))
	err := tx.Create(storeUploadDocRef, uploadContent)
	if err != nil {
		return &ErrUnknown{EntryRef: &ErrStoreUploadRef{StoreId: storeId, UploadId: fmt.Sprint(uploadId)}, Inner: err}
	}
	return nil
}

func getStoreUploadEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, uploadId string) (*StoreUploadEntry, error) {

	storeUploadRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeUploadsCollectionName).Doc(uploadId)
	storeUploadDoc, err := tx.Get(storeUploadRef)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, &ErrEntryNotFound{EntryRef: &ErrStoreUploadRef{StoreId: storeId, UploadId: uploadId}, Inner: err}
		} else {
			return nil, &ErrUnknown{EntryRef: &ErrStoreUploadRef{StoreId: storeId, UploadId: uploadId}, Inner: err}
		}
	}

	var storeUploadEntry StoreUploadEntry
	if err = storeUploadDoc.DataTo(&storeUploadEntry); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, &ErrEntryNotFound{EntryRef: &ErrStoreUploadRef{StoreId: storeId, UploadId: uploadId}, Inner: err}
		} else {
			return nil, &ErrUnknown{EntryRef: &ErrStoreUploadRef{StoreId: storeId, UploadId: uploadId}, Inner: err}
		}
	}

	return &storeUploadEntry, nil
}

func updateStoreUploadEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, uploadId string, storeUploadEntry *StoreUploadEntry) error {

	storeUploadRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeUploadsCollectionName).Doc(uploadId)
	err := tx.Set(storeUploadRef, storeUploadEntry)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return &ErrEntryNotFound{EntryRef: &ErrStoreUploadRef{StoreId: storeId, UploadId: uploadId}, Inner: err}
		} else {
			return &ErrUnknown{EntryRef: &ErrStoreUploadRef{StoreId: storeId, UploadId: uploadId}, Inner: err}
		}
	}

	return nil
}

func updateStoreFileHashEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, fileId string, hashId string, content *StoreFileHashEntry) error {
	storeFileHashDocRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Doc(fileId).Collection(storeFileHashesCollectionName).Doc(hashId)
	err := tx.Set(storeFileHashDocRef, content)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return &ErrEntryNotFound{EntryRef: &ErrStoreFileHashRef{StoreId: storeId, FileId: fileId, HashId: hashId}, Inner: err}
		} else {
			return &ErrUnknown{EntryRef: &ErrStoreFileHashRef{StoreId: storeId, FileId: fileId, HashId: hashId}, Inner: err}
		}
	}
	return nil
}

func createStoreFileHashUploadEntry(client *firestore.Client, tx *firestore.Transaction, storeId string, fileId string, hashId string, uploadId int64, content *StoreFileHashUploadEntry) error {
	storeFileHashUploadDocRef := client.Collection(storesCollectionName).Doc(storeId).Collection(storeFilesCollectionName).Doc(fileId).Collection(storeFileHashesCollectionName).Doc(hashId).Collection(storeFileHashUploadsCollectionName).Doc(fmt.Sprint(uploadId))
	err := tx.Create(storeFileHashUploadDocRef, content)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return &ErrEntryNotFound{EntryRef: &ErrStoreFileHashUploadRef{StoreId: storeId, FileId: fileId, HashId: hashId, UploadId: fmt.Sprint(uploadId)}, Inner: err}
		} else {
			return &ErrUnknown{EntryRef: &ErrStoreFileHashUploadRef{StoreId: storeId, FileId: fileId, HashId: hashId, UploadId: fmt.Sprint(uploadId)}, Inner: err}
		}
	}
	return nil
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
