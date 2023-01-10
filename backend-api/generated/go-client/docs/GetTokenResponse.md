# GetTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | Personal Access Token This token can be used for authentication when accessing non-token related APIs | 
**Description** | **string** | Textual description of token Users fill this in to remind themselves the purpose of a token and/or where it is used | 
**CreationTimestamp** | **string** | Creation timestamp, in RFC3339 format | 

## Methods

### NewGetTokenResponse

`func NewGetTokenResponse(token string, description string, creationTimestamp string, ) *GetTokenResponse`

NewGetTokenResponse instantiates a new GetTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTokenResponseWithDefaults

`func NewGetTokenResponseWithDefaults() *GetTokenResponse`

NewGetTokenResponseWithDefaults instantiates a new GetTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *GetTokenResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetTokenResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetTokenResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetDescription

`func (o *GetTokenResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetTokenResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetTokenResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreationTimestamp

`func (o *GetTokenResponse) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *GetTokenResponse) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *GetTokenResponse) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


