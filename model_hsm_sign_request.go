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

// checks if the HSMSignRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HSMSignRequest{}

// HSMSignRequest Request body representing a sign-data request.
type HSMSignRequest struct {
	// An ethereum address.
	Address string `json:"address" validate:"regexp=^0[xX][a-fA-F0-9]{40}$"`
	// Is the data field an encapsulated EIP-712 typed message?
	IsTyped *bool `json:"isTyped,omitempty"`
	// Data to sign
	Data    string                 `json:"data"`
	ChainId *HSMSignRequestChainId `json:"chainId,omitempty"`
}

type _HSMSignRequest HSMSignRequest

// NewHSMSignRequest instantiates a new HSMSignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHSMSignRequest(address string, data string) *HSMSignRequest {
	this := HSMSignRequest{}
	this.Address = address
	this.Data = data
	return &this
}

// NewHSMSignRequestWithDefaults instantiates a new HSMSignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHSMSignRequestWithDefaults() *HSMSignRequest {
	this := HSMSignRequest{}
	return &this
}

// GetAddress returns the Address field value
func (o *HSMSignRequest) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *HSMSignRequest) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *HSMSignRequest) SetAddress(v string) {
	o.Address = v
}

// GetIsTyped returns the IsTyped field value if set, zero value otherwise.
func (o *HSMSignRequest) GetIsTyped() bool {
	if o == nil || IsNil(o.IsTyped) {
		var ret bool
		return ret
	}
	return *o.IsTyped
}

// GetIsTypedOk returns a tuple with the IsTyped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HSMSignRequest) GetIsTypedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTyped) {
		return nil, false
	}
	return o.IsTyped, true
}

// HasIsTyped returns a boolean if a field has been set.
func (o *HSMSignRequest) HasIsTyped() bool {
	if o != nil && !IsNil(o.IsTyped) {
		return true
	}

	return false
}

// SetIsTyped gets a reference to the given bool and assigns it to the IsTyped field.
func (o *HSMSignRequest) SetIsTyped(v bool) {
	o.IsTyped = &v
}

// GetData returns the Data field value
func (o *HSMSignRequest) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *HSMSignRequest) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *HSMSignRequest) SetData(v string) {
	o.Data = v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *HSMSignRequest) GetChainId() HSMSignRequestChainId {
	if o == nil || IsNil(o.ChainId) {
		var ret HSMSignRequestChainId
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HSMSignRequest) GetChainIdOk() (*HSMSignRequestChainId, bool) {
	if o == nil || IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *HSMSignRequest) HasChainId() bool {
	if o != nil && !IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given HSMSignRequestChainId and assigns it to the ChainId field.
func (o *HSMSignRequest) SetChainId(v HSMSignRequestChainId) {
	o.ChainId = &v
}

func (o HSMSignRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HSMSignRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.IsTyped) {
		toSerialize["isTyped"] = o.IsTyped
	}
	toSerialize["data"] = o.Data
	if !IsNil(o.ChainId) {
		toSerialize["chainId"] = o.ChainId
	}
	return toSerialize, nil
}

func (o *HSMSignRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"data",
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

	varHSMSignRequest := _HSMSignRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHSMSignRequest)

	if err != nil {
		return err
	}

	*o = HSMSignRequest(varHSMSignRequest)

	return err
}

type NullableHSMSignRequest struct {
	value *HSMSignRequest
	isSet bool
}

func (v NullableHSMSignRequest) Get() *HSMSignRequest {
	return v.value
}

func (v *NullableHSMSignRequest) Set(val *HSMSignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHSMSignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHSMSignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHSMSignRequest(val *HSMSignRequest) *NullableHSMSignRequest {
	return &NullableHSMSignRequest{value: val, isSet: true}
}

func (v NullableHSMSignRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHSMSignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
