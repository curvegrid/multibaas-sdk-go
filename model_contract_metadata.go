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

// checks if the ContractMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractMetadata{}

// ContractMetadata struct for ContractMetadata
type ContractMetadata struct {
	// An alias to easily identify and reference the entity in subsequent requests.
	Label string `json:"label" validate:"regexp=^[a-z0-9_-]+$"`
	// The name of the contract.
	Name string `json:"name" validate:"regexp=^[^\\"#$%&''()*+,\\/:;<>?[\\\\\\\\\\\\]^\\\\x60{}~]*$"`
	// The contract version.
	Version string `json:"version" validate:"regexp=^[^\\"#$%&''()*+,\\/:;<>?[\\\\\\\\\\\\]^\\\\x60{}~]*$"`
}

type _ContractMetadata ContractMetadata

// NewContractMetadata instantiates a new ContractMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractMetadata(label string, name string, version string) *ContractMetadata {
	this := ContractMetadata{}
	this.Label = label
	this.Name = name
	this.Version = version
	return &this
}

// NewContractMetadataWithDefaults instantiates a new ContractMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractMetadataWithDefaults() *ContractMetadata {
	this := ContractMetadata{}
	return &this
}

// GetLabel returns the Label field value
func (o *ContractMetadata) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ContractMetadata) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ContractMetadata) SetLabel(v string) {
	o.Label = v
}

// GetName returns the Name field value
func (o *ContractMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContractMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContractMetadata) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *ContractMetadata) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ContractMetadata) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ContractMetadata) SetVersion(v string) {
	o.Version = v
}

func (o ContractMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["name"] = o.Name
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

func (o *ContractMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"name",
		"version",
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

	varContractMetadata := _ContractMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractMetadata)

	if err != nil {
		return err
	}

	*o = ContractMetadata(varContractMetadata)

	return err
}

type NullableContractMetadata struct {
	value *ContractMetadata
	isSet bool
}

func (v NullableContractMetadata) Get() *ContractMetadata {
	return v.value
}

func (v *NullableContractMetadata) Set(val *ContractMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableContractMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableContractMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractMetadata(val *ContractMetadata) *NullableContractMetadata {
	return &NullableContractMetadata{value: val, isSet: true}
}

func (v NullableContractMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
