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

func (s *ApiService) GetStoreUpload(ctx context.Context, uploadId string, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Getting store upload doc")

	var storeUploadEntry *StoreUploadEntry = nil

	err := runDBTransaction(ctx, func(ctx context.Context, client *firestore.Client, tx *firestore.Transaction) error {
		var err error = nil
		storeUploadEntry, err = getStoreUploadEntry(client, tx, storeId, uploadId)
		return err
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Upload %v / %v not found", storeId, uploadId)}), err
		} else {
			return openapi.Response(http.StatusInternalServerError, nil), err
		}
	}

	getStoreUploadResponse := openapi.GetStoreUploadResponse{}
	getStoreUploadResponse.Description = storeUploadEntry.Description
	getStoreUploadResponse.BuildId = storeUploadEntry.BuildId
	getStoreUploadResponse.Timestamp = storeUploadEntry.Timestamp
	// Uploads created before the progress API existed do not have any Status field in the DB
	// These uploads should be interpreted as having status "Unknown"
	if storeUploadEntry.Status != "" {
		getStoreUploadResponse.Status = storeUploadEntry.Status
	} else {
		getStoreUploadResponse.Status = StoreUploadEntry_Status_Unknown
	}

	for _, file := range storeUploadEntry.Files {

		// Uploaded files created before the progress API existed do not have any Status field in the DB
		// These files should be interpreted as having status "Unknown"
		status := FileDBEntry_Status_Unknown
		if file.Status != "" {
			status = file.Status
		}

		getStoreUploadResponse.Files = append(getStoreUploadResponse.Files, openapi.GetFileResponse{
			FileName: file.FileName,
			Hash:     file.Hash,
			Status:   status,
		})
	}

	log.Printf("Response: %v", getStoreUploadResponse)

	return openapi.Response(http.StatusOK, getStoreUploadResponse), nil
}
