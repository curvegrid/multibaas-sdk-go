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

// HSMSignRequestPersonalSignChainId - Optionally lock the message to a specific chain by encoding the chain ID in the signature per EIP-155.
type HSMSignRequestPersonalSignChainId struct {
	Int64  *int64
	String *string
}

// int64AsHSMSignRequestPersonalSignChainId is a convenience function that returns int64 wrapped in HSMSignRequestPersonalSignChainId
func Int64AsHSMSignRequestPersonalSignChainId(v *int64) HSMSignRequestPersonalSignChainId {
	return HSMSignRequestPersonalSignChainId{
		Int64: v,
	}
}

// stringAsHSMSignRequestPersonalSignChainId is a convenience function that returns string wrapped in HSMSignRequestPersonalSignChainId
func StringAsHSMSignRequestPersonalSignChainId(v *string) HSMSignRequestPersonalSignChainId {
	return HSMSignRequestPersonalSignChainId{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HSMSignRequestPersonalSignChainId) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int64
	err = json.Unmarshal(data, &dst.Int64)
	if err == nil {
		jsonInt64, _ := json.Marshal(dst.Int64)
		if string(jsonInt64) == "{}" { // empty struct
			dst.Int64 = nil
		} else {
			match++
		}
	} else {
		dst.Int64 = nil
	}

	// try to unmarshal data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int64 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(HSMSignRequestPersonalSignChainId)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(HSMSignRequestPersonalSignChainId)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HSMSignRequestPersonalSignChainId) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HSMSignRequestPersonalSignChainId) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Int64 != nil {
		return obj.Int64
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableHSMSignRequestPersonalSignChainId struct {
	value *HSMSignRequestPersonalSignChainId
	isSet bool
}

func (v NullableHSMSignRequestPersonalSignChainId) Get() *HSMSignRequestPersonalSignChainId {
	return v.value
}

func (v *NullableHSMSignRequestPersonalSignChainId) Set(val *HSMSignRequestPersonalSignChainId) {
	v.value = val
	v.isSet = true
}

func (v NullableHSMSignRequestPersonalSignChainId) IsSet() bool {
	return v.isSet
}

func (v *NullableHSMSignRequestPersonalSignChainId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHSMSignRequestPersonalSignChainId(val *HSMSignRequestPersonalSignChainId) *NullableHSMSignRequestPersonalSignChainId {
	return &NullableHSMSignRequestPersonalSignChainId{value: val, isSet: true}
}

func (v NullableHSMSignRequestPersonalSignChainId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHSMSignRequestPersonalSignChainId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
