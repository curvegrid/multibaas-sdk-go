# EventQueryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | Pointer to **string** | The rule type, can be omitted if this is the last filter to be applied. | [optional] 
**FieldType** | Pointer to [**FieldType**](FieldType.md) |  | [optional] 
**InputIndex** | Pointer to **NullableInt64** | The field&#39;s index, can be used in place of &#x60;name&#x60;. | [optional] 
**Operator** | Pointer to **string** | The operator on the filter. | [optional] 
**Value** | Pointer to **string** | The value to be compared with. | [optional] 
**Children** | Pointer to [**[]EventQueryFilter**](EventQueryFilter.md) | Other filters to be applied in succession with the rule specified. | [optional] 

## Methods

### NewEventQueryFilter

`func NewEventQueryFilter() *EventQueryFilter`

NewEventQueryFilter instantiates a new EventQueryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventQueryFilterWithDefaults

`func NewEventQueryFilterWithDefaults() *EventQueryFilter`

NewEventQueryFilterWithDefaults instantiates a new EventQueryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *EventQueryFilter) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *EventQueryFilter) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *EventQueryFilter) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *EventQueryFilter) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetFieldType

`func (o *EventQueryFilter) GetFieldType() FieldType`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *EventQueryFilter) GetFieldTypeOk() (*FieldType, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *EventQueryFilter) SetFieldType(v FieldType)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *EventQueryFilter) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetInputIndex

`func (o *EventQueryFilter) GetInputIndex() int64`

GetInputIndex returns the InputIndex field if non-nil, zero value otherwise.

### GetInputIndexOk

`func (o *EventQueryFilter) GetInputIndexOk() (*int64, bool)`

GetInputIndexOk returns a tuple with the InputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputIndex

`func (o *EventQueryFilter) SetInputIndex(v int64)`

SetInputIndex sets InputIndex field to given value.

### HasInputIndex

`func (o *EventQueryFilter) HasInputIndex() bool`

HasInputIndex returns a boolean if a field has been set.

### SetInputIndexNil

`func (o *EventQueryFilter) SetInputIndexNil(b bool)`

 SetInputIndexNil sets the value for InputIndex to be an explicit nil

### UnsetInputIndex
`func (o *EventQueryFilter) UnsetInputIndex()`

UnsetInputIndex ensures that no value is present for InputIndex, not even an explicit nil
### GetOperator

`func (o *EventQueryFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *EventQueryFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *EventQueryFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *EventQueryFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *EventQueryFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EventQueryFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EventQueryFilter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EventQueryFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetChildren

`func (o *EventQueryFilter) GetChildren() []EventQueryFilter`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *EventQueryFilter) GetChildrenOk() (*[]EventQueryFilter, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *EventQueryFilter) SetChildren(v []EventQueryFilter)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *EventQueryFilter) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


