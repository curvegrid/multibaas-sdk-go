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

// checks if the ListEventQueries200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEventQueries200ResponseAllOf{}

// ListEventQueries200ResponseAllOf struct for ListEventQueries200ResponseAllOf
type ListEventQueries200ResponseAllOf struct {
	Result []SavedEventQuery `json:"result"`
}

// NewListEventQueries200ResponseAllOf instantiates a new ListEventQueries200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEventQueries200ResponseAllOf(result []SavedEventQuery) *ListEventQueries200ResponseAllOf {
	this := ListEventQueries200ResponseAllOf{}
	this.Result = result
	return &this
}

// NewListEventQueries200ResponseAllOfWithDefaults instantiates a new ListEventQueries200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEventQueries200ResponseAllOfWithDefaults() *ListEventQueries200ResponseAllOf {
	this := ListEventQueries200ResponseAllOf{}
	return &this
}

// GetResult returns the Result field value
func (o *ListEventQueries200ResponseAllOf) GetResult() []SavedEventQuery {
	if o == nil {
		var ret []SavedEventQuery
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ListEventQueries200ResponseAllOf) GetResultOk() ([]SavedEventQuery, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *ListEventQueries200ResponseAllOf) SetResult(v []SavedEventQuery) {
	o.Result = v
}

func (o ListEventQueries200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEventQueries200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableListEventQueries200ResponseAllOf struct {
	value *ListEventQueries200ResponseAllOf
	isSet bool
}

func (v NullableListEventQueries200ResponseAllOf) Get() *ListEventQueries200ResponseAllOf {
	return v.value
}

func (v *NullableListEventQueries200ResponseAllOf) Set(val *ListEventQueries200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListEventQueries200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListEventQueries200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEventQueries200ResponseAllOf(val *ListEventQueries200ResponseAllOf) *NullableListEventQueries200ResponseAllOf {
	return &NullableListEventQueries200ResponseAllOf{value: val, isSet: true}
}

func (v NullableListEventQueries200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEventQueries200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}