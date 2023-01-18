# GetStoreFileBlobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobIdentifier** | **string** |  | 
**UploadTimestamp** | **string** | Upload timestamp, in RFC3339 format | 
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


