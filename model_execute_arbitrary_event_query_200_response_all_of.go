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

// checks if the ExecuteArbitraryEventQuery200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecuteArbitraryEventQuery200ResponseAllOf{}

// ExecuteArbitraryEventQuery200ResponseAllOf struct for ExecuteArbitraryEventQuery200ResponseAllOf
type ExecuteArbitraryEventQuery200ResponseAllOf struct {
	Result EventQueryResults `json:"result"`
}

// NewExecuteArbitraryEventQuery200ResponseAllOf instantiates a new ExecuteArbitraryEventQuery200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteArbitraryEventQuery200ResponseAllOf(result EventQueryResults) *ExecuteArbitraryEventQuery200ResponseAllOf {
	this := ExecuteArbitraryEventQuery200ResponseAllOf{}
	this.Result = result
	return &this
}

// NewExecuteArbitraryEventQuery200ResponseAllOfWithDefaults instantiates a new ExecuteArbitraryEventQuery200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteArbitraryEventQuery200ResponseAllOfWithDefaults() *ExecuteArbitraryEventQuery200ResponseAllOf {
	this := ExecuteArbitraryEventQuery200ResponseAllOf{}
	return &this
}

// GetResult returns the Result field value
func (o *ExecuteArbitraryEventQuery200ResponseAllOf) GetResult() EventQueryResults {
	if o == nil {
		var ret EventQueryResults
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ExecuteArbitraryEventQuery200ResponseAllOf) GetResultOk() (*EventQueryResults, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *ExecuteArbitraryEventQuery200ResponseAllOf) SetResult(v EventQueryResults) {
	o.Result = v
}

func (o ExecuteArbitraryEventQuery200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecuteArbitraryEventQuery200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableExecuteArbitraryEventQuery200ResponseAllOf struct {
	value *ExecuteArbitraryEventQuery200ResponseAllOf
	isSet bool
}

func (v NullableExecuteArbitraryEventQuery200ResponseAllOf) Get() *ExecuteArbitraryEventQuery200ResponseAllOf {
	return v.value
}

func (v *NullableExecuteArbitraryEventQuery200ResponseAllOf) Set(val *ExecuteArbitraryEventQuery200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteArbitraryEventQuery200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteArbitraryEventQuery200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteArbitraryEventQuery200ResponseAllOf(val *ExecuteArbitraryEventQuery200ResponseAllOf) *NullableExecuteArbitraryEventQuery200ResponseAllOf {
	return &NullableExecuteArbitraryEventQuery200ResponseAllOf{value: val, isSet: true}
}

func (v NullableExecuteArbitraryEventQuery200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteArbitraryEventQuery200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}