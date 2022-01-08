# BackendAPI.Api.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransaction**](DefaultApi.md#createtransaction) | **POST** /transactions | Start a new upload transaction
[**GetTransaction**](DefaultApi.md#gettransaction) | **GET** /transactions/{transactionId} | Fetch a transaction


<a name="createtransaction"></a>
# **CreateTransaction**
> UploadTransactionResponse CreateTransaction (UploadTransactionRequest uploadTransactionRequest)

Start a new upload transaction

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class CreateTransactionExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            var apiInstance = new DefaultApi(config);
            var uploadTransactionRequest = new UploadTransactionRequest(); // UploadTransactionRequest | 

            try
            {
                // Start a new upload transaction
                UploadTransactionResponse result = apiInstance.CreateTransaction(uploadTransactionRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.CreateTransaction: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadTransactionRequest** | [**UploadTransactionRequest**](UploadTransactionRequest.md)|  | 

### Return type

[**UploadTransactionResponse**](UploadTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **401** | Not authorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="gettransaction"></a>
# **GetTransaction**
> GetTransactionResponse GetTransaction (string transactionId)

Fetch a transaction

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class GetTransactionExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            var apiInstance = new DefaultApi(config);
            var transactionId = transactionId_example;  // string | ID of the transaction to fetch

            try
            {
                // Fetch a transaction
                GetTransactionResponse result = apiInstance.GetTransaction(transactionId);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.GetTransaction: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **string**| ID of the transaction to fetch | 

### Return type

[**GetTransactionResponse**](GetTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **401** | Not authorized |  -  |
| **404** | No such transaction |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

