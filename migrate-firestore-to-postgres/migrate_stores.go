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

func MigrateStore(ctx context.Context, firestoreClient *firestore.Client, storeId string) error {

	log.Printf("Migrating store %v", storeId)

	storeDocRef := firestoreClient.Collection(storesCollectionName).Doc(storeId)
	_, err := storeDocRef.Get(ctx)
	if err != nil {
		log.Printf("Get store doc %v failed: %v", storeId, err)
		return err
	}

	tx, err := postgres.BeginDBTransaction(ctx)
	if err != nil {
		log.Printf("Error when beginning DB transaction: %v", err)
		return err
	}

	// Delete old store in postgres

	toBeDeletedStore, err := models.Stores(
		qm.Where(models.StoreTableColumns.Name+" = ?", storeId),
	).One(ctx, tx)
	if (err != nil) && (err != sql.ErrNoRows) {
		log.Printf("Error when locating store %v: %v", storeId, err)
		tx.Rollback()
		return err
	}

	if err != sql.ErrNoRows {

		toBeDeletedStoreId := toBeDeletedStore.StoreID

		// // Delete upload-files referencing uploads in to-be-deleted store

		// _, err = models.StoreUploadFiles(
		// 	qm.Where(models.StoreUploadTableColumns.StoreID+" = ?", toBeDeletedStoreId),
		// ).DeleteAll(ctx, tx)
		// if (err != nil) && (err != sql.ErrNoRows) {
		// 	log.Printf("Error when deleting all uploads referencing storeId %v: %v", toBeDeletedStoreId, err)
		// 	tx.Rollback()
		// 	return err
		// }

		// Delete uploads referencing to-be-deleted store

		_, err = models.StoreUploads(
			qm.Where(models.StoreUploadTableColumns.StoreID+" = ?", toBeDeletedStoreId),
		).DeleteAll(ctx, tx)
		if (err != nil) && (err != sql.ErrNoRows) {
			log.Printf("Error when deleting all uploads referencing storeId %v: %v", toBeDeletedStoreId, err)
			tx.Rollback()
			return err
		}
			
		// Delete already-existing store

		_, err = models.Stores(
			qm.Where(models.StoreTableColumns.StoreID+" = ?", toBeDeletedStoreId),
		).DeleteAll(ctx, tx)
		if (err != nil) && (err != sql.ErrNoRows) {
			log.Printf("Error when deleting store %v: %v", storeId, err)
			tx.Rollback()
			return err
		}
	}


	// Count number of uploads in store

	uploadDocRefIter := firestoreClient.Collection(storesCollectionName).Doc(storeId).Collection(uploadsCollectionName).DocumentRefs(ctx)
	uploadDocRefs, err := uploadDocRefIter.GetAll()
	if err != nil {
		log.Printf("Get all upload doc refs in store %v failed: %v", storeId, err)
		tx.Rollback()
		return err
	}

	// Create new store in postgres, with up-to-date upload count

	var numUploads = len(uploadDocRefs)
	var newStore models.Store
	newStore.Name = storeId
	newStore.NextStoreUploadIndex = numUploads

	err = newStore.Insert(ctx, tx, boil.Infer())
	if err != nil {
		log.Printf("Error when inserting new store %v: %v", storeId, err)
		tx.Rollback()
		return err
	}

	// Create new uploads in store in postgres

	for uploadDocIndex, uploadDocRef := range(uploadDocRefs) {
		uploadId := uploadDocRef.ID

		log.Printf("Processing upload %v", uploadId)

		// Create upload entry

		oldUpload, err := firestore2.GetStoreUploadEntry(ctx, firestoreClient, storeId, uploadId)
		if err != nil {
			log.Printf("Failed fetching store / upload %v / %v: %v", storeId, uploadId, err)
			tx.Rollback()
			return err
		}

		var newUpload models.StoreUpload

		timestamp, err := time.Parse(time.RFC3339, oldUpload.Timestamp)
		if err != nil {
			log.Printf("Unable to parse timestamp %v in store / upload %v / %v: %v", oldUpload.Timestamp, storeId, uploadId, err)
			tx.Rollback()
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

	tx.Commit()

	log.Printf("Migrating store %v done", storeId)

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
		storeId := storeDocRef.ID

		err = MigrateStore(ctx, firestoreClient, storeId)
		if err != nil {
			return err
		}
	}

	return nil
}
