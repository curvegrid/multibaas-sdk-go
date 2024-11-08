# SubmitSignedTransaction200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**SignedTransactionResponse**](SignedTransactionResponse.md) |  | 

## Methods

### NewSubmitSignedTransaction200Response

`func NewSubmitSignedTransaction200Response(status int64, message string, result SignedTransactionResponse, ) *SubmitSignedTransaction200Response`

NewSubmitSignedTransaction200Response instantiates a new SubmitSignedTransaction200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitSignedTransaction200ResponseWithDefaults

`func NewSubmitSignedTransaction200ResponseWithDefaults() *SubmitSignedTransaction200Response`

NewSubmitSignedTransaction200ResponseWithDefaults instantiates a new SubmitSignedTransaction200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SubmitSignedTransaction200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitSignedTransaction200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitSignedTransaction200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *SubmitSignedTransaction200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SubmitSignedTransaction200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SubmitSignedTransaction200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *SubmitSignedTransaction200Response) GetResult() SignedTransactionResponse`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SubmitSignedTransaction200Response) GetResultOk() (*SignedTransactionResponse, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SubmitSignedTransaction200Response) SetResult(v SignedTransactionResponse)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


