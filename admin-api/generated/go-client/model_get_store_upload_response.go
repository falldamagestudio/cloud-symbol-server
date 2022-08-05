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
	Description *string `json:"description,omitempty"`
	BuildId *string `json:"buildId,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Files *[]GetFileResponse `json:"files,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewGetStoreUploadResponse instantiates a new GetStoreUploadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStoreUploadResponse() *GetStoreUploadResponse {
	this := GetStoreUploadResponse{}
	return &this
}

// NewGetStoreUploadResponseWithDefaults instantiates a new GetStoreUploadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStoreUploadResponseWithDefaults() *GetStoreUploadResponse {
	this := GetStoreUploadResponse{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetStoreUploadResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStoreUploadResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetStoreUploadResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetStoreUploadResponse) SetDescription(v string) {
	o.Description = &v
}

// GetBuildId returns the BuildId field value if set, zero value otherwise.
func (o *GetStoreUploadResponse) GetBuildId() string {
	if o == nil || o.BuildId == nil {
		var ret string
		return ret
	}
	return *o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStoreUploadResponse) GetBuildIdOk() (*string, bool) {
	if o == nil || o.BuildId == nil {
		return nil, false
	}
	return o.BuildId, true
}

// HasBuildId returns a boolean if a field has been set.
func (o *GetStoreUploadResponse) HasBuildId() bool {
	if o != nil && o.BuildId != nil {
		return true
	}

	return false
}

// SetBuildId gets a reference to the given string and assigns it to the BuildId field.
func (o *GetStoreUploadResponse) SetBuildId(v string) {
	o.BuildId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetStoreUploadResponse) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStoreUploadResponse) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetStoreUploadResponse) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *GetStoreUploadResponse) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *GetStoreUploadResponse) GetFiles() []GetFileResponse {
	if o == nil || o.Files == nil {
		var ret []GetFileResponse
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStoreUploadResponse) GetFilesOk() (*[]GetFileResponse, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *GetStoreUploadResponse) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []GetFileResponse and assigns it to the Files field.
func (o *GetStoreUploadResponse) SetFiles(v []GetFileResponse) {
	o.Files = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetStoreUploadResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStoreUploadResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetStoreUploadResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetStoreUploadResponse) SetStatus(v string) {
	o.Status = &v
}

func (o GetStoreUploadResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.BuildId != nil {
		toSerialize["buildId"] = o.BuildId
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.Status != nil {
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


