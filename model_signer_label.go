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

// checks if the SignerLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignerLabel{}

// SignerLabel A signer label.
type SignerLabel struct {
	// The label of the signer.
	Label string `json:"label"`
}

type _SignerLabel SignerLabel

// NewSignerLabel instantiates a new SignerLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignerLabel(label string) *SignerLabel {
	this := SignerLabel{}
	this.Label = label
	return &this
}

// NewSignerLabelWithDefaults instantiates a new SignerLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignerLabelWithDefaults() *SignerLabel {
	this := SignerLabel{}
	return &this
}

// GetLabel returns the Label field value
func (o *SignerLabel) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *SignerLabel) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *SignerLabel) SetLabel(v string) {
	o.Label = v
}

func (o SignerLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignerLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

func (o *SignerLabel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
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

	varSignerLabel := _SignerLabel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSignerLabel)

	if err != nil {
		return err
	}

	*o = SignerLabel(varSignerLabel)

	return err
}

type NullableSignerLabel struct {
	value *SignerLabel
	isSet bool
}

func (v NullableSignerLabel) Get() *SignerLabel {
	return v.value
}

func (v *NullableSignerLabel) Set(val *SignerLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableSignerLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableSignerLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignerLabel(val *SignerLabel) *NullableSignerLabel {
	return &NullableSignerLabel{value: val, isSet: true}
}

func (v NullableSignerLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignerLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
