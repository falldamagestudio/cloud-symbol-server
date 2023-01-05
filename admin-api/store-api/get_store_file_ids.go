package store_api

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
	postgres "github.com/falldamagestudio/cloud-symbol-server/admin-api/postgres"
)

func GetStoreFileIds(ctx context.Context, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Getting store file IDs")

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
	if err != nil {
		log.Printf("error while accessing store: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, fmt.Sprintf("error while accessing store: %v", err)), err
	}

	// Fetch IDs of all files within store
	files, err := models.StoreFiles(
		qm.Select(models.StoreFileColumns.FileID),
		qm.Where(models.StoreFileTableColumns.StoreID+" = ?", store.StoreID),
		qm.OrderBy(models.StoreFileColumns.FileID),
	).All(ctx, tx)
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
