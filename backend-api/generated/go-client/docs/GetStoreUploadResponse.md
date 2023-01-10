# GetStoreUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**BuildId** | **string** |  | 
**Timestamp** | **string** |  | 
**Files** | [**[]GetFileResponse**](GetFileResponse.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetStoreUploadResponse

`func NewGetStoreUploadResponse(description string, buildId string, timestamp string, files []GetFileResponse, status string, ) *GetStoreUploadResponse`

NewGetStoreUploadResponse instantiates a new GetStoreUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoreUploadResponseWithDefaults

`func NewGetStoreUploadResponseWithDefaults() *GetStoreUploadResponse`

NewGetStoreUploadResponseWithDefaults instantiates a new GetStoreUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GetStoreUploadResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetStoreUploadResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetStoreUploadResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetBuildId

`func (o *GetStoreUploadResponse) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *GetStoreUploadResponse) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *GetStoreUploadResponse) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.


### GetTimestamp

`func (o *GetStoreUploadResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetStoreUploadResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetStoreUploadResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetFiles

`func (o *GetStoreUploadResponse) GetFiles() []GetFileResponse`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GetStoreUploadResponse) GetFilesOk() (*[]GetFileResponse, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GetStoreUploadResponse) SetFiles(v []GetFileResponse)`

SetFiles sets Files field to given value.


### GetStatus

`func (o *GetStoreUploadResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStoreUploadResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStoreUploadResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


