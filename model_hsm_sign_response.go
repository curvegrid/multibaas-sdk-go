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

// checks if the HSMSignResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HSMSignResponse{}

// HSMSignResponse Response body representing a sign-data response.
type HSMSignResponse struct {
	PublicKey string `json:"publicKey"`
	Signature string `json:"signature"`
}

type _HSMSignResponse HSMSignResponse

// NewHSMSignResponse instantiates a new HSMSignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHSMSignResponse(publicKey string, signature string) *HSMSignResponse {
	this := HSMSignResponse{}
	this.PublicKey = publicKey
	this.Signature = signature
	return &this
}

// NewHSMSignResponseWithDefaults instantiates a new HSMSignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHSMSignResponseWithDefaults() *HSMSignResponse {
	this := HSMSignResponse{}
	return &this
}

// GetPublicKey returns the PublicKey field value
func (o *HSMSignResponse) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *HSMSignResponse) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *HSMSignResponse) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetSignature returns the Signature field value
func (o *HSMSignResponse) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *HSMSignResponse) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *HSMSignResponse) SetSignature(v string) {
	o.Signature = v
}

func (o HSMSignResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HSMSignResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["publicKey"] = o.PublicKey
	toSerialize["signature"] = o.Signature
	return toSerialize, nil
}

func (o *HSMSignResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"publicKey",
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

	varHSMSignResponse := _HSMSignResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHSMSignResponse)

	if err != nil {
		return err
	}

	*o = HSMSignResponse(varHSMSignResponse)

	return err
}

type NullableHSMSignResponse struct {
	value *HSMSignResponse
	isSet bool
}

func (v NullableHSMSignResponse) Get() *HSMSignResponse {
	return v.value
}

func (v *NullableHSMSignResponse) Set(val *HSMSignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHSMSignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHSMSignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHSMSignResponse(val *HSMSignResponse) *NullableHSMSignResponse {
	return &NullableHSMSignResponse{value: val, isSet: true}
}

func (v NullableHSMSignResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHSMSignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
