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

// checks if the AddKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddKey{}

// AddKey Add key request data.
type AddKey struct {
	// The Application ID that will be accessing the Key Vault.
	ClientID string `json:"clientID"`
	// The name of the key.
	KeyName string `json:"keyName"`
	// The version of the key.
	KeyVersion string `json:"keyVersion"`
	// The name given to the vault your key is stored in.
	VaultName string `json:"vaultName"`
}

// NewAddKey instantiates a new AddKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddKey(clientID string, keyName string, keyVersion string, vaultName string) *AddKey {
	this := AddKey{}
	this.ClientID = clientID
	this.KeyName = keyName
	this.KeyVersion = keyVersion
	this.VaultName = vaultName
	return &this
}

// NewAddKeyWithDefaults instantiates a new AddKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddKeyWithDefaults() *AddKey {
	this := AddKey{}
	return &this
}

// GetClientID returns the ClientID field value
func (o *AddKey) GetClientID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientID
}

// GetClientIDOk returns a tuple with the ClientID field value
// and a boolean to check if the value has been set.
func (o *AddKey) GetClientIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientID, true
}

// SetClientID sets field value
func (o *AddKey) SetClientID(v string) {
	o.ClientID = v
}

// GetKeyName returns the KeyName field value
func (o *AddKey) GetKeyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *AddKey) GetKeyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *AddKey) SetKeyName(v string) {
	o.KeyName = v
}

// GetKeyVersion returns the KeyVersion field value
func (o *AddKey) GetKeyVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyVersion
}

// GetKeyVersionOk returns a tuple with the KeyVersion field value
// and a boolean to check if the value has been set.
func (o *AddKey) GetKeyVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyVersion, true
}

// SetKeyVersion sets field value
func (o *AddKey) SetKeyVersion(v string) {
	o.KeyVersion = v
}

// GetVaultName returns the VaultName field value
func (o *AddKey) GetVaultName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultName
}

// GetVaultNameOk returns a tuple with the VaultName field value
// and a boolean to check if the value has been set.
func (o *AddKey) GetVaultNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultName, true
}

// SetVaultName sets field value
func (o *AddKey) SetVaultName(v string) {
	o.VaultName = v
}

func (o AddKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientID"] = o.ClientID
	toSerialize["keyName"] = o.KeyName
	toSerialize["keyVersion"] = o.KeyVersion
	toSerialize["vaultName"] = o.VaultName
	return toSerialize, nil
}

type NullableAddKey struct {
	value *AddKey
	isSet bool
}

func (v NullableAddKey) Get() *AddKey {
	return v.value
}

func (v *NullableAddKey) Set(val *AddKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAddKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAddKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddKey(val *AddKey) *NullableAddKey {
	return &NullableAddKey{value: val, isSet: true}
}

func (v NullableAddKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
