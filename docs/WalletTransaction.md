# WalletTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tx** | [**Transaction**](Transaction.md) |  | 
**Status** | [**TransactionStatus**](TransactionStatus.md) |  | 
**From** | **string** | An ethereum address. | 
**ResubmissionAttempts** | **int64** | The total number of resubmission attempts. | 
**SuccessfulResubmissions** | **int64** | The total number of successful resubmission (added into the transaction pool). | 
**CreatedAt** | **time.Time** | The time the transaction was created. | 
**UpdatedAt** | **time.Time** | The time the transaction was last updated. | 
**Failed** | Pointer to **bool** | Whether the transaction failed when it was included in a block. | [optional] 
**BlockNumber** | Pointer to **int64** | The block number that the transaction was included in. | [optional] 
**BlockHash** | Pointer to **string** | The keccak256 hash as a hex string of 256 bits. | [optional] 

## Methods

### NewWalletTransaction

`func NewWalletTransaction(tx Transaction, status TransactionStatus, from string, resubmissionAttempts int64, successfulResubmissions int64, createdAt time.Time, updatedAt time.Time, ) *WalletTransaction`

NewWalletTransaction instantiates a new WalletTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTransactionWithDefaults

`func NewWalletTransactionWithDefaults() *WalletTransaction`

NewWalletTransactionWithDefaults instantiates a new WalletTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTx

`func (o *WalletTransaction) GetTx() Transaction`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *WalletTransaction) GetTxOk() (*Transaction, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *WalletTransaction) SetTx(v Transaction)`

SetTx sets Tx field to given value.


### GetStatus

`func (o *WalletTransaction) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WalletTransaction) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WalletTransaction) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.


### GetFrom

`func (o *WalletTransaction) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *WalletTransaction) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *WalletTransaction) SetFrom(v string)`

SetFrom sets From field to given value.


### GetResubmissionAttempts

`func (o *WalletTransaction) GetResubmissionAttempts() int64`

GetResubmissionAttempts returns the ResubmissionAttempts field if non-nil, zero value otherwise.

### GetResubmissionAttemptsOk

`func (o *WalletTransaction) GetResubmissionAttemptsOk() (*int64, bool)`

GetResubmissionAttemptsOk returns a tuple with the ResubmissionAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResubmissionAttempts

`func (o *WalletTransaction) SetResubmissionAttempts(v int64)`

SetResubmissionAttempts sets ResubmissionAttempts field to given value.


### GetSuccessfulResubmissions

`func (o *WalletTransaction) GetSuccessfulResubmissions() int64`

GetSuccessfulResubmissions returns the SuccessfulResubmissions field if non-nil, zero value otherwise.

### GetSuccessfulResubmissionsOk

`func (o *WalletTransaction) GetSuccessfulResubmissionsOk() (*int64, bool)`

GetSuccessfulResubmissionsOk returns a tuple with the SuccessfulResubmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulResubmissions

`func (o *WalletTransaction) SetSuccessfulResubmissions(v int64)`

SetSuccessfulResubmissions sets SuccessfulResubmissions field to given value.


### GetCreatedAt

`func (o *WalletTransaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WalletTransaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WalletTransaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WalletTransaction) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WalletTransaction) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WalletTransaction) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFailed

`func (o *WalletTransaction) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *WalletTransaction) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *WalletTransaction) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *WalletTransaction) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetBlockNumber

`func (o *WalletTransaction) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *WalletTransaction) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *WalletTransaction) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *WalletTransaction) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetBlockHash

`func (o *WalletTransaction) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *WalletTransaction) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *WalletTransaction) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *WalletTransaction) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


