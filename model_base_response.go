/*
MultiBaas API

MultiBaas's REST APIv0.

API version: 0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package multibaas

import (
	"encoding/json"
)

// checks if the BaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseResponse{}

// BaseResponse Standard response.
type BaseResponse struct {
	// The status code.
	Status int64 `json:"status"`
	// The human-readable status message.
	Message string `json:"message"`
}

// NewBaseResponse instantiates a new BaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseResponse(status int64, message string) *BaseResponse {
	this := BaseResponse{}
	this.Status = status
	this.Message = message
	return &this
}

// NewBaseResponseWithDefaults instantiates a new BaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseResponseWithDefaults() *BaseResponse {
	this := BaseResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *BaseResponse) GetStatus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BaseResponse) GetStatusOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BaseResponse) SetStatus(v int64) {
	o.Status = v
}

// GetMessage returns the Message field value
func (o *BaseResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *BaseResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *BaseResponse) SetMessage(v string) {
	o.Message = v
}

func (o BaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableBaseResponse struct {
	value *BaseResponse
	isSet bool
}

func (v NullableBaseResponse) Get() *BaseResponse {
	return v.value
}

func (v *NullableBaseResponse) Set(val *BaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseResponse(val *BaseResponse) *NullableBaseResponse {
	return &NullableBaseResponse{value: val, isSet: true}
}

func (v NullableBaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}