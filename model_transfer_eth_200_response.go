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

// checks if the TransferEth200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferEth200Response{}

// TransferEth200Response struct for TransferEth200Response
type TransferEth200Response struct {
	// The status code.
	Status int64 `json:"status"`
	// The human-readable status message.
	Message string            `json:"message"`
	Result  TransactionToSign `json:"result"`
}

type _TransferEth200Response TransferEth200Response

// NewTransferEth200Response instantiates a new TransferEth200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferEth200Response(status int64, message string, result TransactionToSign) *TransferEth200Response {
	this := TransferEth200Response{}
	this.Status = status
	this.Message = message
	this.Result = result
	return &this
}

// NewTransferEth200ResponseWithDefaults instantiates a new TransferEth200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferEth200ResponseWithDefaults() *TransferEth200Response {
	this := TransferEth200Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *TransferEth200Response) GetStatus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransferEth200Response) GetStatusOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransferEth200Response) SetStatus(v int64) {
	o.Status = v
}

// GetMessage returns the Message field value
func (o *TransferEth200Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TransferEth200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *TransferEth200Response) SetMessage(v string) {
	o.Message = v
}

// GetResult returns the Result field value
func (o *TransferEth200Response) GetResult() TransactionToSign {
	if o == nil {
		var ret TransactionToSign
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *TransferEth200Response) GetResultOk() (*TransactionToSign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *TransferEth200Response) SetResult(v TransactionToSign) {
	o.Result = v
}

func (o TransferEth200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["message"] = o.Message
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableTransferEth200Response struct {
	value *TransferEth200Response
	isSet bool
}

func (v NullableTransferEth200Response) Get() *TransferEth200Response {
	return v.value
}

func (v *NullableTransferEth200Response) Set(val *TransferEth200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferEth200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferEth200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferEth200Response(val *TransferEth200Response) *NullableTransferEth200Response {
	return &NullableTransferEth200Response{value: val, isSet: true}
}

func (v NullableTransferEth200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferEth200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
