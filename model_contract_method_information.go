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

// checks if the ContractMethodInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractMethodInformation{}

// ContractMethodInformation The contract method's information returned within the event or transaction.
type ContractMethodInformation struct {
	// The name of the method.
	Name string `json:"name"`
	// The method signature.
	Signature string      `json:"signature"`
	Inputs    []MethodArg `json:"inputs"`
}

type _ContractMethodInformation ContractMethodInformation

// NewContractMethodInformation instantiates a new ContractMethodInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractMethodInformation(name string, signature string, inputs []MethodArg) *ContractMethodInformation {
	this := ContractMethodInformation{}
	this.Name = name
	this.Signature = signature
	this.Inputs = inputs
	return &this
}

// NewContractMethodInformationWithDefaults instantiates a new ContractMethodInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractMethodInformationWithDefaults() *ContractMethodInformation {
	this := ContractMethodInformation{}
	return &this
}

// GetName returns the Name field value
func (o *ContractMethodInformation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContractMethodInformation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContractMethodInformation) SetName(v string) {
	o.Name = v
}

// GetSignature returns the Signature field value
func (o *ContractMethodInformation) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *ContractMethodInformation) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *ContractMethodInformation) SetSignature(v string) {
	o.Signature = v
}

// GetInputs returns the Inputs field value
func (o *ContractMethodInformation) GetInputs() []MethodArg {
	if o == nil {
		var ret []MethodArg
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ContractMethodInformation) GetInputsOk() ([]MethodArg, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *ContractMethodInformation) SetInputs(v []MethodArg) {
	o.Inputs = v
}

func (o ContractMethodInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractMethodInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["signature"] = o.Signature
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

func (o *ContractMethodInformation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"signature",
		"inputs",
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

	varContractMethodInformation := _ContractMethodInformation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractMethodInformation)

	if err != nil {
		return err
	}

	*o = ContractMethodInformation(varContractMethodInformation)

	return err
}

type NullableContractMethodInformation struct {
	value *ContractMethodInformation
	isSet bool
}

func (v NullableContractMethodInformation) Get() *ContractMethodInformation {
	return v.value
}

func (v *NullableContractMethodInformation) Set(val *ContractMethodInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableContractMethodInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableContractMethodInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractMethodInformation(val *ContractMethodInformation) *NullableContractMethodInformation {
	return &NullableContractMethodInformation{value: val, isSet: true}
}

func (v NullableContractMethodInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractMethodInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
