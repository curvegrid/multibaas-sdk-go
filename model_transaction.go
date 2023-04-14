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

// checks if the Transaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transaction{}

// Transaction A transaction from the Ethereum Blockchain.
type Transaction struct {
	// A hex string.
	Type string `json:"type"`
	// A hex string.
	Nonce string `json:"nonce"`
	// A hex string or null.
	GasPrice NullableString `json:"gasPrice,omitempty"`
	// A hex string or null.
	MaxFeePerGas NullableString `json:"maxFeePerGas,omitempty"`
	// A hex string or null.
	MaxPriorityFeePerGas NullableString `json:"maxPriorityFeePerGas,omitempty"`
	// A hex string.
	Gas string `json:"gas"`
	// An ethereum address.
	To NullableString `json:"to"`
	// A hex string or null.
	Value NullableString `json:"value"`
	// A hex string.
	Input string `json:"input"`
	// A hex string.
	V string `json:"v"`
	// A hex string.
	R string `json:"r"`
	// A hex string.
	S string `json:"s"`
	// A hex string.
	ChainId *string `json:"chainId,omitempty"`
	// The keccak256 hash as a hex string of 256 bits.
	Hash string `json:"hash"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(type_ string, nonce string, gas string, to NullableString, value NullableString, input string, v string, r string, s string, hash string) *Transaction {
	this := Transaction{}
	this.Type = type_
	this.Nonce = nonce
	this.Gas = gas
	this.To = to
	this.Value = value
	this.Input = input
	this.V = v
	this.R = r
	this.S = s
	this.Hash = hash
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetType returns the Type field value
func (o *Transaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Transaction) SetType(v string) {
	o.Type = v
}

// GetNonce returns the Nonce field value
func (o *Transaction) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetNonceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *Transaction) SetNonce(v string) {
	o.Nonce = v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetGasPrice() string {
	if o == nil || IsNil(o.GasPrice.Get()) {
		var ret string
		return ret
	}
	return *o.GasPrice.Get()
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetGasPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasPrice.Get(), o.GasPrice.IsSet()
}

// HasGasPrice returns a boolean if a field has been set.
func (o *Transaction) HasGasPrice() bool {
	if o != nil && o.GasPrice.IsSet() {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given NullableString and assigns it to the GasPrice field.
func (o *Transaction) SetGasPrice(v string) {
	o.GasPrice.Set(&v)
}

// SetGasPriceNil sets the value for GasPrice to be an explicit nil
func (o *Transaction) SetGasPriceNil() {
	o.GasPrice.Set(nil)
}

// UnsetGasPrice ensures that no value is present for GasPrice, not even an explicit nil
func (o *Transaction) UnsetGasPrice() {
	o.GasPrice.Unset()
}

// GetMaxFeePerGas returns the MaxFeePerGas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetMaxFeePerGas() string {
	if o == nil || IsNil(o.MaxFeePerGas.Get()) {
		var ret string
		return ret
	}
	return *o.MaxFeePerGas.Get()
}

// GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetMaxFeePerGasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxFeePerGas.Get(), o.MaxFeePerGas.IsSet()
}

// HasMaxFeePerGas returns a boolean if a field has been set.
func (o *Transaction) HasMaxFeePerGas() bool {
	if o != nil && o.MaxFeePerGas.IsSet() {
		return true
	}

	return false
}

// SetMaxFeePerGas gets a reference to the given NullableString and assigns it to the MaxFeePerGas field.
func (o *Transaction) SetMaxFeePerGas(v string) {
	o.MaxFeePerGas.Set(&v)
}

// SetMaxFeePerGasNil sets the value for MaxFeePerGas to be an explicit nil
func (o *Transaction) SetMaxFeePerGasNil() {
	o.MaxFeePerGas.Set(nil)
}

// UnsetMaxFeePerGas ensures that no value is present for MaxFeePerGas, not even an explicit nil
func (o *Transaction) UnsetMaxFeePerGas() {
	o.MaxFeePerGas.Unset()
}

// GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetMaxPriorityFeePerGas() string {
	if o == nil || IsNil(o.MaxPriorityFeePerGas.Get()) {
		var ret string
		return ret
	}
	return *o.MaxPriorityFeePerGas.Get()
}

// GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetMaxPriorityFeePerGasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxPriorityFeePerGas.Get(), o.MaxPriorityFeePerGas.IsSet()
}

// HasMaxPriorityFeePerGas returns a boolean if a field has been set.
func (o *Transaction) HasMaxPriorityFeePerGas() bool {
	if o != nil && o.MaxPriorityFeePerGas.IsSet() {
		return true
	}

	return false
}

// SetMaxPriorityFeePerGas gets a reference to the given NullableString and assigns it to the MaxPriorityFeePerGas field.
func (o *Transaction) SetMaxPriorityFeePerGas(v string) {
	o.MaxPriorityFeePerGas.Set(&v)
}

// SetMaxPriorityFeePerGasNil sets the value for MaxPriorityFeePerGas to be an explicit nil
func (o *Transaction) SetMaxPriorityFeePerGasNil() {
	o.MaxPriorityFeePerGas.Set(nil)
}

// UnsetMaxPriorityFeePerGas ensures that no value is present for MaxPriorityFeePerGas, not even an explicit nil
func (o *Transaction) UnsetMaxPriorityFeePerGas() {
	o.MaxPriorityFeePerGas.Unset()
}

// GetGas returns the Gas field value
func (o *Transaction) GetGas() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gas
}

// GetGasOk returns a tuple with the Gas field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetGasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gas, true
}

// SetGas sets field value
func (o *Transaction) SetGas(v string) {
	o.Gas = v
}

// GetTo returns the To field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetTo() string {
	if o == nil || o.To.Get() == nil {
		var ret string
		return ret
	}

	return *o.To.Get()
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.To.Get(), o.To.IsSet()
}

// SetTo sets field value
func (o *Transaction) SetTo(v string) {
	o.To.Set(&v)
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transaction) GetValue() string {
	if o == nil || o.Value.Get() == nil {
		var ret string
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *Transaction) SetValue(v string) {
	o.Value.Set(&v)
}

// GetInput returns the Input field value
func (o *Transaction) GetInput() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetInputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *Transaction) SetInput(v string) {
	o.Input = v
}

// GetV returns the V field value
func (o *Transaction) GetV() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.V
}

// GetVOk returns a tuple with the V field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetVOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.V, true
}

// SetV sets field value
func (o *Transaction) SetV(v string) {
	o.V = v
}

// GetR returns the R field value
func (o *Transaction) GetR() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.R
}

// GetROk returns a tuple with the R field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetROk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.R, true
}

// SetR sets field value
func (o *Transaction) SetR(v string) {
	o.R = v
}

// GetS returns the S field value
func (o *Transaction) GetS() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.S
}

// GetSOk returns a tuple with the S field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetSOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.S, true
}

// SetS sets field value
func (o *Transaction) SetS(v string) {
	o.S = v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *Transaction) GetChainId() string {
	if o == nil || IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetChainIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *Transaction) HasChainId() bool {
	if o != nil && !IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *Transaction) SetChainId(v string) {
	o.ChainId = &v
}

// GetHash returns the Hash field value
func (o *Transaction) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *Transaction) SetHash(v string) {
	o.Hash = v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["nonce"] = o.Nonce
	if o.GasPrice.IsSet() {
		toSerialize["gasPrice"] = o.GasPrice.Get()
	}
	if o.MaxFeePerGas.IsSet() {
		toSerialize["maxFeePerGas"] = o.MaxFeePerGas.Get()
	}
	if o.MaxPriorityFeePerGas.IsSet() {
		toSerialize["maxPriorityFeePerGas"] = o.MaxPriorityFeePerGas.Get()
	}
	toSerialize["gas"] = o.Gas
	toSerialize["to"] = o.To.Get()
	toSerialize["value"] = o.Value.Get()
	toSerialize["input"] = o.Input
	toSerialize["v"] = o.V
	toSerialize["r"] = o.R
	toSerialize["s"] = o.S
	if !IsNil(o.ChainId) {
		toSerialize["chainId"] = o.ChainId
	}
	toSerialize["hash"] = o.Hash
	return toSerialize, nil
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
