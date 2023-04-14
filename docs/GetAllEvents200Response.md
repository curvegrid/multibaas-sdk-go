# GetAllEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**[]Event**](Event.md) |  | 

## Methods

### NewGetAllEvents200Response

`func NewGetAllEvents200Response(status int64, message string, result []Event, ) *GetAllEvents200Response`

NewGetAllEvents200Response instantiates a new GetAllEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllEvents200ResponseWithDefaults

`func NewGetAllEvents200ResponseWithDefaults() *GetAllEvents200Response`

NewGetAllEvents200ResponseWithDefaults instantiates a new GetAllEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetAllEvents200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAllEvents200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAllEvents200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *GetAllEvents200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetAllEvents200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetAllEvents200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *GetAllEvents200Response) GetResult() []Event`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAllEvents200Response) GetResultOk() (*[]Event, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAllEvents200Response) SetResult(v []Event)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


