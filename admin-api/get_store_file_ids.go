package admin_api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/volatiletech/sqlboiler/v4/boil"
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
	store, err := models.Stores(qm.Where(models.StoreColumns.Name+" = ?", storeId), qm.For("share")).One(ctx, tx)
	if err != nil {
		log.Printf("error while accessing store: %v", err)
		tx.Rollback()
		return openapi.Response(http.StatusInternalServerError, fmt.Sprintf("error while accessing store: %v", err)), err
	}

	boil.DebugMode = true

	// Fetch IDs of all files within store; remove duplicates based on name
	files, err := models.StoreUploadFiles(qm.Select(models.StoreUploadFileColumns.FileID), qm.Where(models.StoreUploadTableColumns.StoreID+" = ?", store.StoreID), qm.InnerJoin("cloud_symbol_server."+models.TableNames.StoreUploads+" as "+models.TableNames.StoreUploads+" on "+models.StoreUploadTableColumns.UploadID+" = "+models.StoreUploadFileTableColumns.UploadID), qm.Distinct(models.StoreUploadFileColumns.FileName+", MAX("+models.StoreUploadFileColumns.FileID+")"), qm.GroupBy(models.StoreUploadFileColumns.FileName), qm.OrderBy("MAX("+models.StoreUploadFileColumns.FileID+") ASC, "+models.StoreUploadFileColumns.FileName)).All(ctx, tx)
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
