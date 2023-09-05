# AuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionByID** | **int64** | The ID of the user who performed the action. | 
**ActionOnID** | Pointer to **int64** | The ID of the user who was acted upon. | [optional] 
**ActionByUserEmail** | **string** | The email of the user who performed the action. | 
**ActionOnUserEmail** | Pointer to **string** | The email of the user who was acted upon. | [optional] 
**Type** | **string** | The type of action that was performed. | 
**CreatedAt** | **time.Time** | The time the action was performed. | 
**ActivityData** | **map[string]interface{}** | The data associated with the action. | 

## Methods

### NewAuditLog

`func NewAuditLog(actionByID int64, actionByUserEmail string, type_ string, createdAt time.Time, activityData map[string]interface{}, ) *AuditLog`

NewAuditLog instantiates a new AuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogWithDefaults

`func NewAuditLogWithDefaults() *AuditLog`

NewAuditLogWithDefaults instantiates a new AuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionByID

`func (o *AuditLog) GetActionByID() int64`

GetActionByID returns the ActionByID field if non-nil, zero value otherwise.

### GetActionByIDOk

`func (o *AuditLog) GetActionByIDOk() (*int64, bool)`

GetActionByIDOk returns a tuple with the ActionByID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionByID

`func (o *AuditLog) SetActionByID(v int64)`

SetActionByID sets ActionByID field to given value.


### GetActionOnID

`func (o *AuditLog) GetActionOnID() int64`

GetActionOnID returns the ActionOnID field if non-nil, zero value otherwise.

### GetActionOnIDOk

`func (o *AuditLog) GetActionOnIDOk() (*int64, bool)`

GetActionOnIDOk returns a tuple with the ActionOnID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOnID

`func (o *AuditLog) SetActionOnID(v int64)`

SetActionOnID sets ActionOnID field to given value.

### HasActionOnID

`func (o *AuditLog) HasActionOnID() bool`

HasActionOnID returns a boolean if a field has been set.

### GetActionByUserEmail

`func (o *AuditLog) GetActionByUserEmail() string`

GetActionByUserEmail returns the ActionByUserEmail field if non-nil, zero value otherwise.

### GetActionByUserEmailOk

`func (o *AuditLog) GetActionByUserEmailOk() (*string, bool)`

GetActionByUserEmailOk returns a tuple with the ActionByUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionByUserEmail

`func (o *AuditLog) SetActionByUserEmail(v string)`

SetActionByUserEmail sets ActionByUserEmail field to given value.


### GetActionOnUserEmail

`func (o *AuditLog) GetActionOnUserEmail() string`

GetActionOnUserEmail returns the ActionOnUserEmail field if non-nil, zero value otherwise.

### GetActionOnUserEmailOk

`func (o *AuditLog) GetActionOnUserEmailOk() (*string, bool)`

GetActionOnUserEmailOk returns a tuple with the ActionOnUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOnUserEmail

`func (o *AuditLog) SetActionOnUserEmail(v string)`

SetActionOnUserEmail sets ActionOnUserEmail field to given value.

### HasActionOnUserEmail

`func (o *AuditLog) HasActionOnUserEmail() bool`

HasActionOnUserEmail returns a boolean if a field has been set.

### GetType

`func (o *AuditLog) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditLog) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditLog) SetType(v string)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *AuditLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetActivityData

`func (o *AuditLog) GetActivityData() map[string]interface{}`

GetActivityData returns the ActivityData field if non-nil, zero value otherwise.

### GetActivityDataOk

`func (o *AuditLog) GetActivityDataOk() (*map[string]interface{}, bool)`

GetActivityDataOk returns a tuple with the ActivityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityData

`func (o *AuditLog) SetActivityData(v map[string]interface{})`

SetActivityData sets ActivityData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


