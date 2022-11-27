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

func (s *ApiService) GetStoreFileIds(ctx context.Context, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Getting store file IDs")

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

	// Fetch IDs of all files within store; remove duplicates based on name
	files, err := models.Files(qm.Select("file_id"), qm.Where("store_id = ?", store.StoreID), qm.Distinct("file_name"), qm.OrderBy("file_id")).All(ctx, tx)
	if err != nil {
		log.Printf("Error while accessing files in store %v : %v", storeId, err)
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing files in store %v : %v", storeId, err)}), err
	}

	// Convert DB query result to a plain array of strings
	var storeFileIds = make([]string, len(files))

	for index, file := range files {
		storeFileIds[index] = strconv.Itoa(file.FileID)
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("unable to commit transaction: %v", err)
		return openapi.Response(http.StatusInternalServerError, fmt.Sprintf("unable to commit transaction: %v", err)), err
	}

	log.Printf("Response: %v", storeFileIds)

	return openapi.Response(http.StatusOK, storeFileIds), nil
}
