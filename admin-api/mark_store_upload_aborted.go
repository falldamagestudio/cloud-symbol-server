package admin_api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
)

func (s *ApiService) MarkStoreUploadAborted(ctx context.Context, uploadId string, storeId string) (openapi.ImplResponse, error) {

	tx, err := BeginDBTransaction(ctx)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	// Locate store in DB, and ensure store remains throughout entire txn
	store, err := models.Stores(qm.Where(models.StoreColumns.Name+" = ?", storeId), qm.For("share")).One(ctx, tx)
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
	upload, err := models.StoreUploads(qm.Where(models.StoreUploadColumns.StoreID+" = ? and "+models.StoreUploadColumns.StoreUploadIndex+" = ?", store.StoreID, uploadId)).One(ctx, tx)
	if err == sql.ErrNoRows {
		log.Printf("Upload %v / %v not found", storeId, uploadId)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Upload %v / %v not found", storeId, uploadId)}), err
	} else if err != nil {
		log.Printf("Error while accessing upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing upload %v / %v: %v", storeId, uploadId, err)}), err
	}

	// Mark upload as aborted
	numRowsAffected, err := models.StoreUploads(qm.Where(models.StoreUploadColumns.UploadID+" = ?", upload.UploadID)).UpdateAll(ctx, tx, models.M{models.StoreUploadColumns.Status: StoreUploadEntry_Status_Aborted})
	if (err == nil) && (numRowsAffected == 0) {
		log.Printf("Upload %v / %v not found", storeId, uploadId)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Upload %v / %v not found", storeId, uploadId)}), err
	} else if (err != nil) || (numRowsAffected != 1) {
		log.Printf("error while accessing upload: %v - numRowsAffected: %v", err, numRowsAffected)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	boil.DebugMode = true

	// Mark store-file-hashes in upload that are unknown/pending as aborted
	storeFileHashes, err := models.StoreFileHashes(
		qm.InnerJoin("cloud_symbol_server."+models.TableNames.StoreUploadFiles+" on "+models.StoreFileHashTableColumns.HashID+" = "+models.StoreUploadFileTableColumns.HashID),
		qm.Where(models.StoreUploadFileColumns.UploadID+" = ?", upload.UploadID),
		qm.AndIn(models.StoreUploadFileTableColumns.Status+" in ?", models.StoreFileHashStatusUnknown, models.StoreFileHashStatusPending),
	).All(ctx, tx)
	if (err != nil) && (err != sql.ErrNoRows) {
		log.Printf("error while accessing files-hashes in upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	_, err = storeFileHashes.UpdateAll(ctx, tx, models.M{models.StoreFileHashColumns.Status: models.StoreFileHashStatusAborted})
	if (err != nil) && (err != sql.ErrNoRows) {
		log.Printf("error while updating files-hashes in upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Mark upload-files in upload that are unknown/pending as aborted
	_, err = models.StoreUploadFiles(qm.Where(models.StoreUploadFileColumns.UploadID+" = ?", upload.UploadID), qm.AndIn(models.StoreUploadFileColumns.Status+" in ?", FileDBEntry_Status_Unknown, FileDBEntry_Status_Pending)).UpdateAll(ctx, tx, models.M{models.StoreUploadFileColumns.Status: FileDBEntry_Status_Aborted})
	if err != nil {
		log.Printf("error while accessing files in upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("error while committing txn: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("Status for %v/%v set to %v", storeId, uploadId, StoreUploadEntry_Status_Aborted)

	return openapi.Response(http.StatusOK, nil), nil
}
