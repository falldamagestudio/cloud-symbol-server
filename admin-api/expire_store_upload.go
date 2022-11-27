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

func (s *ApiService) ExpireStoreUpload(ctx context.Context, uploadId string, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Expiring upload %v/%v", storeId, uploadId)

	storageClient, err := getStorageClient(ctx)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to create storage client"}), err
	}

	type ObjectToDelete struct {
		FileName string
		Hash     string
	}

	objectsToDelete := make([]ObjectToDelete, 0)

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
	upload, err := models.StoreUploads(qm.Where(models.StoreUploadColumns.StoreID+" = ? and "+models.StoreUploadColumns.StoreUploadIndex+" = ?", store.StoreID, uploadId), qm.For("share")).One(ctx, tx)
	if err == sql.ErrNoRows {
		log.Printf("Upload %v / %v not found", storeId, uploadId)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Upload %v / %v not found", storeId, uploadId)}), err
	} else if err != nil {
		log.Printf("Error while accessing upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing upload %v / %v: %v", storeId, uploadId, err)}), err
	}

	// Mark upload as expired
	numRowsAffected, err := models.StoreUploads(qm.Where(models.StoreUploadColumns.UploadID+" = ?", upload.UploadID)).UpdateAll(ctx, tx, models.M{models.StoreUploadColumns.Status: StoreUploadEntry_Status_Expired})
	if (err == nil) && (numRowsAffected == 0) {
		log.Printf("Upload %v / %v not found", storeId, uploadId)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Upload %v / %v not found", storeId, uploadId)}), err
	} else if (err != nil) || (numRowsAffected != 1) {
		log.Printf("error while accessing upload: %v - numRowsAffected: %v", err, numRowsAffected)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Mark all files in upload as expired
	_, err = models.StoreUploadFiles(qm.Where(models.StoreUploadFileColumns.UploadID+" = ?", upload.UploadID)).UpdateAll(ctx, tx, models.M{models.StoreUploadFileColumns.Status: FileDBEntry_Status_Expired})
	if err != nil {
		log.Printf("error while accessing files in upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Fetch all files in upload
	files, err := models.StoreUploadFiles(qm.Where(models.StoreUploadFileColumns.UploadID+" = ?", upload.UploadID), qm.OrderBy(models.StoreUploadFileColumns.UploadFileIndex)).All(ctx, tx)
	if err != nil {
		log.Printf("error while finding files in upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	for _, file := range files {

		// TODO: look at files within the current store only
		numNotExpiredFileRefs, err := models.StoreUploadFiles(qm.Where(models.StoreUploadFileColumns.FileName+" = ? AND "+models.StoreUploadFileColumns.Hash+" = ? AND "+models.StoreUploadFileColumns.Status+" != ?", file.FileName, file.Hash, FileDBEntry_Status_Expired)).Count(ctx, tx)
		if err != nil {
			log.Printf("error while accessing file/hash %v / %v: %v", file.FileName, file.Hash, err)
			tx.Rollback()
			return openapi.Response(http.StatusInternalServerError, nil), err
		}

		if numNotExpiredFileRefs == 0 {
			objectsToDelete = append(objectsToDelete, ObjectToDelete{FileName: file.FileName, Hash: file.Hash})
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("error while committing txn: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("Status for %v/%v set to %v", storeId, uploadId, StoreUploadEntry_Status_Expired)

	for _, objectToDelete := range objectsToDelete {
		err = deleteObjectInStore(ctx, storageClient, storeId, objectToDelete.FileName, objectToDelete.Hash)

		if err != nil {
			log.Printf("Error while deleting %s/%s/%s/%s from store; err = %v", storeId, objectToDelete.FileName, objectToDelete.Hash, objectToDelete.FileName, err)
			return openapi.Response(http.StatusInternalServerError, nil), nil
		}

		log.Printf("Deleted file %s/%s/%s/%s from store", storeId, objectToDelete.FileName, objectToDelete.Hash, objectToDelete.FileName)
	}

	log.Printf("Upload %s/%s has been expired", storeId, uploadId)

	return openapi.Response(http.StatusOK, nil), nil
}
