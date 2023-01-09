package store_api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/sql-db-models"
	postgres "github.com/falldamagestudio/cloud-symbol-server/backend-api/postgres"
)

func GetStoreFiles(ctx context.Context, storeId string, offset int32, limit int32) (openapi.ImplResponse, error) {

	log.Printf("Getting store files; store %v, offset %v, limit %v", storeId, offset, limit)

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

	// Count total number of files in store that match filter query, ignoring pagination
	total, err := models.StoreFiles(
		qm.Distinct(models.StoreFileColumns.FileID),
		qm.Where(models.StoreFileTableColumns.StoreID+" = ?", store.StoreID),
	).Count(ctx, tx)
	if err != nil {
		log.Printf("Error while accessing files in store %v : %v", storeId, err)
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing files in store %v : %v", storeId, err)}), err
	}

	boil.DebugMode = true

	log.Printf("offset: %v", offset)
	log.Printf("limit: %v", limit)

	// Fetch all files within store
	files, err := models.StoreFiles(
		qm.Where(models.StoreFileTableColumns.StoreID+" = ?", store.StoreID),
		qm.OrderBy(models.StoreFileColumns.FileID),
		qm.Offset(int(offset)),
		qm.Limit(int(limit)),
	).All(ctx, tx)
	if err != nil {
		log.Printf("Error while accessing files in store %v : %v", storeId, err)
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing files in store %v : %v", storeId, err)}), err
	}

	// Convert DB query result to a plain array of strings
	var storeFiles = make([]string, len(files))

	for index, file := range files {
		storeFiles[index] = file.FileName
	}

	storeFilesResponse := &openapi.GetStoreFilesResponse{
		Files: storeFiles,
		Pagination: openapi.PaginationResponse{
			Total: int32(total),
		},
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("unable to commit transaction: %v", err)
		return openapi.Response(http.StatusInternalServerError, fmt.Sprintf("unable to commit transaction: %v", err)), err
	}

	log.Printf("Response: %v", storeFilesResponse)

	return openapi.Response(http.StatusOK, storeFilesResponse), nil
}
