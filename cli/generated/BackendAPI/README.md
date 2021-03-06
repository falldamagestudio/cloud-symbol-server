# BackendAPI - the C# library for the Cloud Symbol Server Admin API

This is the API that is used to manage stores and uploads in Cloud Symbol Server

This C# SDK is automatically generated by the [OpenAPI Generator](https://openapi-generator.tech) project:

- API version: 1.0.0
- SDK version: 1.0.0
- Build package: org.openapitools.codegen.languages.CSharpNetCoreClientCodegen

<a name="frameworks-supported"></a>
## Frameworks supported
- .NET Core >=1.0
- .NET Framework >=4.6
- Mono/Xamarin >=vNext

<a name="dependencies"></a>
## Dependencies

- [RestSharp](https://www.nuget.org/packages/RestSharp) - 106.11.7 or later
- [Json.NET](https://www.nuget.org/packages/Newtonsoft.Json/) - 12.0.3 or later
- [JsonSubTypes](https://www.nuget.org/packages/JsonSubTypes/) - 1.8.0 or later
- [System.ComponentModel.Annotations](https://www.nuget.org/packages/System.ComponentModel.Annotations) - 5.0.0 or later

The DLLs included in the package may not be the latest version. We recommend using [NuGet](https://docs.nuget.org/consume/installing-nuget) to obtain the latest version of the packages:
```
Install-Package RestSharp
Install-Package Newtonsoft.Json
Install-Package JsonSubTypes
Install-Package System.ComponentModel.Annotations
```

NOTE: RestSharp versions greater than 105.1.0 have a bug which causes file uploads to fail. See [RestSharp#742](https://github.com/restsharp/RestSharp/issues/742).
NOTE: RestSharp for .Net Core creates a new socket for each api call, which can lead to a socket exhaustion problem. See [RestSharp#1406](https://github.com/restsharp/RestSharp/issues/1406).

<a name="installation"></a>
## Installation
Generate the DLL using your preferred tool (e.g. `dotnet build`)

Then include the DLL (under the `bin` folder) in the C# project, and use the namespaces:
```csharp
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;
```
<a name="usage"></a>
## Usage

To use the API client with a HTTP proxy, setup a `System.Net.WebProxy`
```csharp
Configuration c = new Configuration();
System.Net.WebProxy webProxy = new System.Net.WebProxy("http://myProxyUrl:80/");
webProxy.Credentials = System.Net.CredentialCache.DefaultCredentials;
c.Proxy = webProxy;
```

<a name="getting-started"></a>
## Getting Started

```csharp
using System.Collections.Generic;
using System.Diagnostics;
using BackendAPI.Api;
using BackendAPI.Client;
using BackendAPI.Model;

namespace Example
{
    public class Example
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
            catch (ApiException e)
            {
                Debug.Print("Exception when calling DefaultApi.CreateStore: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }

        }
    }
}
```

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CreateStore**](docs/DefaultApi.md#createstore) | **POST** /stores/{storeId} | Create a new store
*DefaultApi* | [**CreateStoreUpload**](docs/DefaultApi.md#createstoreupload) | **POST** /stores/{storeId}/uploads | Start a new upload
*DefaultApi* | [**DeleteStore**](docs/DefaultApi.md#deletestore) | **DELETE** /stores/{storeId} | Delete an existing store
*DefaultApi* | [**GetStoreUpload**](docs/DefaultApi.md#getstoreupload) | **GET** /stores/{storeId}/uploads/{uploadId} | Fetch an upload
*DefaultApi* | [**GetStoreUploadIds**](docs/DefaultApi.md#getstoreuploadids) | **GET** /stores/{storeId}/uploads | Fetch a list of all uploads in store
*DefaultApi* | [**GetStores**](docs/DefaultApi.md#getstores) | **GET** /stores | Fetch a list of all stores


<a name="documentation-for-models"></a>
## Documentation for Models

 - [Model.CreateStoreUploadRequest](docs/CreateStoreUploadRequest.md)
 - [Model.CreateStoreUploadResponse](docs/CreateStoreUploadResponse.md)
 - [Model.GetFileResponse](docs/GetFileResponse.md)
 - [Model.GetStoreUploadIdsResponse](docs/GetStoreUploadIdsResponse.md)
 - [Model.GetStoreUploadResponse](docs/GetStoreUploadResponse.md)
 - [Model.GetStoresResponse](docs/GetStoresResponse.md)
 - [Model.MessageResponse](docs/MessageResponse.md)
 - [Model.UploadFileRequest](docs/UploadFileRequest.md)
 - [Model.UploadFileResponse](docs/UploadFileResponse.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

<a name="emailAndPat"></a>
### emailAndPat

- **Type**: HTTP basic authentication

