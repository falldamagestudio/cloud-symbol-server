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

func (s *ApiService) MarkStoreUploadCompleted(ctx context.Context, uploadId string, storeId string) (openapi.ImplResponse, error) {

	db := GetDB()
	if db == nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	// Mark upload as completed
	numRowsAffected, err := models.StoreUploads(qm.Where(models.StoreUploadColumns.UploadID+" = ?", uploadId)).UpdateAll(ctx, db, models.M{models.StoreUploadColumns.Status: StoreUploadEntry_Status_Completed})
	if (err == nil) && (numRowsAffected == 0) {
		log.Printf("Upload %v / %v not found", storeId, uploadId)
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Upload %v / %v not found", storeId, uploadId)}), err
	} else if (err != nil) || (numRowsAffected != 1) {
		log.Printf("error while accessing upload: %v - numRowsAffected: %v", err, numRowsAffected)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("Status for %v/%v set to %v", storeId, uploadId, StoreUploadEntry_Status_Completed)

	return openapi.Response(http.StatusOK, nil), nil
}
