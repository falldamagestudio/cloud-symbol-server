# BackendAPI.Api.DefaultApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|--------|--------------|-------------|
| [**CreateStore**](DefaultApi.md#createstore) | **POST** /stores/{storeId} | Create a new store |
| [**CreateStoreUpload**](DefaultApi.md#createstoreupload) | **POST** /stores/{storeId}/uploads | Start a new upload |
| [**CreateToken**](DefaultApi.md#createtoken) | **POST** /tokens | Create a new token for current user |
| [**DeleteStore**](DefaultApi.md#deletestore) | **DELETE** /stores/{storeId} | Delete an existing store |
| [**DeleteToken**](DefaultApi.md#deletetoken) | **DELETE** /tokens/{token} | Delete a token for current user |
| [**ExpireStoreUpload**](DefaultApi.md#expirestoreupload) | **POST** /stores/{storeId}/uploads/{uploadId}/expire | Expire store upload and consider files for GC |
| [**GetStoreFileIds**](DefaultApi.md#getstorefileids) | **GET** /stores/{storeId}/files | Fetch a list of all files in store |
| [**GetStoreUpload**](DefaultApi.md#getstoreupload) | **GET** /stores/{storeId}/uploads/{uploadId} | Fetch an upload |
| [**GetStoreUploadIds**](DefaultApi.md#getstoreuploadids) | **GET** /stores/{storeId}/uploads | Fetch a list of all uploads in store |
| [**GetStores**](DefaultApi.md#getstores) | **GET** /stores | Fetch a list of all stores |
| [**GetToken**](DefaultApi.md#gettoken) | **GET** /tokens/{token} | Fetch a token for current user |
| [**GetTokens**](DefaultApi.md#gettokens) | **GET** /tokens | Fetch a list of all tokens for current user |
| [**MarkStoreUploadAborted**](DefaultApi.md#markstoreuploadaborted) | **POST** /stores/{storeId}/uploads/{uploadId}/aborted | Mark an upload as aborted |
| [**MarkStoreUploadCompleted**](DefaultApi.md#markstoreuploadcompleted) | **POST** /stores/{storeId}/uploads/{uploadId}/completed | Mark an upload as completed |
| [**MarkStoreUploadFileUploaded**](DefaultApi.md#markstoreuploadfileuploaded) | **POST** /stores/{storeId}/uploads/{uploadId}/files/{fileId}/uploaded | Mark a file within an upload as uploaded |
| [**UpdateToken**](DefaultApi.md#updatetoken) | **PUT** /tokens/{token} | Update details of a token for current user |

<a name="createstore"></a>
# **CreateStore**
> void CreateStore (string storeId)

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
            var storeId = "storeId_example";  // string | ID of store to create

            try
            {
                // Create a new store
                apiInstance.CreateStore(storeId);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.CreateStore: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the CreateStoreWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Create a new store
    apiInstance.CreateStoreWithHttpInfo(storeId);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.CreateStoreWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **storeId** | **string** | ID of store to create |  |

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
| **401** | Not authorized |  -  |
| **409** | Store already exists |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="createstoreupload"></a>
# **CreateStoreUpload**
> CreateStoreUploadResponse CreateStoreUpload (string storeId, CreateStoreUploadRequest createStoreUploadRequest)

Start a new upload

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class CreateStoreUploadExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var storeId = "storeId_example";  // string | ID of the store containing the upload
            var createStoreUploadRequest = new CreateStoreUploadRequest(); // CreateStoreUploadRequest | 

            try
            {
                // Start a new upload
                CreateStoreUploadResponse result = apiInstance.CreateStoreUpload(storeId, createStoreUploadRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.CreateStoreUpload: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the CreateStoreUploadWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Start a new upload
    ApiResponse<CreateStoreUploadResponse> response = apiInstance.CreateStoreUploadWithHttpInfo(storeId, createStoreUploadRequest);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.CreateStoreUploadWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **storeId** | **string** | ID of the store containing the upload |  |
| **createStoreUploadRequest** | [**CreateStoreUploadRequest**](CreateStoreUploadRequest.md) |  |  |

### Return type

[**CreateStoreUploadResponse**](CreateStoreUploadResponse.md)

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

<a name="createtoken"></a>
# **CreateToken**
> CreateTokenResponse CreateToken ()

Create a new token for current user

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class CreateTokenExample
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
                // Create a new token for current user
                CreateTokenResponse result = apiInstance.CreateToken();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.CreateToken: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the CreateTokenWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Create a new token for current user
    ApiResponse<CreateTokenResponse> response = apiInstance.CreateTokenWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.CreateTokenWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters
This endpoint does not need any parameter.
### Return type

[**CreateTokenResponse**](CreateTokenResponse.md)

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

<a name="deletestore"></a>
# **DeleteStore**
> void DeleteStore (string storeId)

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
            var storeId = "storeId_example";  // string | ID of store to delete

            try
            {
                // Delete an existing store
                apiInstance.DeleteStore(storeId);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.DeleteStore: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the DeleteStoreWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Delete an existing store
    apiInstance.DeleteStoreWithHttpInfo(storeId);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.DeleteStoreWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **storeId** | **string** | ID of store to delete |  |

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
| **401** | Not authorized |  -  |
| **404** | Store does not exist |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="deletetoken"></a>
# **DeleteToken**
> void DeleteToken (string token)

Delete a token for current user

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class DeleteTokenExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var token = "token_example";  // string | ID of the token to delete

            try
            {
                // Delete a token for current user
                apiInstance.DeleteToken(token);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.DeleteToken: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the DeleteTokenWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Delete a token for current user
    apiInstance.DeleteTokenWithHttpInfo(token);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.DeleteTokenWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **token** | **string** | ID of the token to delete |  |

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
| **404** | Token does not exist |  -  |
| **401** | Not authorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="expirestoreupload"></a>
# **ExpireStoreUpload**
> void ExpireStoreUpload (string uploadId, string storeId)

Expire store upload and consider files for GC

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class ExpireStoreUploadExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var uploadId = "uploadId_example";  // string | ID of the upload to fetch
            var storeId = "storeId_example";  // string | ID of the store containing the upload

            try
            {
                // Expire store upload and consider files for GC
                apiInstance.ExpireStoreUpload(uploadId, storeId);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.ExpireStoreUpload: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the ExpireStoreUploadWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Expire store upload and consider files for GC
    apiInstance.ExpireStoreUploadWithHttpInfo(uploadId, storeId);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.ExpireStoreUploadWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **uploadId** | **string** | ID of the upload to fetch |  |
| **storeId** | **string** | ID of the store containing the upload |  |

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
| **401** | Not authorized |  -  |
| **404** | No such store/upload |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getstorefileids"></a>
# **GetStoreFileIds**
> List&lt;string&gt; GetStoreFileIds (string storeId)

Fetch a list of all files in store

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class GetStoreFileIdsExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var storeId = "storeId_example";  // string | ID of the store containing the files

            try
            {
                // Fetch a list of all files in store
                List<string> result = apiInstance.GetStoreFileIds(storeId);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.GetStoreFileIds: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the GetStoreFileIdsWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Fetch a list of all files in store
    ApiResponse<List<string>> response = apiInstance.GetStoreFileIdsWithHttpInfo(storeId);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.GetStoreFileIdsWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **storeId** | **string** | ID of the store containing the files |  |

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
| **404** | No such store |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getstoreupload"></a>
# **GetStoreUpload**
> GetStoreUploadResponse GetStoreUpload (string uploadId, string storeId)

Fetch an upload

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class GetStoreUploadExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var uploadId = "uploadId_example";  // string | ID of the upload to fetch
            var storeId = "storeId_example";  // string | ID of the store containing the upload

            try
            {
                // Fetch an upload
                GetStoreUploadResponse result = apiInstance.GetStoreUpload(uploadId, storeId);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.GetStoreUpload: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the GetStoreUploadWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Fetch an upload
    ApiResponse<GetStoreUploadResponse> response = apiInstance.GetStoreUploadWithHttpInfo(uploadId, storeId);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.GetStoreUploadWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **uploadId** | **string** | ID of the upload to fetch |  |
| **storeId** | **string** | ID of the store containing the upload |  |

### Return type

[**GetStoreUploadResponse**](GetStoreUploadResponse.md)

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
| **404** | No such store/upload |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getstoreuploadids"></a>
# **GetStoreUploadIds**
> List&lt;string&gt; GetStoreUploadIds (string storeId)

Fetch a list of all uploads in store

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class GetStoreUploadIdsExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var storeId = "storeId_example";  // string | ID of the store containing the uploads

            try
            {
                // Fetch a list of all uploads in store
                List<string> result = apiInstance.GetStoreUploadIds(storeId);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.GetStoreUploadIds: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the GetStoreUploadIdsWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Fetch a list of all uploads in store
    ApiResponse<List<string>> response = apiInstance.GetStoreUploadIdsWithHttpInfo(storeId);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.GetStoreUploadIdsWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **storeId** | **string** | ID of the store containing the uploads |  |

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
| **404** | No such store |  -  |

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
                Debug.Print("Exception when calling DefaultApi.GetStores: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the GetStoresWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Fetch a list of all stores
    ApiResponse<List<string>> response = apiInstance.GetStoresWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.GetStoresWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
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

<a name="gettoken"></a>
# **GetToken**
> GetTokenResponse GetToken (string token)

Fetch a token for current user

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class GetTokenExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var token = "token_example";  // string | ID of the token to fetch

            try
            {
                // Fetch a token for current user
                GetTokenResponse result = apiInstance.GetToken(token);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.GetToken: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the GetTokenWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Fetch a token for current user
    ApiResponse<GetTokenResponse> response = apiInstance.GetTokenWithHttpInfo(token);
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.GetTokenWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **token** | **string** | ID of the token to fetch |  |

### Return type

[**GetTokenResponse**](GetTokenResponse.md)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **404** | Token does not exist |  -  |
| **401** | Not authorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="gettokens"></a>
# **GetTokens**
> List&lt;GetTokenResponse&gt; GetTokens ()

Fetch a list of all tokens for current user

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class GetTokensExample
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
                // Fetch a list of all tokens for current user
                List<GetTokenResponse> result = apiInstance.GetTokens();
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.GetTokens: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the GetTokensWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Fetch a list of all tokens for current user
    ApiResponse<List<GetTokenResponse>> response = apiInstance.GetTokensWithHttpInfo();
    Debug.Write("Status Code: " + response.StatusCode);
    Debug.Write("Response Headers: " + response.Headers);
    Debug.Write("Response Body: " + response.Data);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.GetTokensWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters
This endpoint does not need any parameter.
### Return type

[**List&lt;GetTokenResponse&gt;**](GetTokenResponse.md)

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

<a name="markstoreuploadaborted"></a>
# **MarkStoreUploadAborted**
> void MarkStoreUploadAborted (string uploadId, string storeId)

Mark an upload as aborted

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class MarkStoreUploadAbortedExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var uploadId = "uploadId_example";  // string | ID of the upload to mark as aborted
            var storeId = "storeId_example";  // string | ID of the store containing the upload

            try
            {
                // Mark an upload as aborted
                apiInstance.MarkStoreUploadAborted(uploadId, storeId);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.MarkStoreUploadAborted: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the MarkStoreUploadAbortedWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Mark an upload as aborted
    apiInstance.MarkStoreUploadAbortedWithHttpInfo(uploadId, storeId);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.MarkStoreUploadAbortedWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **uploadId** | **string** | ID of the upload to mark as aborted |  |
| **storeId** | **string** | ID of the store containing the upload |  |

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
| **401** | Not authorized |  -  |
| **404** | No such store/upload |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="markstoreuploadcompleted"></a>
# **MarkStoreUploadCompleted**
> void MarkStoreUploadCompleted (string uploadId, string storeId)

Mark an upload as completed

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class MarkStoreUploadCompletedExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var uploadId = "uploadId_example";  // string | ID of the upload to fetch
            var storeId = "storeId_example";  // string | ID of the store containing the upload

            try
            {
                // Mark an upload as completed
                apiInstance.MarkStoreUploadCompleted(uploadId, storeId);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.MarkStoreUploadCompleted: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the MarkStoreUploadCompletedWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Mark an upload as completed
    apiInstance.MarkStoreUploadCompletedWithHttpInfo(uploadId, storeId);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.MarkStoreUploadCompletedWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **uploadId** | **string** | ID of the upload to fetch |  |
| **storeId** | **string** | ID of the store containing the upload |  |

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
| **401** | Not authorized |  -  |
| **404** | No such store/upload |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="markstoreuploadfileuploaded"></a>
# **MarkStoreUploadFileUploaded**
> void MarkStoreUploadFileUploaded (string uploadId, string storeId, int fileId)

Mark a file within an upload as uploaded

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class MarkStoreUploadFileUploadedExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var uploadId = "uploadId_example";  // string | ID of the upload to fetch
            var storeId = "storeId_example";  // string | ID of the store containing the upload
            var fileId = 56;  // int | Index of the file within the upload that should be marked as uploaded

            try
            {
                // Mark a file within an upload as uploaded
                apiInstance.MarkStoreUploadFileUploaded(uploadId, storeId, fileId);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.MarkStoreUploadFileUploaded: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the MarkStoreUploadFileUploadedWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Mark a file within an upload as uploaded
    apiInstance.MarkStoreUploadFileUploadedWithHttpInfo(uploadId, storeId, fileId);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.MarkStoreUploadFileUploadedWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **uploadId** | **string** | ID of the upload to fetch |  |
| **storeId** | **string** | ID of the store containing the upload |  |
| **fileId** | **int** | Index of the file within the upload that should be marked as uploaded |  |

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
| **401** | Not authorized |  -  |
| **404** | No such store/upload/item |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="updatetoken"></a>
# **UpdateToken**
> void UpdateToken (string token, UpdateTokenRequest updateTokenRequest)

Update details of a token for current user

### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class UpdateTokenExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "http://localhost";
            // Configure HTTP basic authorization: emailAndPat
            config.Username = "YOUR_USERNAME";
            config.Password = "YOUR_PASSWORD";

            var apiInstance = new DefaultApi(config);
            var token = "token_example";  // string | ID of the token to update
            var updateTokenRequest = new UpdateTokenRequest(); // UpdateTokenRequest | 

            try
            {
                // Update details of a token for current user
                apiInstance.UpdateToken(token, updateTokenRequest);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.UpdateToken: " + e.Message);
                Debug.Print("Status Code: " + e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

#### Using the UpdateTokenWithHttpInfo variant
This returns an ApiResponse object which contains the response data, status code and headers.

```csharp
try
{
    // Update details of a token for current user
    apiInstance.UpdateTokenWithHttpInfo(token, updateTokenRequest);
}
catch (ApiException e)
{
    Debug.Print("Exception when calling DefaultApi.UpdateTokenWithHttpInfo: " + e.Message);
    Debug.Print("Status Code: " + e.ErrorCode);
    Debug.Print(e.StackTrace);
}
```

### Parameters

| Name | Type | Description | Notes |
|------|------|-------------|-------|
| **token** | **string** | ID of the token to update |  |
| **updateTokenRequest** | [**UpdateTokenRequest**](UpdateTokenRequest.md) |  |  |

### Return type

void (empty response body)

### Authorization

[emailAndPat](../README.md#emailAndPat)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **404** | Token does not exist |  -  |
| **401** | Not authorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

