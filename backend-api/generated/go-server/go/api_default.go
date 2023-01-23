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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"CreateStore",
			strings.ToUpper("Post"),
			"/stores/{storeId}",
			c.CreateStore,
		},
		{
			"CreateStoreUpload",
			strings.ToUpper("Post"),
			"/stores/{storeId}/uploads",
			c.CreateStoreUpload,
		},
		{
			"CreateToken",
			strings.ToUpper("Post"),
			"/tokens",
			c.CreateToken,
		},
		{
			"DeleteStore",
			strings.ToUpper("Delete"),
			"/stores/{storeId}",
			c.DeleteStore,
		},
		{
			"DeleteToken",
			strings.ToUpper("Delete"),
			"/tokens/{token}",
			c.DeleteToken,
		},
		{
			"ExpireStoreUpload",
			strings.ToUpper("Post"),
			"/stores/{storeId}/uploads/{uploadId}/expire",
			c.ExpireStoreUpload,
		},
		{
			"GetStoreFileBlobDownloadUrl",
			strings.ToUpper("Get"),
			"/stores/{storeId}/files/{fileId}/blobs/{blobId}/getDownloadUrl",
			c.GetStoreFileBlobDownloadUrl,
		},
		{
			"GetStoreFileBlobs",
			strings.ToUpper("Get"),
			"/stores/{storeId}/files/{fileId}/blobs",
			c.GetStoreFileBlobs,
		},
		{
			"GetStoreFiles",
			strings.ToUpper("Get"),
			"/stores/{storeId}/files",
			c.GetStoreFiles,
		},
		{
			"GetStoreUpload",
			strings.ToUpper("Get"),
			"/stores/{storeId}/uploads/{uploadId}",
			c.GetStoreUpload,
		},
		{
			"GetStoreUploads",
			strings.ToUpper("Get"),
			"/stores/{storeId}/uploads",
			c.GetStoreUploads,
		},
		{
			"GetStores",
			strings.ToUpper("Get"),
			"/stores",
			c.GetStores,
		},
		{
			"GetToken",
			strings.ToUpper("Get"),
			"/tokens/{token}",
			c.GetToken,
		},
		{
			"GetTokens",
			strings.ToUpper("Get"),
			"/tokens",
			c.GetTokens,
		},
		{
			"MarkStoreUploadAborted",
			strings.ToUpper("Post"),
			"/stores/{storeId}/uploads/{uploadId}/aborted",
			c.MarkStoreUploadAborted,
		},
		{
			"MarkStoreUploadCompleted",
			strings.ToUpper("Post"),
			"/stores/{storeId}/uploads/{uploadId}/completed",
			c.MarkStoreUploadCompleted,
		},
		{
			"MarkStoreUploadFileUploaded",
			strings.ToUpper("Post"),
			"/stores/{storeId}/uploads/{uploadId}/files/{fileId}/uploaded",
			c.MarkStoreUploadFileUploaded,
		},
		{
			"UpdateToken",
			strings.ToUpper("Put"),
			"/tokens/{token}",
			c.UpdateToken,
		},
	}
}

// CreateStore - Create a new store
func (c *DefaultApiController) CreateStore(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	storeIdParam := params["storeId"]
	
	result, err := c.service.CreateStore(r.Context(), storeIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// CreateStoreUpload - Start a new upload
func (c *DefaultApiController) CreateStoreUpload(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	storeIdParam := params["storeId"]
	
	createStoreUploadRequestParam := CreateStoreUploadRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&createStoreUploadRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCreateStoreUploadRequestRequired(createStoreUploadRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateStoreUpload(r.Context(), storeIdParam, createStoreUploadRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// CreateToken - Create a new token for current user
func (c *DefaultApiController) CreateToken(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.CreateToken(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteStore - Delete an existing store
func (c *DefaultApiController) DeleteStore(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	storeIdParam := params["storeId"]
	
	result, err := c.service.DeleteStore(r.Context(), storeIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteToken - Delete a token for current user
func (c *DefaultApiController) DeleteToken(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tokenParam := params["token"]
	
	result, err := c.service.DeleteToken(r.Context(), tokenParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ExpireStoreUpload - Expire store upload and consider files for GC
func (c *DefaultApiController) ExpireStoreUpload(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uploadIdParam, err := parseInt32Parameter(params["uploadId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	storeIdParam := params["storeId"]
	
	result, err := c.service.ExpireStoreUpload(r.Context(), uploadIdParam, storeIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetStoreFileBlobDownloadUrl - Request download URL for the binary blob associated with a particular store/file/blob-id
func (c *DefaultApiController) GetStoreFileBlobDownloadUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	storeIdParam := params["storeId"]
	
	fileIdParam := params["fileId"]
	
	blobIdParam := params["blobId"]
	
	result, err := c.service.GetStoreFileBlobDownloadUrl(r.Context(), storeIdParam, fileIdParam, blobIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetStoreFileBlobs - Fetch a list of blobs for a specific file in store
func (c *DefaultApiController) GetStoreFileBlobs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	storeIdParam := params["storeId"]
	
	fileIdParam := params["fileId"]
	
	sortParam := query.Get("sort")
	offsetParam, err := parseInt32Parameter(query.Get("offset"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	limitParam, err := parseInt32Parameter(query.Get("limit"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetStoreFileBlobs(r.Context(), storeIdParam, fileIdParam, sortParam, offsetParam, limitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetStoreFiles - Fetch a list of files in store
func (c *DefaultApiController) GetStoreFiles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	storeIdParam := params["storeId"]
	
	sortParam := query.Get("sort")
	offsetParam, err := parseInt32Parameter(query.Get("offset"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	limitParam, err := parseInt32Parameter(query.Get("limit"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetStoreFiles(r.Context(), storeIdParam, sortParam, offsetParam, limitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetStoreUpload - Fetch an upload
func (c *DefaultApiController) GetStoreUpload(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	uploadIdParam, err := parseInt32Parameter(params["uploadId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	storeIdParam := params["storeId"]
	
	sortParam := query.Get("sort")
	result, err := c.service.GetStoreUpload(r.Context(), uploadIdParam, storeIdParam, sortParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetStoreUploads - Fetch a list of uploads in store
func (c *DefaultApiController) GetStoreUploads(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	storeIdParam := params["storeId"]
	
	sortParam := query.Get("sort")
	offsetParam, err := parseInt32Parameter(query.Get("offset"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	limitParam, err := parseInt32Parameter(query.Get("limit"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetStoreUploads(r.Context(), storeIdParam, sortParam, offsetParam, limitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetStores - Fetch a list of all stores
func (c *DefaultApiController) GetStores(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	sortParam := query.Get("sort")
	result, err := c.service.GetStores(r.Context(), sortParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetToken - Fetch a token for current user
func (c *DefaultApiController) GetToken(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tokenParam := params["token"]
	
	result, err := c.service.GetToken(r.Context(), tokenParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTokens - Fetch a list of all tokens for current user
func (c *DefaultApiController) GetTokens(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	sortParam := query.Get("sort")
	result, err := c.service.GetTokens(r.Context(), sortParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// MarkStoreUploadAborted - Mark an upload as aborted
func (c *DefaultApiController) MarkStoreUploadAborted(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uploadIdParam, err := parseInt32Parameter(params["uploadId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	storeIdParam := params["storeId"]
	
	result, err := c.service.MarkStoreUploadAborted(r.Context(), uploadIdParam, storeIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// MarkStoreUploadCompleted - Mark an upload as completed
func (c *DefaultApiController) MarkStoreUploadCompleted(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uploadIdParam, err := parseInt32Parameter(params["uploadId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	storeIdParam := params["storeId"]
	
	result, err := c.service.MarkStoreUploadCompleted(r.Context(), uploadIdParam, storeIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// MarkStoreUploadFileUploaded - Mark a file within an upload as uploaded
func (c *DefaultApiController) MarkStoreUploadFileUploaded(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uploadIdParam, err := parseInt32Parameter(params["uploadId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	storeIdParam := params["storeId"]
	
	fileIdParam, err := parseInt32Parameter(params["fileId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.MarkStoreUploadFileUploaded(r.Context(), uploadIdParam, storeIdParam, fileIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateToken - Update details of a token for current user
func (c *DefaultApiController) UpdateToken(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tokenParam := params["token"]
	
	updateTokenRequestParam := UpdateTokenRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&updateTokenRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUpdateTokenRequestRequired(updateTokenRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateToken(r.Context(), tokenParam, updateTokenRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
