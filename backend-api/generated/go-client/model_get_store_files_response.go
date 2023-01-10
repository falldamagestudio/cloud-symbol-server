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

// GetStoreFilesResponse struct for GetStoreFilesResponse
type GetStoreFilesResponse struct {
	Files []string `json:"files"`
	Pagination PaginationResponse `json:"pagination"`
}

// NewGetStoreFilesResponse instantiates a new GetStoreFilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStoreFilesResponse(files []string, pagination PaginationResponse) *GetStoreFilesResponse {
	this := GetStoreFilesResponse{}
	this.Files = files
	this.Pagination = pagination
	return &this
}

// NewGetStoreFilesResponseWithDefaults instantiates a new GetStoreFilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStoreFilesResponseWithDefaults() *GetStoreFilesResponse {
	this := GetStoreFilesResponse{}
	return &this
}

// GetFiles returns the Files field value
func (o *GetStoreFilesResponse) GetFiles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *GetStoreFilesResponse) GetFilesOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *GetStoreFilesResponse) SetFiles(v []string) {
	o.Files = v
}

// GetPagination returns the Pagination field value
func (o *GetStoreFilesResponse) GetPagination() PaginationResponse {
	if o == nil {
		var ret PaginationResponse
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *GetStoreFilesResponse) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *GetStoreFilesResponse) SetPagination(v PaginationResponse) {
	o.Pagination = v
}

func (o GetStoreFilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["files"] = o.Files
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableGetStoreFilesResponse struct {
	value *GetStoreFilesResponse
	isSet bool
}

func (v NullableGetStoreFilesResponse) Get() *GetStoreFilesResponse {
	return v.value
}

func (v *NullableGetStoreFilesResponse) Set(val *GetStoreFilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStoreFilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStoreFilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStoreFilesResponse(val *GetStoreFilesResponse) *NullableGetStoreFilesResponse {
	return &NullableGetStoreFilesResponse{value: val, isSet: true}
}

func (v NullableGetStoreFilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStoreFilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


