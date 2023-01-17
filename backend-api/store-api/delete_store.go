package store_api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	cloud_storage "github.com/falldamagestudio/cloud-symbol-server/backend-api/cloud_storage"
	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/sql-db-models"
	postgres "github.com/falldamagestudio/cloud-symbol-server/backend-api/postgres"
)

func DeleteStore(ctx context.Context, storeId string) (openapi.ImplResponse, error) {

	storageClient, err := cloud_storage.GetStorageClient(ctx)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to create storage client"}), err
	}

	tx, err := postgres.BeginDBTransaction(ctx)
	if err != nil {
		log.Printf("Err: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	// Locate store in DB, and ensure store remains throughout entire txn
	store, err := models.Stores(
		qm.Where(models.StoreColumns.Name+" = ?", storeId),
		qm.For("update"),
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

	// Delete all upload-files in store
	uploadFiles, err := models.StoreUploadFiles(
		qm.InnerJoin("cloud_symbol_server."+models.TableNames.StoreUploads+" on "+models.StoreUploadTableColumns.UploadID+" = "+models.StoreUploadFileTableColumns.UploadID),
		qm.Where(models.StoreUploadTableColumns.StoreID+" = ?", store.StoreID),
		qm.For("update"),
	).All(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all upload-files in store: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
	_, err = uploadFiles.DeleteAll(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all upload-files in store: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Delete all uploads in store
	_, err = store.StoreUploads().DeleteAll(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all uploads in store: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Delete all file-blobs in store
	storeFileBlobs, err := models.StoreFileBlobs(
		qm.InnerJoin("cloud_symbol_server."+models.TableNames.StoreFiles+" on "+models.StoreFileTableColumns.FileID+" = "+models.StoreFileBlobTableColumns.FileID),
		qm.Where(models.StoreFileTableColumns.StoreID+" = ?", store.StoreID),
		qm.For("update"),
	).All(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all file-blobs in store: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
	_, err = storeFileBlobs.DeleteAll(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all file-blobs in store: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Delete all store-files in store
	storeFiles, err := models.StoreFiles(
		qm.Where(models.StoreFileTableColumns.StoreID+" = ?", store.StoreID),
		qm.For("update"),
	).All(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all store-files in store: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
	_, err = storeFiles.DeleteAll(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all store-files in store: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Delete store
	_, err = store.Delete(ctx, tx)
	if err != nil {
		log.Printf("error while deleting store entry: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("error while committing txn: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Delete all related files in Cloud Storage

	if err = cloud_storage.DeleteAllObjectsInStore(ctx, storageClient, storeId); err != nil {
		if err != nil {
			log.Printf("Unable to delete all documents in collection, err = %v", err)
			return openapi.Response(http.StatusInternalServerError, nil), err
		}
	}

	return openapi.Response(http.StatusOK, nil), err
}
