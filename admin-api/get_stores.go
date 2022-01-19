package admin_api

import (
	"context"
	"log"
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
)

func (s *ApiService) GetStores(context context.Context) (openapi.ImplResponse, error) {

	stores, err := getStoresConfig(context)
	if err != nil {
		log.Printf("Unable to fetch stores: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("Stores: %v", stores)

	return openapi.Response(http.StatusOK, stores), nil
}
