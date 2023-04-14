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

// checks if the SetAddress201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetAddress201Response{}

// SetAddress201Response struct for SetAddress201Response
type SetAddress201Response struct {
	// The status code.
	Status int64 `json:"status"`
	// The human-readable status message.
	Message string  `json:"message"`
	Result  Address `json:"result"`
}

// NewSetAddress201Response instantiates a new SetAddress201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetAddress201Response(status int64, message string, result Address) *SetAddress201Response {
	this := SetAddress201Response{}
	this.Status = status
	this.Message = message
	this.Result = result
	return &this
}

// NewSetAddress201ResponseWithDefaults instantiates a new SetAddress201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetAddress201ResponseWithDefaults() *SetAddress201Response {
	this := SetAddress201Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *SetAddress201Response) GetStatus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SetAddress201Response) GetStatusOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SetAddress201Response) SetStatus(v int64) {
	o.Status = v
}

// GetMessage returns the Message field value
func (o *SetAddress201Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SetAddress201Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SetAddress201Response) SetMessage(v string) {
	o.Message = v
}

// GetResult returns the Result field value
func (o *SetAddress201Response) GetResult() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *SetAddress201Response) GetResultOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *SetAddress201Response) SetResult(v Address) {
	o.Result = v
}

func (o SetAddress201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetAddress201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["message"] = o.Message
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableSetAddress201Response struct {
	value *SetAddress201Response
	isSet bool
}

func (v NullableSetAddress201Response) Get() *SetAddress201Response {
	return v.value
}

func (v *NullableSetAddress201Response) Set(val *SetAddress201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSetAddress201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSetAddress201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetAddress201Response(val *SetAddress201Response) *NullableSetAddress201Response {
	return &NullableSetAddress201Response{value: val, isSet: true}
}

func (v NullableSetAddress201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetAddress201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}