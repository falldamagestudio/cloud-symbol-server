package admin_api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
)

func (s *ApiService) GetStoreUploadIds(context context.Context, storeId string) (openapi.ImplResponse, error) {

	storeDoc, err := getStoreDoc(context, storeId)
	if err != nil {
		log.Printf("Unable to fetch store document for %v, err = %v", storeId, err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to fetch store document for %v", storeId)}), err
	}
	if storeDoc == nil {
		log.Printf("Store %v does not exist", storeId)
		return openapi.Response(http.StatusNotFound, &openapi.MessageResponse{Message: fmt.Sprintf("Store %v does not exist", storeId)}), err
	}

	log.Printf("Getting store upload IDs")
	storeUploadIDs, err := getStoreUploadIds(context, storeId)
	if err != nil {
		log.Printf("Unable to fetch all upload IDs for %v, err = %v", storeId, err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to fetch all upload IDs for %v", storeId)}), err
	}

	log.Printf("Response: %v", storeUploadIDs)

	return openapi.Response(http.StatusOK, storeUploadIDs), nil
}
