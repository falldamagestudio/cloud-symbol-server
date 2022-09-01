# BackendAPI.Api.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStore**](DefaultApi.md#createstore) | **POST** /stores/{storeId} | Create a new store
[**CreateStoreUpload**](DefaultApi.md#createstoreupload) | **POST** /stores/{storeId}/uploads | Start a new upload
[**DeleteStore**](DefaultApi.md#deletestore) | **DELETE** /stores/{storeId} | Delete an existing store
[**ExpireStoreUpload**](DefaultApi.md#expirestoreupload) | **POST** /stores/{storeId}/uploads/{uploadId}/expire | Expire store upload and consider files for GC
[**GetStoreFileIds**](DefaultApi.md#getstorefileids) | **GET** /stores/{storeId}/files | Fetch a list of all files in store
[**GetStoreUpload**](DefaultApi.md#getstoreupload) | **GET** /stores/{storeId}/uploads/{uploadId} | Fetch an upload
[**GetStoreUploadIds**](DefaultApi.md#getstoreuploadids) | **GET** /stores/{storeId}/uploads | Fetch a list of all uploads in store
[**GetStores**](DefaultApi.md#getstores) | **GET** /stores | Fetch a list of all stores
[**MarkStoreUploadAborted**](DefaultApi.md#markstoreuploadaborted) | **POST** /stores/{storeId}/uploads/{uploadId}/aborted | Mark an upload as aborted
[**MarkStoreUploadCompleted**](DefaultApi.md#markstoreuploadcompleted) | **POST** /stores/{storeId}/uploads/{uploadId}/completed | Mark an upload as completed
[**MarkStoreUploadFileUploaded**](DefaultApi.md#markstoreuploadfileuploaded) | **POST** /stores/{storeId}/uploads/{uploadId}/files/{fileId}/uploaded | Mark a file within an upload as uploaded


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
            var storeId = storeId_example;  // string | ID of store to create

            try
            {
                // Create a new store
                apiInstance.CreateStore(storeId);
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
 **storeId** | **string**| ID of store to create | 

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
            var storeId = storeId_example;  // string | ID of the store containing the upload
            var createStoreUploadRequest = new CreateStoreUploadRequest(); // CreateStoreUploadRequest | 

            try
            {
                // Start a new upload
                CreateStoreUploadResponse result = apiInstance.CreateStoreUpload(storeId, createStoreUploadRequest);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.CreateStoreUpload: " + e.Message );
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
 **storeId** | **string**| ID of the store containing the upload | 
 **createStoreUploadRequest** | [**CreateStoreUploadRequest**](CreateStoreUploadRequest.md)|  | 

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
            var storeId = storeId_example;  // string | ID of store to delete

            try
            {
                // Delete an existing store
                apiInstance.DeleteStore(storeId);
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
 **storeId** | **string**| ID of store to delete | 

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
            var uploadId = uploadId_example;  // string | ID of the upload to fetch
            var storeId = storeId_example;  // string | ID of the store containing the upload

            try
            {
                // Expire store upload and consider files for GC
                apiInstance.ExpireStoreUpload(uploadId, storeId);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.ExpireStoreUpload: " + e.Message );
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
 **uploadId** | **string**| ID of the upload to fetch | 
 **storeId** | **string**| ID of the store containing the upload | 

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
> GetStoreFileIdsResponse GetStoreFileIds (string storeId)

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
            var storeId = storeId_example;  // string | ID of the store containing the files

            try
            {
                // Fetch a list of all files in store
                GetStoreFileIdsResponse result = apiInstance.GetStoreFileIds(storeId);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.GetStoreFileIds: " + e.Message );
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
 **storeId** | **string**| ID of the store containing the files | 

### Return type

[**GetStoreFileIdsResponse**](GetStoreFileIdsResponse.md)

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
            var uploadId = uploadId_example;  // string | ID of the upload to fetch
            var storeId = storeId_example;  // string | ID of the store containing the upload

            try
            {
                // Fetch an upload
                GetStoreUploadResponse result = apiInstance.GetStoreUpload(uploadId, storeId);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.GetStoreUpload: " + e.Message );
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
 **uploadId** | **string**| ID of the upload to fetch | 
 **storeId** | **string**| ID of the store containing the upload | 

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
> GetStoreUploadIdsResponse GetStoreUploadIds (string storeId)

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
            var storeId = storeId_example;  // string | ID of the store containing the uploads

            try
            {
                // Fetch a list of all uploads in store
                GetStoreUploadIdsResponse result = apiInstance.GetStoreUploadIds(storeId);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.GetStoreUploadIds: " + e.Message );
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
 **storeId** | **string**| ID of the store containing the uploads | 

### Return type

[**GetStoreUploadIdsResponse**](GetStoreUploadIdsResponse.md)

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
> GetStoresResponse GetStores ()

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
                GetStoresResponse result = apiInstance.GetStores();
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

[**GetStoresResponse**](GetStoresResponse.md)

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
            var uploadId = uploadId_example;  // string | ID of the upload to mark as aborted
            var storeId = storeId_example;  // string | ID of the store containing the upload

            try
            {
                // Mark an upload as aborted
                apiInstance.MarkStoreUploadAborted(uploadId, storeId);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.MarkStoreUploadAborted: " + e.Message );
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
 **uploadId** | **string**| ID of the upload to mark as aborted | 
 **storeId** | **string**| ID of the store containing the upload | 

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
            var uploadId = uploadId_example;  // string | ID of the upload to fetch
            var storeId = storeId_example;  // string | ID of the store containing the upload

            try
            {
                // Mark an upload as completed
                apiInstance.MarkStoreUploadCompleted(uploadId, storeId);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.MarkStoreUploadCompleted: " + e.Message );
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
 **uploadId** | **string**| ID of the upload to fetch | 
 **storeId** | **string**| ID of the store containing the upload | 

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
            var uploadId = uploadId_example;  // string | ID of the upload to fetch
            var storeId = storeId_example;  // string | ID of the store containing the upload
            var fileId = 56;  // int | Index of the file within the upload that should be marked as uploaded

            try
            {
                // Mark a file within an upload as uploaded
                apiInstance.MarkStoreUploadFileUploaded(uploadId, storeId, fileId);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DefaultApi.MarkStoreUploadFileUploaded: " + e.Message );
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
 **uploadId** | **string**| ID of the upload to fetch | 
 **storeId** | **string**| ID of the store containing the upload | 
 **fileId** | **int**| Index of the file within the upload that should be marked as uploaded | 

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

