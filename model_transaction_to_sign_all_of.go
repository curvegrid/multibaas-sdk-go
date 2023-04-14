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

// checks if the TransactionToSignAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionToSignAllOf{}

// TransactionToSignAllOf struct for TransactionToSignAllOf
type TransactionToSignAllOf struct {
	Submitted bool `json:"submitted"`
}

// NewTransactionToSignAllOf instantiates a new TransactionToSignAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionToSignAllOf(submitted bool) *TransactionToSignAllOf {
	this := TransactionToSignAllOf{}
	this.Submitted = submitted
	return &this
}

// NewTransactionToSignAllOfWithDefaults instantiates a new TransactionToSignAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionToSignAllOfWithDefaults() *TransactionToSignAllOf {
	this := TransactionToSignAllOf{}
	return &this
}

// GetSubmitted returns the Submitted field value
func (o *TransactionToSignAllOf) GetSubmitted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Submitted
}

// GetSubmittedOk returns a tuple with the Submitted field value
// and a boolean to check if the value has been set.
func (o *TransactionToSignAllOf) GetSubmittedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Submitted, true
}

// SetSubmitted sets field value
func (o *TransactionToSignAllOf) SetSubmitted(v bool) {
	o.Submitted = v
}

func (o TransactionToSignAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionToSignAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["submitted"] = o.Submitted
	return toSerialize, nil
}

type NullableTransactionToSignAllOf struct {
	value *TransactionToSignAllOf
	isSet bool
}

func (v NullableTransactionToSignAllOf) Get() *TransactionToSignAllOf {
	return v.value
}

func (v *NullableTransactionToSignAllOf) Set(val *TransactionToSignAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionToSignAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionToSignAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionToSignAllOf(val *TransactionToSignAllOf) *NullableTransactionToSignAllOf {
	return &NullableTransactionToSignAllOf{value: val, isSet: true}
}

func (v NullableTransactionToSignAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionToSignAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}