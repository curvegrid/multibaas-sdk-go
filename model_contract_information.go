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

// checks if the ContractInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractInformation{}

// ContractInformation The contract information within the event or transaction.
type ContractInformation struct {
	// An ethereum address.
	Address string `json:"address"`
	// A label.
	AddressLabel string `json:"addressLabel"`
	// The name of the contract.
	Name string `json:"name"`
	// A label.
	Label string `json:"label"`
}

// NewContractInformation instantiates a new ContractInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractInformation(address string, addressLabel string, name string, label string) *ContractInformation {
	this := ContractInformation{}
	this.Address = address
	this.AddressLabel = addressLabel
	this.Name = name
	this.Label = label
	return &this
}

// NewContractInformationWithDefaults instantiates a new ContractInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractInformationWithDefaults() *ContractInformation {
	this := ContractInformation{}
	return &this
}

// GetAddress returns the Address field value
func (o *ContractInformation) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ContractInformation) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ContractInformation) SetAddress(v string) {
	o.Address = v
}

// GetAddressLabel returns the AddressLabel field value
func (o *ContractInformation) GetAddressLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLabel
}

// GetAddressLabelOk returns a tuple with the AddressLabel field value
// and a boolean to check if the value has been set.
func (o *ContractInformation) GetAddressLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLabel, true
}

// SetAddressLabel sets field value
func (o *ContractInformation) SetAddressLabel(v string) {
	o.AddressLabel = v
}

// GetName returns the Name field value
func (o *ContractInformation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContractInformation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContractInformation) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value
func (o *ContractInformation) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ContractInformation) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ContractInformation) SetLabel(v string) {
	o.Label = v
}

func (o ContractInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["addressLabel"] = o.AddressLabel
	toSerialize["name"] = o.Name
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

type NullableContractInformation struct {
	value *ContractInformation
	isSet bool
}

func (v NullableContractInformation) Get() *ContractInformation {
	return v.value
}

func (v *NullableContractInformation) Set(val *ContractInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableContractInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableContractInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractInformation(val *ContractInformation) *NullableContractInformation {
	return &NullableContractInformation{value: val, isSet: true}
}

func (v NullableContractInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}