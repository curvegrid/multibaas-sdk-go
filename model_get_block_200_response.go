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

// checks if the GetBlock200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBlock200Response{}

// GetBlock200Response struct for GetBlock200Response
type GetBlock200Response struct {
	// The status code.
	Status int64 `json:"status"`
	// The human-readable status message.
	Message string `json:"message"`
	Result  Block  `json:"result"`
}

// NewGetBlock200Response instantiates a new GetBlock200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlock200Response(status int64, message string, result Block) *GetBlock200Response {
	this := GetBlock200Response{}
	this.Status = status
	this.Message = message
	this.Result = result
	return &this
}

// NewGetBlock200ResponseWithDefaults instantiates a new GetBlock200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlock200ResponseWithDefaults() *GetBlock200Response {
	this := GetBlock200Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *GetBlock200Response) GetStatus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetBlock200Response) GetStatusOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetBlock200Response) SetStatus(v int64) {
	o.Status = v
}

// GetMessage returns the Message field value
func (o *GetBlock200Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GetBlock200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *GetBlock200Response) SetMessage(v string) {
	o.Message = v
}

// GetResult returns the Result field value
func (o *GetBlock200Response) GetResult() Block {
	if o == nil {
		var ret Block
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetBlock200Response) GetResultOk() (*Block, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *GetBlock200Response) SetResult(v Block) {
	o.Result = v
}

func (o GetBlock200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBlock200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["message"] = o.Message
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableGetBlock200Response struct {
	value *GetBlock200Response
	isSet bool
}

func (v NullableGetBlock200Response) Get() *GetBlock200Response {
	return v.value
}

func (v *NullableGetBlock200Response) Set(val *GetBlock200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlock200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlock200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlock200Response(val *GetBlock200Response) *NullableGetBlock200Response {
	return &NullableGetBlock200Response{value: val, isSet: true}
}

func (v NullableGetBlock200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlock200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
