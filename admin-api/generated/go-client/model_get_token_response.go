/*
Cloud Symbol Server Admin API

This is the API that is used to manage stores and uploads in Cloud Symbol Server

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetTokenResponse struct for GetTokenResponse
type GetTokenResponse struct {
	// Personal Access Token This token can be used for authentication when accessing non-token related APIs
	Token *string `json:"token,omitempty"`
	// Textual description of token Users fill this in to remind themselves the purpose of a token and/or where it is used
	Description *string `json:"description,omitempty"`
	// Creation timestamp, in RFC3339 format
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`
}

// NewGetTokenResponse instantiates a new GetTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTokenResponse() *GetTokenResponse {
	this := GetTokenResponse{}
	return &this
}

// NewGetTokenResponseWithDefaults instantiates a new GetTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTokenResponseWithDefaults() *GetTokenResponse {
	this := GetTokenResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GetTokenResponse) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTokenResponse) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GetTokenResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GetTokenResponse) SetToken(v string) {
	o.Token = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetTokenResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTokenResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetTokenResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetTokenResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *GetTokenResponse) GetCreationTimestamp() string {
	if o == nil || o.CreationTimestamp == nil {
		var ret string
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTokenResponse) GetCreationTimestampOk() (*string, bool) {
	if o == nil || o.CreationTimestamp == nil {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *GetTokenResponse) HasCreationTimestamp() bool {
	if o != nil && o.CreationTimestamp != nil {
		return true
	}

	return false
}

// SetCreationTimestamp gets a reference to the given string and assigns it to the CreationTimestamp field.
func (o *GetTokenResponse) SetCreationTimestamp(v string) {
	o.CreationTimestamp = &v
}

func (o GetTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CreationTimestamp != nil {
		toSerialize["creationTimestamp"] = o.CreationTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableGetTokenResponse struct {
	value *GetTokenResponse
	isSet bool
}

func (v NullableGetTokenResponse) Get() *GetTokenResponse {
	return v.value
}

func (v *NullableGetTokenResponse) Set(val *GetTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTokenResponse(val *GetTokenResponse) *NullableGetTokenResponse {
	return &NullableGetTokenResponse{value: val, isSet: true}
}

func (v NullableGetTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

