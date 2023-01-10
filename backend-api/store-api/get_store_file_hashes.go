package store_api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/sql-db-models"
	postgres "github.com/falldamagestudio/cloud-symbol-server/backend-api/postgres"
)

func GetStoreFileHashes(ctx context.Context, storeId string, fileId string, offset int32, limit int32) (openapi.ImplResponse, error) {

	log.Printf("Getting store file hashes; store %v, file %v, offset %v, limit %v", storeId, fileId, offset, limit)

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

	// Locate file in DB, and ensure file remains throughout entire txn
	file, err := models.StoreFiles(
		qm.Where(models.StoreFileColumns.StoreID+" = ?", store.StoreID),
		qm.And(models.StoreFileColumns.FileName+" = ?", fileId),
		qm.For("share"),
	).One(ctx, tx)
	if err == sql.ErrNoRows {
		log.Printf("Store file %v / %v not found; err = %v", storeId, fileId, err)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Store / file %v / %v not found", storeId, fileId)}), err
	} else if err != nil {
		log.Printf("error while accessing store file: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, fmt.Sprintf("error while accessing store file : %v", err)), err
	}

	// Count total number of hashes in file that match filter query, ignoring pagination
	total, err := models.StoreFileHashes(
		qm.Distinct(models.StoreFileHashColumns.HashID),
		qm.Where(models.StoreFileHashTableColumns.FileID+" = ?", file.FileID),
	).Count(ctx, tx)
	if err != nil {
		log.Printf("Error while accessing hashes in store-file %v/%v : %v", storeId, fileId, err)
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing hashes in store-file %v/%v : %v", storeId, fileId, err)}), err
	}

	// Fetch all hashes within file
	hashes, err := models.StoreFileHashes(
		qm.Where(models.StoreFileHashTableColumns.FileID+" = ?", file.FileID),
		qm.OrderBy(models.StoreFileHashColumns.HashID),
		qm.Offset(int(offset)),
		qm.Limit(int(limit)),
	).All(ctx, tx)
	if err != nil {
		log.Printf("Error while accessing hashes in store-file %v/%v : %v", storeId, fileId, err)
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing hashes in store-file %v/%v : %v", storeId, fileId, err)}), err
	}

	// Convert DB query result to a plain array of strings
	var storeFileHashes = make([]openapi.GetStoreFileHashResponse, len(hashes))

	for index, hash := range hashes {
		storeFileHashes[index] = openapi.GetStoreFileHashResponse{
			Hash: hash.Hash,
			Status: openapi.StoreFileHashStatus(hash.Status),
		}
	}

	storeFileHashesResponse := &openapi.GetStoreFileHashesResponse{
		Hashes: storeFileHashes,
		Pagination: openapi.PaginationResponse{
			Total: int32(total),
		},
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("unable to commit transaction: %v", err)
		return openapi.Response(http.StatusInternalServerError, fmt.Sprintf("unable to commit transaction: %v", err)), err
	}

	log.Printf("Response: %v", storeFileHashesResponse)

	return openapi.Response(http.StatusOK, storeFileHashesResponse), nil
}
