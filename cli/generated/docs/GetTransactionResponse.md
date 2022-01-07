# GetTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**BuildId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]GetFileResponse**](GetFileResponse.md) |  | [optional] 

## Methods

### NewGetTransactionResponse

`func NewGetTransactionResponse() *GetTransactionResponse`

NewGetTransactionResponse instantiates a new GetTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionResponseWithDefaults

`func NewGetTransactionResponseWithDefaults() *GetTransactionResponse`

NewGetTransactionResponseWithDefaults instantiates a new GetTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GetTransactionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetTransactionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetTransactionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetTransactionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBuildId

`func (o *GetTransactionResponse) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *GetTransactionResponse) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *GetTransactionResponse) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *GetTransactionResponse) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetTransactionResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetTransactionResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetTransactionResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetTransactionResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetFiles

`func (o *GetTransactionResponse) GetFiles() []GetFileResponse`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GetTransactionResponse) GetFilesOk() (*[]GetFileResponse, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GetTransactionResponse) SetFiles(v []GetFileResponse)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *GetTransactionResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


