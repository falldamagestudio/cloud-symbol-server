# GetFileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 

## Methods

### NewGetFileResponse

`func NewGetFileResponse() *GetFileResponse`

NewGetFileResponse instantiates a new GetFileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFileResponseWithDefaults

`func NewGetFileResponseWithDefaults() *GetFileResponse`

NewGetFileResponseWithDefaults instantiates a new GetFileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *GetFileResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *GetFileResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *GetFileResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *GetFileResponse) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetHash

`func (o *GetFileResponse) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetFileResponse) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetFileResponse) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GetFileResponse) HasHash() bool`

HasHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


