package admin_api

import (
	"context"
	"log"

	openapi "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/go-server/go"
	models "github.com/falldamagestudio/cloud-symbol-server/admin-api/generated/sql-db-models"
)

func (s *ApiService) ExpireStoreUpload(ctx context.Context, uploadId string, storeId string) (openapi.ImplResponse, error) {

	log.Printf("Expiring upload %v/%v", storeId, uploadId)
	response, err := HandleUploadExpiryOrAbort(ctx, storeId, uploadId, models.StoreUploadStatusExpired)
	return response, err
}
