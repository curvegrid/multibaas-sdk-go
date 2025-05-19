# EventField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The input name. | 
**Value** | **interface{}** |  | 
**Hashed** | **bool** | Has the value been hashed into a keccak256 string? | 
**Type** | **string** | The type of the argument. | 

## Methods

### NewEventField

`func NewEventField(name string, value interface{}, hashed bool, type_ string, ) *EventField`

NewEventField instantiates a new EventField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventFieldWithDefaults

`func NewEventFieldWithDefaults() *EventField`

NewEventFieldWithDefaults instantiates a new EventField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventField) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *EventField) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EventField) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EventField) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *EventField) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *EventField) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetHashed

`func (o *EventField) GetHashed() bool`

GetHashed returns the Hashed field if non-nil, zero value otherwise.

### GetHashedOk

`func (o *EventField) GetHashedOk() (*bool, bool)`

GetHashedOk returns a tuple with the Hashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashed

`func (o *EventField) SetHashed(v bool)`

SetHashed sets Hashed field to given value.


### GetType

`func (o *EventField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventField) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


