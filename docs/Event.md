# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggeredAt** | **time.Time** | The time at which the event was triggered. | 
**Event** | [**EventInformation**](EventInformation.md) |  | 
**Transaction** | [**TransactionInformation**](TransactionInformation.md) |  | 

## Methods

### NewEvent

`func NewEvent(triggeredAt time.Time, event EventInformation, transaction TransactionInformation, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggeredAt

`func (o *Event) GetTriggeredAt() time.Time`

GetTriggeredAt returns the TriggeredAt field if non-nil, zero value otherwise.

### GetTriggeredAtOk

`func (o *Event) GetTriggeredAtOk() (*time.Time, bool)`

GetTriggeredAtOk returns a tuple with the TriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredAt

`func (o *Event) SetTriggeredAt(v time.Time)`

SetTriggeredAt sets TriggeredAt field to given value.


### GetEvent

`func (o *Event) GetEvent() EventInformation`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Event) GetEventOk() (*EventInformation, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Event) SetEvent(v EventInformation)`

SetEvent sets Event field to given value.


### GetTransaction

`func (o *Event) GetTransaction() TransactionInformation`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *Event) GetTransactionOk() (*TransactionInformation, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *Event) SetTransaction(v TransactionInformation)`

SetTransaction sets Transaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


