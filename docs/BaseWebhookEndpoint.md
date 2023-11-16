# BaseWebhookEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL to send the webhook to. | 
**Label** | **string** | A label. | 
**Subscriptions** | [**[]WebhookEventsType**](WebhookEventsType.md) | The events to subscribe to. | 

## Methods

### NewBaseWebhookEndpoint

`func NewBaseWebhookEndpoint(url string, label string, subscriptions []WebhookEventsType, ) *BaseWebhookEndpoint`

NewBaseWebhookEndpoint instantiates a new BaseWebhookEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseWebhookEndpointWithDefaults

`func NewBaseWebhookEndpointWithDefaults() *BaseWebhookEndpoint`

NewBaseWebhookEndpointWithDefaults instantiates a new BaseWebhookEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *BaseWebhookEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BaseWebhookEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BaseWebhookEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLabel

`func (o *BaseWebhookEndpoint) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BaseWebhookEndpoint) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BaseWebhookEndpoint) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubscriptions

`func (o *BaseWebhookEndpoint) GetSubscriptions() []WebhookEventsType`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *BaseWebhookEndpoint) GetSubscriptionsOk() (*[]WebhookEventsType, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *BaseWebhookEndpoint) SetSubscriptions(v []WebhookEventsType)`

SetSubscriptions sets Subscriptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


