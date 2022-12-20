package upload_api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
	helpers "github.com/falldamagestudio/cloud-symbol-server/admin-api/helpers"
)

func HandleUploadExpiryOrAbort(ctx context.Context, storeId string, uploadId string, desiredStatus string) (openapi.ImplResponse, error) {

	storageClient, err := helpers.GetStorageClient(ctx)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to create storage client"}), err
	}

	tx, err := helpers.BeginDBTransaction(ctx)
	if err != nil {
		log.Printf("Err: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	// Locate store in DB, and ensure store remains throughout entire txn
	store, err := models.Stores(
		qm.Where(models.StoreColumns.Name+" = ?", storeId),
		qm.For("share"),
	).One(ctx, tx)
	if err == sql.ErrNoRows {
		log.Printf("Store %v not found; err = %v", storeId, err)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Store %v not found", storeId)}), err
	} else if err != nil {
		log.Printf("error while accessing store: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, fmt.Sprintf("error while accessing store: %v", err)), err
	}

	// Locate upload in DB, and ensure upload remains throughout entire txn
	upload, err := models.StoreUploads(
		qm.Where(models.StoreUploadColumns.StoreID+" = ? and "+models.StoreUploadColumns.StoreUploadIndex+" = ?", store.StoreID, uploadId),
		qm.For("update"),
	).One(ctx, tx)
	if err == sql.ErrNoRows {
		log.Printf("Upload %v / %v not found", storeId, uploadId)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Upload %v / %v not found", storeId, uploadId)}), err
	} else if err != nil {
		log.Printf("Error while accessing upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing upload %v / %v: %v", storeId, uploadId, err)}), err
	}

	// Mark upload as expired/aborted
	numRowsAffected, err := models.StoreUploads(
		qm.Where(models.StoreUploadColumns.UploadID+" = ?", upload.UploadID),
	).UpdateAll(ctx, tx, models.M{models.StoreUploadColumns.Status: desiredStatus})
	if (err == nil) && (numRowsAffected == 0) {
		log.Printf("Upload %v / %v not found", storeId, uploadId)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Upload %v / %v not found", storeId, uploadId)}), err
	} else if (err != nil) || (numRowsAffected != 1) {
		log.Printf("error while accessing upload: %v - numRowsAffected: %v", err, numRowsAffected)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Retrieve IDs of all file-hashes referenced by upload
	var storeFileHashIds []struct {
		HashID int `boil:"hash_id"`
	}
	err = models.NewQuery(
		qm.Distinct(models.StoreUploadFileTableColumns.HashID),
		qm.From("cloud_symbol_server."+models.TableNames.StoreUploadFiles),
		qm.Where(models.StoreUploadFileColumns.UploadID+" = ?", upload.UploadID),
		qm.And(models.StoreUploadFileColumns.HashID+" IS NOT NULL"),
	).Bind(ctx, tx, &storeFileHashIds)
	if err != nil {
		log.Printf("error while retrieveing file-hash IDs in upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Retrieve IDs of all files referenced by upload
	var storeFileIds []struct {
		FileID int `boil:"file_id"`
	}
	err = models.NewQuery(
		qm.Distinct(models.StoreFileHashTableColumns.FileID),
		qm.From("cloud_symbol_server."+models.TableNames.StoreUploadFiles),
		qm.InnerJoin("cloud_symbol_server."+models.TableNames.StoreFileHashes+" on "+models.StoreFileHashTableColumns.HashID+" = "+models.StoreUploadFileTableColumns.HashID),
		qm.Where(models.StoreUploadFileColumns.UploadID+" = ?", upload.UploadID),
	).Bind(ctx, tx, &storeFileIds)
	if err != nil {
		log.Printf("error while retrieveing file-hash IDs in upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Mark upload-files in upload as expired/aborted, and clear their file IDs & hash IDs
	_, err = models.StoreUploadFiles(
		qm.Where(models.StoreUploadFileColumns.UploadID+" = ?", upload.UploadID),
	).UpdateAll(ctx, tx, models.M{
		models.StoreUploadFileColumns.Status: desiredStatus,
		models.StoreUploadFileColumns.HashID: nil,
	})
	if err != nil {
		log.Printf("error while accessing files in upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Locate any file-hashes in the previous set that no longer are referenced by any uploads

	// SELECT store_file_hashes.hash_id
	// FROM cloud_symbol_server.store_file_hashes
	// LEFT JOIN cloud_symbol_server.store_upload_files
	// 	ON store_file_hashes.hash_id = store_upload_files.hash_id
	// WHERE (store_file_hashes.hash_id = ???)
	// 	AND (store_upload_files.hash_id IS NULL);

	storeFileHashIdsArray := make([]interface{}, len(storeFileHashIds))
	for i := range storeFileHashIds {
		storeFileHashIdsArray[i] = storeFileHashIds[i].HashID
	}
	storeFileHashIdsToDelete, err := models.StoreFileHashes(
		qm.Select(models.StoreFileHashTableColumns.HashID+" as "+models.StoreFileHashColumns.HashID),
		qm.LeftOuterJoin("cloud_symbol_server."+models.TableNames.StoreUploadFiles+" on "+models.StoreFileHashTableColumns.HashID+" = "+models.StoreUploadFileTableColumns.HashID),
		qm.WhereIn(models.StoreFileHashTableColumns.HashID+" in ?", storeFileHashIdsArray...),
		qm.And(models.StoreUploadFileTableColumns.HashID+" IS NULL"),
	).All(ctx, tx)
	if err != nil {
		log.Printf("error while retrieveing IDs for unreferenced file-hashes related to upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Create list of objects-to-delete from GCS

	storeFileHashIdsToDeleteArray := make([]interface{}, len(storeFileHashIdsToDelete))
	for i := range storeFileHashIdsToDelete {
		storeFileHashIdsToDeleteArray[i] = storeFileHashIdsToDelete[i].HashID
	}
	var objectsToDelete []struct {
		FileName string `boil:"file_name"`
		Hash     string `boil:"hash"`
	}
	err = models.NewQuery(
		qm.Select(models.StoreFileTableColumns.FileName+" as "+models.StoreFileColumns.FileName, models.StoreFileHashTableColumns.Hash+" as "+models.StoreFileHashColumns.Hash),
		qm.From("cloud_symbol_server."+models.TableNames.StoreFileHashes),
		qm.InnerJoin("cloud_symbol_server."+models.TableNames.StoreFiles+" on "+models.StoreFileTableColumns.FileID+" = "+models.StoreFileHashTableColumns.FileID),
		qm.WhereIn(models.StoreFileHashTableColumns.HashID+" in ?", storeFileHashIdsToDeleteArray...),
	).Bind(ctx, tx, &objectsToDelete)
	if err != nil {
		log.Printf("error while retrieveing file-hashes in upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Delete unreferenced file-hashes

	_, err = storeFileHashIdsToDelete.DeleteAll(ctx, tx)
	if err != nil {
		log.Printf("error while deleting unreferenced file-hashes related to upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Locate any files in the previous set that no longer are referenced by any uploads

	// SELECT store_files.file_id
	// FROM cloud_symbol_server.store_files
	// LEFT JOIN cloud_symbol_server.store_file_hashes
	// 	ON store_files.file_id = store_file_hashes.file_id
	// WHERE (store_files.file_id = ???)
	// 	AND (store_file_hashes.file_id IS NULL);

	storeFileIdsArray := make([]interface{}, len(storeFileIds))
	for i := range storeFileIds {
		storeFileIdsArray[i] = storeFileIds[i].FileID
	}
	storeFileIdsToDelete, err := models.StoreFiles(
		qm.Select(models.StoreFileTableColumns.FileID+" as "+models.StoreFileColumns.FileID),
		qm.LeftOuterJoin("cloud_symbol_server."+models.TableNames.StoreFileHashes+" on "+models.StoreFileTableColumns.FileID+" = "+models.StoreFileHashTableColumns.FileID),
		qm.WhereIn(models.StoreFileTableColumns.FileID+" in ?", storeFileIdsArray...),
		qm.And(models.StoreFileHashTableColumns.FileID+" IS NULL"),
	).All(ctx, tx)
	if err != nil {
		log.Printf("error while retrieveing IDs for unreferenced files related to upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Delete unreferenced files

	_, err = storeFileIdsToDelete.DeleteAll(ctx, tx)
	if err != nil {
		log.Printf("error while deleting unreferenced files related to upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("error while committing txn: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("Status for %v/%v set to %v", storeId, uploadId, models.StoreUploadStatusExpired)

	for _, objectToDelete := range objectsToDelete {
		err = helpers.DeleteObjectInStore(ctx, storageClient, storeId, objectToDelete.FileName, objectToDelete.Hash)

		if err == storage.ErrObjectNotExist {
			log.Printf("file %s/%s/%s/%s not found in store when it was time to delete it", storeId, objectToDelete.FileName, objectToDelete.Hash, objectToDelete.FileName)
		} else if err != nil {
			log.Printf("Error while deleting %s/%s/%s/%s from store; err = %v", storeId, objectToDelete.FileName, objectToDelete.Hash, objectToDelete.FileName, err)
			return openapi.Response(http.StatusInternalServerError, nil), nil
		} else {
			log.Printf("Deleted file %s/%s/%s/%s from store", storeId, objectToDelete.FileName, objectToDelete.Hash, objectToDelete.FileName)
		}
	}

	log.Printf("Upload %s/%s has been expired/aborted", storeId, uploadId)

	return openapi.Response(http.StatusOK, nil), nil
}
