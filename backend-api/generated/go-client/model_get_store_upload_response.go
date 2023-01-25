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

// GetStoreUploadResponse struct for GetStoreUploadResponse
type GetStoreUploadResponse struct {
	UploadId int32 `json:"uploadId"`
	Description string `json:"description"`
	BuildId string `json:"buildId"`
	UploadTimestamp string `json:"uploadTimestamp"`
	ExpiryTimestamp string `json:"expiryTimestamp"`
	Files []GetStoreUploadFileResponse `json:"files"`
	Status StoreUploadStatus `json:"status"`
}

// NewGetStoreUploadResponse instantiates a new GetStoreUploadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStoreUploadResponse(uploadId int32, description string, buildId string, uploadTimestamp string, expiryTimestamp string, files []GetStoreUploadFileResponse, status StoreUploadStatus) *GetStoreUploadResponse {
	this := GetStoreUploadResponse{}
	this.UploadId = uploadId
	this.Description = description
	this.BuildId = buildId
	this.UploadTimestamp = uploadTimestamp
	this.ExpiryTimestamp = expiryTimestamp
	this.Files = files
	this.Status = status
	return &this
}

// NewGetStoreUploadResponseWithDefaults instantiates a new GetStoreUploadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStoreUploadResponseWithDefaults() *GetStoreUploadResponse {
	this := GetStoreUploadResponse{}
	return &this
}

// GetUploadId returns the UploadId field value
func (o *GetStoreUploadResponse) GetUploadId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value
// and a boolean to check if the value has been set.
func (o *GetStoreUploadResponse) GetUploadIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UploadId, true
}

// SetUploadId sets field value
func (o *GetStoreUploadResponse) SetUploadId(v int32) {
	o.UploadId = v
}

// GetDescription returns the Description field value
func (o *GetStoreUploadResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GetStoreUploadResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GetStoreUploadResponse) SetDescription(v string) {
	o.Description = v
}

// GetBuildId returns the BuildId field value
func (o *GetStoreUploadResponse) GetBuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value
// and a boolean to check if the value has been set.
func (o *GetStoreUploadResponse) GetBuildIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BuildId, true
}

// SetBuildId sets field value
func (o *GetStoreUploadResponse) SetBuildId(v string) {
	o.BuildId = v
}

// GetUploadTimestamp returns the UploadTimestamp field value
func (o *GetStoreUploadResponse) GetUploadTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadTimestamp
}

// GetUploadTimestampOk returns a tuple with the UploadTimestamp field value
// and a boolean to check if the value has been set.
func (o *GetStoreUploadResponse) GetUploadTimestampOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UploadTimestamp, true
}

// SetUploadTimestamp sets field value
func (o *GetStoreUploadResponse) SetUploadTimestamp(v string) {
	o.UploadTimestamp = v
}

// GetExpiryTimestamp returns the ExpiryTimestamp field value
func (o *GetStoreUploadResponse) GetExpiryTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiryTimestamp
}

// GetExpiryTimestampOk returns a tuple with the ExpiryTimestamp field value
// and a boolean to check if the value has been set.
func (o *GetStoreUploadResponse) GetExpiryTimestampOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExpiryTimestamp, true
}

// SetExpiryTimestamp sets field value
func (o *GetStoreUploadResponse) SetExpiryTimestamp(v string) {
	o.ExpiryTimestamp = v
}

// GetFiles returns the Files field value
func (o *GetStoreUploadResponse) GetFiles() []GetStoreUploadFileResponse {
	if o == nil {
		var ret []GetStoreUploadFileResponse
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *GetStoreUploadResponse) GetFilesOk() ([]GetStoreUploadFileResponse, bool) {
	if o == nil {
    return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *GetStoreUploadResponse) SetFiles(v []GetStoreUploadFileResponse) {
	o.Files = v
}

// GetStatus returns the Status field value
func (o *GetStoreUploadResponse) GetStatus() StoreUploadStatus {
	if o == nil {
		var ret StoreUploadStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetStoreUploadResponse) GetStatusOk() (*StoreUploadStatus, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetStoreUploadResponse) SetStatus(v StoreUploadStatus) {
	o.Status = v
}

func (o GetStoreUploadResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uploadId"] = o.UploadId
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["buildId"] = o.BuildId
	}
	if true {
		toSerialize["uploadTimestamp"] = o.UploadTimestamp
	}
	if true {
		toSerialize["expiryTimestamp"] = o.ExpiryTimestamp
	}
	if true {
		toSerialize["files"] = o.Files
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGetStoreUploadResponse struct {
	value *GetStoreUploadResponse
	isSet bool
}

func (v NullableGetStoreUploadResponse) Get() *GetStoreUploadResponse {
	return v.value
}

func (v *NullableGetStoreUploadResponse) Set(val *GetStoreUploadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStoreUploadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStoreUploadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStoreUploadResponse(val *GetStoreUploadResponse) *NullableGetStoreUploadResponse {
	return &NullableGetStoreUploadResponse{value: val, isSet: true}
}

func (v NullableGetStoreUploadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStoreUploadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


