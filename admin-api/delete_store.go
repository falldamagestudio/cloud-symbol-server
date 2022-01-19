package admin_api

import (
	"context"
	"log"
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
)

func (s *ApiService) DeleteStore(context context.Context, store string) (openapi.ImplResponse, error) {

	firestoreClient, err := firestoreClient(context)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	_, err = firestoreClient.Collection("stores").Doc(store).Delete(context)
	if err != nil {
		log.Printf("Unable to delete store, err = %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), err
}
