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

// GetStoreUploadFileResponse struct for GetStoreUploadFileResponse
type GetStoreUploadFileResponse struct {
	FileName string `json:"fileName"`
	BlobIdentifier string `json:"blobIdentifier"`
	Status StoreUploadFileStatus `json:"status"`
}

// NewGetStoreUploadFileResponse instantiates a new GetStoreUploadFileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStoreUploadFileResponse(fileName string, blobIdentifier string, status StoreUploadFileStatus) *GetStoreUploadFileResponse {
	this := GetStoreUploadFileResponse{}
	this.FileName = fileName
	this.BlobIdentifier = blobIdentifier
	this.Status = status
	return &this
}

// NewGetStoreUploadFileResponseWithDefaults instantiates a new GetStoreUploadFileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStoreUploadFileResponseWithDefaults() *GetStoreUploadFileResponse {
	this := GetStoreUploadFileResponse{}
	return &this
}

// GetFileName returns the FileName field value
func (o *GetStoreUploadFileResponse) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *GetStoreUploadFileResponse) GetFileNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *GetStoreUploadFileResponse) SetFileName(v string) {
	o.FileName = v
}

// GetBlobIdentifier returns the BlobIdentifier field value
func (o *GetStoreUploadFileResponse) GetBlobIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlobIdentifier
}

// GetBlobIdentifierOk returns a tuple with the BlobIdentifier field value
// and a boolean to check if the value has been set.
func (o *GetStoreUploadFileResponse) GetBlobIdentifierOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BlobIdentifier, true
}

// SetBlobIdentifier sets field value
func (o *GetStoreUploadFileResponse) SetBlobIdentifier(v string) {
	o.BlobIdentifier = v
}

// GetStatus returns the Status field value
func (o *GetStoreUploadFileResponse) GetStatus() StoreUploadFileStatus {
	if o == nil {
		var ret StoreUploadFileStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetStoreUploadFileResponse) GetStatusOk() (*StoreUploadFileStatus, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetStoreUploadFileResponse) SetStatus(v StoreUploadFileStatus) {
	o.Status = v
}

func (o GetStoreUploadFileResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fileName"] = o.FileName
	}
	if true {
		toSerialize["blobIdentifier"] = o.BlobIdentifier
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGetStoreUploadFileResponse struct {
	value *GetStoreUploadFileResponse
	isSet bool
}

func (v NullableGetStoreUploadFileResponse) Get() *GetStoreUploadFileResponse {
	return v.value
}

func (v *NullableGetStoreUploadFileResponse) Set(val *GetStoreUploadFileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStoreUploadFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStoreUploadFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStoreUploadFileResponse(val *GetStoreUploadFileResponse) *NullableGetStoreUploadFileResponse {
	return &NullableGetStoreUploadFileResponse{value: val, isSet: true}
}

func (v NullableGetStoreUploadFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStoreUploadFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


