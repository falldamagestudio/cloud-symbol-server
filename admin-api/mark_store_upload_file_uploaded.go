package admin_api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
)

func (s *ApiService) MarkStoreUploadFileUploaded(context context.Context, uploadId string, storeId string, fileId int32) (openapi.ImplResponse, error) {

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

	log.Printf("Getting store upload ref")
	storeUploadRef, err := getStoreUploadRef(context, storeId, uploadId)
	if err != nil {
		log.Printf("Unable to fetch upload document for %v/%v, err = %v", storeId, uploadId, err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to fetch upload document for %v/%v", storeId, uploadId)}), err
	}

	storeUploadEntry.Files[fileId].Status = FileDBEntry_Status_Uploaded

	_, err = storeUploadRef.Set(context, storeUploadEntry)
	if err != nil {
		log.Printf("Unable to modify status for for %v/%v/%v, err = %v", storeId, uploadId, fileId, err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to modify status document for %v/%v/%v", storeId, uploadId, fileId)}), err
	}

	log.Printf("Status for %v/%v/%v set to %v", storeId, uploadId, fileId, FileDBEntry_Status_Uploaded)

	return openapi.Response(http.StatusOK, nil), nil
}
