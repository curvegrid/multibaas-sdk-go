/*
MultiBaas API

MultiBaas's REST APIv0.

API version: 0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package multibaas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the WebhookEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEvent{}

// WebhookEvent struct for WebhookEvent
type WebhookEvent struct {
	// The ID of the webhook event.
	Id        int64             `json:"id"`
	EventType WebhookEventsType `json:"eventType"`
	// The data associated with the event.
	Data map[string]interface{} `json:"data"`
	// The time the webhook event was created.
	CreatedAt time.Time `json:"createdAt"`
	// The time the webhook event was last updated.
	DeliveredAt *time.Time `json:"deliveredAt,omitempty"`
}

type _WebhookEvent WebhookEvent

// NewWebhookEvent instantiates a new WebhookEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEvent(id int64, eventType WebhookEventsType, data map[string]interface{}, createdAt time.Time) *WebhookEvent {
	this := WebhookEvent{}
	this.Id = id
	this.EventType = eventType
	this.Data = data
	this.CreatedAt = createdAt
	return &this
}

// NewWebhookEventWithDefaults instantiates a new WebhookEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEventWithDefaults() *WebhookEvent {
	this := WebhookEvent{}
	return &this
}

// GetId returns the Id field value
func (o *WebhookEvent) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookEvent) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookEvent) SetId(v int64) {
	o.Id = v
}

// GetEventType returns the EventType field value
func (o *WebhookEvent) GetEventType() WebhookEventsType {
	if o == nil {
		var ret WebhookEventsType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *WebhookEvent) GetEventTypeOk() (*WebhookEventsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *WebhookEvent) SetEventType(v WebhookEventsType) {
	o.EventType = v
}

// GetData returns the Data field value
func (o *WebhookEvent) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WebhookEvent) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *WebhookEvent) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WebhookEvent) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WebhookEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeliveredAt returns the DeliveredAt field value if set, zero value otherwise.
func (o *WebhookEvent) GetDeliveredAt() time.Time {
	if o == nil || IsNil(o.DeliveredAt) {
		var ret time.Time
		return ret
	}
	return *o.DeliveredAt
}

// GetDeliveredAtOk returns a tuple with the DeliveredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookEvent) GetDeliveredAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeliveredAt) {
		return nil, false
	}
	return o.DeliveredAt, true
}

// HasDeliveredAt returns a boolean if a field has been set.
func (o *WebhookEvent) HasDeliveredAt() bool {
	if o != nil && !IsNil(o.DeliveredAt) {
		return true
	}

	return false
}

// SetDeliveredAt gets a reference to the given time.Time and assigns it to the DeliveredAt field.
func (o *WebhookEvent) SetDeliveredAt(v time.Time) {
	o.DeliveredAt = &v
}

func (o WebhookEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["eventType"] = o.EventType
	toSerialize["data"] = o.Data
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.DeliveredAt) {
		toSerialize["deliveredAt"] = o.DeliveredAt
	}
	return toSerialize, nil
}

func (o *WebhookEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"eventType",
		"data",
		"createdAt",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWebhookEvent := _WebhookEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookEvent)

	if err != nil {
		return err
	}

	*o = WebhookEvent(varWebhookEvent)

	return err
}

type NullableWebhookEvent struct {
	value *WebhookEvent
	isSet bool
}

func (v NullableWebhookEvent) Get() *WebhookEvent {
	return v.value
}

func (v *NullableWebhookEvent) Set(val *WebhookEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEvent(val *WebhookEvent) *NullableWebhookEvent {
	return &NullableWebhookEvent{value: val, isSet: true}
}

func (v NullableWebhookEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
