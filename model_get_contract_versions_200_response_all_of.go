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

// checks if the GetContractVersions200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContractVersions200ResponseAllOf{}

// GetContractVersions200ResponseAllOf struct for GetContractVersions200ResponseAllOf
type GetContractVersions200ResponseAllOf struct {
	Result []Contract `json:"result"`
}

// NewGetContractVersions200ResponseAllOf instantiates a new GetContractVersions200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContractVersions200ResponseAllOf(result []Contract) *GetContractVersions200ResponseAllOf {
	this := GetContractVersions200ResponseAllOf{}
	this.Result = result
	return &this
}

// NewGetContractVersions200ResponseAllOfWithDefaults instantiates a new GetContractVersions200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContractVersions200ResponseAllOfWithDefaults() *GetContractVersions200ResponseAllOf {
	this := GetContractVersions200ResponseAllOf{}
	return &this
}

// GetResult returns the Result field value
func (o *GetContractVersions200ResponseAllOf) GetResult() []Contract {
	if o == nil {
		var ret []Contract
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetContractVersions200ResponseAllOf) GetResultOk() ([]Contract, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *GetContractVersions200ResponseAllOf) SetResult(v []Contract) {
	o.Result = v
}

func (o GetContractVersions200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetContractVersions200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableGetContractVersions200ResponseAllOf struct {
	value *GetContractVersions200ResponseAllOf
	isSet bool
}

func (v NullableGetContractVersions200ResponseAllOf) Get() *GetContractVersions200ResponseAllOf {
	return v.value
}

func (v *NullableGetContractVersions200ResponseAllOf) Set(val *GetContractVersions200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContractVersions200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContractVersions200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContractVersions200ResponseAllOf(val *GetContractVersions200ResponseAllOf) *NullableGetContractVersions200ResponseAllOf {
	return &NullableGetContractVersions200ResponseAllOf{value: val, isSet: true}
}

func (v NullableGetContractVersions200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContractVersions200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
