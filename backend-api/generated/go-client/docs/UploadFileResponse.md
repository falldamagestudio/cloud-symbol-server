# UploadFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** |  | 
**BlobIdentifier** | **string** |  | 
**Url** | Pointer to **string** | Short-lived signed URL where the client should upload the file to, or blank if the file already exists in the storage backend | [optional] 

## Methods

### NewUploadFileResponse

`func NewUploadFileResponse(fileName string, blobIdentifier string, ) *UploadFileResponse`

NewUploadFileResponse instantiates a new UploadFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFileResponseWithDefaults

`func NewUploadFileResponseWithDefaults() *UploadFileResponse`

NewUploadFileResponseWithDefaults instantiates a new UploadFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *UploadFileResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *UploadFileResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *UploadFileResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetBlobIdentifier

`func (o *UploadFileResponse) GetBlobIdentifier() string`

GetBlobIdentifier returns the BlobIdentifier field if non-nil, zero value otherwise.

### GetBlobIdentifierOk

`func (o *UploadFileResponse) GetBlobIdentifierOk() (*string, bool)`

GetBlobIdentifierOk returns a tuple with the BlobIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobIdentifier

`func (o *UploadFileResponse) SetBlobIdentifier(v string)`

SetBlobIdentifier sets BlobIdentifier field to given value.


### GetUrl

`func (o *UploadFileResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadFileResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadFileResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UploadFileResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


