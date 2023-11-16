# WebhookEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the webhook event. | 
**EventType** | [**WebhookEventsType**](WebhookEventsType.md) |  | 
**Data** | **map[string]interface{}** | The data associated with the event. | 
**CreatedAt** | **time.Time** | The time the webhook event was created. | 
**DeliveredAt** | Pointer to **time.Time** | The time the webhook event was last updated. | [optional] 

## Methods

### NewWebhookEvent

`func NewWebhookEvent(id int64, eventType WebhookEventsType, data map[string]interface{}, createdAt time.Time, ) *WebhookEvent`

NewWebhookEvent instantiates a new WebhookEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventWithDefaults

`func NewWebhookEventWithDefaults() *WebhookEvent`

NewWebhookEventWithDefaults instantiates a new WebhookEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookEvent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookEvent) SetId(v int64)`

SetId sets Id field to given value.


### GetEventType

`func (o *WebhookEvent) GetEventType() WebhookEventsType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookEvent) GetEventTypeOk() (*WebhookEventsType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookEvent) SetEventType(v WebhookEventsType)`

SetEventType sets EventType field to given value.


### GetData

`func (o *WebhookEvent) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookEvent) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookEvent) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetCreatedAt

`func (o *WebhookEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeliveredAt

`func (o *WebhookEvent) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *WebhookEvent) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *WebhookEvent) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *WebhookEvent) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


