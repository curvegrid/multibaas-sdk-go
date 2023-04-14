# ListHsmWallets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**[]StandaloneWallet**](StandaloneWallet.md) |  | 

## Methods

### NewListHsmWallets200Response

`func NewListHsmWallets200Response(status int64, message string, result []StandaloneWallet, ) *ListHsmWallets200Response`

NewListHsmWallets200Response instantiates a new ListHsmWallets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHsmWallets200ResponseWithDefaults

`func NewListHsmWallets200ResponseWithDefaults() *ListHsmWallets200Response`

NewListHsmWallets200ResponseWithDefaults instantiates a new ListHsmWallets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListHsmWallets200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListHsmWallets200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListHsmWallets200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ListHsmWallets200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListHsmWallets200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListHsmWallets200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *ListHsmWallets200Response) GetResult() []StandaloneWallet`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListHsmWallets200Response) GetResultOk() (*[]StandaloneWallet, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListHsmWallets200Response) SetResult(v []StandaloneWallet)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


