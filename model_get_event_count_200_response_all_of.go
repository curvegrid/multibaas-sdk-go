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

// checks if the GetEventCount200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEventCount200ResponseAllOf{}

// GetEventCount200ResponseAllOf struct for GetEventCount200ResponseAllOf
type GetEventCount200ResponseAllOf struct {
	// The number of events.
	Result int64 `json:"result"`
}

// NewGetEventCount200ResponseAllOf instantiates a new GetEventCount200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventCount200ResponseAllOf(result int64) *GetEventCount200ResponseAllOf {
	this := GetEventCount200ResponseAllOf{}
	this.Result = result
	return &this
}

// NewGetEventCount200ResponseAllOfWithDefaults instantiates a new GetEventCount200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventCount200ResponseAllOfWithDefaults() *GetEventCount200ResponseAllOf {
	this := GetEventCount200ResponseAllOf{}
	return &this
}

// GetResult returns the Result field value
func (o *GetEventCount200ResponseAllOf) GetResult() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetEventCount200ResponseAllOf) GetResultOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *GetEventCount200ResponseAllOf) SetResult(v int64) {
	o.Result = v
}

func (o GetEventCount200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEventCount200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableGetEventCount200ResponseAllOf struct {
	value *GetEventCount200ResponseAllOf
	isSet bool
}

func (v NullableGetEventCount200ResponseAllOf) Get() *GetEventCount200ResponseAllOf {
	return v.value
}

func (v *NullableGetEventCount200ResponseAllOf) Set(val *GetEventCount200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventCount200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventCount200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventCount200ResponseAllOf(val *GetEventCount200ResponseAllOf) *NullableGetEventCount200ResponseAllOf {
	return &NullableGetEventCount200ResponseAllOf{value: val, isSet: true}
}

func (v NullableGetEventCount200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventCount200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}