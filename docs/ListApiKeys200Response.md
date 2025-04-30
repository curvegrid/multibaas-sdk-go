# ListApiKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**[]ApiKey**](ApiKey.md) |  | 

## Methods

### NewListApiKeys200Response

`func NewListApiKeys200Response(status int64, message string, result []ApiKey, ) *ListApiKeys200Response`

NewListApiKeys200Response instantiates a new ListApiKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApiKeys200ResponseWithDefaults

`func NewListApiKeys200ResponseWithDefaults() *ListApiKeys200Response`

NewListApiKeys200ResponseWithDefaults instantiates a new ListApiKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListApiKeys200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListApiKeys200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListApiKeys200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ListApiKeys200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListApiKeys200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListApiKeys200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *ListApiKeys200Response) GetResult() []ApiKey`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListApiKeys200Response) GetResultOk() (*[]ApiKey, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListApiKeys200Response) SetResult(v []ApiKey)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


