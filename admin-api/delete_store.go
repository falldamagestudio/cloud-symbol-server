package admin_api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
)

func (s *ApiService) DeleteStore(ctx context.Context, storeId string) (openapi.ImplResponse, error) {

	storageClient, err := getStorageClient(ctx)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to create storage client"}), err
	}

	db := GetDB()
	if db == nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// Locate store in DB, and ensure store remains throughout entire txn
	store, err := models.Stores(qm.Where(models.StoreColumns.Name+" = ?", storeId), qm.For("update")).One(ctx, tx)
	if err == sql.ErrNoRows {
		log.Printf("Store %v not found; err = %v", storeId, err)
		tx.Rollback()
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Store %v not found", storeId)}), err
	} else if err != nil {
		log.Printf("error while accessing store: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, fmt.Sprintf("error while accessing store: %v", err)), err
	}

	// Delete all files in store
	files, err := models.StoreUploadFiles(qm.Where(models.StoreUploadTableColumns.StoreID+" = ?", store.StoreID), qm.InnerJoin("cloud_symbol_server."+models.TableNames.StoreUploads+" as "+models.TableNames.StoreUploads+" on "+models.StoreUploadTableColumns.UploadID+" = "+models.StoreUploadFileTableColumns.UploadID), qm.For("update")).All(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all files in store: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
	_, err = files.DeleteAll(ctx, tx)
	if err != nil {
		log.Printf("error while deleting all files in store: %v", err)
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

	if err = deleteAllObjectsInStore(ctx, storageClient, storeId); err != nil {
		if err != nil {
			log.Printf("Unable to delete all documents in collection, err = %v", err)
			return openapi.Response(http.StatusInternalServerError, nil), err
		}
	}

	return openapi.Response(http.StatusOK, nil), err
}
