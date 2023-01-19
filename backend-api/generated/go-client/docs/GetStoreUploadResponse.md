# GetStoreUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadId** | **int32** |  | 
**Description** | **string** |  | 
**BuildId** | **string** |  | 
**Timestamp** | **string** |  | 
**Files** | [**[]GetStoreUploadFileResponse**](GetStoreUploadFileResponse.md) |  | 
**Status** | [**StoreUploadStatus**](StoreUploadStatus.md) |  | 

## Methods

### NewGetStoreUploadResponse

`func NewGetStoreUploadResponse(uploadId int32, description string, buildId string, timestamp string, files []GetStoreUploadFileResponse, status StoreUploadStatus, ) *GetStoreUploadResponse`

NewGetStoreUploadResponse instantiates a new GetStoreUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoreUploadResponseWithDefaults

`func NewGetStoreUploadResponseWithDefaults() *GetStoreUploadResponse`

NewGetStoreUploadResponseWithDefaults instantiates a new GetStoreUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadId

`func (o *GetStoreUploadResponse) GetUploadId() int32`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *GetStoreUploadResponse) GetUploadIdOk() (*int32, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *GetStoreUploadResponse) SetUploadId(v int32)`

SetUploadId sets UploadId field to given value.


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

`func (o *GetStoreUploadResponse) GetFiles() []GetStoreUploadFileResponse`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GetStoreUploadResponse) GetFilesOk() (*[]GetStoreUploadFileResponse, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GetStoreUploadResponse) SetFiles(v []GetStoreUploadFileResponse)`

SetFiles sets Files field to given value.


### GetStatus

`func (o *GetStoreUploadResponse) GetStatus() StoreUploadStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStoreUploadResponse) GetStatusOk() (*StoreUploadStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStoreUploadResponse) SetStatus(v StoreUploadStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


