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

// checks if the ListHsmWallets200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHsmWallets200ResponseAllOf{}

// ListHsmWallets200ResponseAllOf struct for ListHsmWallets200ResponseAllOf
type ListHsmWallets200ResponseAllOf struct {
	Result []StandaloneWallet `json:"result"`
}

// NewListHsmWallets200ResponseAllOf instantiates a new ListHsmWallets200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHsmWallets200ResponseAllOf(result []StandaloneWallet) *ListHsmWallets200ResponseAllOf {
	this := ListHsmWallets200ResponseAllOf{}
	this.Result = result
	return &this
}

// NewListHsmWallets200ResponseAllOfWithDefaults instantiates a new ListHsmWallets200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHsmWallets200ResponseAllOfWithDefaults() *ListHsmWallets200ResponseAllOf {
	this := ListHsmWallets200ResponseAllOf{}
	return &this
}

// GetResult returns the Result field value
func (o *ListHsmWallets200ResponseAllOf) GetResult() []StandaloneWallet {
	if o == nil {
		var ret []StandaloneWallet
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ListHsmWallets200ResponseAllOf) GetResultOk() ([]StandaloneWallet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *ListHsmWallets200ResponseAllOf) SetResult(v []StandaloneWallet) {
	o.Result = v
}

func (o ListHsmWallets200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListHsmWallets200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableListHsmWallets200ResponseAllOf struct {
	value *ListHsmWallets200ResponseAllOf
	isSet bool
}

func (v NullableListHsmWallets200ResponseAllOf) Get() *ListHsmWallets200ResponseAllOf {
	return v.value
}

func (v *NullableListHsmWallets200ResponseAllOf) Set(val *ListHsmWallets200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListHsmWallets200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListHsmWallets200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHsmWallets200ResponseAllOf(val *ListHsmWallets200ResponseAllOf) *NullableListHsmWallets200ResponseAllOf {
	return &NullableListHsmWallets200ResponseAllOf{value: val, isSet: true}
}

func (v NullableListHsmWallets200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListHsmWallets200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
