# SignData200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**HSMSignResponse**](HSMSignResponse.md) |  | 

## Methods

### NewSignData200Response

`func NewSignData200Response(status int64, message string, result HSMSignResponse, ) *SignData200Response`

NewSignData200Response instantiates a new SignData200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignData200ResponseWithDefaults

`func NewSignData200ResponseWithDefaults() *SignData200Response`

NewSignData200ResponseWithDefaults instantiates a new SignData200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SignData200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SignData200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SignData200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *SignData200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SignData200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SignData200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *SignData200Response) GetResult() HSMSignResponse`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SignData200Response) GetResultOk() (*HSMSignResponse, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SignData200Response) SetResult(v HSMSignResponse)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


