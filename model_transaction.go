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
	// A hex string or null.
	ChainId NullableString `json:"chainId,omitempty"`
	// A hex string.
	Nonce string `json:"nonce"`
	// An ethereum address.
	To NullableString `json:"to"`
	// A hex string.
	Gas string `json:"gas"`
	// A hex string or null.
	GasPrice NullableString `json:"gasPrice,omitempty"`
	// A hex string or null.
	MaxPriorityFeePerGas NullableString `json:"maxPriorityFeePerGas,omitempty"`
	// A hex string or null.
	MaxFeePerGas NullableString `json:"maxFeePerGas,omitempty"`
	// A hex string or null.
	MaxFeePerBlobGas NullableString `json:"maxFeePerBlobGas,omitempty"`
	// A hex string or null.
	Value NullableString `json:"value"`
	// A hex string.
	Input               string                 `json:"input"`
	AccessList          []AccessTuple          `json:"accessList,omitempty"`
	BlobVersionedHashes []string               `json:"blobVersionedHashes,omitempty"`
	AuthorizationList   []SetCodeAuthorization `json:"authorizationList,omitempty"`
	// A hex string.
	V string `json:"v"`
	// A hex string.
	R string `json:"r"`
	// A hex string.
	S string `json:"s"`
	// A hex string or null.
	YParity     NullableString `json:"yParity,omitempty"`
	Blobs       []string       `json:"blobs,omitempty"`
	Commitments []string       `json:"commitments,omitempty"`
	Proofs      []string       `json:"proofs,omitempty"`
	// The keccak256 hash as a hex string of 256 bits.
	Hash string `json:"hash"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(type_ string, nonce string, to NullableString, gas string, value NullableString, input string, v string, r string, s string, hash string) *Transaction {
	this := Transaction{}
	this.Type = type_
	this.Nonce = nonce
	this.To = to
	this.Gas = gas
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

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetChainId() string {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret string
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *Transaction) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableString and assigns it to the ChainId field.
func (o *Transaction) SetChainId(v string) {
	o.ChainId.Set(&v)
}

// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *Transaction) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *Transaction) UnsetChainId() {
	o.ChainId.Unset()
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

// GetMaxFeePerBlobGas returns the MaxFeePerBlobGas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetMaxFeePerBlobGas() string {
	if o == nil || IsNil(o.MaxFeePerBlobGas.Get()) {
		var ret string
		return ret
	}
	return *o.MaxFeePerBlobGas.Get()
}

// GetMaxFeePerBlobGasOk returns a tuple with the MaxFeePerBlobGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetMaxFeePerBlobGasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxFeePerBlobGas.Get(), o.MaxFeePerBlobGas.IsSet()
}

// HasMaxFeePerBlobGas returns a boolean if a field has been set.
func (o *Transaction) HasMaxFeePerBlobGas() bool {
	if o != nil && o.MaxFeePerBlobGas.IsSet() {
		return true
	}

	return false
}

// SetMaxFeePerBlobGas gets a reference to the given NullableString and assigns it to the MaxFeePerBlobGas field.
func (o *Transaction) SetMaxFeePerBlobGas(v string) {
	o.MaxFeePerBlobGas.Set(&v)
}

// SetMaxFeePerBlobGasNil sets the value for MaxFeePerBlobGas to be an explicit nil
func (o *Transaction) SetMaxFeePerBlobGasNil() {
	o.MaxFeePerBlobGas.Set(nil)
}

// UnsetMaxFeePerBlobGas ensures that no value is present for MaxFeePerBlobGas, not even an explicit nil
func (o *Transaction) UnsetMaxFeePerBlobGas() {
	o.MaxFeePerBlobGas.Unset()
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

// GetAccessList returns the AccessList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetAccessList() []AccessTuple {
	if o == nil {
		var ret []AccessTuple
		return ret
	}
	return o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetAccessListOk() ([]AccessTuple, bool) {
	if o == nil || IsNil(o.AccessList) {
		return nil, false
	}
	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *Transaction) HasAccessList() bool {
	if o != nil && IsNil(o.AccessList) {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given []AccessTuple and assigns it to the AccessList field.
func (o *Transaction) SetAccessList(v []AccessTuple) {
	o.AccessList = v
}

// GetBlobVersionedHashes returns the BlobVersionedHashes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetBlobVersionedHashes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BlobVersionedHashes
}

// GetBlobVersionedHashesOk returns a tuple with the BlobVersionedHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetBlobVersionedHashesOk() ([]string, bool) {
	if o == nil || IsNil(o.BlobVersionedHashes) {
		return nil, false
	}
	return o.BlobVersionedHashes, true
}

// HasBlobVersionedHashes returns a boolean if a field has been set.
func (o *Transaction) HasBlobVersionedHashes() bool {
	if o != nil && IsNil(o.BlobVersionedHashes) {
		return true
	}

	return false
}

// SetBlobVersionedHashes gets a reference to the given []string and assigns it to the BlobVersionedHashes field.
func (o *Transaction) SetBlobVersionedHashes(v []string) {
	o.BlobVersionedHashes = v
}

// GetAuthorizationList returns the AuthorizationList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetAuthorizationList() []SetCodeAuthorization {
	if o == nil {
		var ret []SetCodeAuthorization
		return ret
	}
	return o.AuthorizationList
}

// GetAuthorizationListOk returns a tuple with the AuthorizationList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetAuthorizationListOk() ([]SetCodeAuthorization, bool) {
	if o == nil || IsNil(o.AuthorizationList) {
		return nil, false
	}
	return o.AuthorizationList, true
}

// HasAuthorizationList returns a boolean if a field has been set.
func (o *Transaction) HasAuthorizationList() bool {
	if o != nil && IsNil(o.AuthorizationList) {
		return true
	}

	return false
}

// SetAuthorizationList gets a reference to the given []SetCodeAuthorization and assigns it to the AuthorizationList field.
func (o *Transaction) SetAuthorizationList(v []SetCodeAuthorization) {
	o.AuthorizationList = v
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

// GetYParity returns the YParity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetYParity() string {
	if o == nil || IsNil(o.YParity.Get()) {
		var ret string
		return ret
	}
	return *o.YParity.Get()
}

// GetYParityOk returns a tuple with the YParity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetYParityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.YParity.Get(), o.YParity.IsSet()
}

// HasYParity returns a boolean if a field has been set.
func (o *Transaction) HasYParity() bool {
	if o != nil && o.YParity.IsSet() {
		return true
	}

	return false
}

// SetYParity gets a reference to the given NullableString and assigns it to the YParity field.
func (o *Transaction) SetYParity(v string) {
	o.YParity.Set(&v)
}

// SetYParityNil sets the value for YParity to be an explicit nil
func (o *Transaction) SetYParityNil() {
	o.YParity.Set(nil)
}

// UnsetYParity ensures that no value is present for YParity, not even an explicit nil
func (o *Transaction) UnsetYParity() {
	o.YParity.Unset()
}

// GetBlobs returns the Blobs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetBlobs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Blobs
}

// GetBlobsOk returns a tuple with the Blobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetBlobsOk() ([]string, bool) {
	if o == nil || IsNil(o.Blobs) {
		return nil, false
	}
	return o.Blobs, true
}

// HasBlobs returns a boolean if a field has been set.
func (o *Transaction) HasBlobs() bool {
	if o != nil && IsNil(o.Blobs) {
		return true
	}

	return false
}

// SetBlobs gets a reference to the given []string and assigns it to the Blobs field.
func (o *Transaction) SetBlobs(v []string) {
	o.Blobs = v
}

// GetCommitments returns the Commitments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetCommitments() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Commitments
}

// GetCommitmentsOk returns a tuple with the Commitments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCommitmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Commitments) {
		return nil, false
	}
	return o.Commitments, true
}

// HasCommitments returns a boolean if a field has been set.
func (o *Transaction) HasCommitments() bool {
	if o != nil && IsNil(o.Commitments) {
		return true
	}

	return false
}

// SetCommitments gets a reference to the given []string and assigns it to the Commitments field.
func (o *Transaction) SetCommitments(v []string) {
	o.Commitments = v
}

// GetProofs returns the Proofs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetProofs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Proofs
}

// GetProofsOk returns a tuple with the Proofs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetProofsOk() ([]string, bool) {
	if o == nil || IsNil(o.Proofs) {
		return nil, false
	}
	return o.Proofs, true
}

// HasProofs returns a boolean if a field has been set.
func (o *Transaction) HasProofs() bool {
	if o != nil && IsNil(o.Proofs) {
		return true
	}

	return false
}

// SetProofs gets a reference to the given []string and assigns it to the Proofs field.
func (o *Transaction) SetProofs(v []string) {
	o.Proofs = v
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
	if o.ChainId.IsSet() {
		toSerialize["chainId"] = o.ChainId.Get()
	}
	toSerialize["nonce"] = o.Nonce
	toSerialize["to"] = o.To.Get()
	toSerialize["gas"] = o.Gas
	if o.GasPrice.IsSet() {
		toSerialize["gasPrice"] = o.GasPrice.Get()
	}
	if o.MaxPriorityFeePerGas.IsSet() {
		toSerialize["maxPriorityFeePerGas"] = o.MaxPriorityFeePerGas.Get()
	}
	if o.MaxFeePerGas.IsSet() {
		toSerialize["maxFeePerGas"] = o.MaxFeePerGas.Get()
	}
	if o.MaxFeePerBlobGas.IsSet() {
		toSerialize["maxFeePerBlobGas"] = o.MaxFeePerBlobGas.Get()
	}
	toSerialize["value"] = o.Value.Get()
	toSerialize["input"] = o.Input
	if o.AccessList != nil {
		toSerialize["accessList"] = o.AccessList
	}
	if o.BlobVersionedHashes != nil {
		toSerialize["blobVersionedHashes"] = o.BlobVersionedHashes
	}
	if o.AuthorizationList != nil {
		toSerialize["authorizationList"] = o.AuthorizationList
	}
	toSerialize["v"] = o.V
	toSerialize["r"] = o.R
	toSerialize["s"] = o.S
	if o.YParity.IsSet() {
		toSerialize["yParity"] = o.YParity.Get()
	}
	if o.Blobs != nil {
		toSerialize["blobs"] = o.Blobs
	}
	if o.Commitments != nil {
		toSerialize["commitments"] = o.Commitments
	}
	if o.Proofs != nil {
		toSerialize["proofs"] = o.Proofs
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
