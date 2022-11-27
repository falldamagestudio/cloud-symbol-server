package admin_api

import (
	"context"
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

	// Mark file as uploaded
	numRowsAffected, err := models.StoreUploadFiles(qm.Where(models.StoreUploadFileColumns.UploadID+" = ? AND "+models.StoreUploadFileColumns.UploadFileIndex+" = ?", uploadId, fileId)).UpdateAll(ctx, db, models.M{models.StoreUploadFileColumns.Status: FileDBEntry_Status_Uploaded})
	if (err == nil) && (numRowsAffected == 0) {
		log.Printf("File %v / %v / %v not found", storeId, uploadId, fileId)
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("File %v / %v / %v not found", storeId, uploadId, fileId)}), err
	} else if (err != nil) || (numRowsAffected != 1) {
		log.Printf("error while accessing file: %v - numRowsAffected: %v", err, numRowsAffected)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("Status for %v/%v/%v set to %v", storeId, uploadId, fileId, FileDBEntry_Status_Uploaded)

	return openapi.Response(http.StatusOK, nil), nil
}
