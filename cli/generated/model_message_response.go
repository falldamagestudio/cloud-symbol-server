/*
Cloud Symbol Store Uplaod API

This is the API that is used to upload symbols to Cloud Symbol Store

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MessageResponse struct for MessageResponse
type MessageResponse struct {
	Message *string `json:"message,omitempty"`
}

// NewMessageResponse instantiates a new MessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageResponse() *MessageResponse {
	this := MessageResponse{}
	return &this
}

// NewMessageResponseWithDefaults instantiates a new MessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageResponseWithDefaults() *MessageResponse {
	this := MessageResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *MessageResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MessageResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MessageResponse) SetMessage(v string) {
	o.Message = &v
}

func (o MessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableMessageResponse struct {
	value *MessageResponse
	isSet bool
}

func (v NullableMessageResponse) Get() *MessageResponse {
	return v.value
}

func (v *NullableMessageResponse) Set(val *MessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageResponse(val *MessageResponse) *NullableMessageResponse {
	return &NullableMessageResponse{value: val, isSet: true}
}

func (v NullableMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


