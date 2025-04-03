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

// checks if the ContractLookup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractLookup{}

// ContractLookup The contract lookup item.
type ContractLookup struct {
	// An ethereum address.
	Address string `json:"address"`
	// The name of the contract.
	Name *string `json:"name,omitempty"`
	// The contract ABI JSON string.
	Abi string `json:"abi"`
	// The smart-contract bytecode.
	Bytecode *string `json:"bytecode,omitempty"`
	// The contract's source code.
	Source *string `json:"source,omitempty"`
	// The user documentation JSON string.
	Userdoc *string `json:"userdoc,omitempty"`
	// The developer documentation JSON string.
	Devdoc *string `json:"devdoc,omitempty"`
	// Indicates whether the contract has been verified.
	Verified bool `json:"verified"`
	// The name of the service that provided the contract verification.
	VerifiedSource *string `json:"verifiedSource,omitempty"`
	// The URL to the contract's verification details on the verification service.
	VerifiedLink *string `json:"verifiedLink,omitempty"`
	// Indicates whether the contract is a proxy contract.
	Proxy bool `json:"proxy"`
}

// NewContractLookup instantiates a new ContractLookup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractLookup(address string, abi string, verified bool, proxy bool) *ContractLookup {
	this := ContractLookup{}
	this.Address = address
	this.Abi = abi
	this.Verified = verified
	this.Proxy = proxy
	return &this
}

// NewContractLookupWithDefaults instantiates a new ContractLookup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractLookupWithDefaults() *ContractLookup {
	this := ContractLookup{}
	return &this
}

// GetAddress returns the Address field value
func (o *ContractLookup) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ContractLookup) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ContractLookup) SetAddress(v string) {
	o.Address = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContractLookup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLookup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContractLookup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContractLookup) SetName(v string) {
	o.Name = &v
}

// GetAbi returns the Abi field value
func (o *ContractLookup) GetAbi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Abi
}

// GetAbiOk returns a tuple with the Abi field value
// and a boolean to check if the value has been set.
func (o *ContractLookup) GetAbiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Abi, true
}

// SetAbi sets field value
func (o *ContractLookup) SetAbi(v string) {
	o.Abi = v
}

// GetBytecode returns the Bytecode field value if set, zero value otherwise.
func (o *ContractLookup) GetBytecode() string {
	if o == nil || IsNil(o.Bytecode) {
		var ret string
		return ret
	}
	return *o.Bytecode
}

// GetBytecodeOk returns a tuple with the Bytecode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLookup) GetBytecodeOk() (*string, bool) {
	if o == nil || IsNil(o.Bytecode) {
		return nil, false
	}
	return o.Bytecode, true
}

// HasBytecode returns a boolean if a field has been set.
func (o *ContractLookup) HasBytecode() bool {
	if o != nil && !IsNil(o.Bytecode) {
		return true
	}

	return false
}

// SetBytecode gets a reference to the given string and assigns it to the Bytecode field.
func (o *ContractLookup) SetBytecode(v string) {
	o.Bytecode = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ContractLookup) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLookup) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ContractLookup) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ContractLookup) SetSource(v string) {
	o.Source = &v
}

// GetUserdoc returns the Userdoc field value if set, zero value otherwise.
func (o *ContractLookup) GetUserdoc() string {
	if o == nil || IsNil(o.Userdoc) {
		var ret string
		return ret
	}
	return *o.Userdoc
}

// GetUserdocOk returns a tuple with the Userdoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLookup) GetUserdocOk() (*string, bool) {
	if o == nil || IsNil(o.Userdoc) {
		return nil, false
	}
	return o.Userdoc, true
}

// HasUserdoc returns a boolean if a field has been set.
func (o *ContractLookup) HasUserdoc() bool {
	if o != nil && !IsNil(o.Userdoc) {
		return true
	}

	return false
}

// SetUserdoc gets a reference to the given string and assigns it to the Userdoc field.
func (o *ContractLookup) SetUserdoc(v string) {
	o.Userdoc = &v
}

// GetDevdoc returns the Devdoc field value if set, zero value otherwise.
func (o *ContractLookup) GetDevdoc() string {
	if o == nil || IsNil(o.Devdoc) {
		var ret string
		return ret
	}
	return *o.Devdoc
}

// GetDevdocOk returns a tuple with the Devdoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLookup) GetDevdocOk() (*string, bool) {
	if o == nil || IsNil(o.Devdoc) {
		return nil, false
	}
	return o.Devdoc, true
}

// HasDevdoc returns a boolean if a field has been set.
func (o *ContractLookup) HasDevdoc() bool {
	if o != nil && !IsNil(o.Devdoc) {
		return true
	}

	return false
}

// SetDevdoc gets a reference to the given string and assigns it to the Devdoc field.
func (o *ContractLookup) SetDevdoc(v string) {
	o.Devdoc = &v
}

// GetVerified returns the Verified field value
func (o *ContractLookup) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *ContractLookup) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *ContractLookup) SetVerified(v bool) {
	o.Verified = v
}

// GetVerifiedSource returns the VerifiedSource field value if set, zero value otherwise.
func (o *ContractLookup) GetVerifiedSource() string {
	if o == nil || IsNil(o.VerifiedSource) {
		var ret string
		return ret
	}
	return *o.VerifiedSource
}

// GetVerifiedSourceOk returns a tuple with the VerifiedSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLookup) GetVerifiedSourceOk() (*string, bool) {
	if o == nil || IsNil(o.VerifiedSource) {
		return nil, false
	}
	return o.VerifiedSource, true
}

// HasVerifiedSource returns a boolean if a field has been set.
func (o *ContractLookup) HasVerifiedSource() bool {
	if o != nil && !IsNil(o.VerifiedSource) {
		return true
	}

	return false
}

// SetVerifiedSource gets a reference to the given string and assigns it to the VerifiedSource field.
func (o *ContractLookup) SetVerifiedSource(v string) {
	o.VerifiedSource = &v
}

// GetVerifiedLink returns the VerifiedLink field value if set, zero value otherwise.
func (o *ContractLookup) GetVerifiedLink() string {
	if o == nil || IsNil(o.VerifiedLink) {
		var ret string
		return ret
	}
	return *o.VerifiedLink
}

// GetVerifiedLinkOk returns a tuple with the VerifiedLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLookup) GetVerifiedLinkOk() (*string, bool) {
	if o == nil || IsNil(o.VerifiedLink) {
		return nil, false
	}
	return o.VerifiedLink, true
}

// HasVerifiedLink returns a boolean if a field has been set.
func (o *ContractLookup) HasVerifiedLink() bool {
	if o != nil && !IsNil(o.VerifiedLink) {
		return true
	}

	return false
}

// SetVerifiedLink gets a reference to the given string and assigns it to the VerifiedLink field.
func (o *ContractLookup) SetVerifiedLink(v string) {
	o.VerifiedLink = &v
}

// GetProxy returns the Proxy field value
func (o *ContractLookup) GetProxy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value
// and a boolean to check if the value has been set.
func (o *ContractLookup) GetProxyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proxy, true
}

// SetProxy sets field value
func (o *ContractLookup) SetProxy(v bool) {
	o.Proxy = v
}

func (o ContractLookup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractLookup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["abi"] = o.Abi
	if !IsNil(o.Bytecode) {
		toSerialize["bytecode"] = o.Bytecode
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Userdoc) {
		toSerialize["userdoc"] = o.Userdoc
	}
	if !IsNil(o.Devdoc) {
		toSerialize["devdoc"] = o.Devdoc
	}
	toSerialize["verified"] = o.Verified
	if !IsNil(o.VerifiedSource) {
		toSerialize["verifiedSource"] = o.VerifiedSource
	}
	if !IsNil(o.VerifiedLink) {
		toSerialize["verifiedLink"] = o.VerifiedLink
	}
	toSerialize["proxy"] = o.Proxy
	return toSerialize, nil
}

type NullableContractLookup struct {
	value *ContractLookup
	isSet bool
}

func (v NullableContractLookup) Get() *ContractLookup {
	return v.value
}

func (v *NullableContractLookup) Set(val *ContractLookup) {
	v.value = val
	v.isSet = true
}

func (v NullableContractLookup) IsSet() bool {
	return v.isSet
}

func (v *NullableContractLookup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractLookup(val *ContractLookup) *NullableContractLookup {
	return &NullableContractLookup{value: val, isSet: true}
}

func (v NullableContractLookup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractLookup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
