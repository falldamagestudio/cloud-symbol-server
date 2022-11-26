package admin_api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
)

func (s *ApiService) GetStoreUpload(ctx context.Context, uploadId string, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Getting store upload")

	db := GetDB()
	if db == nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	upload, err := models.Uploads(qm.Where("upload_id = ?", uploadId)).One(ctx, db)
	if err == sql.ErrNoRows {
		log.Printf("Upload %v / %v not found", storeId, uploadId)
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Upload %v / %v not found", storeId, uploadId)}), err
	} else if err != nil {
		log.Printf("Error while accessing upload %v / %v: %v", storeId, uploadId, err)
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing upload %v / %v: %v", storeId, uploadId, err)}), err
	}

	files, err := models.Files(qm.Where("upload_id = ?", uploadId), qm.OrderBy("upload_file_index")).All(ctx, db)
	if err != nil {
		log.Printf("Error while accessing files of upload %v / %v: %v", storeId, uploadId, err)
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing files of upload %v / %v: %v", storeId, uploadId, err)}), err
	}

	getStoreUploadResponse := openapi.GetStoreUploadResponse{}
	getStoreUploadResponse.Description = upload.Description
	getStoreUploadResponse.BuildId = upload.Build
	getStoreUploadResponse.Timestamp = upload.Timestamp.Format(time.RFC3339)
	// Uploads created before the progress API existed do not have any Status field in the DB
	// These uploads should be interpreted as having status "Unknown"
	if upload.Status != "" {
		getStoreUploadResponse.Status = upload.Status
	} else {
		getStoreUploadResponse.Status = StoreUploadEntry_Status_Unknown
	}

	for _, file := range files {

		// Uploaded files created before the progress API existed do not have any Status field in the DB
		// These files should be interpreted as having status "Unknown"
		status := FileDBEntry_Status_Unknown
		if file.Status != "" {
			status = file.Status
		}

		getStoreUploadResponse.Files = append(getStoreUploadResponse.Files, openapi.GetFileResponse{
			FileName: file.FileName,
			Hash:     file.Hash,
			Status:   status,
		})
	}

	log.Printf("Response: %v", getStoreUploadResponse)

	return openapi.Response(http.StatusOK, getStoreUploadResponse), nil
}
