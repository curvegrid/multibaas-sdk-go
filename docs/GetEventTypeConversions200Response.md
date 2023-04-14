# GetEventTypeConversions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**EventTypeConversionOptions**](EventTypeConversionOptions.md) |  | 

## Methods

### NewGetEventTypeConversions200Response

`func NewGetEventTypeConversions200Response(status int64, message string, result EventTypeConversionOptions, ) *GetEventTypeConversions200Response`

NewGetEventTypeConversions200Response instantiates a new GetEventTypeConversions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventTypeConversions200ResponseWithDefaults

`func NewGetEventTypeConversions200ResponseWithDefaults() *GetEventTypeConversions200Response`

NewGetEventTypeConversions200ResponseWithDefaults instantiates a new GetEventTypeConversions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetEventTypeConversions200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEventTypeConversions200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEventTypeConversions200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *GetEventTypeConversions200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetEventTypeConversions200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetEventTypeConversions200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *GetEventTypeConversions200Response) GetResult() EventTypeConversionOptions`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetEventTypeConversions200Response) GetResultOk() (*EventTypeConversionOptions, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetEventTypeConversions200Response) SetResult(v EventTypeConversionOptions)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


