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

func GetStoreUploads(ctx context.Context, storeId string, offset int32, limit int32) (openapi.ImplResponse, error) {

	log.Printf("Getting store uploads for store %v, offset %v, limit %v", storeId, offset, limit)

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

	// Count total number of uploads in store that match filter query, ignoring pagination
	total, err := models.StoreUploads(
		qm.Distinct(models.StoreUploadColumns.UploadID),
		qm.Where(models.StoreUploadTableColumns.StoreID+" = ?", store.StoreID),
	).Count(ctx, tx)
	if err != nil {
		log.Printf("Error while accessing uploads in store %v : %v", storeId, err)
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing uploads in store %v : %v", storeId, err)}), err
	}

	// Fetch all uploads within store
	uploads, err := models.StoreUploads(
		qm.Where(models.StoreUploadTableColumns.StoreID+" = ?", store.StoreID),
		qm.OrderBy(models.StoreUploadColumns.UploadID),
		qm.Offset(int(offset)),
		qm.Limit(int(limit)),
	).All(ctx, tx)
	if err != nil {
		log.Printf("Error while accessing uploads in store %v : %v", storeId, err)
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing uploads in store %v : %v", storeId, err)}), err
	}

	// Fetch all files related to selected uploads

	var uploadIds = make([]interface{}, len(uploads))
	for uploadIndex, upload := range uploads {
		uploadIds[uploadIndex] = upload.UploadID
	}

	var uploadFiles []models.StoreUploadFile

	err = models.NewQuery(
		qm.Select("cloud_symbol_server."+models.TableNames.StoreUploadFiles+".*"),
		qm.From("cloud_symbol_server."+models.TableNames.StoreUploads),
		qm.InnerJoin("cloud_symbol_server."+models.TableNames.StoreUploadFiles+" on "+models.StoreUploadTableColumns.UploadID+" = "+models.StoreUploadFileTableColumns.UploadID),
		qm.Where(models.StoreUploadTableColumns.StoreID+" = ?", store.StoreID),
		qm.AndIn(models.StoreUploadTableColumns.UploadID+" IN ?", uploadIds...),
	).Bind(ctx, tx, &uploadFiles)
	if err != nil {
		log.Printf("Error while accessing upload-files in store %v : %v", storeId, err)
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing upload-files in store %v : %v", storeId, err)}), err
	}

	// Convert DB query result to API result

	// Count number of files mapping to each UploadId
	var uploadIdToFileCount = map[int]int{}

	for _, file := range uploadFiles {
		uploadId := file.UploadID.Int
		if !file.UploadID.Valid {
			log.Printf("Error while processing upload-file in store %v : %v", storeId, err)
			return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing upload-files in store %v : %v", storeId, err)}), err
		}

		_, ok := uploadIdToFileCount[uploadId]
		if !ok {
			uploadIdToFileCount[uploadId] = 0
		}

		uploadIdToFileCount[uploadId]++
	}

	// Create per-upload files result data structures
	var uploadIdToFiles = map[int][]openapi.GetStoreUploadFileResponse{}
	for uploadId := range uploadIdToFileCount {
		numFiles := uploadIdToFileCount[uploadId]
		uploadIdToFiles[uploadId] = make([]openapi.GetStoreUploadFileResponse, numFiles)
	}

	// Populate per-upload files result data structures
	for _, file := range uploadFiles {
		uploadId := file.UploadID.Int
		uploadFileIndex := file.UploadFileIndex

		// Uploaded files created before the progress API existed do not have any Status field in the DB
		// These files will be interpreted as being completed
		status := models.StoreUploadFileStatusCompleted
		if file.Status != "" {
			status = file.Status
		}

		uploadIdToFiles[uploadId][uploadFileIndex] = openapi.GetStoreUploadFileResponse{
			FileName:       file.FileName,
			BlobIdentifier: file.FileBlobIdentifier,
			Status:         openapi.StoreUploadFileStatus(status),
		}
	}

	// Create per-upload results
	var storeUploads = make([]openapi.GetStoreUploadResponse, len(uploads))

	for uploadIndex, upload := range uploads {
		var files = uploadIdToFiles[upload.UploadID]

		storeUploads[uploadIndex] = openapi.GetStoreUploadResponse{
			UploadId:    int32(upload.StoreUploadIndex),
			Description: upload.Description,
			BuildId:     upload.Build,
			Timestamp:   upload.Timestamp.Format(time.RFC3339),
			Files:       files,
			Status:      openapi.StoreUploadStatus(upload.Status),
		}
	}

	storeUploadsResponse := &openapi.GetStoreUploadsResponse{
		Uploads: storeUploads,
		Pagination: openapi.PaginationResponse{
			Total: int32(total),
		},
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("unable to commit transaction: %v", err)
		return openapi.Response(http.StatusInternalServerError, fmt.Sprintf("unable to commit transaction: %v", err)), err
	}

	log.Printf("Response: %v", storeUploads)

	return openapi.Response(http.StatusOK, storeUploadsResponse), nil
}
