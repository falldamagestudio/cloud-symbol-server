# GetStoreFileHashResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StoreFileHashStatus**](StoreFileHashStatus.md) |  | [optional] 

## Methods

### NewGetStoreFileHashResponse

`func NewGetStoreFileHashResponse() *GetStoreFileHashResponse`

NewGetStoreFileHashResponse instantiates a new GetStoreFileHashResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStoreFileHashResponseWithDefaults

`func NewGetStoreFileHashResponseWithDefaults() *GetStoreFileHashResponse`

NewGetStoreFileHashResponseWithDefaults instantiates a new GetStoreFileHashResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *GetStoreFileHashResponse) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetStoreFileHashResponse) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetStoreFileHashResponse) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GetStoreFileHashResponse) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetStatus

`func (o *GetStoreFileHashResponse) GetStatus() StoreFileHashStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStoreFileHashResponse) GetStatusOk() (*StoreFileHashStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStoreFileHashResponse) SetStatus(v StoreFileHashStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetStoreFileHashResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


