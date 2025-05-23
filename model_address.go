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

// checks if the Address type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Address{}

// Address An address details.
type Address struct {
	// An alias to easily identify and reference addresses.
	Alias string `json:"alias" validate:"regexp=^[a-z0-9_-]+$"`
	// An ethereum address.
	Address string   `json:"address" validate:"regexp=^0[xX][a-fA-F0-9]{40}$"`
	Balance *string  `json:"balance,omitempty"`
	Chain   string   `json:"chain"`
	Modules []string `json:"modules"`
	// The next transaction nonce for this address (obtained from the blockchain node).
	Nonce *int64 `json:"nonce,omitempty"`
	// The next transaction nonce for this address when using the nonce management feature.
	LocalNonce     *int64             `json:"localNonce,omitempty"`
	CodeAt         *string            `json:"codeAt,omitempty"`
	Contracts      []ContractMetadata `json:"contracts"`
	ContractLookup []ContractLookup   `json:"contractLookup,omitempty"`
}

type _Address Address

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress(alias string, address string, chain string, modules []string, contracts []ContractMetadata) *Address {
	this := Address{}
	this.Alias = alias
	this.Address = address
	this.Chain = chain
	this.Modules = modules
	this.Contracts = contracts
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetAlias returns the Alias field value
func (o *Address) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *Address) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *Address) SetAlias(v string) {
	o.Alias = v
}

// GetAddress returns the Address field value
func (o *Address) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Address) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *Address) SetAddress(v string) {
	o.Address = v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *Address) GetBalance() string {
	if o == nil || IsNil(o.Balance) {
		var ret string
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *Address) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given string and assigns it to the Balance field.
func (o *Address) SetBalance(v string) {
	o.Balance = &v
}

// GetChain returns the Chain field value
func (o *Address) GetChain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *Address) GetChainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *Address) SetChain(v string) {
	o.Chain = v
}

// GetModules returns the Modules field value
func (o *Address) GetModules() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value
// and a boolean to check if the value has been set.
func (o *Address) GetModulesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modules, true
}

// SetModules sets field value
func (o *Address) SetModules(v []string) {
	o.Modules = v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *Address) GetNonce() int64 {
	if o == nil || IsNil(o.Nonce) {
		var ret int64
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetNonceOk() (*int64, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *Address) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given int64 and assigns it to the Nonce field.
func (o *Address) SetNonce(v int64) {
	o.Nonce = &v
}

// GetLocalNonce returns the LocalNonce field value if set, zero value otherwise.
func (o *Address) GetLocalNonce() int64 {
	if o == nil || IsNil(o.LocalNonce) {
		var ret int64
		return ret
	}
	return *o.LocalNonce
}

// GetLocalNonceOk returns a tuple with the LocalNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetLocalNonceOk() (*int64, bool) {
	if o == nil || IsNil(o.LocalNonce) {
		return nil, false
	}
	return o.LocalNonce, true
}

// HasLocalNonce returns a boolean if a field has been set.
func (o *Address) HasLocalNonce() bool {
	if o != nil && !IsNil(o.LocalNonce) {
		return true
	}

	return false
}

// SetLocalNonce gets a reference to the given int64 and assigns it to the LocalNonce field.
func (o *Address) SetLocalNonce(v int64) {
	o.LocalNonce = &v
}

// GetCodeAt returns the CodeAt field value if set, zero value otherwise.
func (o *Address) GetCodeAt() string {
	if o == nil || IsNil(o.CodeAt) {
		var ret string
		return ret
	}
	return *o.CodeAt
}

// GetCodeAtOk returns a tuple with the CodeAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCodeAtOk() (*string, bool) {
	if o == nil || IsNil(o.CodeAt) {
		return nil, false
	}
	return o.CodeAt, true
}

// HasCodeAt returns a boolean if a field has been set.
func (o *Address) HasCodeAt() bool {
	if o != nil && !IsNil(o.CodeAt) {
		return true
	}

	return false
}

// SetCodeAt gets a reference to the given string and assigns it to the CodeAt field.
func (o *Address) SetCodeAt(v string) {
	o.CodeAt = &v
}

// GetContracts returns the Contracts field value
func (o *Address) GetContracts() []ContractMetadata {
	if o == nil {
		var ret []ContractMetadata
		return ret
	}

	return o.Contracts
}

// GetContractsOk returns a tuple with the Contracts field value
// and a boolean to check if the value has been set.
func (o *Address) GetContractsOk() ([]ContractMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contracts, true
}

// SetContracts sets field value
func (o *Address) SetContracts(v []ContractMetadata) {
	o.Contracts = v
}

// GetContractLookup returns the ContractLookup field value if set, zero value otherwise.
func (o *Address) GetContractLookup() []ContractLookup {
	if o == nil || IsNil(o.ContractLookup) {
		var ret []ContractLookup
		return ret
	}
	return o.ContractLookup
}

// GetContractLookupOk returns a tuple with the ContractLookup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetContractLookupOk() ([]ContractLookup, bool) {
	if o == nil || IsNil(o.ContractLookup) {
		return nil, false
	}
	return o.ContractLookup, true
}

// HasContractLookup returns a boolean if a field has been set.
func (o *Address) HasContractLookup() bool {
	if o != nil && !IsNil(o.ContractLookup) {
		return true
	}

	return false
}

// SetContractLookup gets a reference to the given []ContractLookup and assigns it to the ContractLookup field.
func (o *Address) SetContractLookup(v []ContractLookup) {
	o.ContractLookup = v
}

func (o Address) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alias"] = o.Alias
	toSerialize["address"] = o.Address
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	toSerialize["chain"] = o.Chain
	toSerialize["modules"] = o.Modules
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.LocalNonce) {
		toSerialize["localNonce"] = o.LocalNonce
	}
	if !IsNil(o.CodeAt) {
		toSerialize["codeAt"] = o.CodeAt
	}
	toSerialize["contracts"] = o.Contracts
	if !IsNil(o.ContractLookup) {
		toSerialize["contractLookup"] = o.ContractLookup
	}
	return toSerialize, nil
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
