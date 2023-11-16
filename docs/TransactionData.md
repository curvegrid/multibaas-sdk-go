# TransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Transaction**](Transaction.md) |  | 
**IsPending** | **bool** | Whether the transaction has been included yet. | 
**From** | **string** | An ethereum address. | 
**BlockHash** | Pointer to **string** | The keccak256 hash as a hex string of 256 bits. | [optional] 
**BlockNumber** | Pointer to **string** | The transaction block number. | [optional] 
**Contract** | Pointer to [**ContractInformation**](ContractInformation.md) |  | [optional] 
**Method** | Pointer to [**ContractMethodInformation**](ContractMethodInformation.md) |  | [optional] 

## Methods

### NewTransactionData

`func NewTransactionData(data Transaction, isPending bool, from string, ) *TransactionData`

NewTransactionData instantiates a new TransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDataWithDefaults

`func NewTransactionDataWithDefaults() *TransactionData`

NewTransactionDataWithDefaults instantiates a new TransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransactionData) GetData() Transaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionData) GetDataOk() (*Transaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionData) SetData(v Transaction)`

SetData sets Data field to given value.


### GetIsPending

`func (o *TransactionData) GetIsPending() bool`

GetIsPending returns the IsPending field if non-nil, zero value otherwise.

### GetIsPendingOk

`func (o *TransactionData) GetIsPendingOk() (*bool, bool)`

GetIsPendingOk returns a tuple with the IsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPending

`func (o *TransactionData) SetIsPending(v bool)`

SetIsPending sets IsPending field to given value.


### GetFrom

`func (o *TransactionData) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TransactionData) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TransactionData) SetFrom(v string)`

SetFrom sets From field to given value.


### GetBlockHash

`func (o *TransactionData) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TransactionData) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TransactionData) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *TransactionData) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetBlockNumber

`func (o *TransactionData) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *TransactionData) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *TransactionData) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *TransactionData) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetContract

`func (o *TransactionData) GetContract() ContractInformation`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *TransactionData) GetContractOk() (*ContractInformation, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *TransactionData) SetContract(v ContractInformation)`

SetContract sets Contract field to given value.

### HasContract

`func (o *TransactionData) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetMethod

`func (o *TransactionData) GetMethod() ContractMethodInformation`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TransactionData) GetMethodOk() (*ContractMethodInformation, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TransactionData) SetMethod(v ContractMethodInformation)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TransactionData) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


