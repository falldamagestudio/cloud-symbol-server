# CreateStoreUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Files** | [**[]UploadFileResponse**](UploadFileResponse.md) |  | 

## Methods

### NewCreateStoreUploadResponse

`func NewCreateStoreUploadResponse(id string, files []UploadFileResponse, ) *CreateStoreUploadResponse`

NewCreateStoreUploadResponse instantiates a new CreateStoreUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStoreUploadResponseWithDefaults

`func NewCreateStoreUploadResponseWithDefaults() *CreateStoreUploadResponse`

NewCreateStoreUploadResponseWithDefaults instantiates a new CreateStoreUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateStoreUploadResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateStoreUploadResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateStoreUploadResponse) SetId(v string)`

SetId sets Id field to given value.


### GetFiles

`func (o *CreateStoreUploadResponse) GetFiles() []UploadFileResponse`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CreateStoreUploadResponse) GetFilesOk() (*[]UploadFileResponse, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CreateStoreUploadResponse) SetFiles(v []UploadFileResponse)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


