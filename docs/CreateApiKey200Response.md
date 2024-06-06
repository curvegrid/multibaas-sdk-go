# CreateApiKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**APIKeyWithSecret**](APIKeyWithSecret.md) |  | 

## Methods

### NewCreateApiKey200Response

`func NewCreateApiKey200Response(status int64, message string, result APIKeyWithSecret, ) *CreateApiKey200Response`

NewCreateApiKey200Response instantiates a new CreateApiKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKey200ResponseWithDefaults

`func NewCreateApiKey200ResponseWithDefaults() *CreateApiKey200Response`

NewCreateApiKey200ResponseWithDefaults instantiates a new CreateApiKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateApiKey200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateApiKey200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateApiKey200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *CreateApiKey200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateApiKey200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateApiKey200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *CreateApiKey200Response) GetResult() APIKeyWithSecret`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateApiKey200Response) GetResultOk() (*APIKeyWithSecret, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateApiKey200Response) SetResult(v APIKeyWithSecret)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


