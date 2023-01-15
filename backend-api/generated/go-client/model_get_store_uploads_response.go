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

// GetStoreUploadsResponse struct for GetStoreUploadsResponse
type GetStoreUploadsResponse struct {
	Uploads []GetStoreUploadResponse `json:"uploads"`
	Pagination PaginationResponse `json:"pagination"`
}

// NewGetStoreUploadsResponse instantiates a new GetStoreUploadsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStoreUploadsResponse(uploads []GetStoreUploadResponse, pagination PaginationResponse) *GetStoreUploadsResponse {
	this := GetStoreUploadsResponse{}
	this.Uploads = uploads
	this.Pagination = pagination
	return &this
}

// NewGetStoreUploadsResponseWithDefaults instantiates a new GetStoreUploadsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStoreUploadsResponseWithDefaults() *GetStoreUploadsResponse {
	this := GetStoreUploadsResponse{}
	return &this
}

// GetUploads returns the Uploads field value
func (o *GetStoreUploadsResponse) GetUploads() []GetStoreUploadResponse {
	if o == nil {
		var ret []GetStoreUploadResponse
		return ret
	}

	return o.Uploads
}

// GetUploadsOk returns a tuple with the Uploads field value
// and a boolean to check if the value has been set.
func (o *GetStoreUploadsResponse) GetUploadsOk() ([]GetStoreUploadResponse, bool) {
	if o == nil {
    return nil, false
	}
	return o.Uploads, true
}

// SetUploads sets field value
func (o *GetStoreUploadsResponse) SetUploads(v []GetStoreUploadResponse) {
	o.Uploads = v
}

// GetPagination returns the Pagination field value
func (o *GetStoreUploadsResponse) GetPagination() PaginationResponse {
	if o == nil {
		var ret PaginationResponse
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *GetStoreUploadsResponse) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *GetStoreUploadsResponse) SetPagination(v PaginationResponse) {
	o.Pagination = v
}

func (o GetStoreUploadsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uploads"] = o.Uploads
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableGetStoreUploadsResponse struct {
	value *GetStoreUploadsResponse
	isSet bool
}

func (v NullableGetStoreUploadsResponse) Get() *GetStoreUploadsResponse {
	return v.value
}

func (v *NullableGetStoreUploadsResponse) Set(val *GetStoreUploadsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStoreUploadsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStoreUploadsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStoreUploadsResponse(val *GetStoreUploadsResponse) *NullableGetStoreUploadsResponse {
	return &NullableGetStoreUploadsResponse{value: val, isSet: true}
}

func (v NullableGetStoreUploadsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStoreUploadsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

