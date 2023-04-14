# CreateHsmKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**AzureWallet**](AzureWallet.md) |  | 

## Methods

### NewCreateHsmKey200Response

`func NewCreateHsmKey200Response(status int64, message string, result AzureWallet, ) *CreateHsmKey200Response`

NewCreateHsmKey200Response instantiates a new CreateHsmKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHsmKey200ResponseWithDefaults

`func NewCreateHsmKey200ResponseWithDefaults() *CreateHsmKey200Response`

NewCreateHsmKey200ResponseWithDefaults instantiates a new CreateHsmKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateHsmKey200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateHsmKey200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateHsmKey200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *CreateHsmKey200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateHsmKey200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateHsmKey200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *CreateHsmKey200Response) GetResult() AzureWallet`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateHsmKey200Response) GetResultOk() (*AzureWallet, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateHsmKey200Response) SetResult(v AzureWallet)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


