# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStore**](DefaultApi.md#CreateStore) | **Post** /stores/{storeId} | Create a new store
[**CreateStoreUpload**](DefaultApi.md#CreateStoreUpload) | **Post** /stores/{storeId}/uploads | Start a new upload
[**DeleteStore**](DefaultApi.md#DeleteStore) | **Delete** /stores/{storeId} | Delete an existing store
[**GetStoreUpload**](DefaultApi.md#GetStoreUpload) | **Get** /stores/{storeId}/uploads/{uploadId} | Fetch an upload
[**GetStoreUploads**](DefaultApi.md#GetStoreUploads) | **Get** /stores/{storeId}/uploads | Fetch a list of all uploads in store
[**GetStores**](DefaultApi.md#GetStores) | **Get** /stores | Fetch a list of all stores



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateStore(context.Background(), storeId).Execute()
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateStoreUpload(context.Background(), storeId).CreateStoreUploadRequest(createStoreUploadRequest).Execute()
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteStore(context.Background(), storeId).Execute()
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


## GetStoreUpload

> GetStoreUploadResponse GetStoreUpload(ctx, uploadId, storeId).Execute()

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
    uploadId := "uploadId_example" // string | ID of the upload to fetch
    storeId := "storeId_example" // string | ID of the store containing the upload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetStoreUpload(context.Background(), uploadId, storeId).Execute()
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
**uploadId** | **string** | ID of the upload to fetch | 
**storeId** | **string** | ID of the store containing the upload | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

> GetStoreUploadsResponse GetStoreUploads(ctx, storeId).Execute()

Fetch a list of all uploads in store

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetStoreUploads(context.Background(), storeId).Execute()
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

> GetStoresResponse GetStores(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetStores(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStores`: GetStoresResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoresRequest struct via the builder pattern


### Return type

[**GetStoresResponse**](GetStoresResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

