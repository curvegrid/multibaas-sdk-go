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

// checks if the MethodCallPreviewResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MethodCallPreviewResponseAllOf{}

// MethodCallPreviewResponseAllOf struct for MethodCallPreviewResponseAllOf
type MethodCallPreviewResponseAllOf struct {
	// The function call inputs.
	Input []interface{} `json:"input"`
	// The function call output.
	Output interface{} `json:"output"`
}

// NewMethodCallPreviewResponseAllOf instantiates a new MethodCallPreviewResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMethodCallPreviewResponseAllOf(input []interface{}, output interface{}) *MethodCallPreviewResponseAllOf {
	this := MethodCallPreviewResponseAllOf{}
	this.Input = input
	this.Output = output
	return &this
}

// NewMethodCallPreviewResponseAllOfWithDefaults instantiates a new MethodCallPreviewResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMethodCallPreviewResponseAllOfWithDefaults() *MethodCallPreviewResponseAllOf {
	this := MethodCallPreviewResponseAllOf{}
	return &this
}

// GetInput returns the Input field value
func (o *MethodCallPreviewResponseAllOf) GetInput() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *MethodCallPreviewResponseAllOf) GetInputOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Input, true
}

// SetInput sets field value
func (o *MethodCallPreviewResponseAllOf) SetInput(v []interface{}) {
	o.Input = v
}

// GetOutput returns the Output field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *MethodCallPreviewResponseAllOf) GetOutput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Output
}

// GetOutputOk returns a tuple with the Output field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MethodCallPreviewResponseAllOf) GetOutputOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return &o.Output, true
}

// SetOutput sets field value
func (o *MethodCallPreviewResponseAllOf) SetOutput(v interface{}) {
	o.Output = v
}

func (o MethodCallPreviewResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MethodCallPreviewResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["input"] = o.Input
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	return toSerialize, nil
}

type NullableMethodCallPreviewResponseAllOf struct {
	value *MethodCallPreviewResponseAllOf
	isSet bool
}

func (v NullableMethodCallPreviewResponseAllOf) Get() *MethodCallPreviewResponseAllOf {
	return v.value
}

func (v *NullableMethodCallPreviewResponseAllOf) Set(val *MethodCallPreviewResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMethodCallPreviewResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMethodCallPreviewResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMethodCallPreviewResponseAllOf(val *MethodCallPreviewResponseAllOf) *NullableMethodCallPreviewResponseAllOf {
	return &NullableMethodCallPreviewResponseAllOf{value: val, isSet: true}
}

func (v NullableMethodCallPreviewResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMethodCallPreviewResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
