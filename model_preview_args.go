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

// checks if the PreviewArgs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreviewArgs{}

// PreviewArgs Ephemeral configuration for previewing the effect of a Type Conversion on a contract function call.
type PreviewArgs struct {
	// Only preview the effect of a Type Conversion on the inputs. Only applicable for write function calls, where the output is an unsigned transaction.
	InputsOnly bool `json:"inputsOnly"`
	// Type Conversion information for the function inputs. The number of inputs must match that of the actual function inputs. The parameter is a contract function argument where only the type conversion information is used.
	Inputs []ContractABIMethodArgument `json:"inputs"`
	// Type Conversion information for the function outputs. The number of outputs must match that of the actual function outputs. The parameter is a contract function argument where only the type conversion information is used.
	Outputs []ContractABIMethodArgument `json:"outputs"`
}

type _PreviewArgs PreviewArgs

// NewPreviewArgs instantiates a new PreviewArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviewArgs(inputsOnly bool, inputs []ContractABIMethodArgument, outputs []ContractABIMethodArgument) *PreviewArgs {
	this := PreviewArgs{}
	this.InputsOnly = inputsOnly
	this.Inputs = inputs
	this.Outputs = outputs
	return &this
}

// NewPreviewArgsWithDefaults instantiates a new PreviewArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewArgsWithDefaults() *PreviewArgs {
	this := PreviewArgs{}
	return &this
}

// GetInputsOnly returns the InputsOnly field value
func (o *PreviewArgs) GetInputsOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InputsOnly
}

// GetInputsOnlyOk returns a tuple with the InputsOnly field value
// and a boolean to check if the value has been set.
func (o *PreviewArgs) GetInputsOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputsOnly, true
}

// SetInputsOnly sets field value
func (o *PreviewArgs) SetInputsOnly(v bool) {
	o.InputsOnly = v
}

// GetInputs returns the Inputs field value
func (o *PreviewArgs) GetInputs() []ContractABIMethodArgument {
	if o == nil {
		var ret []ContractABIMethodArgument
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *PreviewArgs) GetInputsOk() ([]ContractABIMethodArgument, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *PreviewArgs) SetInputs(v []ContractABIMethodArgument) {
	o.Inputs = v
}

// GetOutputs returns the Outputs field value
func (o *PreviewArgs) GetOutputs() []ContractABIMethodArgument {
	if o == nil {
		var ret []ContractABIMethodArgument
		return ret
	}

	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value
// and a boolean to check if the value has been set.
func (o *PreviewArgs) GetOutputsOk() ([]ContractABIMethodArgument, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outputs, true
}

// SetOutputs sets field value
func (o *PreviewArgs) SetOutputs(v []ContractABIMethodArgument) {
	o.Outputs = v
}

func (o PreviewArgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputsOnly"] = o.InputsOnly
	toSerialize["inputs"] = o.Inputs
	toSerialize["outputs"] = o.Outputs
	return toSerialize, nil
}

type NullablePreviewArgs struct {
	value *PreviewArgs
	isSet bool
}

func (v NullablePreviewArgs) Get() *PreviewArgs {
	return v.value
}

func (v *NullablePreviewArgs) Set(val *PreviewArgs) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviewArgs) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviewArgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviewArgs(val *PreviewArgs) *NullablePreviewArgs {
	return &NullablePreviewArgs{value: val, isSet: true}
}

func (v NullablePreviewArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviewArgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
