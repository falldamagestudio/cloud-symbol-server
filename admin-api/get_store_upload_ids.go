package admin_api

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
)

func (s *ApiService) GetStoreUploadIds(ctx context.Context, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Getting store upload IDs")

	var storeUploadIds []string = nil

	err := runDBOperation(ctx, func(ctx context.Context, client *firestore.Client) error {
		var err error = nil
		storeUploadIds, err = getStoreUploadIds(ctx, client, storeId)
		return err
	})
	if err != nil {
		log.Printf("Unable to fetch all upload IDs for %v, err = %v", storeId, err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("Response: %v", storeUploadIds)

	return openapi.Response(http.StatusOK, storeUploadIds), nil
}
