# SavedEventQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Label** | **string** | An event query label. | 
**Query** | [**EventQuery**](EventQuery.md) |  | 
**IsSystem** | **bool** | Specifies if this a system-generated query which is not modifiable by the user. | 

## Methods

### NewSavedEventQuery

`func NewSavedEventQuery(id int64, label string, query EventQuery, isSystem bool, ) *SavedEventQuery`

NewSavedEventQuery instantiates a new SavedEventQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedEventQueryWithDefaults

`func NewSavedEventQueryWithDefaults() *SavedEventQuery`

NewSavedEventQueryWithDefaults instantiates a new SavedEventQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SavedEventQuery) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SavedEventQuery) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SavedEventQuery) SetId(v int64)`

SetId sets Id field to given value.


### GetLabel

`func (o *SavedEventQuery) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SavedEventQuery) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SavedEventQuery) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetQuery

`func (o *SavedEventQuery) GetQuery() EventQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SavedEventQuery) GetQueryOk() (*EventQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SavedEventQuery) SetQuery(v EventQuery)`

SetQuery sets Query field to given value.


### GetIsSystem

`func (o *SavedEventQuery) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *SavedEventQuery) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *SavedEventQuery) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


