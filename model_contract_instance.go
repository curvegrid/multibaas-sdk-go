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

// checks if the ContractInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractInstance{}

// ContractInstance A contract instance
type ContractInstance struct {
	Label string `json:"label"`
	// An ethereum address.
	Address string `json:"address"`
}

// NewContractInstance instantiates a new ContractInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractInstance(label string, address string) *ContractInstance {
	this := ContractInstance{}
	this.Label = label
	this.Address = address
	return &this
}

// NewContractInstanceWithDefaults instantiates a new ContractInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractInstanceWithDefaults() *ContractInstance {
	this := ContractInstance{}
	return &this
}

// GetLabel returns the Label field value
func (o *ContractInstance) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ContractInstance) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ContractInstance) SetLabel(v string) {
	o.Label = v
}

// GetAddress returns the Address field value
func (o *ContractInstance) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ContractInstance) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ContractInstance) SetAddress(v string) {
	o.Address = v
}

func (o ContractInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

type NullableContractInstance struct {
	value *ContractInstance
	isSet bool
}

func (v NullableContractInstance) Get() *ContractInstance {
	return v.value
}

func (v *NullableContractInstance) Set(val *ContractInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableContractInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableContractInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractInstance(val *ContractInstance) *NullableContractInstance {
	return &NullableContractInstance{value: val, isSet: true}
}

func (v NullableContractInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
