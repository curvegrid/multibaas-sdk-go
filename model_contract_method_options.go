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

// checks if the ContractMethodOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractMethodOptions{}

// ContractMethodOptions Type conversion options for a function.
type ContractMethodOptions struct {
	// The function signature.
	Signature *string `json:"signature,omitempty"`
	// List of function input parameters.
	Inputs []ContractParameter `json:"inputs"`
	// List of function output parameters.
	Outputs []ContractParameter `json:"outputs,omitempty"`
}

// NewContractMethodOptions instantiates a new ContractMethodOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractMethodOptions(inputs []ContractParameter) *ContractMethodOptions {
	this := ContractMethodOptions{}
	this.Inputs = inputs
	return &this
}

// NewContractMethodOptionsWithDefaults instantiates a new ContractMethodOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractMethodOptionsWithDefaults() *ContractMethodOptions {
	this := ContractMethodOptions{}
	return &this
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *ContractMethodOptions) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractMethodOptions) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *ContractMethodOptions) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *ContractMethodOptions) SetSignature(v string) {
	o.Signature = &v
}

// GetInputs returns the Inputs field value
func (o *ContractMethodOptions) GetInputs() []ContractParameter {
	if o == nil {
		var ret []ContractParameter
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ContractMethodOptions) GetInputsOk() ([]ContractParameter, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *ContractMethodOptions) SetInputs(v []ContractParameter) {
	o.Inputs = v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *ContractMethodOptions) GetOutputs() []ContractParameter {
	if o == nil || IsNil(o.Outputs) {
		var ret []ContractParameter
		return ret
	}
	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractMethodOptions) GetOutputsOk() ([]ContractParameter, bool) {
	if o == nil || IsNil(o.Outputs) {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *ContractMethodOptions) HasOutputs() bool {
	if o != nil && !IsNil(o.Outputs) {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given []ContractParameter and assigns it to the Outputs field.
func (o *ContractMethodOptions) SetOutputs(v []ContractParameter) {
	o.Outputs = v
}

func (o ContractMethodOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractMethodOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	toSerialize["inputs"] = o.Inputs
	if !IsNil(o.Outputs) {
		toSerialize["outputs"] = o.Outputs
	}
	return toSerialize, nil
}

type NullableContractMethodOptions struct {
	value *ContractMethodOptions
	isSet bool
}

func (v NullableContractMethodOptions) Get() *ContractMethodOptions {
	return v.value
}

func (v *NullableContractMethodOptions) Set(val *ContractMethodOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableContractMethodOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableContractMethodOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractMethodOptions(val *ContractMethodOptions) *NullableContractMethodOptions {
	return &NullableContractMethodOptions{value: val, isSet: true}
}

func (v NullableContractMethodOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractMethodOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
