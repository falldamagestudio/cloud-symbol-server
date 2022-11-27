package admin_api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
)

func (s *ApiService) MarkStoreUploadFileUploaded(ctx context.Context, uploadId string, storeId string, fileId int32) (openapi.ImplResponse, error) {

	db := GetDB()
	if db == nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
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

	// Mark file as uploaded
	numRowsAffected, err := models.StoreUploadFiles(qm.Where(models.StoreUploadFileColumns.UploadID+" = ? AND "+models.StoreUploadFileColumns.UploadFileIndex+" = ?", upload.UploadID, fileId)).UpdateAll(ctx, db, models.M{models.StoreUploadFileColumns.Status: FileDBEntry_Status_Uploaded})
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

	log.Printf("Status for %v/%v/%v set to %v", storeId, uploadId, fileId, FileDBEntry_Status_Uploaded)

	return openapi.Response(http.StatusOK, nil), nil
}
