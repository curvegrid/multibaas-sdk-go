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

// checks if the TransactionReceiptData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionReceiptData{}

// TransactionReceiptData struct for TransactionReceiptData
type TransactionReceiptData struct {
	// A hex string.
	Type *string `json:"type,omitempty"`
	// A hex string.
	Root string `json:"root"`
	// A hex string.
	Status string `json:"status"`
	// A hex string.
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	// A hex string.
	LogsBloom string `json:"logsBloom"`
	Logs      []Log  `json:"logs"`
	// The keccak256 hash as a hex string of 256 bits.
	TransactionHash string `json:"transactionHash"`
	// An ethereum address.
	ContractAddress string `json:"contractAddress"`
	// A hex string.
	GasUsed string `json:"gasUsed"`
	// A hex string.
	EffectiveGasPrice string `json:"effectiveGasPrice"`
	// A hex string.
	BlockNumber string `json:"blockNumber"`
	// A hex string.
	TransactionIndex string `json:"transactionIndex"`
	// The keccak256 hash as a hex string of 256 bits.
	BlockHash string `json:"blockHash"`
}

// NewTransactionReceiptData instantiates a new TransactionReceiptData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionReceiptData(root string, status string, cumulativeGasUsed string, logsBloom string, logs []Log, transactionHash string, contractAddress string, gasUsed string, effectiveGasPrice string, blockNumber string, transactionIndex string, blockHash string) *TransactionReceiptData {
	this := TransactionReceiptData{}
	this.Root = root
	this.Status = status
	this.CumulativeGasUsed = cumulativeGasUsed
	this.LogsBloom = logsBloom
	this.Logs = logs
	this.TransactionHash = transactionHash
	this.ContractAddress = contractAddress
	this.GasUsed = gasUsed
	this.EffectiveGasPrice = effectiveGasPrice
	this.BlockNumber = blockNumber
	this.TransactionIndex = transactionIndex
	this.BlockHash = blockHash
	return &this
}

// NewTransactionReceiptDataWithDefaults instantiates a new TransactionReceiptData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionReceiptDataWithDefaults() *TransactionReceiptData {
	this := TransactionReceiptData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransactionReceiptData) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionReceiptData) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransactionReceiptData) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransactionReceiptData) SetType(v string) {
	o.Type = &v
}

// GetRoot returns the Root field value
func (o *TransactionReceiptData) GetRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Root
}

// GetRootOk returns a tuple with the Root field value
// and a boolean to check if the value has been set.
func (o *TransactionReceiptData) GetRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Root, true
}

// SetRoot sets field value
func (o *TransactionReceiptData) SetRoot(v string) {
	o.Root = v
}

// GetStatus returns the Status field value
func (o *TransactionReceiptData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransactionReceiptData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransactionReceiptData) SetStatus(v string) {
	o.Status = v
}

// GetCumulativeGasUsed returns the CumulativeGasUsed field value
func (o *TransactionReceiptData) GetCumulativeGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CumulativeGasUsed
}

// GetCumulativeGasUsedOk returns a tuple with the CumulativeGasUsed field value
// and a boolean to check if the value has been set.
func (o *TransactionReceiptData) GetCumulativeGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CumulativeGasUsed, true
}

// SetCumulativeGasUsed sets field value
func (o *TransactionReceiptData) SetCumulativeGasUsed(v string) {
	o.CumulativeGasUsed = v
}

// GetLogsBloom returns the LogsBloom field value
func (o *TransactionReceiptData) GetLogsBloom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogsBloom
}

// GetLogsBloomOk returns a tuple with the LogsBloom field value
// and a boolean to check if the value has been set.
func (o *TransactionReceiptData) GetLogsBloomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogsBloom, true
}

// SetLogsBloom sets field value
func (o *TransactionReceiptData) SetLogsBloom(v string) {
	o.LogsBloom = v
}

// GetLogs returns the Logs field value
func (o *TransactionReceiptData) GetLogs() []Log {
	if o == nil {
		var ret []Log
		return ret
	}

	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *TransactionReceiptData) GetLogsOk() ([]Log, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logs, true
}

// SetLogs sets field value
func (o *TransactionReceiptData) SetLogs(v []Log) {
	o.Logs = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *TransactionReceiptData) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *TransactionReceiptData) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *TransactionReceiptData) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetContractAddress returns the ContractAddress field value
func (o *TransactionReceiptData) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *TransactionReceiptData) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *TransactionReceiptData) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetGasUsed returns the GasUsed field value
func (o *TransactionReceiptData) GetGasUsed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *TransactionReceiptData) GetGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *TransactionReceiptData) SetGasUsed(v string) {
	o.GasUsed = v
}

// GetEffectiveGasPrice returns the EffectiveGasPrice field value
func (o *TransactionReceiptData) GetEffectiveGasPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EffectiveGasPrice
}

// GetEffectiveGasPriceOk returns a tuple with the EffectiveGasPrice field value
// and a boolean to check if the value has been set.
func (o *TransactionReceiptData) GetEffectiveGasPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveGasPrice, true
}

// SetEffectiveGasPrice sets field value
func (o *TransactionReceiptData) SetEffectiveGasPrice(v string) {
	o.EffectiveGasPrice = v
}

// GetBlockNumber returns the BlockNumber field value
func (o *TransactionReceiptData) GetBlockNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value
// and a boolean to check if the value has been set.
func (o *TransactionReceiptData) GetBlockNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockNumber, true
}

// SetBlockNumber sets field value
func (o *TransactionReceiptData) SetBlockNumber(v string) {
	o.BlockNumber = v
}

// GetTransactionIndex returns the TransactionIndex field value
func (o *TransactionReceiptData) GetTransactionIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionIndex
}

// GetTransactionIndexOk returns a tuple with the TransactionIndex field value
// and a boolean to check if the value has been set.
func (o *TransactionReceiptData) GetTransactionIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionIndex, true
}

// SetTransactionIndex sets field value
func (o *TransactionReceiptData) SetTransactionIndex(v string) {
	o.TransactionIndex = v
}

// GetBlockHash returns the BlockHash field value
func (o *TransactionReceiptData) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *TransactionReceiptData) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *TransactionReceiptData) SetBlockHash(v string) {
	o.BlockHash = v
}

func (o TransactionReceiptData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionReceiptData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["root"] = o.Root
	toSerialize["status"] = o.Status
	toSerialize["cumulativeGasUsed"] = o.CumulativeGasUsed
	toSerialize["logsBloom"] = o.LogsBloom
	toSerialize["logs"] = o.Logs
	toSerialize["transactionHash"] = o.TransactionHash
	toSerialize["contractAddress"] = o.ContractAddress
	toSerialize["gasUsed"] = o.GasUsed
	toSerialize["effectiveGasPrice"] = o.EffectiveGasPrice
	toSerialize["blockNumber"] = o.BlockNumber
	toSerialize["transactionIndex"] = o.TransactionIndex
	toSerialize["blockHash"] = o.BlockHash
	return toSerialize, nil
}

type NullableTransactionReceiptData struct {
	value *TransactionReceiptData
	isSet bool
}

func (v NullableTransactionReceiptData) Get() *TransactionReceiptData {
	return v.value
}

func (v *NullableTransactionReceiptData) Set(val *TransactionReceiptData) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionReceiptData) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionReceiptData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionReceiptData(val *TransactionReceiptData) *NullableTransactionReceiptData {
	return &NullableTransactionReceiptData{value: val, isSet: true}
}

func (v NullableTransactionReceiptData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionReceiptData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}