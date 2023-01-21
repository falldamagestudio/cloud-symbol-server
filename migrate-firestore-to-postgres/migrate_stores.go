package migrate_firestore_to_postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	firestore2 "github.com/falldamagestudio/cloud-symbol-server/migrate-firestore-to-postgres/firestore2"
	postgres "github.com/falldamagestudio/cloud-symbol-server/migrate-firestore-to-postgres/postgres"
	models "github.com/falldamagestudio/cloud-symbol-server/migrate-firestore-to-postgres/generated/sql-db-models"
)

const (
	storesCollectionName = "stores"
	uploadsCollectionName = "uploads"
)

func deleteOldStore(ctx context.Context, tx *sql.Tx, storeName string) error {

	// Retrieve ID of store

	store, err := models.Stores(
		qm.Where(models.StoreTableColumns.Name+" = ?", storeName),
	).One(ctx, tx)

	if err == sql.ErrNoRows {
		// Store does not exist; do nothing
		return nil
	} else if err != nil {
		log.Printf("Error when locating store %v: %v", storeName, err)
		return err
	}

	storeId := store.StoreID

	// Delete upload-files referencing uploads in to-be-deleted store

	storeUploadFiles, err := models.StoreUploadFiles(
		qm.InnerJoin("cloud_symbol_server."+models.TableNames.StoreUploads+" on "+models.StoreUploadTableColumns.UploadID+" = "+models.StoreUploadFileTableColumns.UploadID),
		qm.Where(models.StoreUploadTableColumns.StoreID+" = ?", storeId),
		qm.For("update"),
		).All(ctx, tx)
	if (err != nil) && (err != sql.ErrNoRows) {
		log.Printf("Error when finding all upload-files referencing store %v: %v", storeName, err)
		return err
	}
	_, err = storeUploadFiles.DeleteAll(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all upload-files in store %v: %v", storeName, err)
		return err
	}

	// Delete uploads referencing to-be-deleted store

	_, err = models.StoreUploads(
		qm.Where(models.StoreUploadTableColumns.StoreID+" = ?", storeId),
	).DeleteAll(ctx, tx)
	if (err != nil) && (err != sql.ErrNoRows) {
		log.Printf("Error when deleting all uploads referencing store %v: %v", storeName, err)
		return err
	}

	// Delete all file-blobs in store
	storeFileBlobs, err := models.StoreFileBlobs(
		qm.InnerJoin("cloud_symbol_server."+models.TableNames.StoreFiles+" on "+models.StoreFileTableColumns.FileID+" = "+models.StoreFileBlobTableColumns.FileID),
		qm.Where(models.StoreFileTableColumns.StoreID+" = ?", store.StoreID),
		qm.For("update"),
	).All(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all file-blobs in store %v: %v", storeName, err)
		return err
	}
	_, err = storeFileBlobs.DeleteAll(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all file-blobs in store %v: %v", storeName, err)
		return err
	}

	// Delete all store-files in store
	storeFiles, err := models.StoreFiles(
		qm.Where(models.StoreFileTableColumns.StoreID+" = ?", store.StoreID),
		qm.For("update"),
	).All(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all store-files in store %v: %v", storeName, err)
		return err
	}
	_, err = storeFiles.DeleteAll(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all store-files in store %v: %v", storeName, err)
		return err
	}

	// Delete already-existing store

	_, err = models.Stores(
		qm.Where(models.StoreTableColumns.StoreID+" = ?", storeId),
	).DeleteAll(ctx, tx)
	if (err != nil) && (err != sql.ErrNoRows) {
		log.Printf("Error when deleting store %v: %v", storeName, err)
		return err
	}

	return nil
}

func replicateStoreFromFirestoreToPostgres(ctx context.Context, firestoreClient *firestore.Client, tx *sql.Tx, storeName string) error {

	// Count number of uploads in store

	uploadDocRefIter := firestoreClient.Collection(storesCollectionName).Doc(storeName).Collection(uploadsCollectionName).DocumentRefs(ctx)
	uploadDocRefs, err := uploadDocRefIter.GetAll()
	if err != nil {
		log.Printf("Get all upload doc refs in store %v failed: %v", storeName, err)
		return err
	}

	// Create new store in postgres, with up-to-date upload count

	var numUploads = len(uploadDocRefs)
	var newStore models.Store
	newStore.Name = storeName
	newStore.NextStoreUploadIndex = numUploads

	err = newStore.Insert(ctx, tx, boil.Infer())
	if err != nil {
		log.Printf("Error when inserting new store %v: %v", storeName, err)
		return err
	}

	// Create new uploads in store in postgres

	for uploadDocIndex, uploadDocRef := range(uploadDocRefs) {
		uploadId := uploadDocRef.ID

		log.Printf("Processing upload %v", uploadId)

		// Create upload entry

		oldUpload, err := firestore2.GetStoreUploadEntry(ctx, firestoreClient, storeName, uploadId)
		if err != nil {
			log.Printf("Failed fetching store / upload %v / %v: %v", storeName, uploadId, err)
			return err
		}

		var newUpload models.StoreUpload

		timestamp, err := time.Parse(time.RFC3339, oldUpload.Timestamp)
		if err != nil {
			log.Printf("Unable to parse timestamp %v in store / upload %v / %v: %v", oldUpload.Timestamp, storeName, uploadId, err)
			return err
		}

		newUpload.StoreID = null.IntFrom(newStore.StoreID)
		newUpload.StoreUploadIndex = uploadDocIndex
		newUpload.Description = oldUpload.Description
		newUpload.Build = oldUpload.BuildId
		newUpload.Timestamp = timestamp
		newUpload.Status = oldUpload.Status

		newUpload.Insert(ctx, tx, boil.Infer())

		// Create upload-files

		for oldUploadFileIndex, oldUploadFile := range(oldUpload.Files) {
			var newUploadFile models.StoreUploadFile

			newUploadFile.UploadID = null.IntFrom(newUpload.UploadID)
			// TODO: assign blob ID for those uploads that were successful, and have not yet been expired
			// newUploadFile.BlobID = ...
			newUploadFile.UploadFileIndex = oldUploadFileIndex
			newUploadFile.Status = oldUploadFile.Status
			newUploadFile.FileName = oldUploadFile.FileName
			newUploadFile.FileBlobIdentifier = oldUploadFile.Hash

			newUploadFile.Insert(ctx, tx, boil.Infer())
		}
	}

	return nil
}

func MigrateStore(ctx context.Context, firestoreClient *firestore.Client, storeName string) error {

	log.Printf("Migrating store %v", storeName)

	storeDocRef := firestoreClient.Collection(storesCollectionName).Doc(storeName)
	_, err := storeDocRef.Get(ctx)
	if err != nil {
		log.Printf("Get store %v doc failed: %v", storeName, err)
		return err
	}

	tx, err := postgres.BeginDBTransaction(ctx)
	if err != nil {
		log.Printf("Error when beginning DB transaction: %v", err)
		return err
	}

	// Delete old store in postgres

	if err = deleteOldStore(ctx, tx, storeName); err != nil {
		tx.Rollback()
		return err
	}

	// Replicate store from Firestore to Postgres

	if err = replicateStoreFromFirestoreToPostgres(ctx, firestoreClient, tx, storeName); err != nil {
		tx.Rollback()
		return err		
	}

	tx.Commit()

	log.Printf("Migrating store %v done", storeName)

	return nil
}

func MigrateStores(ctx context.Context, firestoreClient *firestore.Client) error {

	documentRefIter := firestoreClient.Collection(storesCollectionName).DocumentRefs(ctx)

	storeDocRefs, err := documentRefIter.GetAll()
	if err != nil {
		log.Printf("Get all store doc refs failed: %v", err)
		return err
	}

	for _, storeDocRef := range(storeDocRefs) {
		storeName := storeDocRef.ID

		err = MigrateStore(ctx, firestoreClient, storeName)
		if err != nil {
			return err
		}
	}

	return nil
}
