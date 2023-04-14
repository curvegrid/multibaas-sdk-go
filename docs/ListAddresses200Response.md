# ListAddresses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**[]AddressLabel**](AddressLabel.md) |  | 

## Methods

### NewListAddresses200Response

`func NewListAddresses200Response(status int64, message string, result []AddressLabel, ) *ListAddresses200Response`

NewListAddresses200Response instantiates a new ListAddresses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAddresses200ResponseWithDefaults

`func NewListAddresses200ResponseWithDefaults() *ListAddresses200Response`

NewListAddresses200ResponseWithDefaults instantiates a new ListAddresses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListAddresses200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListAddresses200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListAddresses200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ListAddresses200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListAddresses200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListAddresses200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *ListAddresses200Response) GetResult() []AddressLabel`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListAddresses200Response) GetResultOk() (*[]AddressLabel, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListAddresses200Response) SetResult(v []AddressLabel)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


