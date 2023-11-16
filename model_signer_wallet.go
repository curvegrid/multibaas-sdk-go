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

// checks if the SignerWallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignerWallet{}

// SignerWallet A signer wallet.
type SignerWallet struct {
	// The type of the signer.
	Type string `json:"type"`
	// An ethereum address.
	Wallet string `json:"wallet"`
	// An ethereum address.
	Signer string `json:"signer"`
	// The label of the signer.
	Label string `json:"label"`
}

// NewSignerWallet instantiates a new SignerWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignerWallet(type_ string, wallet string, signer string, label string) *SignerWallet {
	this := SignerWallet{}
	this.Type = type_
	this.Wallet = wallet
	this.Signer = signer
	this.Label = label
	return &this
}

// NewSignerWalletWithDefaults instantiates a new SignerWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignerWalletWithDefaults() *SignerWallet {
	this := SignerWallet{}
	return &this
}

// GetType returns the Type field value
func (o *SignerWallet) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SignerWallet) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SignerWallet) SetType(v string) {
	o.Type = v
}

// GetWallet returns the Wallet field value
func (o *SignerWallet) GetWallet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value
// and a boolean to check if the value has been set.
func (o *SignerWallet) GetWalletOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wallet, true
}

// SetWallet sets field value
func (o *SignerWallet) SetWallet(v string) {
	o.Wallet = v
}

// GetSigner returns the Signer field value
func (o *SignerWallet) GetSigner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signer
}

// GetSignerOk returns a tuple with the Signer field value
// and a boolean to check if the value has been set.
func (o *SignerWallet) GetSignerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signer, true
}

// SetSigner sets field value
func (o *SignerWallet) SetSigner(v string) {
	o.Signer = v
}

// GetLabel returns the Label field value
func (o *SignerWallet) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *SignerWallet) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *SignerWallet) SetLabel(v string) {
	o.Label = v
}

func (o SignerWallet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignerWallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["wallet"] = o.Wallet
	toSerialize["signer"] = o.Signer
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

type NullableSignerWallet struct {
	value *SignerWallet
	isSet bool
}

func (v NullableSignerWallet) Get() *SignerWallet {
	return v.value
}

func (v *NullableSignerWallet) Set(val *SignerWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableSignerWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableSignerWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignerWallet(val *SignerWallet) *NullableSignerWallet {
	return &NullableSignerWallet{value: val, isSet: true}
}

func (v NullableSignerWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignerWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
