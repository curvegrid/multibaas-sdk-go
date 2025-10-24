# PlanLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the limit. | 
**Limit** | **NullableInt64** | The limit value. Null means unlimited. | 
**Count** | Pointer to **int64** | The current count for this limit. | [optional] 

## Methods

### NewPlanLimit

`func NewPlanLimit(name string, limit NullableInt64, ) *PlanLimit`

NewPlanLimit instantiates a new PlanLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanLimitWithDefaults

`func NewPlanLimitWithDefaults() *PlanLimit`

NewPlanLimitWithDefaults instantiates a new PlanLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlanLimit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanLimit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanLimit) SetName(v string)`

SetName sets Name field to given value.


### GetLimit

`func (o *PlanLimit) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PlanLimit) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PlanLimit) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### SetLimitNil

`func (o *PlanLimit) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *PlanLimit) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetCount

`func (o *PlanLimit) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PlanLimit) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PlanLimit) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PlanLimit) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


