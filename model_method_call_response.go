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

// checks if the MethodCallResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MethodCallResponse{}

// MethodCallResponse The result of a function call.
type MethodCallResponse struct {
	PostMethodResponse
	// The function call output.
	Output interface{} `json:"output"`
}

type _MethodCallResponse MethodCallResponse

// NewMethodCallResponse instantiates a new MethodCallResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMethodCallResponse(output interface{}, kind string) *MethodCallResponse {
	this := MethodCallResponse{}
	this.Kind = kind
	this.Output = output
	return &this
}

// NewMethodCallResponseWithDefaults instantiates a new MethodCallResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMethodCallResponseWithDefaults() *MethodCallResponse {
	this := MethodCallResponse{}
	return &this
}

// GetOutput returns the Output field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *MethodCallResponse) GetOutput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Output
}

// GetOutputOk returns a tuple with the Output field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MethodCallResponse) GetOutputOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return &o.Output, true
}

// SetOutput sets field value
func (o *MethodCallResponse) SetOutput(v interface{}) {
	o.Output = v
}

func (o MethodCallResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MethodCallResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPostMethodResponse, errPostMethodResponse := json.Marshal(o.PostMethodResponse)
	if errPostMethodResponse != nil {
		return map[string]interface{}{}, errPostMethodResponse
	}
	errPostMethodResponse = json.Unmarshal([]byte(serializedPostMethodResponse), &toSerialize)
	if errPostMethodResponse != nil {
		return map[string]interface{}{}, errPostMethodResponse
	}
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	return toSerialize, nil
}

func (o *MethodCallResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varMethodCallResponse := _MethodCallResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMethodCallResponse)

	if err != nil {
		return err
	}

	*o = MethodCallResponse(varMethodCallResponse)

	return err
}

type NullableMethodCallResponse struct {
	value *MethodCallResponse
	isSet bool
}

func (v NullableMethodCallResponse) Get() *MethodCallResponse {
	return v.value
}

func (v *NullableMethodCallResponse) Set(val *MethodCallResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMethodCallResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMethodCallResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMethodCallResponse(val *MethodCallResponse) *NullableMethodCallResponse {
	return &NullableMethodCallResponse{value: val, isSet: true}
}

func (v NullableMethodCallResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMethodCallResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
