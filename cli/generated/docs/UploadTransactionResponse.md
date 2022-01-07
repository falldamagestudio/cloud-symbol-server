# UploadTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]UploadFileResponse**](UploadFileResponse.md) |  | [optional] 

## Methods

### NewUploadTransactionResponse

`func NewUploadTransactionResponse() *UploadTransactionResponse`

NewUploadTransactionResponse instantiates a new UploadTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadTransactionResponseWithDefaults

`func NewUploadTransactionResponseWithDefaults() *UploadTransactionResponse`

NewUploadTransactionResponseWithDefaults instantiates a new UploadTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UploadTransactionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UploadTransactionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UploadTransactionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UploadTransactionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFiles

`func (o *UploadTransactionResponse) GetFiles() []UploadFileResponse`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *UploadTransactionResponse) GetFilesOk() (*[]UploadFileResponse, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *UploadTransactionResponse) SetFiles(v []UploadFileResponse)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *UploadTransactionResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


