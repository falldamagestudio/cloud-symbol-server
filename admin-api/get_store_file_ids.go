package admin_api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ApiService) GetStoreFileIds(ctx context.Context, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Getting store file IDs")

	var storeFileIds []string = nil

	err := runDBTransaction(ctx, func(ctx context.Context, client *firestore.Client, tx *firestore.Transaction) error {
		var err error = nil
		_, err = getStoreEntry(client, tx, storeId)
		if err != nil {
			return err
		}

		storeFileIds, err = getStoreFileIds(ctx, client, storeId)
		return err
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Printf("Store %v not found: %v", storeId, err)
			return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Store %v not found", storeId)}), err
		} else {
			log.Printf("Error while fetching file IDs for store %v: %v", storeId, err)
			return openapi.Response(http.StatusInternalServerError, nil), err
		}
	}

	log.Printf("Response: %v", storeFileIds)

	return openapi.Response(http.StatusOK, storeFileIds), nil
}