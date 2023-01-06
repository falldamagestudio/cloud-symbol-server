# GetStoreUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**BuildId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]GetFileResponse**](GetFileResponse.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetStoreUploadResponse

`func NewGetStoreUploadResponse() *GetStoreUploadResponse`

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

### HasDescription

`func (o *GetStoreUploadResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

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

### HasBuildId

`func (o *GetStoreUploadResponse) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

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

### HasTimestamp

`func (o *GetStoreUploadResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

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

### HasFiles

`func (o *GetStoreUploadResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

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

### HasStatus

`func (o *GetStoreUploadResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


