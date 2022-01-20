package admin_api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
)

func (s *ApiService) GetTransaction(context context.Context, storeId string, transactionId string) (openapi.ImplResponse, error) {

	storeDoc, err := getStoreDoc(context, storeId)
	if err != nil {
		log.Printf("Unable to fetch store document for %v, err = %v", storeId, err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to fetch store document for %v", storeId)}), err
	}
	if storeDoc == nil {
		log.Printf("Store %v does not exist", storeId)
		return openapi.Response(http.StatusNotFound, &openapi.MessageResponse{Message: fmt.Sprintf("Store %v does not exist", storeId)}), err
	}

	log.Printf("Getting transaction doc")
	transactionDoc, err := getTransactionDoc(context, storeId, transactionId)
	if err != nil {
		log.Printf("Unable to fetch transaction document for %v/%v, err = %v", storeId, transactionId, err)
		return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: fmt.Sprintf("Unable to fetch transaction document for %v/%v", storeId, transactionId)}), err
	}

	log.Printf("Extracting transaction doc data")
	transactionDBEntry := TransactionDBEntry{}
	if err = transactionDoc.DataTo(&transactionDBEntry); err != nil {
		log.Printf("Extracting transaction doc data failed")
		return openapi.Response(http.StatusOK, &openapi.MessageResponse{Message: "Error while extracting contents of doc"}), err
	}

	getTransactionResponse := openapi.GetTransactionResponse{}
	getTransactionResponse.Description = transactionDBEntry.Description
	getTransactionResponse.BuildId = transactionDBEntry.BuildId
	getTransactionResponse.Timestamp = transactionDBEntry.Timestamp

	for _, file := range transactionDBEntry.Files {

		getTransactionResponse.Files = append(getTransactionResponse.Files, openapi.GetFileResponse{
			FileName: file.FileName,
			Hash:     file.Hash,
		})
	}

	log.Printf("Response: %v", getTransactionResponse)

	return openapi.Response(http.StatusOK, getTransactionResponse), nil
}
