package upload_api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/sql-db-models"
	postgres "github.com/falldamagestudio/cloud-symbol-server/backend-api/postgres"
)

func GetStoreUploadIds(ctx context.Context, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Getting store upload IDs")

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

	// Fetch IDs of all uploads within store
	uploads, err := models.StoreUploads(
		qm.Select(models.StoreUploadColumns.StoreUploadIndex),
		qm.Where(models.StoreUploadColumns.StoreID+" = ?", store.StoreID),
		qm.OrderBy(models.StoreUploadColumns.StoreUploadIndex),
	).All(ctx, tx)
	if err != nil {
		log.Printf("Error while accessing uploads of store %v : %v", storeId, err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing uploads of store %v : %v", storeId, err)}), err
	}

	// Convert DB query result to a plain array of strings
	var storeUploadIds = make([]string, len(uploads))

	for index, upload := range uploads {
		storeUploadIds[index] = strconv.Itoa(upload.StoreUploadIndex)
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("unable to commit transaction: %v", err)
		return openapi.Response(http.StatusInternalServerError, fmt.Sprintf("unable to commit transaction: %v", err)), err
	}

	log.Printf("Response: %v", storeUploadIds)

	return openapi.Response(http.StatusOK, storeUploadIds), nil
}
