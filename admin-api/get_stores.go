package admin_api

import (
	"context"
	"log"
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go"
)

func (s *ApiService) GetStores(context context.Context) (openapi.ImplResponse, error) {

	stores, err := getStoresConfig()
	if err != nil {
		log.Printf("Unable to get stores config: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to get stores config"}), err
	}

	log.Printf("Stores: %v", stores)

	return openapi.Response(http.StatusOK, stores), nil
}
