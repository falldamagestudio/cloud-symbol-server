# CreateStoreUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseProgressApi** | Pointer to **bool** | When present and set to true, the client will provide progress updates; Legacy clients will create an upload, then upload the required files to GCS, without progress/completion callbacks | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**BuildId** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]UploadFileRequest**](UploadFileRequest.md) |  | [optional] 

## Methods

### NewCreateStoreUploadRequest

`func NewCreateStoreUploadRequest() *CreateStoreUploadRequest`

NewCreateStoreUploadRequest instantiates a new CreateStoreUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStoreUploadRequestWithDefaults

`func NewCreateStoreUploadRequestWithDefaults() *CreateStoreUploadRequest`

NewCreateStoreUploadRequestWithDefaults instantiates a new CreateStoreUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseProgressApi

`func (o *CreateStoreUploadRequest) GetUseProgressApi() bool`

GetUseProgressApi returns the UseProgressApi field if non-nil, zero value otherwise.

### GetUseProgressApiOk

`func (o *CreateStoreUploadRequest) GetUseProgressApiOk() (*bool, bool)`

GetUseProgressApiOk returns a tuple with the UseProgressApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProgressApi

`func (o *CreateStoreUploadRequest) SetUseProgressApi(v bool)`

SetUseProgressApi sets UseProgressApi field to given value.

### HasUseProgressApi

`func (o *CreateStoreUploadRequest) HasUseProgressApi() bool`

HasUseProgressApi returns a boolean if a field has been set.

### GetDescription

`func (o *CreateStoreUploadRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStoreUploadRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStoreUploadRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateStoreUploadRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBuildId

`func (o *CreateStoreUploadRequest) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *CreateStoreUploadRequest) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *CreateStoreUploadRequest) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *CreateStoreUploadRequest) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### GetFiles

`func (o *CreateStoreUploadRequest) GetFiles() []UploadFileRequest`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CreateStoreUploadRequest) GetFilesOk() (*[]UploadFileRequest, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CreateStoreUploadRequest) SetFiles(v []UploadFileRequest)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CreateStoreUploadRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


