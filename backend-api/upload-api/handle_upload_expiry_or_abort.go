package upload_api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	cloud_storage "github.com/falldamagestudio/cloud-symbol-server/backend-api/cloud_storage"
	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/sql-db-models"
	postgres "github.com/falldamagestudio/cloud-symbol-server/backend-api/postgres"
)

func HandleUploadExpiryOrAbort(ctx context.Context, storeId string, uploadId int32, desiredStatus string) (openapi.ImplResponse, error) {

	storageClient, err := cloud_storage.GetStorageClient(ctx)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to create storage client"}), err
	}

	tx, err := postgres.BeginDBTransaction(ctx)
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

	// Mark upload as expired/aborted and set the expiry timestamp
	numRowsAffected, err := models.StoreUploads(
		qm.Where(models.StoreUploadColumns.UploadID+" = ?", upload.UploadID),
	).UpdateAll(ctx, tx, models.M{
		models.StoreUploadColumns.Status: desiredStatus,
		models.StoreUploadColumns.ExpiryTimestamp: time.Now(),
	})
	if (err == nil) && (numRowsAffected == 0) {
		log.Printf("Upload %v / %v not found", storeId, uploadId)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Upload %v / %v not found", storeId, uploadId)}), err
	} else if (err != nil) || (numRowsAffected != 1) {
		log.Printf("error while accessing upload: %v - numRowsAffected: %v", err, numRowsAffected)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Retrieve IDs of all file-blobs referenced by upload
	var storeFileBlobIds []struct {
		BlobID int `boil:"blob_id"`
	}
	err = models.NewQuery(
		qm.Distinct(models.StoreUploadFileTableColumns.BlobID),
		qm.From("cloud_symbol_server."+models.TableNames.StoreUploadFiles),
		qm.Where(models.StoreUploadFileColumns.UploadID+" = ?", upload.UploadID),
		qm.And(models.StoreUploadFileColumns.BlobID+" IS NOT NULL"),
	).Bind(ctx, tx, &storeFileBlobIds)
	if err != nil {
		log.Printf("error while retrieveing file-blob IDs in upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Retrieve IDs of all files referenced by upload
	var storeFileIds []struct {
		FileID int `boil:"file_id"`
	}
	err = models.NewQuery(
		qm.Distinct(models.StoreFileBlobTableColumns.FileID),
		qm.From("cloud_symbol_server."+models.TableNames.StoreUploadFiles),
		qm.InnerJoin("cloud_symbol_server."+models.TableNames.StoreFileBlobs+" on "+models.StoreFileBlobTableColumns.BlobID+" = "+models.StoreUploadFileTableColumns.BlobID),
		qm.Where(models.StoreUploadFileColumns.UploadID+" = ?", upload.UploadID),
	).Bind(ctx, tx, &storeFileIds)
	if err != nil {
		log.Printf("error while retrieveing file-blob IDs in upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Mark upload-files in upload as expired/aborted, and clear their file IDs & blob IDs
	_, err = models.StoreUploadFiles(
		qm.Where(models.StoreUploadFileColumns.UploadID+" = ?", upload.UploadID),
	).UpdateAll(ctx, tx, models.M{
		models.StoreUploadFileColumns.Status: desiredStatus,
		models.StoreUploadFileColumns.BlobID: nil,
	})
	if err != nil {
		log.Printf("error while accessing files in upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Locate any file-blobs in the previous set that no longer are referenced by any uploads

	// SELECT store_file_blobs.blob_id
	// FROM cloud_symbol_server.store_file_blobs
	// LEFT JOIN cloud_symbol_server.store_upload_files
	// 	ON store_file_blobs.blob_id = store_upload_files.blob_id
	// WHERE (store_file_blobs.blob_id = ???)
	// 	AND (store_upload_files.blob_id IS NULL);

	storeFileBlobIdsArray := make([]interface{}, len(storeFileBlobIds))
	for i := range storeFileBlobIds {
		storeFileBlobIdsArray[i] = storeFileBlobIds[i].BlobID
	}
	storeFileBlobIdsToDelete, err := models.StoreFileBlobs(
		qm.Select(models.StoreFileBlobTableColumns.BlobID+" as "+models.StoreFileBlobColumns.BlobID),
		qm.LeftOuterJoin("cloud_symbol_server."+models.TableNames.StoreUploadFiles+" on "+models.StoreFileBlobTableColumns.BlobID+" = "+models.StoreUploadFileTableColumns.BlobID),
		qm.WhereIn(models.StoreFileBlobTableColumns.BlobID+" in ?", storeFileBlobIdsArray...),
		qm.And(models.StoreUploadFileTableColumns.BlobID+" IS NULL"),
	).All(ctx, tx)
	if err != nil {
		log.Printf("error while retrieveing IDs for unreferenced file-blobs related to upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Create list of objects-to-delete from GCS

	storeFileBlobIdsToDeleteArray := make([]interface{}, len(storeFileBlobIdsToDelete))
	for i := range storeFileBlobIdsToDelete {
		storeFileBlobIdsToDeleteArray[i] = storeFileBlobIdsToDelete[i].BlobID
	}
	var objectsToDelete []struct {
		FileName       string `boil:"file_name"`
		BlobIdentifier string `boil:"blob_identifier"`
	}
	err = models.NewQuery(
		qm.Select(models.StoreFileTableColumns.FileName+" as "+models.StoreFileColumns.FileName, models.StoreFileBlobTableColumns.BlobIdentifier+" as "+models.StoreFileBlobColumns.BlobIdentifier),
		qm.From("cloud_symbol_server."+models.TableNames.StoreFileBlobs),
		qm.InnerJoin("cloud_symbol_server."+models.TableNames.StoreFiles+" on "+models.StoreFileTableColumns.FileID+" = "+models.StoreFileBlobTableColumns.FileID),
		qm.WhereIn(models.StoreFileBlobTableColumns.BlobID+" in ?", storeFileBlobIdsToDeleteArray...),
	).Bind(ctx, tx, &objectsToDelete)
	if err != nil {
		log.Printf("error while retrieveing file-blobs in upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Delete unreferenced file-blobs

	_, err = storeFileBlobIdsToDelete.DeleteAll(ctx, tx)
	if err != nil {
		log.Printf("error while deleting unreferenced file-blobs related to upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Locate any files in the previous set that no longer are referenced by any uploads

	// SELECT store_files.file_id
	// FROM cloud_symbol_server.store_files
	// LEFT JOIN cloud_symbol_server.store_file_blobs
	// 	ON store_files.file_id = store_file_blobs.file_id
	// WHERE (store_files.file_id = ???)
	// 	AND (store_file_blobs.file_id IS NULL);

	storeFileIdsArray := make([]interface{}, len(storeFileIds))
	for i := range storeFileIds {
		storeFileIdsArray[i] = storeFileIds[i].FileID
	}
	storeFileIdsToDelete, err := models.StoreFiles(
		qm.Select(models.StoreFileTableColumns.FileID+" as "+models.StoreFileColumns.FileID),
		qm.LeftOuterJoin("cloud_symbol_server."+models.TableNames.StoreFileBlobs+" on "+models.StoreFileTableColumns.FileID+" = "+models.StoreFileBlobTableColumns.FileID),
		qm.WhereIn(models.StoreFileTableColumns.FileID+" in ?", storeFileIdsArray...),
		qm.And(models.StoreFileBlobTableColumns.FileID+" IS NULL"),
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
		err = cloud_storage.DeleteObjectInStore(ctx, storageClient, storeId, objectToDelete.FileName, objectToDelete.BlobIdentifier)

		if err == storage.ErrObjectNotExist {
			log.Printf("file %s/%s/%s/%s not found in store when it was time to delete it", storeId, objectToDelete.FileName, objectToDelete.BlobIdentifier, objectToDelete.FileName)
		} else if err != nil {
			log.Printf("Error while deleting %s/%s/%s/%s from store; err = %v", storeId, objectToDelete.FileName, objectToDelete.BlobIdentifier, objectToDelete.FileName, err)
			return openapi.Response(http.StatusInternalServerError, nil), nil
		} else {
			log.Printf("Deleted file %s/%s/%s/%s from store", storeId, objectToDelete.FileName, objectToDelete.BlobIdentifier, objectToDelete.FileName)
		}
	}

	log.Printf("Upload %s/%s has been expired/aborted", storeId, uploadId)

	return openapi.Response(http.StatusOK, nil), nil
}
