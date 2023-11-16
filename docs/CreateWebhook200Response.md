# CreateWebhook200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**WebhookEndpoint**](WebhookEndpoint.md) |  | 

## Methods

### NewCreateWebhook200Response

`func NewCreateWebhook200Response(status int64, message string, result WebhookEndpoint, ) *CreateWebhook200Response`

NewCreateWebhook200Response instantiates a new CreateWebhook200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhook200ResponseWithDefaults

`func NewCreateWebhook200ResponseWithDefaults() *CreateWebhook200Response`

NewCreateWebhook200ResponseWithDefaults instantiates a new CreateWebhook200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateWebhook200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateWebhook200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateWebhook200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *CreateWebhook200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateWebhook200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateWebhook200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *CreateWebhook200Response) GetResult() WebhookEndpoint`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateWebhook200Response) GetResultOk() (*WebhookEndpoint, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateWebhook200Response) SetResult(v WebhookEndpoint)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


