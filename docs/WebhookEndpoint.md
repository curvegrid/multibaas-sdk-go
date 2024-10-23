# WebhookEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL to send the webhook to. | 
**Label** | **string** | An alias to easily identify and reference the entity in subsequent requests. | 
**Subscriptions** | [**[]WebhookEventsType**](WebhookEventsType.md) | The events to subscribe to. | 
**Id** | **int64** |  | 
**NextAttempt** | Pointer to **time.Time** | The time the next attempt will be made. | [optional] 
**LastAttempt** | Pointer to **time.Time** | The time the last attempt was made. | [optional] 
**FailedCalls** | **int64** | The number of failed webhook endpoint calls since the last successful call. | 
**LastError** | Pointer to **string** | The last error received from the webhook endpoint. | [optional] 
**CreatedAt** | **time.Time** | The time the webhook was created. | 
**UpdatedAt** | **time.Time** | The time the webhook was last updated. | 
**Secret** | **string** | The secret key used to sign the webhook. | 

## Methods

### NewWebhookEndpoint

`func NewWebhookEndpoint(url string, label string, subscriptions []WebhookEventsType, id int64, failedCalls int64, createdAt time.Time, updatedAt time.Time, secret string, ) *WebhookEndpoint`

NewWebhookEndpoint instantiates a new WebhookEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEndpointWithDefaults

`func NewWebhookEndpointWithDefaults() *WebhookEndpoint`

NewWebhookEndpointWithDefaults instantiates a new WebhookEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLabel

`func (o *WebhookEndpoint) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WebhookEndpoint) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WebhookEndpoint) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubscriptions

`func (o *WebhookEndpoint) GetSubscriptions() []WebhookEventsType`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *WebhookEndpoint) GetSubscriptionsOk() (*[]WebhookEventsType, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *WebhookEndpoint) SetSubscriptions(v []WebhookEventsType)`

SetSubscriptions sets Subscriptions field to given value.


### GetId

`func (o *WebhookEndpoint) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookEndpoint) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookEndpoint) SetId(v int64)`

SetId sets Id field to given value.


### GetNextAttempt

`func (o *WebhookEndpoint) GetNextAttempt() time.Time`

GetNextAttempt returns the NextAttempt field if non-nil, zero value otherwise.

### GetNextAttemptOk

`func (o *WebhookEndpoint) GetNextAttemptOk() (*time.Time, bool)`

GetNextAttemptOk returns a tuple with the NextAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAttempt

`func (o *WebhookEndpoint) SetNextAttempt(v time.Time)`

SetNextAttempt sets NextAttempt field to given value.

### HasNextAttempt

`func (o *WebhookEndpoint) HasNextAttempt() bool`

HasNextAttempt returns a boolean if a field has been set.

### GetLastAttempt

`func (o *WebhookEndpoint) GetLastAttempt() time.Time`

GetLastAttempt returns the LastAttempt field if non-nil, zero value otherwise.

### GetLastAttemptOk

`func (o *WebhookEndpoint) GetLastAttemptOk() (*time.Time, bool)`

GetLastAttemptOk returns a tuple with the LastAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttempt

`func (o *WebhookEndpoint) SetLastAttempt(v time.Time)`

SetLastAttempt sets LastAttempt field to given value.

### HasLastAttempt

`func (o *WebhookEndpoint) HasLastAttempt() bool`

HasLastAttempt returns a boolean if a field has been set.

### GetFailedCalls

`func (o *WebhookEndpoint) GetFailedCalls() int64`

GetFailedCalls returns the FailedCalls field if non-nil, zero value otherwise.

### GetFailedCallsOk

`func (o *WebhookEndpoint) GetFailedCallsOk() (*int64, bool)`

GetFailedCallsOk returns a tuple with the FailedCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCalls

`func (o *WebhookEndpoint) SetFailedCalls(v int64)`

SetFailedCalls sets FailedCalls field to given value.


### GetLastError

`func (o *WebhookEndpoint) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *WebhookEndpoint) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *WebhookEndpoint) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *WebhookEndpoint) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookEndpoint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookEndpoint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookEndpoint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WebhookEndpoint) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookEndpoint) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookEndpoint) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSecret

`func (o *WebhookEndpoint) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookEndpoint) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookEndpoint) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


