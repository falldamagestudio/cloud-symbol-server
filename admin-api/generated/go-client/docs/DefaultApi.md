# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransaction**](DefaultApi.md#CreateTransaction) | **Post** /stores/{storeId}/transactions | Start a new upload transaction
[**GetStores**](DefaultApi.md#GetStores) | **Get** /stores | Fetch a list of all stores
[**GetTransaction**](DefaultApi.md#GetTransaction) | **Get** /stores/{storeId}/transactions/{transactionId} | Fetch a transaction



## CreateTransaction

> UploadTransactionResponse CreateTransaction(ctx, storeId).UploadTransactionRequest(uploadTransactionRequest).Execute()

Start a new upload transaction

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
    storeId := "storeId_example" // string | ID of the store containing the transactions
    uploadTransactionRequest := *openapiclient.NewUploadTransactionRequest() // UploadTransactionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateTransaction(context.Background(), storeId).UploadTransactionRequest(uploadTransactionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransaction`: UploadTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | ID of the store containing the transactions | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uploadTransactionRequest** | [**UploadTransactionRequest**](UploadTransactionRequest.md) |  | 

### Return type

[**UploadTransactionResponse**](UploadTransactionResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStores

> []string GetStores(ctx).Execute()

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
    // response from `GetStores`: []string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoresRequest struct via the builder pattern


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


## GetTransaction

> GetTransactionResponse GetTransaction(ctx, transactionId, storeId).Execute()

Fetch a transaction

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
    transactionId := "transactionId_example" // string | ID of the transaction to fetch
    storeId := "storeId_example" // string | ID of the store containing the transaction

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTransaction(context.Background(), transactionId, storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransaction`: GetTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | ID of the transaction to fetch | 
**storeId** | **string** | ID of the store containing the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetTransactionResponse**](GetTransactionResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
