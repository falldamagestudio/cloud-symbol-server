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

func (s *ApiService) MarkStoreUploadAborted(ctx context.Context, uploadId string, storeId string) (openapi.ImplResponse, error) {

	db := GetDB()
	if db == nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Mark upload as aborted
	numRowsAffected, err := models.Uploads(qm.Where("upload_id = ?", uploadId)).UpdateAll(ctx, tx, models.M{"status": StoreUploadEntry_Status_Aborted})
	if (err == nil) && (numRowsAffected == 0) {
		log.Printf("Upload %v / %v not found", storeId, uploadId)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Upload %v / %v not found", storeId, uploadId)}), err
	} else if (err != nil) || (numRowsAffected != 1) {
		log.Printf("error while accessing upload: %v - numRowsAffected: %v", err, numRowsAffected)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Mark files in upload that are unknown/pending as aborted
	_, err = models.Files(qm.Where("upload_id = ?, uploadId"), qm.AndIn("status in ?", FileDBEntry_Status_Unknown, FileDBEntry_Status_Pending)).UpdateAll(ctx, tx, models.M{"status": FileDBEntry_Status_Aborted})
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
