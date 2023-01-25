# CreateStoreUploadFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** |  | 
**BlobIdentifier** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**StoreFileBlobType**](StoreFileBlobType.md) |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**ContentHash** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateStoreUploadFileRequest

`func NewCreateStoreUploadFileRequest(fileName string, ) *CreateStoreUploadFileRequest`

NewCreateStoreUploadFileRequest instantiates a new CreateStoreUploadFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStoreUploadFileRequestWithDefaults

`func NewCreateStoreUploadFileRequestWithDefaults() *CreateStoreUploadFileRequest`

NewCreateStoreUploadFileRequestWithDefaults instantiates a new CreateStoreUploadFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *CreateStoreUploadFileRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CreateStoreUploadFileRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CreateStoreUploadFileRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetBlobIdentifier

`func (o *CreateStoreUploadFileRequest) GetBlobIdentifier() string`

GetBlobIdentifier returns the BlobIdentifier field if non-nil, zero value otherwise.

### GetBlobIdentifierOk

`func (o *CreateStoreUploadFileRequest) GetBlobIdentifierOk() (*string, bool)`

GetBlobIdentifierOk returns a tuple with the BlobIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobIdentifier

`func (o *CreateStoreUploadFileRequest) SetBlobIdentifier(v string)`

SetBlobIdentifier sets BlobIdentifier field to given value.

### HasBlobIdentifier

`func (o *CreateStoreUploadFileRequest) HasBlobIdentifier() bool`

HasBlobIdentifier returns a boolean if a field has been set.

### GetType

`func (o *CreateStoreUploadFileRequest) GetType() StoreFileBlobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateStoreUploadFileRequest) GetTypeOk() (*StoreFileBlobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateStoreUploadFileRequest) SetType(v StoreFileBlobType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateStoreUploadFileRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *CreateStoreUploadFileRequest) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateStoreUploadFileRequest) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateStoreUploadFileRequest) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateStoreUploadFileRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetContentHash

`func (o *CreateStoreUploadFileRequest) GetContentHash() string`

GetContentHash returns the ContentHash field if non-nil, zero value otherwise.

### GetContentHashOk

`func (o *CreateStoreUploadFileRequest) GetContentHashOk() (*string, bool)`

GetContentHashOk returns a tuple with the ContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHash

`func (o *CreateStoreUploadFileRequest) SetContentHash(v string)`

SetContentHash sets ContentHash field to given value.

### HasContentHash

`func (o *CreateStoreUploadFileRequest) HasContentHash() bool`

HasContentHash returns a boolean if a field has been set.

### GetHash

`func (o *CreateStoreUploadFileRequest) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CreateStoreUploadFileRequest) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CreateStoreUploadFileRequest) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *CreateStoreUploadFileRequest) HasHash() bool`

HasHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


