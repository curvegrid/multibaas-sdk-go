# EventQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]EventQueryEvent**](EventQueryEvent.md) |  | 
**GroupBy** | Pointer to **string** | The results will be grouped according to this field. An aggregator for non Group By fields must be specified if groupBy is specified. | [optional] 
**OrderBy** | Pointer to **string** | The results will be ordered according to this field. | [optional] 
**Order** | Pointer to **string** | Specify ascending or descending order, the default is \&quot;ASC\&quot;. | [optional] 

## Methods

### NewEventQuery

`func NewEventQuery(events []EventQueryEvent, ) *EventQuery`

NewEventQuery instantiates a new EventQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventQueryWithDefaults

`func NewEventQueryWithDefaults() *EventQuery`

NewEventQueryWithDefaults instantiates a new EventQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EventQuery) GetEvents() []EventQueryEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventQuery) GetEventsOk() (*[]EventQueryEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventQuery) SetEvents(v []EventQueryEvent)`

SetEvents sets Events field to given value.


### GetGroupBy

`func (o *EventQuery) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *EventQuery) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *EventQuery) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *EventQuery) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetOrderBy

`func (o *EventQuery) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *EventQuery) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *EventQuery) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *EventQuery) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetOrder

`func (o *EventQuery) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *EventQuery) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *EventQuery) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *EventQuery) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


