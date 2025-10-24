# ListInvites200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**[]Invite**](Invite.md) |  | 

## Methods

### NewListInvites200Response

`func NewListInvites200Response(status int64, message string, result []Invite, ) *ListInvites200Response`

NewListInvites200Response instantiates a new ListInvites200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvites200ResponseWithDefaults

`func NewListInvites200ResponseWithDefaults() *ListInvites200Response`

NewListInvites200ResponseWithDefaults instantiates a new ListInvites200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListInvites200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListInvites200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListInvites200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ListInvites200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListInvites200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListInvites200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *ListInvites200Response) GetResult() []Invite`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListInvites200Response) GetResultOk() (*[]Invite, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListInvites200Response) SetResult(v []Invite)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


