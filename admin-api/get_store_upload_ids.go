package admin_api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
)

func (s *ApiService) GetStoreUploadIds(ctx context.Context, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Getting store upload IDs")

	db := GetDB()
	if db == nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Locate store in DB, and ensure store remains throughout entire txn
	store, err := models.Stores(qm.Where("name = ?", storeId), qm.For("share")).One(ctx, tx)
	if err != nil {
		log.Printf("error while accessing store: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, fmt.Sprintf("error while accessing store: %v", err)), err
	}

	// Fetch IDs of all uploads within store
	uploads, err := models.Uploads(qm.Select("upload_id"), qm.Where("store_id = ?", store.StoreID), qm.OrderBy("upload_id")).All(ctx, tx)
	if err != nil {
		log.Printf("Error while accessing uploads of store %v : %v", storeId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing uploads of store %v : %v", storeId, err)}), err
	}

	// Convert DB query result to a plain array of strings
	var storeUploadIds = make([]string, len(uploads))

	for index, upload := range uploads {
		storeUploadIds[index] = strconv.Itoa(upload.UploadID)
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("unable to commit transaction: %v", err)
		return openapi.Response(http.StatusInternalServerError, fmt.Sprintf("unable to commit transaction: %v", err)), err
	}

	log.Printf("Response: %v", storeUploadIds)

	return openapi.Response(http.StatusOK, storeUploadIds), nil
}
