package backend_api

import (
	"context"

	openapi "github.com/falldamagestudio/cloud-symbol-server/backend-api/generated/go-server/go"
	store_api "github.com/falldamagestudio/cloud-symbol-server/backend-api/store-api"
	token_api "github.com/falldamagestudio/cloud-symbol-server/backend-api/token-api"
	upload_api "github.com/falldamagestudio/cloud-symbol-server/backend-api/upload-api"
)

// store API

func (s *ApiService) CreateStore(ctx context.Context, storeId string) (openapi.ImplResponse, error) {
	return store_api.CreateStore(ctx, storeId)
}

func (s *ApiService) DeleteStore(ctx context.Context, storeId string) (openapi.ImplResponse, error) {
	return store_api.DeleteStore(ctx, storeId)
}

func (s *ApiService) GetStoreFiles(ctx context.Context, storeId string, offset int32, limit int32) (openapi.ImplResponse, error) {
	return store_api.GetStoreFiles(ctx, storeId, offset, limit)
}

func (s *ApiService) GetStores(ctx context.Context) (openapi.ImplResponse, error) {
	return store_api.GetStores(ctx)
}

func (s *ApiService) GetStoreFileBlobs(ctx context.Context, storeId string, fileId string, offset int32, limit int32) (openapi.ImplResponse, error) {
	return store_api.GetStoreFileBlobs(ctx, storeId, fileId, offset, limit)
}

func (s *ApiService) GetStoreFileBlobDownloadUrl(ctx context.Context, storeId string, fileId string, blobId string) (openapi.ImplResponse, error) {
	return store_api.GetStoreFileBlobDownloadUrl(ctx, storeId, fileId, blobId)
}

// Upload API

func (s *ApiService) CreateStoreUpload(ctx context.Context, storeId string, createStoreUploadRequest openapi.CreateStoreUploadRequest) (openapi.ImplResponse, error) {
	return upload_api.CreateStoreUpload(ctx, storeId, createStoreUploadRequest)
}

func (s *ApiService) ExpireStoreUpload(ctx context.Context, uploadId int32, storeId string) (openapi.ImplResponse, error) {
	return upload_api.ExpireStoreUpload(ctx, uploadId, storeId)
}

func (s *ApiService) GetStoreUploads(ctx context.Context, storeId string, offset int32, limit int32) (openapi.ImplResponse, error) {
	return upload_api.GetStoreUploads(ctx, storeId, offset, limit)
}

func (s *ApiService) GetStoreUpload(ctx context.Context, uploadId int32, storeId string) (openapi.ImplResponse, error) {
	return upload_api.GetStoreUpload(ctx, uploadId, storeId)
}

func (s *ApiService) MarkStoreUploadAborted(ctx context.Context, uploadId int32, storeId string) (openapi.ImplResponse, error) {
	return upload_api.MarkStoreUploadAborted(ctx, uploadId, storeId)
}

func (s *ApiService) MarkStoreUploadCompleted(ctx context.Context, uploadId int32, storeId string) (openapi.ImplResponse, error) {
	return upload_api.MarkStoreUploadCompleted(ctx, uploadId, storeId)
}

func (s *ApiService) MarkStoreUploadFileUploaded(ctx context.Context, uploadId int32, storeId string, fileId int32) (openapi.ImplResponse, error) {
	return upload_api.MarkStoreUploadFileUploaded(ctx, uploadId, storeId, fileId)
}

// Token API

func (s *ApiService) CreateToken(ctx context.Context) (openapi.ImplResponse, error) {
	return token_api.CreateToken(ctx)
}

func (s *ApiService) DeleteToken(ctx context.Context, token string) (openapi.ImplResponse, error) {
	return token_api.DeleteToken(ctx, token)
}

func (s *ApiService) GetToken(ctx context.Context, token string) (openapi.ImplResponse, error) {
	return token_api.GetToken(ctx, token)
}

func (s *ApiService) GetTokens(ctx context.Context) (openapi.ImplResponse, error) {
	return token_api.GetTokens(ctx)
}

func (s *ApiService) UpdateToken(ctx context.Context, token string, updateTokenRequest openapi.UpdateTokenRequest) (openapi.ImplResponse, error) {
	return token_api.UpdateToken(ctx, token, updateTokenRequest)
}
