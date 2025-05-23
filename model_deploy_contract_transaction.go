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

// checks if the DeployContractTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeployContractTransaction{}

// DeployContractTransaction The transaction returned when you deploy a contracts.
type DeployContractTransaction struct {
	Tx        TransactionToSignTx `json:"tx"`
	Submitted bool                `json:"submitted"`
	DeployAt  *string             `json:"deployAt,omitempty"`
	// An alias to easily identify and reference the entity in subsequent requests.
	Label *string `json:"label,omitempty" validate:"regexp=^[a-z0-9_-]+$"`
}

type _DeployContractTransaction DeployContractTransaction

// NewDeployContractTransaction instantiates a new DeployContractTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployContractTransaction(tx TransactionToSignTx, submitted bool) *DeployContractTransaction {
	this := DeployContractTransaction{}
	this.Tx = tx
	this.Submitted = submitted
	return &this
}

// NewDeployContractTransactionWithDefaults instantiates a new DeployContractTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployContractTransactionWithDefaults() *DeployContractTransaction {
	this := DeployContractTransaction{}
	return &this
}

// GetTx returns the Tx field value
func (o *DeployContractTransaction) GetTx() TransactionToSignTx {
	if o == nil {
		var ret TransactionToSignTx
		return ret
	}

	return o.Tx
}

// GetTxOk returns a tuple with the Tx field value
// and a boolean to check if the value has been set.
func (o *DeployContractTransaction) GetTxOk() (*TransactionToSignTx, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tx, true
}

// SetTx sets field value
func (o *DeployContractTransaction) SetTx(v TransactionToSignTx) {
	o.Tx = v
}

// GetSubmitted returns the Submitted field value
func (o *DeployContractTransaction) GetSubmitted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Submitted
}

// GetSubmittedOk returns a tuple with the Submitted field value
// and a boolean to check if the value has been set.
func (o *DeployContractTransaction) GetSubmittedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Submitted, true
}

// SetSubmitted sets field value
func (o *DeployContractTransaction) SetSubmitted(v bool) {
	o.Submitted = v
}

// GetDeployAt returns the DeployAt field value if set, zero value otherwise.
func (o *DeployContractTransaction) GetDeployAt() string {
	if o == nil || IsNil(o.DeployAt) {
		var ret string
		return ret
	}
	return *o.DeployAt
}

// GetDeployAtOk returns a tuple with the DeployAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployContractTransaction) GetDeployAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeployAt) {
		return nil, false
	}
	return o.DeployAt, true
}

// HasDeployAt returns a boolean if a field has been set.
func (o *DeployContractTransaction) HasDeployAt() bool {
	if o != nil && !IsNil(o.DeployAt) {
		return true
	}

	return false
}

// SetDeployAt gets a reference to the given string and assigns it to the DeployAt field.
func (o *DeployContractTransaction) SetDeployAt(v string) {
	o.DeployAt = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DeployContractTransaction) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployContractTransaction) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DeployContractTransaction) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DeployContractTransaction) SetLabel(v string) {
	o.Label = &v
}

func (o DeployContractTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx"] = o.Tx
	toSerialize["submitted"] = o.Submitted
	if !IsNil(o.DeployAt) {
		toSerialize["deployAt"] = o.DeployAt
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableDeployContractTransaction struct {
	value *DeployContractTransaction
	isSet bool
}

func (v NullableDeployContractTransaction) Get() *DeployContractTransaction {
	return v.value
}

func (v *NullableDeployContractTransaction) Set(val *DeployContractTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployContractTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployContractTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployContractTransaction(val *DeployContractTransaction) *NullableDeployContractTransaction {
	return &NullableDeployContractTransaction{value: val, isSet: true}
}

func (v NullableDeployContractTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployContractTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
