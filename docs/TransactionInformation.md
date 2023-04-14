# TransactionInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | An ethereum address. | 
**TxData** | **string** | A hex string. | 
**TxHash** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**TxIndexInBlock** | **int64** | The location of the transaction in the block. | 
**BlockHash** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**BlockNumber** | **int64** | The transaction block number. | 
**Contract** | [**ContractInformation**](ContractInformation.md) |  | 
**Method** | [**ContractMethodInformation**](ContractMethodInformation.md) |  | 

## Methods

### NewTransactionInformation

`func NewTransactionInformation(from string, txData string, txHash string, txIndexInBlock int64, blockHash string, blockNumber int64, contract ContractInformation, method ContractMethodInformation, ) *TransactionInformation`

NewTransactionInformation instantiates a new TransactionInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionInformationWithDefaults

`func NewTransactionInformationWithDefaults() *TransactionInformation`

NewTransactionInformationWithDefaults instantiates a new TransactionInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *TransactionInformation) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TransactionInformation) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TransactionInformation) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTxData

`func (o *TransactionInformation) GetTxData() string`

GetTxData returns the TxData field if non-nil, zero value otherwise.

### GetTxDataOk

`func (o *TransactionInformation) GetTxDataOk() (*string, bool)`

GetTxDataOk returns a tuple with the TxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxData

`func (o *TransactionInformation) SetTxData(v string)`

SetTxData sets TxData field to given value.


### GetTxHash

`func (o *TransactionInformation) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *TransactionInformation) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *TransactionInformation) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetTxIndexInBlock

`func (o *TransactionInformation) GetTxIndexInBlock() int64`

GetTxIndexInBlock returns the TxIndexInBlock field if non-nil, zero value otherwise.

### GetTxIndexInBlockOk

`func (o *TransactionInformation) GetTxIndexInBlockOk() (*int64, bool)`

GetTxIndexInBlockOk returns a tuple with the TxIndexInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndexInBlock

`func (o *TransactionInformation) SetTxIndexInBlock(v int64)`

SetTxIndexInBlock sets TxIndexInBlock field to given value.


### GetBlockHash

`func (o *TransactionInformation) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TransactionInformation) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TransactionInformation) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetBlockNumber

`func (o *TransactionInformation) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *TransactionInformation) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *TransactionInformation) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.


### GetContract

`func (o *TransactionInformation) GetContract() ContractInformation`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *TransactionInformation) GetContractOk() (*ContractInformation, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *TransactionInformation) SetContract(v ContractInformation)`

SetContract sets Contract field to given value.


### GetMethod

`func (o *TransactionInformation) GetMethod() ContractMethodInformation`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TransactionInformation) GetMethodOk() (*ContractMethodInformation, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TransactionInformation) SetMethod(v ContractMethodInformation)`

SetMethod sets Method field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


