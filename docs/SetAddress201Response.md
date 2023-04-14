# SetAddress201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**Address**](Address.md) |  | 

## Methods

### NewSetAddress201Response

`func NewSetAddress201Response(status int64, message string, result Address, ) *SetAddress201Response`

NewSetAddress201Response instantiates a new SetAddress201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAddress201ResponseWithDefaults

`func NewSetAddress201ResponseWithDefaults() *SetAddress201Response`

NewSetAddress201ResponseWithDefaults instantiates a new SetAddress201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SetAddress201Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SetAddress201Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SetAddress201Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *SetAddress201Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SetAddress201Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SetAddress201Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *SetAddress201Response) GetResult() Address`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SetAddress201Response) GetResultOk() (*Address, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SetAddress201Response) SetResult(v Address)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


