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

func (s *ApiService) GetStores(ctx context.Context) (openapi.ImplResponse, error) {

	log.Printf("Getting store names")

	db := GetDB()
	if db == nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	// Fetch names of all stores
	stores, err := models.Stores(qm.Select("name"), qm.OrderBy("store_id")).All(ctx, db)
	if err != nil {
		log.Printf("Error while accessing stores: %v", err)
		return openapi.Response(http.StatusInternalServerError, openapi.MessageResponse{Message: fmt.Sprintf("Error while accessing stores: %v", err)}), err
	}

	// Convert DB query result to a plain array of strings
	var storeIds = make([]string, len(stores))

	for index, store := range stores {
		storeIds[index] = store.Name
	}

	log.Printf("Stores: %v", storeIds)

	return openapi.Response(http.StatusOK, storeIds), nil
}
