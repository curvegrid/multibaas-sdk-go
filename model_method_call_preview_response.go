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

// checks if the MethodCallPreviewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MethodCallPreviewResponse{}

// MethodCallPreviewResponse The result of a preview function arguments call.
type MethodCallPreviewResponse struct {
	PostMethodResponse
	// The function call inputs.
	Input []interface{} `json:"input"`
	// The function call output.
	Output interface{} `json:"output"`
}

type _MethodCallPreviewResponse MethodCallPreviewResponse

// NewMethodCallPreviewResponse instantiates a new MethodCallPreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMethodCallPreviewResponse(input []interface{}, output interface{}, kind string) *MethodCallPreviewResponse {
	this := MethodCallPreviewResponse{}
	this.Kind = kind
	this.Input = input
	this.Output = output
	return &this
}

// NewMethodCallPreviewResponseWithDefaults instantiates a new MethodCallPreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMethodCallPreviewResponseWithDefaults() *MethodCallPreviewResponse {
	this := MethodCallPreviewResponse{}
	return &this
}

// GetInput returns the Input field value
func (o *MethodCallPreviewResponse) GetInput() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *MethodCallPreviewResponse) GetInputOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Input, true
}

// SetInput sets field value
func (o *MethodCallPreviewResponse) SetInput(v []interface{}) {
	o.Input = v
}

// GetOutput returns the Output field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *MethodCallPreviewResponse) GetOutput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Output
}

// GetOutputOk returns a tuple with the Output field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MethodCallPreviewResponse) GetOutputOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return &o.Output, true
}

// SetOutput sets field value
func (o *MethodCallPreviewResponse) SetOutput(v interface{}) {
	o.Output = v
}

func (o MethodCallPreviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MethodCallPreviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPostMethodResponse, errPostMethodResponse := json.Marshal(o.PostMethodResponse)
	if errPostMethodResponse != nil {
		return map[string]interface{}{}, errPostMethodResponse
	}
	errPostMethodResponse = json.Unmarshal([]byte(serializedPostMethodResponse), &toSerialize)
	if errPostMethodResponse != nil {
		return map[string]interface{}{}, errPostMethodResponse
	}
	toSerialize["input"] = o.Input
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	return toSerialize, nil
}

func (o *MethodCallPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"input",
		"output",
		"kind",
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

	varMethodCallPreviewResponse := _MethodCallPreviewResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMethodCallPreviewResponse)

	if err != nil {
		return err
	}

	*o = MethodCallPreviewResponse(varMethodCallPreviewResponse)

	return err
}

type NullableMethodCallPreviewResponse struct {
	value *MethodCallPreviewResponse
	isSet bool
}

func (v NullableMethodCallPreviewResponse) Get() *MethodCallPreviewResponse {
	return v.value
}

func (v *NullableMethodCallPreviewResponse) Set(val *MethodCallPreviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMethodCallPreviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMethodCallPreviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMethodCallPreviewResponse(val *MethodCallPreviewResponse) *NullableMethodCallPreviewResponse {
	return &NullableMethodCallPreviewResponse{value: val, isSet: true}
}

func (v NullableMethodCallPreviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMethodCallPreviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
