# EventQueryField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FieldType**](FieldType.md) |  | 
**Name** | Pointer to **string** | The field name. Either &#x60;name&#x60; or &#x60;inputIndex&#x60; is required if &#x60;fieldType &#x3D;&#x3D; \&quot;input\&quot;&#x60;. | [optional] 
**InputIndex** | Pointer to **NullableInt64** | The field&#39;s index, can be used in place of &#x60;name&#x60;. | [optional] 
**Alias** | Pointer to **string** | The name that will be used to return the field. | [optional] 
**Aggregator** | Pointer to **NullableString** | The type of aggregation to perform on the field. | [optional] 

## Methods

### NewEventQueryField

`func NewEventQueryField(type_ FieldType, ) *EventQueryField`

NewEventQueryField instantiates a new EventQueryField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventQueryFieldWithDefaults

`func NewEventQueryFieldWithDefaults() *EventQueryField`

NewEventQueryFieldWithDefaults instantiates a new EventQueryField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventQueryField) GetType() FieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventQueryField) GetTypeOk() (*FieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventQueryField) SetType(v FieldType)`

SetType sets Type field to given value.


### GetName

`func (o *EventQueryField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventQueryField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventQueryField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventQueryField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInputIndex

`func (o *EventQueryField) GetInputIndex() int64`

GetInputIndex returns the InputIndex field if non-nil, zero value otherwise.

### GetInputIndexOk

`func (o *EventQueryField) GetInputIndexOk() (*int64, bool)`

GetInputIndexOk returns a tuple with the InputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputIndex

`func (o *EventQueryField) SetInputIndex(v int64)`

SetInputIndex sets InputIndex field to given value.

### HasInputIndex

`func (o *EventQueryField) HasInputIndex() bool`

HasInputIndex returns a boolean if a field has been set.

### SetInputIndexNil

`func (o *EventQueryField) SetInputIndexNil(b bool)`

 SetInputIndexNil sets the value for InputIndex to be an explicit nil

### UnsetInputIndex
`func (o *EventQueryField) UnsetInputIndex()`

UnsetInputIndex ensures that no value is present for InputIndex, not even an explicit nil
### GetAlias

`func (o *EventQueryField) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *EventQueryField) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *EventQueryField) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *EventQueryField) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAggregator

`func (o *EventQueryField) GetAggregator() string`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *EventQueryField) GetAggregatorOk() (*string, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *EventQueryField) SetAggregator(v string)`

SetAggregator sets Aggregator field to given value.

### HasAggregator

`func (o *EventQueryField) HasAggregator() bool`

HasAggregator returns a boolean if a field has been set.

### SetAggregatorNil

`func (o *EventQueryField) SetAggregatorNil(b bool)`

 SetAggregatorNil sets the value for Aggregator to be an explicit nil

### UnsetAggregator
`func (o *EventQueryField) UnsetAggregator()`

UnsetAggregator ensures that no value is present for Aggregator, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


