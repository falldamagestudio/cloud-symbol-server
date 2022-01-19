# BackendAPI.Api.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStore**](DefaultApi.md#createstore) | **POST** /stores/{store} | Create a new store
[**CreateTransaction**](DefaultApi.md#createtransaction) | **POST** /stores/{storeId}/transactions | Start a new upload transaction
[**DeleteStore**](DefaultApi.md#deletestore) | **DELETE** /stores/{store} | Delete an existing store
[**GetStores**](DefaultApi.md#getstores) | **GET** /stores | Fetch a list of all stores
[**GetTransaction**](DefaultApi.md#gettransaction) | **GET** /stores/{storeId}/transactions/{transactionId} | Fetch a transaction


<a name="createstore"></a>
# **CreateStore**
> void CreateStore (string store)

Create a new store

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class CreateStoreExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var store = store_example;  // string | ID of store to create

            try
            {
                // Create a new store
                apiInstance.CreateStore(store);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.CreateStore: " + e.Message );
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
 **store** | **string**| ID of store to create | 

### Return type

void (empty response body)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **409** | Store already exists |  -  |
| **401** | Not authorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="createtransaction"></a>
# **CreateTransaction**
> UploadTransactionResponse CreateTransaction (string storeId, UploadTransactionRequest uploadTransactionRequest)

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
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var storeId = storeId_example;  // string | ID of the store containing the transactions
            var uploadTransactionRequest = new UploadTransactionRequest(); // UploadTransactionRequest | 

            try
            {
                // Start a new upload transaction
                UploadTransactionResponse result = apiInstance.CreateTransaction(storeId, uploadTransactionRequest);
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
 **storeId** | **string**| ID of the store containing the transactions | 
 **uploadTransactionRequest** | [**UploadTransactionRequest**](UploadTransactionRequest.md)|  | 

### Return type

[**UploadTransactionResponse**](UploadTransactionResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **401** | Not authorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="deletestore"></a>
# **DeleteStore**
> void DeleteStore (string store)

Delete an existing store

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class DeleteStoreExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var store = store_example;  // string | ID of store to delete

            try
            {
                // Delete an existing store
                apiInstance.DeleteStore(store);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.DeleteStore: " + e.Message );
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
 **store** | **string**| ID of store to delete | 

### Return type

void (empty response body)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **404** | Store does not exist |  -  |
| **401** | Not authorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getstores"></a>
# **GetStores**
> List&lt;string&gt; GetStores ()

Fetch a list of all stores

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class GetStoresExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);

            try
            {
                // Fetch a list of all stores
                List<string> result = apiInstance.GetStores();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.GetStores: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

**List<string>**

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **401** | Not authorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="gettransaction"></a>
# **GetTransaction**
> GetTransactionResponse GetTransaction (string transactionId, string storeId)

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
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var transactionId = transactionId_example;  // string | ID of the transaction to fetch
            var storeId = storeId_example;  // string | ID of the store containing the transaction

            try
            {
                // Fetch a transaction
                GetTransactionResponse result = apiInstance.GetTransaction(transactionId, storeId);
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
 **storeId** | **string**| ID of the store containing the transaction | 

### Return type

[**GetTransactionResponse**](GetTransactionResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **401** | Not authorized |  -  |
| **404** | No such store/transaction |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

