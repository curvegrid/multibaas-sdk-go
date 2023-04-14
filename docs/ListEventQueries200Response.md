# ListEventQueries200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**[]SavedEventQuery**](SavedEventQuery.md) |  | 

## Methods

### NewListEventQueries200Response

`func NewListEventQueries200Response(status int64, message string, result []SavedEventQuery, ) *ListEventQueries200Response`

NewListEventQueries200Response instantiates a new ListEventQueries200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEventQueries200ResponseWithDefaults

`func NewListEventQueries200ResponseWithDefaults() *ListEventQueries200Response`

NewListEventQueries200ResponseWithDefaults instantiates a new ListEventQueries200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListEventQueries200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListEventQueries200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListEventQueries200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ListEventQueries200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListEventQueries200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListEventQueries200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *ListEventQueries200Response) GetResult() []SavedEventQuery`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListEventQueries200Response) GetResultOk() (*[]SavedEventQuery, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListEventQueries200Response) SetResult(v []SavedEventQuery)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


