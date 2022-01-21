package admin_api

import (
	"context"
	"log"
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ApiService) CreateStore(context context.Context, store string) (openapi.ImplResponse, error) {

	firestoreClient, err := firestoreClient(context)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to talk to database"}), err
	}

	_, err = firestoreClient.Collection("stores").Doc(store).Create(context, &StoreEntry{LatestTransactionId: -1})
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			log.Printf("Store already exists, err = %v", err)
			return openapi.Response(http.StatusConflict, &openapi.MessageResponse{Message: "Store already exists"}), err
		} else {
			log.Printf("Unable to create store, err = %v", err)
			return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to create store"}), err
		}
	}

	return openapi.Response(http.StatusOK, nil), err
}
