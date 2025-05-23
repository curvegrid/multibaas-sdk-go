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

// checks if the Block type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Block{}

// Block A block in the Ethereum blockchain.
type Block struct {
	// The keccak256 hash as a hex string of 256 bits.
	Hash         string        `json:"hash" validate:"regexp=^(0x[0-9a-f]{64}|0X[0-9A-F]{64})$"`
	Difficulty   string        `json:"difficulty"`
	GasLimit     int64         `json:"gasLimit"`
	Number       string        `json:"number"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	// The keccak256 hash as a hex string of 256 bits.
	ReceiptsRoot string `json:"receiptsRoot" validate:"regexp=^(0x[0-9a-f]{64}|0X[0-9A-F]{64})$"`
	// The keccak256 hash as a hex string of 256 bits.
	ParentHash string `json:"parentHash" validate:"regexp=^(0x[0-9a-f]{64}|0X[0-9A-F]{64})$"`
	// The keccak256 hash as a hex string of 256 bits.
	Sha3Uncles string `json:"sha3Uncles" validate:"regexp=^(0x[0-9a-f]{64}|0X[0-9A-F]{64})$"`
	// An ethereum address.
	Miner string `json:"miner" validate:"regexp=^0[xX][a-fA-F0-9]{40}$"`
	// The keccak256 hash as a hex string of 256 bits.
	StateRoot string `json:"stateRoot" validate:"regexp=^(0x[0-9a-f]{64}|0X[0-9A-F]{64})$"`
	// The keccak256 hash as a hex string of 256 bits.
	TransactionsRoot string `json:"transactionsRoot" validate:"regexp=^(0x[0-9a-f]{64}|0X[0-9A-F]{64})$"`
	// A hex string.
	LogsBloom string `json:"logsBloom" validate:"regexp=^(0x[0-9a-f]*|0X[0-9A-F]*)$"`
	GasUsed   int32  `json:"gasUsed"`
	// A hex string.
	Nonce string `json:"nonce" validate:"regexp=^(0x[0-9a-f]*|0X[0-9A-F]*)$"`
	// The keccak256 hash as a hex string of 256 bits.
	MixHash       string  `json:"mixHash" validate:"regexp=^(0x[0-9a-f]{64}|0X[0-9A-F]{64})$"`
	ExtraData     string  `json:"extraData"`
	BaseFeePerGas *string `json:"baseFeePerGas,omitempty"`
}

type _Block Block

// NewBlock instantiates a new Block object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlock(hash string, difficulty string, gasLimit int64, number string, timestamp int64, transactions []Transaction, receiptsRoot string, parentHash string, sha3Uncles string, miner string, stateRoot string, transactionsRoot string, logsBloom string, gasUsed int32, nonce string, mixHash string, extraData string) *Block {
	this := Block{}
	this.Hash = hash
	this.Difficulty = difficulty
	this.GasLimit = gasLimit
	this.Number = number
	this.Timestamp = timestamp
	this.Transactions = transactions
	this.ReceiptsRoot = receiptsRoot
	this.ParentHash = parentHash
	this.Sha3Uncles = sha3Uncles
	this.Miner = miner
	this.StateRoot = stateRoot
	this.TransactionsRoot = transactionsRoot
	this.LogsBloom = logsBloom
	this.GasUsed = gasUsed
	this.Nonce = nonce
	this.MixHash = mixHash
	this.ExtraData = extraData
	return &this
}

// NewBlockWithDefaults instantiates a new Block object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockWithDefaults() *Block {
	this := Block{}
	return &this
}

// GetHash returns the Hash field value
func (o *Block) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *Block) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *Block) SetHash(v string) {
	o.Hash = v
}

// GetDifficulty returns the Difficulty field value
func (o *Block) GetDifficulty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Difficulty
}

// GetDifficultyOk returns a tuple with the Difficulty field value
// and a boolean to check if the value has been set.
func (o *Block) GetDifficultyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Difficulty, true
}

// SetDifficulty sets field value
func (o *Block) SetDifficulty(v string) {
	o.Difficulty = v
}

// GetGasLimit returns the GasLimit field value
func (o *Block) GetGasLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GasLimit
}

// GetGasLimitOk returns a tuple with the GasLimit field value
// and a boolean to check if the value has been set.
func (o *Block) GetGasLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasLimit, true
}

// SetGasLimit sets field value
func (o *Block) SetGasLimit(v int64) {
	o.GasLimit = v
}

// GetNumber returns the Number field value
func (o *Block) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *Block) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *Block) SetNumber(v string) {
	o.Number = v
}

// GetTimestamp returns the Timestamp field value
func (o *Block) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Block) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Block) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetTransactions returns the Transactions field value
func (o *Block) GetTransactions() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *Block) GetTransactionsOk() ([]Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transactions, true
}

// SetTransactions sets field value
func (o *Block) SetTransactions(v []Transaction) {
	o.Transactions = v
}

// GetReceiptsRoot returns the ReceiptsRoot field value
func (o *Block) GetReceiptsRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiptsRoot
}

// GetReceiptsRootOk returns a tuple with the ReceiptsRoot field value
// and a boolean to check if the value has been set.
func (o *Block) GetReceiptsRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiptsRoot, true
}

// SetReceiptsRoot sets field value
func (o *Block) SetReceiptsRoot(v string) {
	o.ReceiptsRoot = v
}

// GetParentHash returns the ParentHash field value
func (o *Block) GetParentHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentHash
}

// GetParentHashOk returns a tuple with the ParentHash field value
// and a boolean to check if the value has been set.
func (o *Block) GetParentHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentHash, true
}

// SetParentHash sets field value
func (o *Block) SetParentHash(v string) {
	o.ParentHash = v
}

// GetSha3Uncles returns the Sha3Uncles field value
func (o *Block) GetSha3Uncles() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha3Uncles
}

// GetSha3UnclesOk returns a tuple with the Sha3Uncles field value
// and a boolean to check if the value has been set.
func (o *Block) GetSha3UnclesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha3Uncles, true
}

// SetSha3Uncles sets field value
func (o *Block) SetSha3Uncles(v string) {
	o.Sha3Uncles = v
}

// GetMiner returns the Miner field value
func (o *Block) GetMiner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Miner
}

// GetMinerOk returns a tuple with the Miner field value
// and a boolean to check if the value has been set.
func (o *Block) GetMinerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Miner, true
}

// SetMiner sets field value
func (o *Block) SetMiner(v string) {
	o.Miner = v
}

// GetStateRoot returns the StateRoot field value
func (o *Block) GetStateRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateRoot
}

// GetStateRootOk returns a tuple with the StateRoot field value
// and a boolean to check if the value has been set.
func (o *Block) GetStateRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateRoot, true
}

// SetStateRoot sets field value
func (o *Block) SetStateRoot(v string) {
	o.StateRoot = v
}

// GetTransactionsRoot returns the TransactionsRoot field value
func (o *Block) GetTransactionsRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionsRoot
}

// GetTransactionsRootOk returns a tuple with the TransactionsRoot field value
// and a boolean to check if the value has been set.
func (o *Block) GetTransactionsRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionsRoot, true
}

// SetTransactionsRoot sets field value
func (o *Block) SetTransactionsRoot(v string) {
	o.TransactionsRoot = v
}

// GetLogsBloom returns the LogsBloom field value
func (o *Block) GetLogsBloom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogsBloom
}

// GetLogsBloomOk returns a tuple with the LogsBloom field value
// and a boolean to check if the value has been set.
func (o *Block) GetLogsBloomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogsBloom, true
}

// SetLogsBloom sets field value
func (o *Block) SetLogsBloom(v string) {
	o.LogsBloom = v
}

// GetGasUsed returns the GasUsed field value
func (o *Block) GetGasUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *Block) GetGasUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *Block) SetGasUsed(v int32) {
	o.GasUsed = v
}

// GetNonce returns the Nonce field value
func (o *Block) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *Block) GetNonceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *Block) SetNonce(v string) {
	o.Nonce = v
}

// GetMixHash returns the MixHash field value
func (o *Block) GetMixHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MixHash
}

// GetMixHashOk returns a tuple with the MixHash field value
// and a boolean to check if the value has been set.
func (o *Block) GetMixHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MixHash, true
}

// SetMixHash sets field value
func (o *Block) SetMixHash(v string) {
	o.MixHash = v
}

// GetExtraData returns the ExtraData field value
func (o *Block) GetExtraData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value
// and a boolean to check if the value has been set.
func (o *Block) GetExtraDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtraData, true
}

// SetExtraData sets field value
func (o *Block) SetExtraData(v string) {
	o.ExtraData = v
}

// GetBaseFeePerGas returns the BaseFeePerGas field value if set, zero value otherwise.
func (o *Block) GetBaseFeePerGas() string {
	if o == nil || IsNil(o.BaseFeePerGas) {
		var ret string
		return ret
	}
	return *o.BaseFeePerGas
}

// GetBaseFeePerGasOk returns a tuple with the BaseFeePerGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Block) GetBaseFeePerGasOk() (*string, bool) {
	if o == nil || IsNil(o.BaseFeePerGas) {
		return nil, false
	}
	return o.BaseFeePerGas, true
}

// HasBaseFeePerGas returns a boolean if a field has been set.
func (o *Block) HasBaseFeePerGas() bool {
	if o != nil && !IsNil(o.BaseFeePerGas) {
		return true
	}

	return false
}

// SetBaseFeePerGas gets a reference to the given string and assigns it to the BaseFeePerGas field.
func (o *Block) SetBaseFeePerGas(v string) {
	o.BaseFeePerGas = &v
}

func (o Block) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hash"] = o.Hash
	toSerialize["difficulty"] = o.Difficulty
	toSerialize["gasLimit"] = o.GasLimit
	toSerialize["number"] = o.Number
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["transactions"] = o.Transactions
	toSerialize["receiptsRoot"] = o.ReceiptsRoot
	toSerialize["parentHash"] = o.ParentHash
	toSerialize["sha3Uncles"] = o.Sha3Uncles
	toSerialize["miner"] = o.Miner
	toSerialize["stateRoot"] = o.StateRoot
	toSerialize["transactionsRoot"] = o.TransactionsRoot
	toSerialize["logsBloom"] = o.LogsBloom
	toSerialize["gasUsed"] = o.GasUsed
	toSerialize["nonce"] = o.Nonce
	toSerialize["mixHash"] = o.MixHash
	toSerialize["extraData"] = o.ExtraData
	if !IsNil(o.BaseFeePerGas) {
		toSerialize["baseFeePerGas"] = o.BaseFeePerGas
	}
	return toSerialize, nil
}

type NullableBlock struct {
	value *Block
	isSet bool
}

func (v NullableBlock) Get() *Block {
	return v.value
}

func (v *NullableBlock) Set(val *Block) {
	v.value = val
	v.isSet = true
}

func (v NullableBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlock(val *Block) *NullableBlock {
	return &NullableBlock{value: val, isSet: true}
}

func (v NullableBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
