# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStore**](DefaultApi.md#CreateStore) | **Post** /stores/{storeId} | Create a new store
[**CreateStoreUpload**](DefaultApi.md#CreateStoreUpload) | **Post** /stores/{storeId}/uploads | Start a new upload
[**DeleteStore**](DefaultApi.md#DeleteStore) | **Delete** /stores/{storeId} | Delete an existing store
[**ExpireStoreUpload**](DefaultApi.md#ExpireStoreUpload) | **Post** /stores/{storeId}/uploads/{uploadId}/expire | Expire store upload and consider files for GC
[**GetStoreFileIds**](DefaultApi.md#GetStoreFileIds) | **Get** /stores/{storeId}/files | Fetch a list of all files in store
[**GetStoreUpload**](DefaultApi.md#GetStoreUpload) | **Get** /stores/{storeId}/uploads/{uploadId} | Fetch an upload
[**GetStoreUploadIds**](DefaultApi.md#GetStoreUploadIds) | **Get** /stores/{storeId}/uploads | Fetch a list of all uploads in store
[**GetStores**](DefaultApi.md#GetStores) | **Get** /stores | Fetch a list of all stores
[**MarkStoreUploadAborted**](DefaultApi.md#MarkStoreUploadAborted) | **Post** /stores/{storeId}/uploads/{uploadId}/aborted | Mark an upload as aborted
[**MarkStoreUploadCompleted**](DefaultApi.md#MarkStoreUploadCompleted) | **Post** /stores/{storeId}/uploads/{uploadId}/completed | Mark an upload as completed
[**MarkStoreUploadFileUploaded**](DefaultApi.md#MarkStoreUploadFileUploaded) | **Post** /stores/{storeId}/uploads/{uploadId}/files/{fileId}/uploaded | Mark a file within an upload as uploaded



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
    uploadId := "uploadId_example" // string | ID of the upload to fetch
    storeId := "storeId_example" // string | ID of the store containing the upload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ExpireStoreUpload(context.Background(), uploadId, storeId).Execute()
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
**uploadId** | **string** | ID of the upload to fetch | 
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


## GetStoreFileIds

> GetStoreFileIdsResponse GetStoreFileIds(ctx, storeId).Execute()

Fetch a list of all files in store

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetStoreFileIds(context.Background(), storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStoreFileIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoreFileIds`: GetStoreFileIdsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStoreFileIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | ID of the store containing the files | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreFileIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetStoreFileIdsResponse**](GetStoreFileIdsResponse.md)

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


## GetStoreUploadIds

> GetStoreUploadIdsResponse GetStoreUploadIds(ctx, storeId).Execute()

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
    resp, r, err := api_client.DefaultApi.GetStoreUploadIds(context.Background(), storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStoreUploadIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoreUploadIds`: GetStoreUploadIdsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStoreUploadIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | ID of the store containing the uploads | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreUploadIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetStoreUploadIdsResponse**](GetStoreUploadIdsResponse.md)

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
    uploadId := "uploadId_example" // string | ID of the upload to mark as aborted
    storeId := "storeId_example" // string | ID of the store containing the upload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.MarkStoreUploadAborted(context.Background(), uploadId, storeId).Execute()
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
**uploadId** | **string** | ID of the upload to mark as aborted | 
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
    uploadId := "uploadId_example" // string | ID of the upload to fetch
    storeId := "storeId_example" // string | ID of the store containing the upload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.MarkStoreUploadCompleted(context.Background(), uploadId, storeId).Execute()
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
**uploadId** | **string** | ID of the upload to fetch | 
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
    uploadId := "uploadId_example" // string | ID of the upload to fetch
    storeId := "storeId_example" // string | ID of the store containing the upload
    fileId := int32(56) // int32 | Index of the file within the upload that should be marked as uploaded

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.MarkStoreUploadFileUploaded(context.Background(), uploadId, storeId, fileId).Execute()
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
**uploadId** | **string** | ID of the upload to fetch | 
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

