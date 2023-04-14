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

// checks if the SignedTransactionSubmission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignedTransactionSubmission{}

// SignedTransactionSubmission The object used to receive a pre-signed raw transaction.
type SignedTransactionSubmission struct {
	// The pre-signed raw transaction.
	SignedTx string `json:"signedTx"`
}

// NewSignedTransactionSubmission instantiates a new SignedTransactionSubmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignedTransactionSubmission(signedTx string) *SignedTransactionSubmission {
	this := SignedTransactionSubmission{}
	this.SignedTx = signedTx
	return &this
}

// NewSignedTransactionSubmissionWithDefaults instantiates a new SignedTransactionSubmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignedTransactionSubmissionWithDefaults() *SignedTransactionSubmission {
	this := SignedTransactionSubmission{}
	return &this
}

// GetSignedTx returns the SignedTx field value
func (o *SignedTransactionSubmission) GetSignedTx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignedTx
}

// GetSignedTxOk returns a tuple with the SignedTx field value
// and a boolean to check if the value has been set.
func (o *SignedTransactionSubmission) GetSignedTxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedTx, true
}

// SetSignedTx sets field value
func (o *SignedTransactionSubmission) SetSignedTx(v string) {
	o.SignedTx = v
}

func (o SignedTransactionSubmission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignedTransactionSubmission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signedTx"] = o.SignedTx
	return toSerialize, nil
}

type NullableSignedTransactionSubmission struct {
	value *SignedTransactionSubmission
	isSet bool
}

func (v NullableSignedTransactionSubmission) Get() *SignedTransactionSubmission {
	return v.value
}

func (v *NullableSignedTransactionSubmission) Set(val *SignedTransactionSubmission) {
	v.value = val
	v.isSet = true
}

func (v NullableSignedTransactionSubmission) IsSet() bool {
	return v.isSet
}

func (v *NullableSignedTransactionSubmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignedTransactionSubmission(val *SignedTransactionSubmission) *NullableSignedTransactionSubmission {
	return &NullableSignedTransactionSubmission{value: val, isSet: true}
}

func (v NullableSignedTransactionSubmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignedTransactionSubmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
