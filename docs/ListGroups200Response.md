# ListGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**[]Group**](Group.md) |  | 

## Methods

### NewListGroups200Response

`func NewListGroups200Response(status int64, message string, result []Group, ) *ListGroups200Response`

NewListGroups200Response instantiates a new ListGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGroups200ResponseWithDefaults

`func NewListGroups200ResponseWithDefaults() *ListGroups200Response`

NewListGroups200ResponseWithDefaults instantiates a new ListGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListGroups200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListGroups200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListGroups200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ListGroups200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListGroups200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListGroups200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *ListGroups200Response) GetResult() []Group`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListGroups200Response) GetResultOk() (*[]Group, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListGroups200Response) SetResult(v []Group)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


