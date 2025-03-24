/*
MultiBaas API

MultiBaas's REST APIv0.

API version: 0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package multibaas

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ContractInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractInstance{}

// ContractInstance A contract instance
type ContractInstance struct {
	Alias string `json:"alias"`
	// An ethereum address.
	Address string `json:"address" validate:"regexp=^0[xX][a-fA-F0-9]{40}$"`
}

type _ContractInstance ContractInstance

// NewContractInstance instantiates a new ContractInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractInstance(alias string, address string) *ContractInstance {
	this := ContractInstance{}
	this.Alias = alias
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

// GetAlias returns the Alias field value
func (o *ContractInstance) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *ContractInstance) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *ContractInstance) SetAlias(v string) {
	o.Alias = v
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
	toSerialize["alias"] = o.Alias
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

func (o *ContractInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alias",
		"address",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varContractInstance := _ContractInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractInstance)

	if err != nil {
		return err
	}

	*o = ContractInstance(varContractInstance)

	return err
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
