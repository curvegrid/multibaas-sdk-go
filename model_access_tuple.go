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

// checks if the AccessTuple type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTuple{}

// AccessTuple An access tuple representing an address and its storage keys.
type AccessTuple struct {
	// An ethereum address.
	Address     NullableString `json:"address" validate:"regexp=^0[xX][a-fA-F0-9]{40}$"`
	StorageKeys []string       `json:"storageKeys"`
}

type _AccessTuple AccessTuple

// NewAccessTuple instantiates a new AccessTuple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTuple(address NullableString, storageKeys []string) *AccessTuple {
	this := AccessTuple{}
	this.Address = address
	this.StorageKeys = storageKeys
	return &this
}

// NewAccessTupleWithDefaults instantiates a new AccessTuple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTupleWithDefaults() *AccessTuple {
	this := AccessTuple{}
	return &this
}

// GetAddress returns the Address field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccessTuple) GetAddress() string {
	if o == nil || o.Address.Get() == nil {
		var ret string
		return ret
	}

	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessTuple) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// SetAddress sets field value
func (o *AccessTuple) SetAddress(v string) {
	o.Address.Set(&v)
}

// GetStorageKeys returns the StorageKeys field value
func (o *AccessTuple) GetStorageKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.StorageKeys
}

// GetStorageKeysOk returns a tuple with the StorageKeys field value
// and a boolean to check if the value has been set.
func (o *AccessTuple) GetStorageKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageKeys, true
}

// SetStorageKeys sets field value
func (o *AccessTuple) SetStorageKeys(v []string) {
	o.StorageKeys = v
}

func (o AccessTuple) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTuple) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address.Get()
	toSerialize["storageKeys"] = o.StorageKeys
	return toSerialize, nil
}

func (o *AccessTuple) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"storageKeys",
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

	varAccessTuple := _AccessTuple{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessTuple)

	if err != nil {
		return err
	}

	*o = AccessTuple(varAccessTuple)

	return err
}

type NullableAccessTuple struct {
	value *AccessTuple
	isSet bool
}

func (v NullableAccessTuple) Get() *AccessTuple {
	return v.value
}

func (v *NullableAccessTuple) Set(val *AccessTuple) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTuple) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTuple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTuple(val *AccessTuple) *NullableAccessTuple {
	return &NullableAccessTuple{value: val, isSet: true}
}

func (v NullableAccessTuple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTuple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
