# EventQueryEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | **string** | The name of an event. | 
**Select** | [**[]EventQueryField**](EventQueryField.md) | The query information about all the fields to select from an event. | 
**Filter** | Pointer to [**EventQueryFilter**](EventQueryFilter.md) |  | [optional] 

## Methods

### NewEventQueryEvent

`func NewEventQueryEvent(eventName string, select_ []EventQueryField, ) *EventQueryEvent`

NewEventQueryEvent instantiates a new EventQueryEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventQueryEventWithDefaults

`func NewEventQueryEventWithDefaults() *EventQueryEvent`

NewEventQueryEventWithDefaults instantiates a new EventQueryEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *EventQueryEvent) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *EventQueryEvent) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *EventQueryEvent) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetSelect

`func (o *EventQueryEvent) GetSelect() []EventQueryField`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *EventQueryEvent) GetSelectOk() (*[]EventQueryField, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *EventQueryEvent) SetSelect(v []EventQueryField)`

SetSelect sets Select field to given value.


### GetFilter

`func (o *EventQueryEvent) GetFilter() EventQueryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *EventQueryEvent) GetFilterOk() (*EventQueryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *EventQueryEvent) SetFilter(v EventQueryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *EventQueryEvent) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


