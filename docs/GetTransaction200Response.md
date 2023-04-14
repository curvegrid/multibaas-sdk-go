# GetTransaction200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**TransactionData**](TransactionData.md) |  | 

## Methods

### NewGetTransaction200Response

`func NewGetTransaction200Response(status int64, message string, result TransactionData, ) *GetTransaction200Response`

NewGetTransaction200Response instantiates a new GetTransaction200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransaction200ResponseWithDefaults

`func NewGetTransaction200ResponseWithDefaults() *GetTransaction200Response`

NewGetTransaction200ResponseWithDefaults instantiates a new GetTransaction200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetTransaction200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTransaction200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTransaction200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *GetTransaction200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetTransaction200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetTransaction200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *GetTransaction200Response) GetResult() TransactionData`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetTransaction200Response) GetResultOk() (*TransactionData, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetTransaction200Response) SetResult(v TransactionData)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


