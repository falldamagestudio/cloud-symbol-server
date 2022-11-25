package admin_api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
)

func (s *ApiService) DeleteStore(ctx context.Context, storeId string) (openapi.ImplResponse, error) {

	storageClient, err := getStorageClient(ctx)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to create storage client"}), err
	}

	// Validate that store exists

	store, err := sqlGetStore(ctx, storeId)
	if err != nil || store == nil {
		log.Printf("Store %v not found; err = %v", storeId, err)
		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Store %v not found", storeId)}), err
	}

	// if err = runDBTransaction(ctx, func(ctx context.Context, client *firestore.Client, tx *firestore.Transaction) error {
	// 	var err error = nil
	// 	_, err = getStoreEntry(client, tx, storeId)
	// 	return err
	// }); err != nil {
	// 	if status.Code(err) == codes.NotFound {
	// 		log.Printf("Store %v not found; err = %v", storeId, err)
	// 		return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Store %v not found", storeId)}), err
	// 	} else {
	// 		return openapi.Response(http.StatusInternalServerError, nil), err
	// 	}
	// }

	// Delete all related files in Cloud Storage

	if err = deleteAllObjectsInStore(ctx, storageClient, storeId); err != nil {
		if err != nil {
			log.Printf("Unable to delete all documents in collection, err = %v", err)
			return openapi.Response(http.StatusInternalServerError, nil), err
		}
	}

	// Delete all related documents in Firestore

	if err = sqlDeleteStore(ctx, storeId); err != nil {
		log.Printf("Unable to delete store, err = %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	// if err = deleteDocument(ctx, firestoreClient, firestoreClient.Collection(storesCollectionName).Doc(storeId), true); err != nil {
	// 	if err != nil {
	// 		log.Printf("Unable to delete document + child documents, err = %v", err)
	// 		return openapi.Response(http.StatusInternalServerError, nil), err
	// 	}
	// }

	return openapi.Response(http.StatusOK, nil), err
}
