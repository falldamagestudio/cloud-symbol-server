package upload_api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/sql-db-models"
	postgres "github.com/falldamagestudio/cloud-symbol-server/backend-api/postgres"
)

func MarkStoreUploadFileUploaded(ctx context.Context, uploadId int32, storeId string, fileId int32) (openapi.ImplResponse, error) {

	tx, err := postgres.BeginDBTransaction(ctx)
	if err != nil {
		log.Printf("Err: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	// Locate store in DB, and ensure store remains throughout entire txn
	store, err := models.Stores(
		qm.Where(models.StoreColumns.Name+" = ?", storeId), qm.For("share"),
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

	// Mark store-file-blob as present
	storeFileBlob, err := models.StoreFileBlobs(
		qm.InnerJoin("cloud_symbol_server."+models.TableNames.StoreUploadFiles+" on "+models.StoreFileBlobTableColumns.BlobID+" = "+models.StoreUploadFileTableColumns.BlobID),
		qm.Where(models.StoreUploadFileColumns.UploadID+" = ? AND "+models.StoreUploadFileColumns.UploadFileIndex+" = ?", upload.UploadID, fileId),
		qm.For("update"),
	).One(ctx, tx)
	if err == sql.ErrNoRows {
		log.Printf("File-blob for %v / %v / %v not found", storeId, uploadId, fileId)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("File %v / %v / %v not found", storeId, uploadId, fileId)}), err
	} else if err != nil {
		log.Printf("error while locating file-blob: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	storeFileBlob.Status = models.StoreFileBlobStatusPresent
	numRowsAffected, err := storeFileBlob.Update(ctx, tx, boil.Whitelist(models.StoreFileBlobColumns.Status))
	if (err != nil) || (numRowsAffected != 1) {
		log.Printf("error while updating file-blob: %v - numRowsAffected: %v", err, numRowsAffected)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Mark upload-file as completed
	numRowsAffected, err = models.StoreUploadFiles(
		qm.Where(models.StoreUploadFileColumns.UploadID+" = ? AND "+models.StoreUploadFileColumns.UploadFileIndex+" = ?", upload.UploadID, fileId),
	).UpdateAll(ctx, tx, models.M{models.StoreUploadFileColumns.Status: models.StoreUploadFileStatusCompleted})
	if (err == nil) && (numRowsAffected == 0) {
		log.Printf("File %v / %v / %v not found", storeId, uploadId, fileId)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("File %v / %v / %v not found", storeId, uploadId, fileId)}), err
	} else if (err != nil) || (numRowsAffected != 1) {
		log.Printf("error while accessing file: %v - numRowsAffected: %v", err, numRowsAffected)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("error while committing txn: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("Status for %v/%v/%v set to %v", storeId, uploadId, fileId, models.StoreUploadFileStatusCompleted)

	return openapi.Response(http.StatusOK, nil), nil
}
