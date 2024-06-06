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

// checks if the CountWalletTransactions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountWalletTransactions200Response{}

// CountWalletTransactions200Response struct for CountWalletTransactions200Response
type CountWalletTransactions200Response struct {
	// The status code.
	Status int64 `json:"status"`
	// The human-readable status message.
	Message string `json:"message"`
	// The transaction count.
	Result int64 `json:"result"`
}

type _CountWalletTransactions200Response CountWalletTransactions200Response

// NewCountWalletTransactions200Response instantiates a new CountWalletTransactions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountWalletTransactions200Response(status int64, message string, result int64) *CountWalletTransactions200Response {
	this := CountWalletTransactions200Response{}
	this.Status = status
	this.Message = message
	this.Result = result
	return &this
}

// NewCountWalletTransactions200ResponseWithDefaults instantiates a new CountWalletTransactions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountWalletTransactions200ResponseWithDefaults() *CountWalletTransactions200Response {
	this := CountWalletTransactions200Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *CountWalletTransactions200Response) GetStatus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CountWalletTransactions200Response) GetStatusOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CountWalletTransactions200Response) SetStatus(v int64) {
	o.Status = v
}

// GetMessage returns the Message field value
func (o *CountWalletTransactions200Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CountWalletTransactions200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CountWalletTransactions200Response) SetMessage(v string) {
	o.Message = v
}

// GetResult returns the Result field value
func (o *CountWalletTransactions200Response) GetResult() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CountWalletTransactions200Response) GetResultOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CountWalletTransactions200Response) SetResult(v int64) {
	o.Result = v
}

func (o CountWalletTransactions200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountWalletTransactions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["message"] = o.Message
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *CountWalletTransactions200Response) UnmarshalJSON(data []byte) (err error) {
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

	varCountWalletTransactions200Response := _CountWalletTransactions200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCountWalletTransactions200Response)

	if err != nil {
		return err
	}

	*o = CountWalletTransactions200Response(varCountWalletTransactions200Response)

	return err
}

type NullableCountWalletTransactions200Response struct {
	value *CountWalletTransactions200Response
	isSet bool
}

func (v NullableCountWalletTransactions200Response) Get() *CountWalletTransactions200Response {
	return v.value
}

func (v *NullableCountWalletTransactions200Response) Set(val *CountWalletTransactions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCountWalletTransactions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCountWalletTransactions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountWalletTransactions200Response(val *CountWalletTransactions200Response) *NullableCountWalletTransactions200Response {
	return &NullableCountWalletTransactions200Response{value: val, isSet: true}
}

func (v NullableCountWalletTransactions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountWalletTransactions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
