# UploadTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**BuildId** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]UploadFileRequest**](UploadFileRequest.md) |  | [optional] 

## Methods

### NewUploadTransactionRequest

`func NewUploadTransactionRequest() *UploadTransactionRequest`

NewUploadTransactionRequest instantiates a new UploadTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadTransactionRequestWithDefaults

`func NewUploadTransactionRequestWithDefaults() *UploadTransactionRequest`

NewUploadTransactionRequestWithDefaults instantiates a new UploadTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UploadTransactionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UploadTransactionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UploadTransactionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UploadTransactionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBuildId

`func (o *UploadTransactionRequest) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *UploadTransactionRequest) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *UploadTransactionRequest) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *UploadTransactionRequest) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### GetFiles

`func (o *UploadTransactionRequest) GetFiles() []UploadFileRequest`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *UploadTransactionRequest) GetFilesOk() (*[]UploadFileRequest, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *UploadTransactionRequest) SetFiles(v []UploadFileRequest)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *UploadTransactionRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


