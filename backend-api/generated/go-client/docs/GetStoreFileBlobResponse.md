# GetStoreFileBlobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobIdentifier** | **string** |  | 
**UploadTimestamp** | **string** | Upload timestamp, in RFC3339 format | 
**Type** | Pointer to [**StoreFileBlobType**](StoreFileBlobType.md) |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**ContentHash** | Pointer to **string** |  | [optional] 
**Status** | [**StoreFileBlobStatus**](StoreFileBlobStatus.md) |  | 

## Methods

### NewGetStoreFileBlobResponse

`func NewGetStoreFileBlobResponse(blobIdentifier string, uploadTimestamp string, status StoreFileBlobStatus, ) *GetStoreFileBlobResponse`

NewGetStoreFileBlobResponse instantiates a new GetStoreFileBlobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoreFileBlobResponseWithDefaults

`func NewGetStoreFileBlobResponseWithDefaults() *GetStoreFileBlobResponse`

NewGetStoreFileBlobResponseWithDefaults instantiates a new GetStoreFileBlobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobIdentifier

`func (o *GetStoreFileBlobResponse) GetBlobIdentifier() string`

GetBlobIdentifier returns the BlobIdentifier field if non-nil, zero value otherwise.

### GetBlobIdentifierOk

`func (o *GetStoreFileBlobResponse) GetBlobIdentifierOk() (*string, bool)`

GetBlobIdentifierOk returns a tuple with the BlobIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobIdentifier

`func (o *GetStoreFileBlobResponse) SetBlobIdentifier(v string)`

SetBlobIdentifier sets BlobIdentifier field to given value.


### GetUploadTimestamp

`func (o *GetStoreFileBlobResponse) GetUploadTimestamp() string`

GetUploadTimestamp returns the UploadTimestamp field if non-nil, zero value otherwise.

### GetUploadTimestampOk

`func (o *GetStoreFileBlobResponse) GetUploadTimestampOk() (*string, bool)`

GetUploadTimestampOk returns a tuple with the UploadTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTimestamp

`func (o *GetStoreFileBlobResponse) SetUploadTimestamp(v string)`

SetUploadTimestamp sets UploadTimestamp field to given value.


### GetType

`func (o *GetStoreFileBlobResponse) GetType() StoreFileBlobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetStoreFileBlobResponse) GetTypeOk() (*StoreFileBlobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetStoreFileBlobResponse) SetType(v StoreFileBlobType)`

SetType sets Type field to given value.

### HasType

`func (o *GetStoreFileBlobResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *GetStoreFileBlobResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetStoreFileBlobResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetStoreFileBlobResponse) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetStoreFileBlobResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetContentHash

`func (o *GetStoreFileBlobResponse) GetContentHash() string`

GetContentHash returns the ContentHash field if non-nil, zero value otherwise.

### GetContentHashOk

`func (o *GetStoreFileBlobResponse) GetContentHashOk() (*string, bool)`

GetContentHashOk returns a tuple with the ContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHash

`func (o *GetStoreFileBlobResponse) SetContentHash(v string)`

SetContentHash sets ContentHash field to given value.

### HasContentHash

`func (o *GetStoreFileBlobResponse) HasContentHash() bool`

HasContentHash returns a boolean if a field has been set.

### GetStatus

`func (o *GetStoreFileBlobResponse) GetStatus() StoreFileBlobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStoreFileBlobResponse) GetStatusOk() (*StoreFileBlobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStoreFileBlobResponse) SetStatus(v StoreFileBlobStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


