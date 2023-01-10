# GetStoreFileHashesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hashes** | Pointer to [**[]GetStoreFileHashResponse**](GetStoreFileHashResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewGetStoreFileHashesResponse

`func NewGetStoreFileHashesResponse() *GetStoreFileHashesResponse`

NewGetStoreFileHashesResponse instantiates a new GetStoreFileHashesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoreFileHashesResponseWithDefaults

`func NewGetStoreFileHashesResponseWithDefaults() *GetStoreFileHashesResponse`

NewGetStoreFileHashesResponseWithDefaults instantiates a new GetStoreFileHashesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashes

`func (o *GetStoreFileHashesResponse) GetHashes() []GetStoreFileHashResponse`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *GetStoreFileHashesResponse) GetHashesOk() (*[]GetStoreFileHashResponse, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashes

`func (o *GetStoreFileHashesResponse) SetHashes(v []GetStoreFileHashResponse)`

SetHashes sets Hashes field to given value.

### HasHashes

`func (o *GetStoreFileHashesResponse) HasHashes() bool`

HasHashes returns a boolean if a field has been set.

### GetPagination

`func (o *GetStoreFileHashesResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetStoreFileHashesResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetStoreFileHashesResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetStoreFileHashesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


