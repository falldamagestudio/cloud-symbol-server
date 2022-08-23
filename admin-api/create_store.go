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

func (s *ApiService) CreateStore(ctx context.Context, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Creating store")

	if err := runDBTransaction(ctx, func(ctx context.Context, client *firestore.Client, tx *firestore.Transaction) error {
		err := createStoreEntry(client, tx, storeId, &StoreEntry{LatestUploadId: -1})
		return err

	}); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			log.Printf("Store %v already exists; err = %v", storeId, err)
			return openapi.Response(http.StatusConflict, openapi.MessageResponse{Message: fmt.Sprintf("Store %v already exists", storeId)}), err
		} else {
			log.Printf("CreateStore err = %v", err)
			return openapi.Response(http.StatusInternalServerError, nil), err
		}
	}

	return openapi.Response(http.StatusOK, nil), nil
}
