# GetStoreUploadsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uploads** | [**[]GetStoreUploadResponse**](GetStoreUploadResponse.md) |  | 
**Pagination** | [**PaginationResponse**](PaginationResponse.md) |  | 

## Methods

### NewGetStoreUploadsResponse

`func NewGetStoreUploadsResponse(uploads []GetStoreUploadResponse, pagination PaginationResponse, ) *GetStoreUploadsResponse`

NewGetStoreUploadsResponse instantiates a new GetStoreUploadsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoreUploadsResponseWithDefaults

`func NewGetStoreUploadsResponseWithDefaults() *GetStoreUploadsResponse`

NewGetStoreUploadsResponseWithDefaults instantiates a new GetStoreUploadsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploads

`func (o *GetStoreUploadsResponse) GetUploads() []GetStoreUploadResponse`

GetUploads returns the Uploads field if non-nil, zero value otherwise.

### GetUploadsOk

`func (o *GetStoreUploadsResponse) GetUploadsOk() (*[]GetStoreUploadResponse, bool)`

GetUploadsOk returns a tuple with the Uploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploads

`func (o *GetStoreUploadsResponse) SetUploads(v []GetStoreUploadResponse)`

SetUploads sets Uploads field to given value.


### GetPagination

`func (o *GetStoreUploadsResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetStoreUploadsResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetStoreUploadsResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


