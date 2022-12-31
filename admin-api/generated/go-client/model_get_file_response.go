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

// GetFileResponse struct for GetFileResponse
type GetFileResponse struct {
	FileName *string `json:"fileName,omitempty"`
	Hash *string `json:"hash,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewGetFileResponse instantiates a new GetFileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFileResponse() *GetFileResponse {
	this := GetFileResponse{}
	return &this
}

// NewGetFileResponseWithDefaults instantiates a new GetFileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFileResponseWithDefaults() *GetFileResponse {
	this := GetFileResponse{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *GetFileResponse) GetFileName() string {
	if o == nil || isNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFileResponse) GetFileNameOk() (*string, bool) {
	if o == nil || isNil(o.FileName) {
    return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *GetFileResponse) HasFileName() bool {
	if o != nil && !isNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *GetFileResponse) SetFileName(v string) {
	o.FileName = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *GetFileResponse) GetHash() string {
	if o == nil || isNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFileResponse) GetHashOk() (*string, bool) {
	if o == nil || isNil(o.Hash) {
    return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *GetFileResponse) HasHash() bool {
	if o != nil && !isNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *GetFileResponse) SetHash(v string) {
	o.Hash = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetFileResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFileResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetFileResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetFileResponse) SetStatus(v string) {
	o.Status = &v
}

func (o GetFileResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !isNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGetFileResponse struct {
	value *GetFileResponse
	isSet bool
}

func (v NullableGetFileResponse) Get() *GetFileResponse {
	return v.value
}

func (v *NullableGetFileResponse) Set(val *GetFileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFileResponse(val *GetFileResponse) *NullableGetFileResponse {
	return &NullableGetFileResponse{value: val, isSet: true}
}

func (v NullableGetFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


