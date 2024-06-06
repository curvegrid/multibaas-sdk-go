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

// checks if the EventQueryField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventQueryField{}

// EventQueryField A single event field's query information.
type EventQueryField struct {
	Type FieldType `json:"type"`
	// The field name. Either `name` or `inputIndex` is required if `fieldType == \"input\"`.
	Name *string `json:"name,omitempty"`
	// The field's index, can be used in place of `name`.
	InputIndex NullableInt64 `json:"inputIndex,omitempty"`
	// The name that will be used to return the field.
	Alias *string `json:"alias,omitempty"`
	// The type of aggregation to perform on the field.
	Aggregator NullableString `json:"aggregator,omitempty"`
}

type _EventQueryField EventQueryField

// NewEventQueryField instantiates a new EventQueryField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventQueryField(type_ FieldType) *EventQueryField {
	this := EventQueryField{}
	this.Type = type_
	return &this
}

// NewEventQueryFieldWithDefaults instantiates a new EventQueryField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventQueryFieldWithDefaults() *EventQueryField {
	this := EventQueryField{}
	return &this
}

// GetType returns the Type field value
func (o *EventQueryField) GetType() FieldType {
	if o == nil {
		var ret FieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventQueryField) GetTypeOk() (*FieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventQueryField) SetType(v FieldType) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EventQueryField) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventQueryField) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EventQueryField) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EventQueryField) SetName(v string) {
	o.Name = &v
}

// GetInputIndex returns the InputIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventQueryField) GetInputIndex() int64 {
	if o == nil || IsNil(o.InputIndex.Get()) {
		var ret int64
		return ret
	}
	return *o.InputIndex.Get()
}

// GetInputIndexOk returns a tuple with the InputIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventQueryField) GetInputIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputIndex.Get(), o.InputIndex.IsSet()
}

// HasInputIndex returns a boolean if a field has been set.
func (o *EventQueryField) HasInputIndex() bool {
	if o != nil && o.InputIndex.IsSet() {
		return true
	}

	return false
}

// SetInputIndex gets a reference to the given NullableInt64 and assigns it to the InputIndex field.
func (o *EventQueryField) SetInputIndex(v int64) {
	o.InputIndex.Set(&v)
}

// SetInputIndexNil sets the value for InputIndex to be an explicit nil
func (o *EventQueryField) SetInputIndexNil() {
	o.InputIndex.Set(nil)
}

// UnsetInputIndex ensures that no value is present for InputIndex, not even an explicit nil
func (o *EventQueryField) UnsetInputIndex() {
	o.InputIndex.Unset()
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *EventQueryField) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventQueryField) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *EventQueryField) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *EventQueryField) SetAlias(v string) {
	o.Alias = &v
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventQueryField) GetAggregator() string {
	if o == nil || IsNil(o.Aggregator.Get()) {
		var ret string
		return ret
	}
	return *o.Aggregator.Get()
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventQueryField) GetAggregatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aggregator.Get(), o.Aggregator.IsSet()
}

// HasAggregator returns a boolean if a field has been set.
func (o *EventQueryField) HasAggregator() bool {
	if o != nil && o.Aggregator.IsSet() {
		return true
	}

	return false
}

// SetAggregator gets a reference to the given NullableString and assigns it to the Aggregator field.
func (o *EventQueryField) SetAggregator(v string) {
	o.Aggregator.Set(&v)
}

// SetAggregatorNil sets the value for Aggregator to be an explicit nil
func (o *EventQueryField) SetAggregatorNil() {
	o.Aggregator.Set(nil)
}

// UnsetAggregator ensures that no value is present for Aggregator, not even an explicit nil
func (o *EventQueryField) UnsetAggregator() {
	o.Aggregator.Unset()
}

func (o EventQueryField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventQueryField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.InputIndex.IsSet() {
		toSerialize["inputIndex"] = o.InputIndex.Get()
	}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if o.Aggregator.IsSet() {
		toSerialize["aggregator"] = o.Aggregator.Get()
	}
	return toSerialize, nil
}

func (o *EventQueryField) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varEventQueryField := _EventQueryField{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventQueryField)

	if err != nil {
		return err
	}

	*o = EventQueryField(varEventQueryField)

	return err
}

type NullableEventQueryField struct {
	value *EventQueryField
	isSet bool
}

func (v NullableEventQueryField) Get() *EventQueryField {
	return v.value
}

func (v *NullableEventQueryField) Set(val *EventQueryField) {
	v.value = val
	v.isSet = true
}

func (v NullableEventQueryField) IsSet() bool {
	return v.isSet
}

func (v *NullableEventQueryField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventQueryField(val *EventQueryField) *NullableEventQueryField {
	return &NullableEventQueryField{value: val, isSet: true}
}

func (v NullableEventQueryField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventQueryField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
