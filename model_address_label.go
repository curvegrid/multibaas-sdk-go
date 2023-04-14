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

// checks if the AddressLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressLabel{}

// AddressLabel An address and it's label.
type AddressLabel struct {
	// A label.
	Label string `json:"label"`
	// An ethereum address.
	Address string `json:"address"`
}

// NewAddressLabel instantiates a new AddressLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressLabel(label string, address string) *AddressLabel {
	this := AddressLabel{}
	this.Label = label
	this.Address = address
	return &this
}

// NewAddressLabelWithDefaults instantiates a new AddressLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressLabelWithDefaults() *AddressLabel {
	this := AddressLabel{}
	return &this
}

// GetLabel returns the Label field value
func (o *AddressLabel) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *AddressLabel) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *AddressLabel) SetLabel(v string) {
	o.Label = v
}

// GetAddress returns the Address field value
func (o *AddressLabel) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AddressLabel) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AddressLabel) SetAddress(v string) {
	o.Address = v
}

func (o AddressLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

type NullableAddressLabel struct {
	value *AddressLabel
	isSet bool
}

func (v NullableAddressLabel) Get() *AddressLabel {
	return v.value
}

func (v *NullableAddressLabel) Set(val *AddressLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressLabel(val *AddressLabel) *NullableAddressLabel {
	return &NullableAddressLabel{value: val, isSet: true}
}

func (v NullableAddressLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
