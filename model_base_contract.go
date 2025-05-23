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

// checks if the BaseContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseContract{}

// BaseContract A contract.
type BaseContract struct {
	// An alias to easily identify and reference the entity in subsequent requests.
	Label string `json:"label" validate:"regexp=^[a-z0-9_-]+$"`
	// The name of the contract.
	ContractName string `json:"contractName" validate:"regexp=^[^\\"#$%&''()*+,\\/:;<>?[\\\\\\\\\\\\]^\\\\x60{}~]*$"`
	// The contract version.
	Version string `json:"version" validate:"regexp=^[^\\"#$%&''()*+,\\/:;<>?[\\\\\\\\\\\\]^\\\\x60{}~]*$"`
	// The smart-contract bytecode.
	Bin *string `json:"bin,omitempty" validate:"regexp=^(0x[0-9a-f]*|0X[0-9A-F]*)$"`
	// The contract raw ABI JSON string.
	RawAbi string `json:"rawAbi"`
	// The user documentation JSON string.
	UserDoc *string `json:"userDoc,omitempty"`
	// The developer documentation JSON string.
	DeveloperDoc *string `json:"developerDoc,omitempty"`
	// The contract metadata JSON string.
	Metadata   *string `json:"metadata,omitempty"`
	IsFavorite *bool   `json:"isFavorite,omitempty"`
}

type _BaseContract BaseContract

// NewBaseContract instantiates a new BaseContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseContract(label string, contractName string, version string, rawAbi string) *BaseContract {
	this := BaseContract{}
	this.Label = label
	this.ContractName = contractName
	this.Version = version
	this.RawAbi = rawAbi
	return &this
}

// NewBaseContractWithDefaults instantiates a new BaseContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseContractWithDefaults() *BaseContract {
	this := BaseContract{}
	return &this
}

// GetLabel returns the Label field value
func (o *BaseContract) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *BaseContract) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *BaseContract) SetLabel(v string) {
	o.Label = v
}

// GetContractName returns the ContractName field value
func (o *BaseContract) GetContractName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value
// and a boolean to check if the value has been set.
func (o *BaseContract) GetContractNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractName, true
}

// SetContractName sets field value
func (o *BaseContract) SetContractName(v string) {
	o.ContractName = v
}

// GetVersion returns the Version field value
func (o *BaseContract) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *BaseContract) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *BaseContract) SetVersion(v string) {
	o.Version = v
}

// GetBin returns the Bin field value if set, zero value otherwise.
func (o *BaseContract) GetBin() string {
	if o == nil || IsNil(o.Bin) {
		var ret string
		return ret
	}
	return *o.Bin
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContract) GetBinOk() (*string, bool) {
	if o == nil || IsNil(o.Bin) {
		return nil, false
	}
	return o.Bin, true
}

// HasBin returns a boolean if a field has been set.
func (o *BaseContract) HasBin() bool {
	if o != nil && !IsNil(o.Bin) {
		return true
	}

	return false
}

// SetBin gets a reference to the given string and assigns it to the Bin field.
func (o *BaseContract) SetBin(v string) {
	o.Bin = &v
}

// GetRawAbi returns the RawAbi field value
func (o *BaseContract) GetRawAbi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RawAbi
}

// GetRawAbiOk returns a tuple with the RawAbi field value
// and a boolean to check if the value has been set.
func (o *BaseContract) GetRawAbiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawAbi, true
}

// SetRawAbi sets field value
func (o *BaseContract) SetRawAbi(v string) {
	o.RawAbi = v
}

// GetUserDoc returns the UserDoc field value if set, zero value otherwise.
func (o *BaseContract) GetUserDoc() string {
	if o == nil || IsNil(o.UserDoc) {
		var ret string
		return ret
	}
	return *o.UserDoc
}

// GetUserDocOk returns a tuple with the UserDoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContract) GetUserDocOk() (*string, bool) {
	if o == nil || IsNil(o.UserDoc) {
		return nil, false
	}
	return o.UserDoc, true
}

// HasUserDoc returns a boolean if a field has been set.
func (o *BaseContract) HasUserDoc() bool {
	if o != nil && !IsNil(o.UserDoc) {
		return true
	}

	return false
}

// SetUserDoc gets a reference to the given string and assigns it to the UserDoc field.
func (o *BaseContract) SetUserDoc(v string) {
	o.UserDoc = &v
}

// GetDeveloperDoc returns the DeveloperDoc field value if set, zero value otherwise.
func (o *BaseContract) GetDeveloperDoc() string {
	if o == nil || IsNil(o.DeveloperDoc) {
		var ret string
		return ret
	}
	return *o.DeveloperDoc
}

// GetDeveloperDocOk returns a tuple with the DeveloperDoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContract) GetDeveloperDocOk() (*string, bool) {
	if o == nil || IsNil(o.DeveloperDoc) {
		return nil, false
	}
	return o.DeveloperDoc, true
}

// HasDeveloperDoc returns a boolean if a field has been set.
func (o *BaseContract) HasDeveloperDoc() bool {
	if o != nil && !IsNil(o.DeveloperDoc) {
		return true
	}

	return false
}

// SetDeveloperDoc gets a reference to the given string and assigns it to the DeveloperDoc field.
func (o *BaseContract) SetDeveloperDoc(v string) {
	o.DeveloperDoc = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BaseContract) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContract) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BaseContract) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *BaseContract) SetMetadata(v string) {
	o.Metadata = &v
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise.
func (o *BaseContract) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite) {
		var ret bool
		return ret
	}
	return *o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContract) GetIsFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFavorite) {
		return nil, false
	}
	return o.IsFavorite, true
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *BaseContract) HasIsFavorite() bool {
	if o != nil && !IsNil(o.IsFavorite) {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given bool and assigns it to the IsFavorite field.
func (o *BaseContract) SetIsFavorite(v bool) {
	o.IsFavorite = &v
}

func (o BaseContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["contractName"] = o.ContractName
	toSerialize["version"] = o.Version
	if !IsNil(o.Bin) {
		toSerialize["bin"] = o.Bin
	}
	toSerialize["rawAbi"] = o.RawAbi
	if !IsNil(o.UserDoc) {
		toSerialize["userDoc"] = o.UserDoc
	}
	if !IsNil(o.DeveloperDoc) {
		toSerialize["developerDoc"] = o.DeveloperDoc
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.IsFavorite) {
		toSerialize["isFavorite"] = o.IsFavorite
	}
	return toSerialize, nil
}

type NullableBaseContract struct {
	value *BaseContract
	isSet bool
}

func (v NullableBaseContract) Get() *BaseContract {
	return v.value
}

func (v *NullableBaseContract) Set(val *BaseContract) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseContract) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseContract(val *BaseContract) *NullableBaseContract {
	return &NullableBaseContract{value: val, isSet: true}
}

func (v NullableBaseContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
