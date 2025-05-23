/*
MultiBaas API

MultiBaas's REST APIv0.

API version: 0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package multibaas

import (
	"encoding/json"
	"time"
)

// checks if the WalletTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletTransaction{}

// WalletTransaction struct for WalletTransaction
type WalletTransaction struct {
	Tx     Transaction       `json:"tx"`
	Status TransactionStatus `json:"status"`
	// An ethereum address.
	From string `json:"from" validate:"regexp=^0[xX][a-fA-F0-9]{40}$"`
	// The total number of resubmission attempts.
	ResubmissionAttempts int64 `json:"resubmissionAttempts"`
	// The total number of successful resubmission (added into the transaction pool).
	SuccessfulResubmissions int64 `json:"successfulResubmissions"`
	// The time the transaction was created.
	CreatedAt time.Time `json:"createdAt"`
	// The time the transaction was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
	// Whether the transaction failed when it was included in a block.
	Failed *bool `json:"failed,omitempty"`
	// The block number that the transaction was included in.
	BlockNumber *int64 `json:"blockNumber,omitempty"`
	// The keccak256 hash as a hex string of 256 bits.
	BlockHash *string `json:"blockHash,omitempty" validate:"regexp=^(0x[0-9a-f]{64}|0X[0-9A-F]{64})$"`
}

type _WalletTransaction WalletTransaction

// NewWalletTransaction instantiates a new WalletTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransaction(tx Transaction, status TransactionStatus, from string, resubmissionAttempts int64, successfulResubmissions int64, createdAt time.Time, updatedAt time.Time) *WalletTransaction {
	this := WalletTransaction{}
	this.Tx = tx
	this.Status = status
	this.From = from
	this.ResubmissionAttempts = resubmissionAttempts
	this.SuccessfulResubmissions = successfulResubmissions
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewWalletTransactionWithDefaults instantiates a new WalletTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionWithDefaults() *WalletTransaction {
	this := WalletTransaction{}
	return &this
}

// GetTx returns the Tx field value
func (o *WalletTransaction) GetTx() Transaction {
	if o == nil {
		var ret Transaction
		return ret
	}

	return o.Tx
}

// GetTxOk returns a tuple with the Tx field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetTxOk() (*Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tx, true
}

// SetTx sets field value
func (o *WalletTransaction) SetTx(v Transaction) {
	o.Tx = v
}

// GetStatus returns the Status field value
func (o *WalletTransaction) GetStatus() TransactionStatus {
	if o == nil {
		var ret TransactionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetStatusOk() (*TransactionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WalletTransaction) SetStatus(v TransactionStatus) {
	o.Status = v
}

// GetFrom returns the From field value
func (o *WalletTransaction) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *WalletTransaction) SetFrom(v string) {
	o.From = v
}

// GetResubmissionAttempts returns the ResubmissionAttempts field value
func (o *WalletTransaction) GetResubmissionAttempts() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ResubmissionAttempts
}

// GetResubmissionAttemptsOk returns a tuple with the ResubmissionAttempts field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetResubmissionAttemptsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResubmissionAttempts, true
}

// SetResubmissionAttempts sets field value
func (o *WalletTransaction) SetResubmissionAttempts(v int64) {
	o.ResubmissionAttempts = v
}

// GetSuccessfulResubmissions returns the SuccessfulResubmissions field value
func (o *WalletTransaction) GetSuccessfulResubmissions() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SuccessfulResubmissions
}

// GetSuccessfulResubmissionsOk returns a tuple with the SuccessfulResubmissions field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetSuccessfulResubmissionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessfulResubmissions, true
}

// SetSuccessfulResubmissions sets field value
func (o *WalletTransaction) SetSuccessfulResubmissions(v int64) {
	o.SuccessfulResubmissions = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WalletTransaction) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WalletTransaction) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WalletTransaction) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WalletTransaction) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *WalletTransaction) GetFailed() bool {
	if o == nil || IsNil(o.Failed) {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *WalletTransaction) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *WalletTransaction) SetFailed(v bool) {
	o.Failed = &v
}

// GetBlockNumber returns the BlockNumber field value if set, zero value otherwise.
func (o *WalletTransaction) GetBlockNumber() int64 {
	if o == nil || IsNil(o.BlockNumber) {
		var ret int64
		return ret
	}
	return *o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetBlockNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.BlockNumber) {
		return nil, false
	}
	return o.BlockNumber, true
}

// HasBlockNumber returns a boolean if a field has been set.
func (o *WalletTransaction) HasBlockNumber() bool {
	if o != nil && !IsNil(o.BlockNumber) {
		return true
	}

	return false
}

// SetBlockNumber gets a reference to the given int64 and assigns it to the BlockNumber field.
func (o *WalletTransaction) SetBlockNumber(v int64) {
	o.BlockNumber = &v
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise.
func (o *WalletTransaction) GetBlockHash() string {
	if o == nil || IsNil(o.BlockHash) {
		var ret string
		return ret
	}
	return *o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransaction) GetBlockHashOk() (*string, bool) {
	if o == nil || IsNil(o.BlockHash) {
		return nil, false
	}
	return o.BlockHash, true
}

// HasBlockHash returns a boolean if a field has been set.
func (o *WalletTransaction) HasBlockHash() bool {
	if o != nil && !IsNil(o.BlockHash) {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given string and assigns it to the BlockHash field.
func (o *WalletTransaction) SetBlockHash(v string) {
	o.BlockHash = &v
}

func (o WalletTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx"] = o.Tx
	toSerialize["status"] = o.Status
	toSerialize["from"] = o.From
	toSerialize["resubmissionAttempts"] = o.ResubmissionAttempts
	toSerialize["successfulResubmissions"] = o.SuccessfulResubmissions
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.BlockNumber) {
		toSerialize["blockNumber"] = o.BlockNumber
	}
	if !IsNil(o.BlockHash) {
		toSerialize["blockHash"] = o.BlockHash
	}
	return toSerialize, nil
}

type NullableWalletTransaction struct {
	value *WalletTransaction
	isSet bool
}

func (v NullableWalletTransaction) Get() *WalletTransaction {
	return v.value
}

func (v *NullableWalletTransaction) Set(val *WalletTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransaction(val *WalletTransaction) *NullableWalletTransaction {
	return &NullableWalletTransaction{value: val, isSet: true}
}

func (v NullableWalletTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
