package upload_api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/sql-db-models"
	postgres "github.com/falldamagestudio/cloud-symbol-server/backend-api/postgres"
)

func GetStoreUpload(ctx context.Context, uploadId string, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Getting store upload")

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
		log.Printf("Error while accessing store %v: %v", storeId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing store %v: %v", storeId, err)}), err
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

	// Locate upload-files in DB
	uploadFiles, err := models.StoreUploadFiles(
		qm.Where(models.StoreUploadFileColumns.UploadID+" = ?", upload.UploadID),
		qm.OrderBy(models.StoreUploadFileColumns.UploadFileIndex),
	).All(ctx, tx)
	if err != nil {
		log.Printf("Error while accessing files of upload %v / %v: %v", storeId, uploadId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing files of upload %v / %v: %v", storeId, uploadId, err)}), err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("error while committing txn: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	getStoreUploadResponse := openapi.GetStoreUploadResponse{}
	getStoreUploadResponse.Description = upload.Description
	getStoreUploadResponse.BuildId = upload.Build
	getStoreUploadResponse.Timestamp = upload.Timestamp.Format(time.RFC3339)
	// Uploads created before the progress API existed do not have any Status field in the DB
	// These uploads will be assumed to be completed
	if upload.Status != "" {
		getStoreUploadResponse.Status = upload.Status
	} else {
		getStoreUploadResponse.Status = models.StoreUploadStatusCompleted
	}

	for _, file := range uploadFiles {

		// Uploaded files created before the progress API existed do not have any Status field in the DB
		// These files will be interpreted as being completed
		status := models.StoreUploadFileStatusCompleted
		if file.Status != "" {
			status = file.Status
		}

		getStoreUploadResponse.Files = append(getStoreUploadResponse.Files, openapi.GetFileResponse{
			FileName: file.FileName,
			Hash:     file.FileHash,
			Status:   status,
		})
	}

	log.Printf("Response: %v", getStoreUploadResponse)

	return openapi.Response(http.StatusOK, getStoreUploadResponse), nil
}
