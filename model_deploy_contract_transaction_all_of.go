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

// checks if the DeployContractTransactionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeployContractTransactionAllOf{}

// DeployContractTransactionAllOf struct for DeployContractTransactionAllOf
type DeployContractTransactionAllOf struct {
	DeployAt *string `json:"deployAt,omitempty"`
	// A label.
	Label *string `json:"label,omitempty"`
}

// NewDeployContractTransactionAllOf instantiates a new DeployContractTransactionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployContractTransactionAllOf() *DeployContractTransactionAllOf {
	this := DeployContractTransactionAllOf{}
	return &this
}

// NewDeployContractTransactionAllOfWithDefaults instantiates a new DeployContractTransactionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployContractTransactionAllOfWithDefaults() *DeployContractTransactionAllOf {
	this := DeployContractTransactionAllOf{}
	return &this
}

// GetDeployAt returns the DeployAt field value if set, zero value otherwise.
func (o *DeployContractTransactionAllOf) GetDeployAt() string {
	if o == nil || IsNil(o.DeployAt) {
		var ret string
		return ret
	}
	return *o.DeployAt
}

// GetDeployAtOk returns a tuple with the DeployAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployContractTransactionAllOf) GetDeployAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeployAt) {
		return nil, false
	}
	return o.DeployAt, true
}

// HasDeployAt returns a boolean if a field has been set.
func (o *DeployContractTransactionAllOf) HasDeployAt() bool {
	if o != nil && !IsNil(o.DeployAt) {
		return true
	}

	return false
}

// SetDeployAt gets a reference to the given string and assigns it to the DeployAt field.
func (o *DeployContractTransactionAllOf) SetDeployAt(v string) {
	o.DeployAt = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DeployContractTransactionAllOf) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployContractTransactionAllOf) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DeployContractTransactionAllOf) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DeployContractTransactionAllOf) SetLabel(v string) {
	o.Label = &v
}

func (o DeployContractTransactionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeployContractTransactionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeployAt) {
		toSerialize["deployAt"] = o.DeployAt
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableDeployContractTransactionAllOf struct {
	value *DeployContractTransactionAllOf
	isSet bool
}

func (v NullableDeployContractTransactionAllOf) Get() *DeployContractTransactionAllOf {
	return v.value
}

func (v *NullableDeployContractTransactionAllOf) Set(val *DeployContractTransactionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployContractTransactionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployContractTransactionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployContractTransactionAllOf(val *DeployContractTransactionAllOf) *NullableDeployContractTransactionAllOf {
	return &NullableDeployContractTransactionAllOf{value: val, isSet: true}
}

func (v NullableDeployContractTransactionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployContractTransactionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
