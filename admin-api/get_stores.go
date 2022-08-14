package admin_api

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
)

func (s *ApiService) GetStores(ctx context.Context) (openapi.ImplResponse, error) {

	log.Printf("Getting store upload doc")

	var storeIds []string = nil

	err := runDBOperation(ctx, func(ctx context.Context, client *firestore.Client) error {
		var err error = nil
		storeIds, err = getStoreIds(ctx, client)
		return err
	})
	if err != nil {
		log.Printf("Error while fetching stores: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("Stores: %v", storeIds)

	return openapi.Response(http.StatusOK, storeIds), nil
}
