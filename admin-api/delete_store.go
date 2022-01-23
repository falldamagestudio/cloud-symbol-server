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

	storageClient, _, err := getStorageClient(context)
	if err != nil {
		log.Printf("Unable to create storageClient: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to create storage client"}), err
	}

	symbolStoreBucketName, err := getSymbolStoreBucketName()
	if err != nil {
		log.Printf("Unable to determine symbol store bucket name: %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to determine symbol store bucket name"}), err
	}

	storeDoc, err := getStoreDoc(context, storeId)
	if storeDoc == nil && err == nil {
		log.Printf("Store does not exist")
		return openapi.Response(http.StatusNotFound, &openapi.MessageResponse{Message: "Store does not exist"}), nil
	} else if err != nil {
		log.Printf("Unable to fetch store doc, err = %v", err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Unable to fetch store doc"}), err
	}

	if err = deleteAllObjectsInStore(context, storageClient, symbolStoreBucketName, storeId); err != nil {
		if err != nil {
			log.Printf("Unable to delete all documents in collection, err = %v", err)
			return openapi.Response(http.StatusInternalServerError, nil), err
		}
	}

	if err = deleteAllDocumentsInCollection(context, firestoreClient, firestoreClient.Collection(storesCollectionName).Doc(storeId).Collection(storeUploadsCollectionName), 100); err != nil {
		if err != nil {
			log.Printf("Unable to delete all documents in collection, err = %v", err)
			return openapi.Response(http.StatusInternalServerError, nil), err
		}
	}

	_, err = firestoreClient.Collection(storesCollectionName).Doc(storeId).Delete(context)
	if err != nil {
		log.Printf("Unable to delete store, err = %v", err)
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	return openapi.Response(http.StatusOK, nil), err
}
