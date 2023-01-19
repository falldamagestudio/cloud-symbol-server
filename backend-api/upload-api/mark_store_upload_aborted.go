package upload_api

import (
	"context"
	"log"

	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/sql-db-models"
)

func MarkStoreUploadAborted(ctx context.Context, uploadId int32, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Aborting upload %v/%v", storeId, uploadId)
	response, err := HandleUploadExpiryOrAbort(ctx, storeId, uploadId, models.StoreUploadStatusAborted)
	return response, err
}
