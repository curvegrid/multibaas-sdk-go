/*
MultiBaas API

MultiBaas's REST APIv0.

API version: 0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package multibaas

import (
	"encoding/json"
	"fmt"
)

// FieldType The field type in a select or filter.
type FieldType string

// List of FieldType
const (
	INPUT                  FieldType = "input"
	CONTRACT_LABEL         FieldType = "contract_label"
	CONTRACT_NAME          FieldType = "contract_name"
	CONTRACT_ADDRESS       FieldType = "contract_address"
	CONTRACT_ADDRESS_ALIAS FieldType = "contract_address_alias"
	BLOCK_NUMBER           FieldType = "block_number"
	TRIGGERED_AT           FieldType = "triggered_at"
	EVENT_SIGNATURE        FieldType = "event_signature"
	BLOCK_HASH             FieldType = "block_hash"
	TX_HASH                FieldType = "tx_hash"
	TX_FROM                FieldType = "tx_from"
)

// All allowed values of FieldType enum
var AllowedFieldTypeEnumValues = []FieldType{
	"input",
	"contract_label",
	"contract_name",
	"contract_address",
	"contract_address_alias",
	"block_number",
	"triggered_at",
	"event_signature",
	"block_hash",
	"tx_hash",
	"tx_from",
}

func (v *FieldType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldType(value)
	for _, existing := range AllowedFieldTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldType", value)
}

// NewFieldTypeFromValue returns a pointer to a valid FieldType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldTypeFromValue(v string) (*FieldType, error) {
	ev := FieldType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldType: valid values are %v", v, AllowedFieldTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldType) IsValid() bool {
	for _, existing := range AllowedFieldTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FieldType value
func (v FieldType) Ptr() *FieldType {
	return &v
}

type NullableFieldType struct {
	value *FieldType
	isSet bool
}

func (v NullableFieldType) Get() *FieldType {
	return v.value
}

func (v *NullableFieldType) Set(val *FieldType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldType(val *FieldType) *NullableFieldType {
	return &NullableFieldType{value: val, isSet: true}
}

func (v NullableFieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
