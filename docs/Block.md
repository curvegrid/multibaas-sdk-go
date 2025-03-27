# Block

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**Difficulty** | **string** |  | 
**GasLimit** | **int64** |  | 
**Number** | **string** |  | 
**Timestamp** | **int64** |  | 
**Transactions** | [**[]Transaction**](Transaction.md) |  | 
**ReceiptsRoot** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**ParentHash** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**Sha3Uncles** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**Miner** | **string** | An ethereum address. | 
**StateRoot** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**TransactionsRoot** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**LogsBloom** | **string** | A hex string. | 
**GasUsed** | **int32** |  | 
**Nonce** | **string** | A hex string. | 
**MixHash** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**ExtraData** | **string** |  | 
**BaseFeePerGas** | Pointer to **string** |  | [optional] 

## Methods

### NewBlock

`func NewBlock(hash string, difficulty string, gasLimit int64, number string, timestamp int64, transactions []Transaction, receiptsRoot string, parentHash string, sha3Uncles string, miner string, stateRoot string, transactionsRoot string, logsBloom string, gasUsed int32, nonce string, mixHash string, extraData string, ) *Block`

NewBlock instantiates a new Block object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockWithDefaults

`func NewBlockWithDefaults() *Block`

NewBlockWithDefaults instantiates a new Block object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *Block) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Block) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Block) SetHash(v string)`

SetHash sets Hash field to given value.


### GetDifficulty

`func (o *Block) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *Block) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *Block) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.


### GetGasLimit

`func (o *Block) GetGasLimit() int64`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *Block) GetGasLimitOk() (*int64, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *Block) SetGasLimit(v int64)`

SetGasLimit sets GasLimit field to given value.


### GetNumber

`func (o *Block) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Block) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Block) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetTimestamp

`func (o *Block) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Block) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Block) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetTransactions

`func (o *Block) GetTransactions() []Transaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Block) GetTransactionsOk() (*[]Transaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Block) SetTransactions(v []Transaction)`

SetTransactions sets Transactions field to given value.


### GetReceiptsRoot

`func (o *Block) GetReceiptsRoot() string`

GetReceiptsRoot returns the ReceiptsRoot field if non-nil, zero value otherwise.

### GetReceiptsRootOk

`func (o *Block) GetReceiptsRootOk() (*string, bool)`

GetReceiptsRootOk returns a tuple with the ReceiptsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptsRoot

`func (o *Block) SetReceiptsRoot(v string)`

SetReceiptsRoot sets ReceiptsRoot field to given value.


### GetParentHash

`func (o *Block) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *Block) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *Block) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.


### GetSha3Uncles

`func (o *Block) GetSha3Uncles() string`

GetSha3Uncles returns the Sha3Uncles field if non-nil, zero value otherwise.

### GetSha3UnclesOk

`func (o *Block) GetSha3UnclesOk() (*string, bool)`

GetSha3UnclesOk returns a tuple with the Sha3Uncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha3Uncles

`func (o *Block) SetSha3Uncles(v string)`

SetSha3Uncles sets Sha3Uncles field to given value.


### GetMiner

`func (o *Block) GetMiner() string`

GetMiner returns the Miner field if non-nil, zero value otherwise.

### GetMinerOk

`func (o *Block) GetMinerOk() (*string, bool)`

GetMinerOk returns a tuple with the Miner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiner

`func (o *Block) SetMiner(v string)`

SetMiner sets Miner field to given value.


### GetStateRoot

`func (o *Block) GetStateRoot() string`

GetStateRoot returns the StateRoot field if non-nil, zero value otherwise.

### GetStateRootOk

`func (o *Block) GetStateRootOk() (*string, bool)`

GetStateRootOk returns a tuple with the StateRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRoot

`func (o *Block) SetStateRoot(v string)`

SetStateRoot sets StateRoot field to given value.


### GetTransactionsRoot

`func (o *Block) GetTransactionsRoot() string`

GetTransactionsRoot returns the TransactionsRoot field if non-nil, zero value otherwise.

### GetTransactionsRootOk

`func (o *Block) GetTransactionsRootOk() (*string, bool)`

GetTransactionsRootOk returns a tuple with the TransactionsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsRoot

`func (o *Block) SetTransactionsRoot(v string)`

SetTransactionsRoot sets TransactionsRoot field to given value.


### GetLogsBloom

`func (o *Block) GetLogsBloom() string`

GetLogsBloom returns the LogsBloom field if non-nil, zero value otherwise.

### GetLogsBloomOk

`func (o *Block) GetLogsBloomOk() (*string, bool)`

GetLogsBloomOk returns a tuple with the LogsBloom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsBloom

`func (o *Block) SetLogsBloom(v string)`

SetLogsBloom sets LogsBloom field to given value.


### GetGasUsed

`func (o *Block) GetGasUsed() int32`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *Block) GetGasUsedOk() (*int32, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *Block) SetGasUsed(v int32)`

SetGasUsed sets GasUsed field to given value.


### GetNonce

`func (o *Block) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *Block) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *Block) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetMixHash

`func (o *Block) GetMixHash() string`

GetMixHash returns the MixHash field if non-nil, zero value otherwise.

### GetMixHashOk

`func (o *Block) GetMixHashOk() (*string, bool)`

GetMixHashOk returns a tuple with the MixHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixHash

`func (o *Block) SetMixHash(v string)`

SetMixHash sets MixHash field to given value.


### GetExtraData

`func (o *Block) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *Block) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *Block) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.


### GetBaseFeePerGas

`func (o *Block) GetBaseFeePerGas() string`

GetBaseFeePerGas returns the BaseFeePerGas field if non-nil, zero value otherwise.

### GetBaseFeePerGasOk

`func (o *Block) GetBaseFeePerGasOk() (*string, bool)`

GetBaseFeePerGasOk returns a tuple with the BaseFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFeePerGas

`func (o *Block) SetBaseFeePerGas(v string)`

SetBaseFeePerGas sets BaseFeePerGas field to given value.

### HasBaseFeePerGas

`func (o *Block) HasBaseFeePerGas() bool`

HasBaseFeePerGas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


