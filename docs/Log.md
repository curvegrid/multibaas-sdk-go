# Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | An ethereum address. | 
**Topics** | **[]string** | A list of topics provided by the contract. | 
**Data** | **string** | A hex string. | 
**BlockNumber** | **string** | A hex string. | 
**TransactionHash** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**TransactionIndex** | **string** | A hex string. | 
**BlockHash** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**LogIndex** | **string** | A hex string. | 
**Removed** | **bool** | True if this log was reverted due to a chain reorganization. | 

## Methods

### NewLog

`func NewLog(address string, topics []string, data string, blockNumber string, transactionHash string, transactionIndex string, blockHash string, logIndex string, removed bool, ) *Log`

NewLog instantiates a new Log object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogWithDefaults

`func NewLogWithDefaults() *Log`

NewLogWithDefaults instantiates a new Log object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Log) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Log) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Log) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTopics

`func (o *Log) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *Log) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *Log) SetTopics(v []string)`

SetTopics sets Topics field to given value.


### GetData

`func (o *Log) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Log) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Log) SetData(v string)`

SetData sets Data field to given value.


### GetBlockNumber

`func (o *Log) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *Log) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *Log) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.


### GetTransactionHash

`func (o *Log) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *Log) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *Log) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetTransactionIndex

`func (o *Log) GetTransactionIndex() string`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *Log) GetTransactionIndexOk() (*string, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *Log) SetTransactionIndex(v string)`

SetTransactionIndex sets TransactionIndex field to given value.


### GetBlockHash

`func (o *Log) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *Log) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *Log) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetLogIndex

`func (o *Log) GetLogIndex() string`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *Log) GetLogIndexOk() (*string, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *Log) SetLogIndex(v string)`

SetLogIndex sets LogIndex field to given value.


### GetRemoved

`func (o *Log) GetRemoved() bool`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *Log) GetRemovedOk() (*bool, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *Log) SetRemoved(v bool)`

SetRemoved sets Removed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


