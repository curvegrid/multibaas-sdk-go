# EventIndexingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsProcessingPastLogs** | **bool** | When true, the event indexer is syncing historical events from past blocks in order to catch up to the current state. | 
**LatestBlockNumber** | **int64** | The most recent block number that has been processed for event indexing. | 
**LatestBlockHash** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**StartBlockNumber** | **int64** | The block number from which event indexing began. | 
**StartBlockHash** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**UpdatedAt** | **string** | The timestamp of the last update to the indexing status. | 

## Methods

### NewEventIndexingStatus

`func NewEventIndexingStatus(isProcessingPastLogs bool, latestBlockNumber int64, latestBlockHash string, startBlockNumber int64, startBlockHash string, updatedAt string, ) *EventIndexingStatus`

NewEventIndexingStatus instantiates a new EventIndexingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventIndexingStatusWithDefaults

`func NewEventIndexingStatusWithDefaults() *EventIndexingStatus`

NewEventIndexingStatusWithDefaults instantiates a new EventIndexingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsProcessingPastLogs

`func (o *EventIndexingStatus) GetIsProcessingPastLogs() bool`

GetIsProcessingPastLogs returns the IsProcessingPastLogs field if non-nil, zero value otherwise.

### GetIsProcessingPastLogsOk

`func (o *EventIndexingStatus) GetIsProcessingPastLogsOk() (*bool, bool)`

GetIsProcessingPastLogsOk returns a tuple with the IsProcessingPastLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProcessingPastLogs

`func (o *EventIndexingStatus) SetIsProcessingPastLogs(v bool)`

SetIsProcessingPastLogs sets IsProcessingPastLogs field to given value.


### GetLatestBlockNumber

`func (o *EventIndexingStatus) GetLatestBlockNumber() int64`

GetLatestBlockNumber returns the LatestBlockNumber field if non-nil, zero value otherwise.

### GetLatestBlockNumberOk

`func (o *EventIndexingStatus) GetLatestBlockNumberOk() (*int64, bool)`

GetLatestBlockNumberOk returns a tuple with the LatestBlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestBlockNumber

`func (o *EventIndexingStatus) SetLatestBlockNumber(v int64)`

SetLatestBlockNumber sets LatestBlockNumber field to given value.


### GetLatestBlockHash

`func (o *EventIndexingStatus) GetLatestBlockHash() string`

GetLatestBlockHash returns the LatestBlockHash field if non-nil, zero value otherwise.

### GetLatestBlockHashOk

`func (o *EventIndexingStatus) GetLatestBlockHashOk() (*string, bool)`

GetLatestBlockHashOk returns a tuple with the LatestBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestBlockHash

`func (o *EventIndexingStatus) SetLatestBlockHash(v string)`

SetLatestBlockHash sets LatestBlockHash field to given value.


### GetStartBlockNumber

`func (o *EventIndexingStatus) GetStartBlockNumber() int64`

GetStartBlockNumber returns the StartBlockNumber field if non-nil, zero value otherwise.

### GetStartBlockNumberOk

`func (o *EventIndexingStatus) GetStartBlockNumberOk() (*int64, bool)`

GetStartBlockNumberOk returns a tuple with the StartBlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartBlockNumber

`func (o *EventIndexingStatus) SetStartBlockNumber(v int64)`

SetStartBlockNumber sets StartBlockNumber field to given value.


### GetStartBlockHash

`func (o *EventIndexingStatus) GetStartBlockHash() string`

GetStartBlockHash returns the StartBlockHash field if non-nil, zero value otherwise.

### GetStartBlockHashOk

`func (o *EventIndexingStatus) GetStartBlockHashOk() (*string, bool)`

GetStartBlockHashOk returns a tuple with the StartBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartBlockHash

`func (o *EventIndexingStatus) SetStartBlockHash(v string)`

SetStartBlockHash sets StartBlockHash field to given value.


### GetUpdatedAt

`func (o *EventIndexingStatus) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EventIndexingStatus) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EventIndexingStatus) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


