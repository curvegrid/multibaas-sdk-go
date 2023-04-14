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

// checks if the HSMData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HSMData{}

// HSMData Response body for returning HSM Data.
type HSMData struct {
	Configuration AzureAccount `json:"configuration"`
	// An array of Azure Hardware Wallets.
	Wallets []AzureHardwareWallet `json:"wallets"`
}

// NewHSMData instantiates a new HSMData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHSMData(configuration AzureAccount, wallets []AzureHardwareWallet) *HSMData {
	this := HSMData{}
	this.Configuration = configuration
	this.Wallets = wallets
	return &this
}

// NewHSMDataWithDefaults instantiates a new HSMData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHSMDataWithDefaults() *HSMData {
	this := HSMData{}
	return &this
}

// GetConfiguration returns the Configuration field value
func (o *HSMData) GetConfiguration() AzureAccount {
	if o == nil {
		var ret AzureAccount
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *HSMData) GetConfigurationOk() (*AzureAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *HSMData) SetConfiguration(v AzureAccount) {
	o.Configuration = v
}

// GetWallets returns the Wallets field value
func (o *HSMData) GetWallets() []AzureHardwareWallet {
	if o == nil {
		var ret []AzureHardwareWallet
		return ret
	}

	return o.Wallets
}

// GetWalletsOk returns a tuple with the Wallets field value
// and a boolean to check if the value has been set.
func (o *HSMData) GetWalletsOk() ([]AzureHardwareWallet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Wallets, true
}

// SetWallets sets field value
func (o *HSMData) SetWallets(v []AzureHardwareWallet) {
	o.Wallets = v
}

func (o HSMData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HSMData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configuration"] = o.Configuration
	toSerialize["wallets"] = o.Wallets
	return toSerialize, nil
}

type NullableHSMData struct {
	value *HSMData
	isSet bool
}

func (v NullableHSMData) Get() *HSMData {
	return v.value
}

func (v *NullableHSMData) Set(val *HSMData) {
	v.value = val
	v.isSet = true
}

func (v NullableHSMData) IsSet() bool {
	return v.isSet
}

func (v *NullableHSMData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHSMData(val *HSMData) *NullableHSMData {
	return &NullableHSMData{value: val, isSet: true}
}

func (v NullableHSMData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHSMData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}