# GetStoreFilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to **[]string** |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](paginationResponse.md) |  | [optional] 

## Methods

### NewGetStoreFilesResponse

`func NewGetStoreFilesResponse() *GetStoreFilesResponse`

NewGetStoreFilesResponse instantiates a new GetStoreFilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoreFilesResponseWithDefaults

`func NewGetStoreFilesResponseWithDefaults() *GetStoreFilesResponse`

NewGetStoreFilesResponseWithDefaults instantiates a new GetStoreFilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *GetStoreFilesResponse) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GetStoreFilesResponse) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GetStoreFilesResponse) SetFiles(v []string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *GetStoreFilesResponse) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetPagination

`func (o *GetStoreFilesResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetStoreFilesResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetStoreFilesResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetStoreFilesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


