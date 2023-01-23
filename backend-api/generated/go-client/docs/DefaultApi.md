# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStore**](DefaultApi.md#CreateStore) | **Post** /stores/{storeId} | Create a new store
[**CreateStoreUpload**](DefaultApi.md#CreateStoreUpload) | **Post** /stores/{storeId}/uploads | Start a new upload
[**CreateToken**](DefaultApi.md#CreateToken) | **Post** /tokens | Create a new token for current user
[**DeleteStore**](DefaultApi.md#DeleteStore) | **Delete** /stores/{storeId} | Delete an existing store
[**DeleteToken**](DefaultApi.md#DeleteToken) | **Delete** /tokens/{token} | Delete a token for current user
[**ExpireStoreUpload**](DefaultApi.md#ExpireStoreUpload) | **Post** /stores/{storeId}/uploads/{uploadId}/expire | Expire store upload and consider files for GC
[**GetStoreFileBlobDownloadUrl**](DefaultApi.md#GetStoreFileBlobDownloadUrl) | **Get** /stores/{storeId}/files/{fileId}/blobs/{blobId}/getDownloadUrl | Request download URL for the binary blob associated with a particular store/file/blob-id
[**GetStoreFileBlobs**](DefaultApi.md#GetStoreFileBlobs) | **Get** /stores/{storeId}/files/{fileId}/blobs | Fetch a list of blobs for a specific file in store
[**GetStoreFiles**](DefaultApi.md#GetStoreFiles) | **Get** /stores/{storeId}/files | Fetch a list of files in store
[**GetStoreUpload**](DefaultApi.md#GetStoreUpload) | **Get** /stores/{storeId}/uploads/{uploadId} | Fetch an upload
[**GetStoreUploads**](DefaultApi.md#GetStoreUploads) | **Get** /stores/{storeId}/uploads | Fetch a list of uploads in store
[**GetStores**](DefaultApi.md#GetStores) | **Get** /stores | Fetch a list of all stores
[**GetToken**](DefaultApi.md#GetToken) | **Get** /tokens/{token} | Fetch a token for current user
[**GetTokens**](DefaultApi.md#GetTokens) | **Get** /tokens | Fetch a list of all tokens for current user
[**MarkStoreUploadAborted**](DefaultApi.md#MarkStoreUploadAborted) | **Post** /stores/{storeId}/uploads/{uploadId}/aborted | Mark an upload as aborted
[**MarkStoreUploadCompleted**](DefaultApi.md#MarkStoreUploadCompleted) | **Post** /stores/{storeId}/uploads/{uploadId}/completed | Mark an upload as completed
[**MarkStoreUploadFileUploaded**](DefaultApi.md#MarkStoreUploadFileUploaded) | **Post** /stores/{storeId}/uploads/{uploadId}/files/{fileId}/uploaded | Mark a file within an upload as uploaded
[**UpdateToken**](DefaultApi.md#UpdateToken) | **Put** /tokens/{token} | Update details of a token for current user



## CreateStore

> CreateStore(ctx, storeId).Execute()

Create a new store

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeId := "storeId_example" // string | ID of store to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateStore(context.Background(), storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | ID of store to create | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStoreUpload

> CreateStoreUploadResponse CreateStoreUpload(ctx, storeId).CreateStoreUploadRequest(createStoreUploadRequest).Execute()

Start a new upload

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeId := "storeId_example" // string | ID of the store containing the upload
    createStoreUploadRequest := *openapiclient.NewCreateStoreUploadRequest() // CreateStoreUploadRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateStoreUpload(context.Background(), storeId).CreateStoreUploadRequest(createStoreUploadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStoreUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStoreUpload`: CreateStoreUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateStoreUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | ID of the store containing the upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoreUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createStoreUploadRequest** | [**CreateStoreUploadRequest**](CreateStoreUploadRequest.md) |  | 

### Return type

[**CreateStoreUploadResponse**](CreateStoreUploadResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateToken

> CreateTokenResponse CreateToken(ctx).Execute()

Create a new token for current user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: CreateTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


### Return type

[**CreateTokenResponse**](CreateTokenResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStore

> DeleteStore(ctx, storeId).Execute()

Delete an existing store

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeId := "storeId_example" // string | ID of store to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteStore(context.Background(), storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | ID of store to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToken

> DeleteToken(ctx, token).Execute()

Delete a token for current user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string | ID of the token to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | ID of the token to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpireStoreUpload

> ExpireStoreUpload(ctx, uploadId, storeId).Execute()

Expire store upload and consider files for GC

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uploadId := int32(56) // int32 | ID of the upload to fetch
    storeId := "storeId_example" // string | ID of the store containing the upload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ExpireStoreUpload(context.Background(), uploadId, storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExpireStoreUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadId** | **int32** | ID of the upload to fetch | 
**storeId** | **string** | ID of the store containing the upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpireStoreUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoreFileBlobDownloadUrl

> GetStoreFileBlobDownloadUrlResponse GetStoreFileBlobDownloadUrl(ctx, storeId, fileId, blobId).Execute()

Request download URL for the binary blob associated with a particular store/file/blob-id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeId := "storeId_example" // string | ID of the store containing the file
    fileId := "fileId_example" // string | ID of the file
    blobId := "blobId_example" // string | ID of the blob

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStoreFileBlobDownloadUrl(context.Background(), storeId, fileId, blobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStoreFileBlobDownloadUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoreFileBlobDownloadUrl`: GetStoreFileBlobDownloadUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStoreFileBlobDownloadUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | ID of the store containing the file | 
**fileId** | **string** | ID of the file | 
**blobId** | **string** | ID of the blob | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreFileBlobDownloadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetStoreFileBlobDownloadUrlResponse**](GetStoreFileBlobDownloadUrlResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoreFileBlobs

> GetStoreFileBlobsResponse GetStoreFileBlobs(ctx, storeId, fileId).Sort(sort).Offset(offset).Limit(limit).Execute()

Fetch a list of blobs for a specific file in store

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeId := "storeId_example" // string | ID of the store containing the file
    fileId := "fileId_example" // string | ID of the file
    sort := "name" // string | Sort the content of the response according to a given key, potentially in descending order (optional)
    offset := int32(56) // int32 | How many entries to skip (used for pagination of results) (optional) (default to 0)
    limit := int32(56) // int32 | Max number of results to return (used for pagination of results) (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStoreFileBlobs(context.Background(), storeId, fileId).Sort(sort).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStoreFileBlobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoreFileBlobs`: GetStoreFileBlobsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStoreFileBlobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | ID of the store containing the file | 
**fileId** | **string** | ID of the file | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreFileBlobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sort** | **string** | Sort the content of the response according to a given key, potentially in descending order | 
 **offset** | **int32** | How many entries to skip (used for pagination of results) | [default to 0]
 **limit** | **int32** | Max number of results to return (used for pagination of results) | [default to 25]

### Return type

[**GetStoreFileBlobsResponse**](GetStoreFileBlobsResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoreFiles

> GetStoreFilesResponse GetStoreFiles(ctx, storeId).Sort(sort).Offset(offset).Limit(limit).Execute()

Fetch a list of files in store

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeId := "storeId_example" // string | ID of the store containing the files
    sort := "name" // string | Sort the content of the response according to a given key, potentially in descending order (optional)
    offset := int32(56) // int32 | How many entries to skip (used for pagination of results) (optional) (default to 0)
    limit := int32(56) // int32 | Max number of results to return (used for pagination of results) (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStoreFiles(context.Background(), storeId).Sort(sort).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStoreFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoreFiles`: GetStoreFilesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStoreFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | ID of the store containing the files | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | Sort the content of the response according to a given key, potentially in descending order | 
 **offset** | **int32** | How many entries to skip (used for pagination of results) | [default to 0]
 **limit** | **int32** | Max number of results to return (used for pagination of results) | [default to 25]

### Return type

[**GetStoreFilesResponse**](GetStoreFilesResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoreUpload

> GetStoreUploadResponse GetStoreUpload(ctx, uploadId, storeId).Sort(sort).Execute()

Fetch an upload

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uploadId := int32(56) // int32 | ID of the upload to fetch
    storeId := "storeId_example" // string | ID of the store containing the upload
    sort := "name" // string | Sort the content of the response according to a given key, potentially in descending order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStoreUpload(context.Background(), uploadId, storeId).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStoreUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoreUpload`: GetStoreUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStoreUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadId** | **int32** | ID of the upload to fetch | 
**storeId** | **string** | ID of the store containing the upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sort** | **string** | Sort the content of the response according to a given key, potentially in descending order | 

### Return type

[**GetStoreUploadResponse**](GetStoreUploadResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoreUploads

> GetStoreUploadsResponse GetStoreUploads(ctx, storeId).Sort(sort).Offset(offset).Limit(limit).Execute()

Fetch a list of uploads in store

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    storeId := "storeId_example" // string | ID of the store containing the uploads
    sort := "name" // string | Sort the content of the response according to a given key, potentially in descending order (optional)
    offset := int32(56) // int32 | How many entries to skip (used for pagination of results) (optional) (default to 0)
    limit := int32(56) // int32 | Max number of results to return (used for pagination of results) (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStoreUploads(context.Background(), storeId).Sort(sort).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStoreUploads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoreUploads`: GetStoreUploadsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStoreUploads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | ID of the store containing the uploads | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreUploadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | Sort the content of the response according to a given key, potentially in descending order | 
 **offset** | **int32** | How many entries to skip (used for pagination of results) | [default to 0]
 **limit** | **int32** | Max number of results to return (used for pagination of results) | [default to 25]

### Return type

[**GetStoreUploadsResponse**](GetStoreUploadsResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStores

> []string GetStores(ctx).Sort(sort).Execute()

Fetch a list of all stores

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sort := "name" // string | Sort the content of the response according to a given key, potentially in descending order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStores(context.Background()).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStores`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort the content of the response according to a given key, potentially in descending order | 

### Return type

**[]string**

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToken

> GetTokenResponse GetToken(ctx, token).Execute()

Fetch a token for current user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string | ID of the token to fetch

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetToken`: GetTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | ID of the token to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTokenResponse**](GetTokenResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokens

> []GetTokenResponse GetTokens(ctx).Sort(sort).Execute()

Fetch a list of all tokens for current user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sort := "name" // string | Sort the content of the response according to a given key, potentially in descending order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTokens(context.Background()).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokens`: []GetTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort the content of the response according to a given key, potentially in descending order | 

### Return type

[**[]GetTokenResponse**](GetTokenResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkStoreUploadAborted

> MarkStoreUploadAborted(ctx, uploadId, storeId).Execute()

Mark an upload as aborted

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uploadId := int32(56) // int32 | ID of the upload to mark as aborted
    storeId := "storeId_example" // string | ID of the store containing the upload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MarkStoreUploadAborted(context.Background(), uploadId, storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MarkStoreUploadAborted``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadId** | **int32** | ID of the upload to mark as aborted | 
**storeId** | **string** | ID of the store containing the upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkStoreUploadAbortedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkStoreUploadCompleted

> MarkStoreUploadCompleted(ctx, uploadId, storeId).Execute()

Mark an upload as completed

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uploadId := int32(56) // int32 | ID of the upload to fetch
    storeId := "storeId_example" // string | ID of the store containing the upload

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MarkStoreUploadCompleted(context.Background(), uploadId, storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MarkStoreUploadCompleted``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadId** | **int32** | ID of the upload to fetch | 
**storeId** | **string** | ID of the store containing the upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkStoreUploadCompletedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkStoreUploadFileUploaded

> MarkStoreUploadFileUploaded(ctx, uploadId, storeId, fileId).Execute()

Mark a file within an upload as uploaded

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uploadId := int32(56) // int32 | ID of the upload to fetch
    storeId := "storeId_example" // string | ID of the store containing the upload
    fileId := int32(56) // int32 | Index of the file within the upload that should be marked as uploaded

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MarkStoreUploadFileUploaded(context.Background(), uploadId, storeId, fileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MarkStoreUploadFileUploaded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadId** | **int32** | ID of the upload to fetch | 
**storeId** | **string** | ID of the store containing the upload | 
**fileId** | **int32** | Index of the file within the upload that should be marked as uploaded | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkStoreUploadFileUploadedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateToken

> UpdateToken(ctx, token).UpdateTokenRequest(updateTokenRequest).Execute()

Update details of a token for current user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    token := "token_example" // string | ID of the token to update
    updateTokenRequest := *openapiclient.NewUpdateTokenRequest("Used for local testing purposes") // UpdateTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateToken(context.Background(), token).UpdateTokenRequest(updateTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | ID of the token to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTokenRequest** | [**UpdateTokenRequest**](UpdateTokenRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

