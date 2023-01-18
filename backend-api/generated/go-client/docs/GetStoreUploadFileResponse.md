# GetStoreUploadFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** |  | 
**BlobIdentifier** | **string** |  | 
**Status** | [**StoreUploadFileStatus**](StoreUploadFileStatus.md) |  | 

## Methods

### NewGetStoreUploadFileResponse

`func NewGetStoreUploadFileResponse(fileName string, blobIdentifier string, status StoreUploadFileStatus, ) *GetStoreUploadFileResponse`

NewGetStoreUploadFileResponse instantiates a new GetStoreUploadFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoreUploadFileResponseWithDefaults

`func NewGetStoreUploadFileResponseWithDefaults() *GetStoreUploadFileResponse`

NewGetStoreUploadFileResponseWithDefaults instantiates a new GetStoreUploadFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *GetStoreUploadFileResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *GetStoreUploadFileResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *GetStoreUploadFileResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetBlobIdentifier

`func (o *GetStoreUploadFileResponse) GetBlobIdentifier() string`

GetBlobIdentifier returns the BlobIdentifier field if non-nil, zero value otherwise.

### GetBlobIdentifierOk

`func (o *GetStoreUploadFileResponse) GetBlobIdentifierOk() (*string, bool)`

GetBlobIdentifierOk returns a tuple with the BlobIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobIdentifier

`func (o *GetStoreUploadFileResponse) SetBlobIdentifier(v string)`

SetBlobIdentifier sets BlobIdentifier field to given value.


### GetStatus

`func (o *GetStoreUploadFileResponse) GetStatus() StoreUploadFileStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStoreUploadFileResponse) GetStatusOk() (*StoreUploadFileStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStoreUploadFileResponse) SetStatus(v StoreUploadFileStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


