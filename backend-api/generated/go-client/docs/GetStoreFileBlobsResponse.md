# GetStoreFileBlobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blobs** | [**[]GetStoreFileBlobResponse**](GetStoreFileBlobResponse.md) |  | 
**Pagination** | [**PaginationResponse**](PaginationResponse.md) |  | 

## Methods

### NewGetStoreFileBlobsResponse

`func NewGetStoreFileBlobsResponse(blobs []GetStoreFileBlobResponse, pagination PaginationResponse, ) *GetStoreFileBlobsResponse`

NewGetStoreFileBlobsResponse instantiates a new GetStoreFileBlobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoreFileBlobsResponseWithDefaults

`func NewGetStoreFileBlobsResponseWithDefaults() *GetStoreFileBlobsResponse`

NewGetStoreFileBlobsResponseWithDefaults instantiates a new GetStoreFileBlobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobs

`func (o *GetStoreFileBlobsResponse) GetBlobs() []GetStoreFileBlobResponse`

GetBlobs returns the Blobs field if non-nil, zero value otherwise.

### GetBlobsOk

`func (o *GetStoreFileBlobsResponse) GetBlobsOk() (*[]GetStoreFileBlobResponse, bool)`

GetBlobsOk returns a tuple with the Blobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobs

`func (o *GetStoreFileBlobsResponse) SetBlobs(v []GetStoreFileBlobResponse)`

SetBlobs sets Blobs field to given value.


### GetPagination

`func (o *GetStoreFileBlobsResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetStoreFileBlobsResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetStoreFileBlobsResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


