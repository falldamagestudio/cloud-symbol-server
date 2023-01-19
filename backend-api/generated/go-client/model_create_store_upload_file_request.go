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

// CreateStoreUploadFileRequest struct for CreateStoreUploadFileRequest
type CreateStoreUploadFileRequest struct {
	FileName string `json:"fileName"`
	BlobIdentifier string `json:"blobIdentifier"`
	Type *StoreFileBlobType `json:"type,omitempty"`
	Size *int64 `json:"size,omitempty"`
	ContentHash *string `json:"contentHash,omitempty"`
}

// NewCreateStoreUploadFileRequest instantiates a new CreateStoreUploadFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStoreUploadFileRequest(fileName string, blobIdentifier string) *CreateStoreUploadFileRequest {
	this := CreateStoreUploadFileRequest{}
	this.FileName = fileName
	this.BlobIdentifier = blobIdentifier
	return &this
}

// NewCreateStoreUploadFileRequestWithDefaults instantiates a new CreateStoreUploadFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStoreUploadFileRequestWithDefaults() *CreateStoreUploadFileRequest {
	this := CreateStoreUploadFileRequest{}
	return &this
}

// GetFileName returns the FileName field value
func (o *CreateStoreUploadFileRequest) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *CreateStoreUploadFileRequest) GetFileNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *CreateStoreUploadFileRequest) SetFileName(v string) {
	o.FileName = v
}

// GetBlobIdentifier returns the BlobIdentifier field value
func (o *CreateStoreUploadFileRequest) GetBlobIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlobIdentifier
}

// GetBlobIdentifierOk returns a tuple with the BlobIdentifier field value
// and a boolean to check if the value has been set.
func (o *CreateStoreUploadFileRequest) GetBlobIdentifierOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BlobIdentifier, true
}

// SetBlobIdentifier sets field value
func (o *CreateStoreUploadFileRequest) SetBlobIdentifier(v string) {
	o.BlobIdentifier = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateStoreUploadFileRequest) GetType() StoreFileBlobType {
	if o == nil || isNil(o.Type) {
		var ret StoreFileBlobType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStoreUploadFileRequest) GetTypeOk() (*StoreFileBlobType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateStoreUploadFileRequest) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given StoreFileBlobType and assigns it to the Type field.
func (o *CreateStoreUploadFileRequest) SetType(v StoreFileBlobType) {
	o.Type = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CreateStoreUploadFileRequest) GetSize() int64 {
	if o == nil || isNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStoreUploadFileRequest) GetSizeOk() (*int64, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CreateStoreUploadFileRequest) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *CreateStoreUploadFileRequest) SetSize(v int64) {
	o.Size = &v
}

// GetContentHash returns the ContentHash field value if set, zero value otherwise.
func (o *CreateStoreUploadFileRequest) GetContentHash() string {
	if o == nil || isNil(o.ContentHash) {
		var ret string
		return ret
	}
	return *o.ContentHash
}

// GetContentHashOk returns a tuple with the ContentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStoreUploadFileRequest) GetContentHashOk() (*string, bool) {
	if o == nil || isNil(o.ContentHash) {
    return nil, false
	}
	return o.ContentHash, true
}

// HasContentHash returns a boolean if a field has been set.
func (o *CreateStoreUploadFileRequest) HasContentHash() bool {
	if o != nil && !isNil(o.ContentHash) {
		return true
	}

	return false
}

// SetContentHash gets a reference to the given string and assigns it to the ContentHash field.
func (o *CreateStoreUploadFileRequest) SetContentHash(v string) {
	o.ContentHash = &v
}

func (o CreateStoreUploadFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fileName"] = o.FileName
	}
	if true {
		toSerialize["blobIdentifier"] = o.BlobIdentifier
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.ContentHash) {
		toSerialize["contentHash"] = o.ContentHash
	}
	return json.Marshal(toSerialize)
}

type NullableCreateStoreUploadFileRequest struct {
	value *CreateStoreUploadFileRequest
	isSet bool
}

func (v NullableCreateStoreUploadFileRequest) Get() *CreateStoreUploadFileRequest {
	return v.value
}

func (v *NullableCreateStoreUploadFileRequest) Set(val *CreateStoreUploadFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStoreUploadFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStoreUploadFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStoreUploadFileRequest(val *CreateStoreUploadFileRequest) *NullableCreateStoreUploadFileRequest {
	return &NullableCreateStoreUploadFileRequest{value: val, isSet: true}
}

func (v NullableCreateStoreUploadFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStoreUploadFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


