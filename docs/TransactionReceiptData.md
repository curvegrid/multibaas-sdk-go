# TransactionReceiptData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A hex string. | [optional] 
**Root** | **string** | A hex string. | 
**Status** | **string** | A hex string. | 
**CumulativeGasUsed** | **string** | A hex string. | 
**LogsBloom** | **string** | A hex string. | 
**Logs** | [**[]Log**](Log.md) |  | 
**TransactionHash** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**ContractAddress** | **string** | An ethereum address. | 
**GasUsed** | **string** | A hex string. | 
**EffectiveGasPrice** | **string** | A hex string. | 
**BlockNumber** | **string** | A hex string. | 
**TransactionIndex** | **string** | A hex string. | 
**BlockHash** | **string** | The keccak256 hash as a hex string of 256 bits. | 

## Methods

### NewTransactionReceiptData

`func NewTransactionReceiptData(root string, status string, cumulativeGasUsed string, logsBloom string, logs []Log, transactionHash string, contractAddress string, gasUsed string, effectiveGasPrice string, blockNumber string, transactionIndex string, blockHash string, ) *TransactionReceiptData`

NewTransactionReceiptData instantiates a new TransactionReceiptData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionReceiptDataWithDefaults

`func NewTransactionReceiptDataWithDefaults() *TransactionReceiptData`

NewTransactionReceiptDataWithDefaults instantiates a new TransactionReceiptData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionReceiptData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionReceiptData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionReceiptData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionReceiptData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRoot

`func (o *TransactionReceiptData) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *TransactionReceiptData) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *TransactionReceiptData) SetRoot(v string)`

SetRoot sets Root field to given value.


### GetStatus

`func (o *TransactionReceiptData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionReceiptData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionReceiptData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCumulativeGasUsed

`func (o *TransactionReceiptData) GetCumulativeGasUsed() string`

GetCumulativeGasUsed returns the CumulativeGasUsed field if non-nil, zero value otherwise.

### GetCumulativeGasUsedOk

`func (o *TransactionReceiptData) GetCumulativeGasUsedOk() (*string, bool)`

GetCumulativeGasUsedOk returns a tuple with the CumulativeGasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeGasUsed

`func (o *TransactionReceiptData) SetCumulativeGasUsed(v string)`

SetCumulativeGasUsed sets CumulativeGasUsed field to given value.


### GetLogsBloom

`func (o *TransactionReceiptData) GetLogsBloom() string`

GetLogsBloom returns the LogsBloom field if non-nil, zero value otherwise.

### GetLogsBloomOk

`func (o *TransactionReceiptData) GetLogsBloomOk() (*string, bool)`

GetLogsBloomOk returns a tuple with the LogsBloom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsBloom

`func (o *TransactionReceiptData) SetLogsBloom(v string)`

SetLogsBloom sets LogsBloom field to given value.


### GetLogs

`func (o *TransactionReceiptData) GetLogs() []Log`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *TransactionReceiptData) GetLogsOk() (*[]Log, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *TransactionReceiptData) SetLogs(v []Log)`

SetLogs sets Logs field to given value.


### GetTransactionHash

`func (o *TransactionReceiptData) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *TransactionReceiptData) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *TransactionReceiptData) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetContractAddress

`func (o *TransactionReceiptData) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *TransactionReceiptData) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *TransactionReceiptData) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetGasUsed

`func (o *TransactionReceiptData) GetGasUsed() string`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *TransactionReceiptData) GetGasUsedOk() (*string, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *TransactionReceiptData) SetGasUsed(v string)`

SetGasUsed sets GasUsed field to given value.


### GetEffectiveGasPrice

`func (o *TransactionReceiptData) GetEffectiveGasPrice() string`

GetEffectiveGasPrice returns the EffectiveGasPrice field if non-nil, zero value otherwise.

### GetEffectiveGasPriceOk

`func (o *TransactionReceiptData) GetEffectiveGasPriceOk() (*string, bool)`

GetEffectiveGasPriceOk returns a tuple with the EffectiveGasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveGasPrice

`func (o *TransactionReceiptData) SetEffectiveGasPrice(v string)`

SetEffectiveGasPrice sets EffectiveGasPrice field to given value.


### GetBlockNumber

`func (o *TransactionReceiptData) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *TransactionReceiptData) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *TransactionReceiptData) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.


### GetTransactionIndex

`func (o *TransactionReceiptData) GetTransactionIndex() string`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *TransactionReceiptData) GetTransactionIndexOk() (*string, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *TransactionReceiptData) SetTransactionIndex(v string)`

SetTransactionIndex sets TransactionIndex field to given value.


### GetBlockHash

`func (o *TransactionReceiptData) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TransactionReceiptData) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TransactionReceiptData) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


