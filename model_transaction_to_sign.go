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

// checks if the TransactionToSign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionToSign{}

// TransactionToSign A transaction to be signed.
type TransactionToSign struct {
	Tx        TransactionToSignTx `json:"tx"`
	Submitted bool                `json:"submitted"`
}

type _TransactionToSign TransactionToSign

// NewTransactionToSign instantiates a new TransactionToSign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionToSign(tx TransactionToSignTx, submitted bool) *TransactionToSign {
	this := TransactionToSign{}
	this.Tx = tx
	this.Submitted = submitted
	return &this
}

// NewTransactionToSignWithDefaults instantiates a new TransactionToSign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionToSignWithDefaults() *TransactionToSign {
	this := TransactionToSign{}
	return &this
}

// GetTx returns the Tx field value
func (o *TransactionToSign) GetTx() TransactionToSignTx {
	if o == nil {
		var ret TransactionToSignTx
		return ret
	}

	return o.Tx
}

// GetTxOk returns a tuple with the Tx field value
// and a boolean to check if the value has been set.
func (o *TransactionToSign) GetTxOk() (*TransactionToSignTx, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tx, true
}

// SetTx sets field value
func (o *TransactionToSign) SetTx(v TransactionToSignTx) {
	o.Tx = v
}

// GetSubmitted returns the Submitted field value
func (o *TransactionToSign) GetSubmitted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Submitted
}

// GetSubmittedOk returns a tuple with the Submitted field value
// and a boolean to check if the value has been set.
func (o *TransactionToSign) GetSubmittedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Submitted, true
}

// SetSubmitted sets field value
func (o *TransactionToSign) SetSubmitted(v bool) {
	o.Submitted = v
}

func (o TransactionToSign) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionToSign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx"] = o.Tx
	toSerialize["submitted"] = o.Submitted
	return toSerialize, nil
}

func (o *TransactionToSign) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tx",
		"submitted",
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

	varTransactionToSign := _TransactionToSign{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionToSign)

	if err != nil {
		return err
	}

	*o = TransactionToSign(varTransactionToSign)

	return err
}

type NullableTransactionToSign struct {
	value *TransactionToSign
	isSet bool
}

func (v NullableTransactionToSign) Get() *TransactionToSign {
	return v.value
}

func (v *NullableTransactionToSign) Set(val *TransactionToSign) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionToSign) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionToSign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionToSign(val *TransactionToSign) *NullableTransactionToSign {
	return &NullableTransactionToSign{value: val, isSet: true}
}

func (v NullableTransactionToSign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionToSign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
