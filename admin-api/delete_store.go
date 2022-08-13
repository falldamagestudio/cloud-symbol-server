package admin_api

import (
	"context"
	"log"
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
)

func (s *ApiService) DeleteStore(context context.Context, storeId string) (openapi.ImplResponse, error) {

	firestoreClient, err := firestoreClient(context)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	storageClient, err := getStorageClient(context)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to create storage client"}), err
	}

	symbolStoreBucketName, err := getSymbolStoreBucketName()
	if err != nil {
		log.Printf("Unable to determine symbol store bucket name: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to determine symbol store bucket name"}), err
	}

	if err = deleteAllObjectsInStore(context, storageClient, symbolStoreBucketName, storeId); err != nil {
		if err != nil {
			log.Printf("Unable to delete all documents in collection, err = %v", err)
			return openapi.Response(http.StatusInternalServerError, nil), err
		}
	}

	if err = deleteDocument(context, firestoreClient, firestoreClient.Collection(storesCollectionName).Doc(storeId), true); err != nil {
		if err != nil {
			log.Printf("Unable to delete document + child documents, err = %v", err)
			return openapi.Response(http.StatusInternalServerError, nil), err
		}
	}

	return openapi.Response(http.StatusOK, nil), err
}
