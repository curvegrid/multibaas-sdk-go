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

// checks if the GetAllEvents200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllEvents200ResponseAllOf{}

// GetAllEvents200ResponseAllOf struct for GetAllEvents200ResponseAllOf
type GetAllEvents200ResponseAllOf struct {
	Result []Event `json:"result"`
}

// NewGetAllEvents200ResponseAllOf instantiates a new GetAllEvents200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllEvents200ResponseAllOf(result []Event) *GetAllEvents200ResponseAllOf {
	this := GetAllEvents200ResponseAllOf{}
	this.Result = result
	return &this
}

// NewGetAllEvents200ResponseAllOfWithDefaults instantiates a new GetAllEvents200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllEvents200ResponseAllOfWithDefaults() *GetAllEvents200ResponseAllOf {
	this := GetAllEvents200ResponseAllOf{}
	return &this
}

// GetResult returns the Result field value
func (o *GetAllEvents200ResponseAllOf) GetResult() []Event {
	if o == nil {
		var ret []Event
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetAllEvents200ResponseAllOf) GetResultOk() ([]Event, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *GetAllEvents200ResponseAllOf) SetResult(v []Event) {
	o.Result = v
}

func (o GetAllEvents200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllEvents200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableGetAllEvents200ResponseAllOf struct {
	value *GetAllEvents200ResponseAllOf
	isSet bool
}

func (v NullableGetAllEvents200ResponseAllOf) Get() *GetAllEvents200ResponseAllOf {
	return v.value
}

func (v *NullableGetAllEvents200ResponseAllOf) Set(val *GetAllEvents200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllEvents200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllEvents200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllEvents200ResponseAllOf(val *GetAllEvents200ResponseAllOf) *NullableGetAllEvents200ResponseAllOf {
	return &NullableGetAllEvents200ResponseAllOf{value: val, isSet: true}
}

func (v NullableGetAllEvents200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllEvents200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
