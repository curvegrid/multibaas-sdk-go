/*
MultiBaas API

MultiBaas's REST APIv0.

API version: 0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package multibaas

import (
	"encoding/json"
)

// checks if the CORSOrigin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CORSOrigin{}

// CORSOrigin CORS Origin
type CORSOrigin struct {
	Id *int64 `json:"id,omitempty"`
	// The CORS Origin
	Origin *string `json:"origin,omitempty"`
}

// NewCORSOrigin instantiates a new CORSOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCORSOrigin() *CORSOrigin {
	this := CORSOrigin{}
	return &this
}

// NewCORSOriginWithDefaults instantiates a new CORSOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCORSOriginWithDefaults() *CORSOrigin {
	this := CORSOrigin{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CORSOrigin) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CORSOrigin) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CORSOrigin) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CORSOrigin) SetId(v int64) {
	o.Id = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *CORSOrigin) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CORSOrigin) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *CORSOrigin) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *CORSOrigin) SetOrigin(v string) {
	o.Origin = &v
}

func (o CORSOrigin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	return toSerialize, nil
}

type NullableCORSOrigin struct {
	value *CORSOrigin
	isSet bool
}

func (v NullableCORSOrigin) Get() *CORSOrigin {
	return v.value
}

func (v *NullableCORSOrigin) Set(val *CORSOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableCORSOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableCORSOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCORSOrigin(val *CORSOrigin) *NullableCORSOrigin {
	return &NullableCORSOrigin{value: val, isSet: true}
}

func (v NullableCORSOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCORSOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
