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

// checks if the CreateApiKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateApiKeyRequest{}

// CreateApiKeyRequest struct for CreateApiKeyRequest
type CreateApiKeyRequest struct {
	// An alias to easily identify and reference the entity in subsequent requests.
	Label    string  `json:"label" validate:"regexp=^[a-z0-9_-]+$"`
	GroupIDs []int64 `json:"groupIDs,omitempty"`
}

type _CreateApiKeyRequest CreateApiKeyRequest

// NewCreateApiKeyRequest instantiates a new CreateApiKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiKeyRequest(label string) *CreateApiKeyRequest {
	this := CreateApiKeyRequest{}
	this.Label = label
	return &this
}

// NewCreateApiKeyRequestWithDefaults instantiates a new CreateApiKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiKeyRequestWithDefaults() *CreateApiKeyRequest {
	this := CreateApiKeyRequest{}
	return &this
}

// GetLabel returns the Label field value
func (o *CreateApiKeyRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CreateApiKeyRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CreateApiKeyRequest) SetLabel(v string) {
	o.Label = v
}

// GetGroupIDs returns the GroupIDs field value if set, zero value otherwise.
func (o *CreateApiKeyRequest) GetGroupIDs() []int64 {
	if o == nil || IsNil(o.GroupIDs) {
		var ret []int64
		return ret
	}
	return o.GroupIDs
}

// GetGroupIDsOk returns a tuple with the GroupIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiKeyRequest) GetGroupIDsOk() ([]int64, bool) {
	if o == nil || IsNil(o.GroupIDs) {
		return nil, false
	}
	return o.GroupIDs, true
}

// HasGroupIDs returns a boolean if a field has been set.
func (o *CreateApiKeyRequest) HasGroupIDs() bool {
	if o != nil && !IsNil(o.GroupIDs) {
		return true
	}

	return false
}

// SetGroupIDs gets a reference to the given []int64 and assigns it to the GroupIDs field.
func (o *CreateApiKeyRequest) SetGroupIDs(v []int64) {
	o.GroupIDs = v
}

func (o CreateApiKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateApiKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	if !IsNil(o.GroupIDs) {
		toSerialize["groupIDs"] = o.GroupIDs
	}
	return toSerialize, nil
}

func (o *CreateApiKeyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
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

	varCreateApiKeyRequest := _CreateApiKeyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateApiKeyRequest)

	if err != nil {
		return err
	}

	*o = CreateApiKeyRequest(varCreateApiKeyRequest)

	return err
}

type NullableCreateApiKeyRequest struct {
	value *CreateApiKeyRequest
	isSet bool
}

func (v NullableCreateApiKeyRequest) Get() *CreateApiKeyRequest {
	return v.value
}

func (v *NullableCreateApiKeyRequest) Set(val *CreateApiKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiKeyRequest(val *CreateApiKeyRequest) *NullableCreateApiKeyRequest {
	return &NullableCreateApiKeyRequest{value: val, isSet: true}
}

func (v NullableCreateApiKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
