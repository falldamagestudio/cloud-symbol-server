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

func (s *ApiService) MarkStoreUploadFileUploaded(ctx context.Context, uploadId string, storeId string, fileId int32) (openapi.ImplResponse, error) {

	err := runDBTransaction(ctx, func(ctx context.Context, client *firestore.Client, tx *firestore.Transaction) error {
		storeUploadEntry, err := getStoreUploadEntry(client, tx, storeId, uploadId)
		if err != nil {
			return err
		}

		storeUploadEntry.Files[fileId].Status = FileDBEntry_Status_Uploaded

		err = updateStoreUploadEntry(client, tx, storeId, uploadId, storeUploadEntry)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return openapi.Response(http.StatusNotFound, openapi.MessageResponse{Message: fmt.Sprintf("Upload %v / %v not found", storeId, uploadId)}), err
		}
		return openapi.Response(http.StatusInternalServerError, nil), err
	}

	log.Printf("Status for %v/%v/%v set to %v", storeId, uploadId, fileId, FileDBEntry_Status_Uploaded)

	return openapi.Response(http.StatusOK, nil), nil
}
