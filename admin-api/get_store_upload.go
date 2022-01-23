package admin_api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
)

func (s *ApiService) GetStoreUpload(context context.Context, uploadId string, storeId string) (openapi.ImplResponse, error) {

	storeDoc, err := getStoreDoc(context, storeId)
	if err != nil {
		log.Printf("Unable to fetch store document for %v, err = %v", storeId, err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to fetch store document for %v", storeId)}), err
	}
	if storeDoc == nil {
		log.Printf("Store %v does not exist", storeId)
		return openapi.Response(http.StatusNotFound, &openapi.MessageResponse{Message: fmt.Sprintf("Store %v does not exist", storeId)}), err
	}

	log.Printf("Getting store upload doc")
	storeUploadDoc, err := getStoreUploadDoc(context, storeId, uploadId)
	if err != nil {
		log.Printf("Unable to fetch upload document for %v/%v, err = %v", storeId, uploadId, err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to fetch upload document for %v/%v", storeId, uploadId)}), err
	}
	if storeUploadDoc == nil {
		log.Printf("Upload doc %v/%v does not exist", storeId, uploadId)
		return openapi.Response(http.StatusNotFound, &openapi.MessageResponse{Message: fmt.Sprintf("Upload %v/%v does not exist", storeId, uploadId)}), err
	}

	log.Printf("Extracting upload doc data")
	var storeUploadEntry StoreUploadEntry
	if err = storeUploadDoc.DataTo(&storeUploadEntry); err != nil {
		log.Printf("Extracting upload doc data failed")
		return openapi.Response(http.StatusOK, &openapi.MessageResponse{Message: "Error while extracting contents of doc"}), err
	}

	getStoreUploadResponse := openapi.GetStoreUploadResponse{}
	getStoreUploadResponse.Description = storeUploadEntry.Description
	getStoreUploadResponse.BuildId = storeUploadEntry.BuildId
	getStoreUploadResponse.Timestamp = storeUploadEntry.Timestamp

	for _, file := range storeUploadEntry.Files {

		getStoreUploadResponse.Files = append(getStoreUploadResponse.Files, openapi.GetFileResponse{
			FileName: file.FileName,
			Hash:     file.Hash,
		})
	}

	log.Printf("Response: %v", getStoreUploadResponse)

	return openapi.Response(http.StatusOK, getStoreUploadResponse), nil
}
