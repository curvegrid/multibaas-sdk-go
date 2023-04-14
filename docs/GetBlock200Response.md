# GetBlock200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**Block**](Block.md) |  | 

## Methods

### NewGetBlock200Response

`func NewGetBlock200Response(status int64, message string, result Block, ) *GetBlock200Response`

NewGetBlock200Response instantiates a new GetBlock200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlock200ResponseWithDefaults

`func NewGetBlock200ResponseWithDefaults() *GetBlock200Response`

NewGetBlock200ResponseWithDefaults instantiates a new GetBlock200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetBlock200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBlock200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBlock200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *GetBlock200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetBlock200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetBlock200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *GetBlock200Response) GetResult() Block`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetBlock200Response) GetResultOk() (*Block, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetBlock200Response) SetResult(v Block)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


