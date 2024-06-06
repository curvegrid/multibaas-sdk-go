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

// checks if the APIKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIKey{}

// APIKey An API key.
type APIKey struct {
	// A label.
	Label string `json:"label"`
	Id    int64  `json:"id"`
	// The time the API key was created.
	CreatedAt time.Time `json:"createdAt"`
	// The time the API key was last used.
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	// The ID of the user that created the API key.
	CreatedBy int64 `json:"createdBy"`
	// The signature of the API key.
	Signature string `json:"signature"`
}

type _APIKey APIKey

// NewAPIKey instantiates a new APIKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIKey(label string, id int64, createdAt time.Time, createdBy int64, signature string) *APIKey {
	this := APIKey{}
	this.Label = label
	this.Id = id
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Signature = signature
	return &this
}

// NewAPIKeyWithDefaults instantiates a new APIKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIKeyWithDefaults() *APIKey {
	this := APIKey{}
	return &this
}

// GetLabel returns the Label field value
func (o *APIKey) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *APIKey) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *APIKey) SetLabel(v string) {
	o.Label = v
}

// GetId returns the Id field value
func (o *APIKey) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *APIKey) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *APIKey) SetId(v int64) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *APIKey) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *APIKey) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *APIKey) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *APIKey) GetLastUsedAt() time.Time {
	if o == nil || IsNil(o.LastUsedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIKey) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUsedAt) {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *APIKey) HasLastUsedAt() bool {
	if o != nil && !IsNil(o.LastUsedAt) {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *APIKey) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *APIKey) GetCreatedBy() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *APIKey) GetCreatedByOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *APIKey) SetCreatedBy(v int64) {
	o.CreatedBy = v
}

// GetSignature returns the Signature field value
func (o *APIKey) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *APIKey) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *APIKey) SetSignature(v string) {
	o.Signature = v
}

func (o APIKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["id"] = o.Id
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.LastUsedAt) {
		toSerialize["lastUsedAt"] = o.LastUsedAt
	}
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["signature"] = o.Signature
	return toSerialize, nil
}

func (o *APIKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"id",
		"createdAt",
		"createdBy",
		"signature",
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

	varAPIKey := _APIKey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAPIKey)

	if err != nil {
		return err
	}

	*o = APIKey(varAPIKey)

	return err
}

type NullableAPIKey struct {
	value *APIKey
	isSet bool
}

func (v NullableAPIKey) Get() *APIKey {
	return v.value
}

func (v *NullableAPIKey) Set(val *APIKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIKey(val *APIKey) *NullableAPIKey {
	return &NullableAPIKey{value: val, isSet: true}
}

func (v NullableAPIKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
