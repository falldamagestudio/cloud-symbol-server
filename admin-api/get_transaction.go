package admin_api

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ApiService) GetTransaction(context context.Context, storeId string, transactionId string) (openapi.ImplResponse, error) {

	log.Printf("Getting transaction doc")
	transactionDoc, err := getTransactionDoc(context, storeId, transactionId)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Printf("Doc %v does not exist", transactionId)
			return openapi.Response(http.StatusNotFound, &openapi.MessageResponse{Message: "Doc does not exist"}), err
		} else {
			log.Printf("Getting transaction doc failed")
			return openapi.Response(http.StatusInternalServerError, &openapi.MessageResponse{Message: "Error while fetching doc"}), err
		}
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

func getTransactionDoc(context context.Context, storeId string, transactionId string) (*firestore.DocumentSnapshot, error) {

	firestoreClient, err := firestoreClient(context)
	if err != nil {
		log.Printf("Unable to talk to database: %v", err)
		return nil, err
	}

	transactionDoc, err := firestoreClient.Collection("stores").Doc(storeId).Collection("transactions").Doc(transactionId).Get(context)

	if err != nil {
		log.Printf("Unable to fetch transaction, err = %v", err)
		return nil, err
	}

	return transactionDoc, nil
}
