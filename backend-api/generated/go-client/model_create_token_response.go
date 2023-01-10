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

// CreateTokenResponse struct for CreateTokenResponse
type CreateTokenResponse struct {
	// Personal Access Token
	Token string `json:"token"`
}

// NewCreateTokenResponse instantiates a new CreateTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTokenResponse(token string) *CreateTokenResponse {
	this := CreateTokenResponse{}
	this.Token = token
	return &this
}

// NewCreateTokenResponseWithDefaults instantiates a new CreateTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTokenResponseWithDefaults() *CreateTokenResponse {
	this := CreateTokenResponse{}
	return &this
}

// GetToken returns the Token field value
func (o *CreateTokenResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CreateTokenResponse) GetTokenOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CreateTokenResponse) SetToken(v string) {
	o.Token = v
}

func (o CreateTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTokenResponse struct {
	value *CreateTokenResponse
	isSet bool
}

func (v NullableCreateTokenResponse) Get() *CreateTokenResponse {
	return v.value
}

func (v *NullableCreateTokenResponse) Set(val *CreateTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTokenResponse(val *CreateTokenResponse) *NullableCreateTokenResponse {
	return &NullableCreateTokenResponse{value: val, isSet: true}
}

func (v NullableCreateTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


