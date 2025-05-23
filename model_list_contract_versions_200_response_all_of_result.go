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

// checks if the ListContractVersions200ResponseAllOfResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListContractVersions200ResponseAllOfResult{}

// ListContractVersions200ResponseAllOfResult struct for ListContractVersions200ResponseAllOfResult
type ListContractVersions200ResponseAllOfResult struct {
	// An alias to easily identify and reference the entity in subsequent requests.
	Label    string   `json:"label" validate:"regexp=^[a-z0-9_-]+$"`
	Versions []string `json:"versions"`
}

type _ListContractVersions200ResponseAllOfResult ListContractVersions200ResponseAllOfResult

// NewListContractVersions200ResponseAllOfResult instantiates a new ListContractVersions200ResponseAllOfResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListContractVersions200ResponseAllOfResult(label string, versions []string) *ListContractVersions200ResponseAllOfResult {
	this := ListContractVersions200ResponseAllOfResult{}
	this.Label = label
	this.Versions = versions
	return &this
}

// NewListContractVersions200ResponseAllOfResultWithDefaults instantiates a new ListContractVersions200ResponseAllOfResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListContractVersions200ResponseAllOfResultWithDefaults() *ListContractVersions200ResponseAllOfResult {
	this := ListContractVersions200ResponseAllOfResult{}
	return &this
}

// GetLabel returns the Label field value
func (o *ListContractVersions200ResponseAllOfResult) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ListContractVersions200ResponseAllOfResult) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ListContractVersions200ResponseAllOfResult) SetLabel(v string) {
	o.Label = v
}

// GetVersions returns the Versions field value
func (o *ListContractVersions200ResponseAllOfResult) GetVersions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *ListContractVersions200ResponseAllOfResult) GetVersionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Versions, true
}

// SetVersions sets field value
func (o *ListContractVersions200ResponseAllOfResult) SetVersions(v []string) {
	o.Versions = v
}

func (o ListContractVersions200ResponseAllOfResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["versions"] = o.Versions
	return toSerialize, nil
}

type NullableListContractVersions200ResponseAllOfResult struct {
	value *ListContractVersions200ResponseAllOfResult
	isSet bool
}

func (v NullableListContractVersions200ResponseAllOfResult) Get() *ListContractVersions200ResponseAllOfResult {
	return v.value
}

func (v *NullableListContractVersions200ResponseAllOfResult) Set(val *ListContractVersions200ResponseAllOfResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListContractVersions200ResponseAllOfResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListContractVersions200ResponseAllOfResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListContractVersions200ResponseAllOfResult(val *ListContractVersions200ResponseAllOfResult) *NullableListContractVersions200ResponseAllOfResult {
	return &NullableListContractVersions200ResponseAllOfResult{value: val, isSet: true}
}

func (v NullableListContractVersions200ResponseAllOfResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListContractVersions200ResponseAllOfResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
