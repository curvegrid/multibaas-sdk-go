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

// checks if the ListContracts200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListContracts200Response{}

// ListContracts200Response struct for ListContracts200Response
type ListContracts200Response struct {
	// The status code.
	Status int64 `json:"status"`
	// The human-readable status message.
	Message string             `json:"message"`
	Result  []ContractOverview `json:"result"`
}

type _ListContracts200Response ListContracts200Response

// NewListContracts200Response instantiates a new ListContracts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListContracts200Response(status int64, message string, result []ContractOverview) *ListContracts200Response {
	this := ListContracts200Response{}
	this.Status = status
	this.Message = message
	this.Result = result
	return &this
}

// NewListContracts200ResponseWithDefaults instantiates a new ListContracts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListContracts200ResponseWithDefaults() *ListContracts200Response {
	this := ListContracts200Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *ListContracts200Response) GetStatus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ListContracts200Response) GetStatusOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ListContracts200Response) SetStatus(v int64) {
	o.Status = v
}

// GetMessage returns the Message field value
func (o *ListContracts200Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ListContracts200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ListContracts200Response) SetMessage(v string) {
	o.Message = v
}

// GetResult returns the Result field value
func (o *ListContracts200Response) GetResult() []ContractOverview {
	if o == nil {
		var ret []ContractOverview
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ListContracts200Response) GetResultOk() ([]ContractOverview, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *ListContracts200Response) SetResult(v []ContractOverview) {
	o.Result = v
}

func (o ListContracts200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["message"] = o.Message
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableListContracts200Response struct {
	value *ListContracts200Response
	isSet bool
}

func (v NullableListContracts200Response) Get() *ListContracts200Response {
	return v.value
}

func (v *NullableListContracts200Response) Set(val *ListContracts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListContracts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListContracts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListContracts200Response(val *ListContracts200Response) *NullableListContracts200Response {
	return &NullableListContracts200Response{value: val, isSet: true}
}

func (v NullableListContracts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListContracts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
