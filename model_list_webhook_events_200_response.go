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
)

// checks if the ListWebhookEvents200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWebhookEvents200Response{}

// ListWebhookEvents200Response struct for ListWebhookEvents200Response
type ListWebhookEvents200Response struct {
	// The status code.
	Status int64 `json:"status"`
	// The human-readable status message.
	Message string         `json:"message"`
	Result  []WebhookEvent `json:"result"`
}

type _ListWebhookEvents200Response ListWebhookEvents200Response

// NewListWebhookEvents200Response instantiates a new ListWebhookEvents200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWebhookEvents200Response(status int64, message string, result []WebhookEvent) *ListWebhookEvents200Response {
	this := ListWebhookEvents200Response{}
	this.Status = status
	this.Message = message
	this.Result = result
	return &this
}

// NewListWebhookEvents200ResponseWithDefaults instantiates a new ListWebhookEvents200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWebhookEvents200ResponseWithDefaults() *ListWebhookEvents200Response {
	this := ListWebhookEvents200Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *ListWebhookEvents200Response) GetStatus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ListWebhookEvents200Response) GetStatusOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ListWebhookEvents200Response) SetStatus(v int64) {
	o.Status = v
}

// GetMessage returns the Message field value
func (o *ListWebhookEvents200Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ListWebhookEvents200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ListWebhookEvents200Response) SetMessage(v string) {
	o.Message = v
}

// GetResult returns the Result field value
func (o *ListWebhookEvents200Response) GetResult() []WebhookEvent {
	if o == nil {
		var ret []WebhookEvent
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ListWebhookEvents200Response) GetResultOk() ([]WebhookEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *ListWebhookEvents200Response) SetResult(v []WebhookEvent) {
	o.Result = v
}

func (o ListWebhookEvents200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWebhookEvents200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["message"] = o.Message
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *ListWebhookEvents200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"message",
		"result",
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

	varListWebhookEvents200Response := _ListWebhookEvents200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListWebhookEvents200Response)

	if err != nil {
		return err
	}

	*o = ListWebhookEvents200Response(varListWebhookEvents200Response)

	return err
}

type NullableListWebhookEvents200Response struct {
	value *ListWebhookEvents200Response
	isSet bool
}

func (v NullableListWebhookEvents200Response) Get() *ListWebhookEvents200Response {
	return v.value
}

func (v *NullableListWebhookEvents200Response) Set(val *ListWebhookEvents200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListWebhookEvents200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListWebhookEvents200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWebhookEvents200Response(val *ListWebhookEvents200Response) *NullableListWebhookEvents200Response {
	return &NullableListWebhookEvents200Response{value: val, isSet: true}
}

func (v NullableListWebhookEvents200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWebhookEvents200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
