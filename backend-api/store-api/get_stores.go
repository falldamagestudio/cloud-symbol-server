package store_api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/sql-db-models"
	postgres "github.com/falldamagestudio/cloud-symbol-server/backend-api/postgres"
)

var getStoresSortOptions = map[string]string{
	"":      models.StoreColumns.Name,
	"name":  models.StoreColumns.Name,
	"-name": models.StoreColumns.Name + " desc",
}

func GetStores(ctx context.Context, sort string) (openapi.ImplResponse, error) {

	log.Printf("Getting store names, sort %v", sort)

	// Handle sorting
	orderByOption := ""
	if sortOption, ok := getStoresSortOptions[sort]; ok {
		orderByOption = sortOption
	} else {
		log.Printf("Unsupported sort option: %v", sort)
		return openapi.Response(http.StatusBadRequest, openapi.MessageResponse{Message: fmt.Sprintf("Unsupported sort option: %v", sort)}), nil
	}

	db := postgres.GetDB()
	if db == nil {
		return openapi.Response(http.StatusInternalServerError, nil), errors.New("no DB")
	}

	// Fetch names of all stores
	stores, err := models.Stores(
		qm.Select(models.StoreColumns.Name),
		qm.OrderBy(orderByOption),
	).All(ctx, db)
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
