package admin_api

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
)

func (s *ApiService) MarkStoreUploadCompleted(ctx context.Context, uploadId string, storeId string) (openapi.ImplResponse, error) {

	err := runDBTransaction(ctx, func(ctx context.Context, client *firestore.Client, tx *firestore.Transaction) error {
		storeUploadEntry, err := getStoreUploadEntry(client, tx, storeId, uploadId)
		if err != nil {
			return err
		}

		storeUploadEntry.Status = StoreUploadEntry_Status_Completed

		err = updateStoreUploadEntry(client, tx, storeId, uploadId, storeUploadEntry)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		// TOOD: Smarter error handling
		return openapi.Response(http.StatusBadRequest, nil), err
	}

	log.Printf("Status for %v/%v set to %v", storeId, uploadId, StoreUploadEntry_Status_Completed)

	return openapi.Response(http.StatusOK, nil), nil
}
