/*
 * Cloud Symbol Server Admin API
 *
 * This is the API that is used to manage stores and uploads in Cloud Symbol Server
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface { 
	CreateStore(http.ResponseWriter, *http.Request)
	CreateStoreUpload(http.ResponseWriter, *http.Request)
	CreateToken(http.ResponseWriter, *http.Request)
	DeleteStore(http.ResponseWriter, *http.Request)
	DeleteToken(http.ResponseWriter, *http.Request)
	ExpireStoreUpload(http.ResponseWriter, *http.Request)
	GetStoreFileBlobDownloadUrl(http.ResponseWriter, *http.Request)
	GetStoreFileBlobs(http.ResponseWriter, *http.Request)
	GetStoreFiles(http.ResponseWriter, *http.Request)
	GetStoreUpload(http.ResponseWriter, *http.Request)
	GetStoreUploads(http.ResponseWriter, *http.Request)
	GetStores(http.ResponseWriter, *http.Request)
	GetToken(http.ResponseWriter, *http.Request)
	GetTokens(http.ResponseWriter, *http.Request)
	MarkStoreUploadAborted(http.ResponseWriter, *http.Request)
	MarkStoreUploadCompleted(http.ResponseWriter, *http.Request)
	MarkStoreUploadFileUploaded(http.ResponseWriter, *http.Request)
	UpdateToken(http.ResponseWriter, *http.Request)
}


// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultApiServicer interface { 
	CreateStore(context.Context, string) (ImplResponse, error)
	CreateStoreUpload(context.Context, string, CreateStoreUploadRequest) (ImplResponse, error)
	CreateToken(context.Context) (ImplResponse, error)
	DeleteStore(context.Context, string) (ImplResponse, error)
	DeleteToken(context.Context, string) (ImplResponse, error)
	ExpireStoreUpload(context.Context, int32, string) (ImplResponse, error)
	GetStoreFileBlobDownloadUrl(context.Context, string, string, string) (ImplResponse, error)
	GetStoreFileBlobs(context.Context, string, string, int32, int32) (ImplResponse, error)
	GetStoreFiles(context.Context, string, int32, int32) (ImplResponse, error)
	GetStoreUpload(context.Context, int32, string) (ImplResponse, error)
	GetStoreUploads(context.Context, string, int32, int32) (ImplResponse, error)
	GetStores(context.Context) (ImplResponse, error)
	GetToken(context.Context, string) (ImplResponse, error)
	GetTokens(context.Context) (ImplResponse, error)
	MarkStoreUploadAborted(context.Context, int32, string) (ImplResponse, error)
	MarkStoreUploadCompleted(context.Context, int32, string) (ImplResponse, error)
	MarkStoreUploadFileUploaded(context.Context, int32, string, int32) (ImplResponse, error)
	UpdateToken(context.Context, string, UpdateTokenRequest) (ImplResponse, error)
}
