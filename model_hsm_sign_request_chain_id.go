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

// HSMSignRequestChainId - struct for HSMSignRequestChainId
type HSMSignRequestChainId struct {
	Int64  *int64
	String *string
}

// int64AsHSMSignRequestChainId is a convenience function that returns int64 wrapped in HSMSignRequestChainId
func Int64AsHSMSignRequestChainId(v *int64) HSMSignRequestChainId {
	return HSMSignRequestChainId{
		Int64: v,
	}
}

// stringAsHSMSignRequestChainId is a convenience function that returns string wrapped in HSMSignRequestChainId
func StringAsHSMSignRequestChainId(v *string) HSMSignRequestChainId {
	return HSMSignRequestChainId{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HSMSignRequestChainId) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(HSMSignRequestChainId)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(HSMSignRequestChainId)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HSMSignRequestChainId) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HSMSignRequestChainId) GetActualInstance() interface{} {
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

type NullableHSMSignRequestChainId struct {
	value *HSMSignRequestChainId
	isSet bool
}

func (v NullableHSMSignRequestChainId) Get() *HSMSignRequestChainId {
	return v.value
}

func (v *NullableHSMSignRequestChainId) Set(val *HSMSignRequestChainId) {
	v.value = val
	v.isSet = true
}

func (v NullableHSMSignRequestChainId) IsSet() bool {
	return v.isSet
}

func (v *NullableHSMSignRequestChainId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHSMSignRequestChainId(val *HSMSignRequestChainId) *NullableHSMSignRequestChainId {
	return &NullableHSMSignRequestChainId{value: val, isSet: true}
}

func (v NullableHSMSignRequestChainId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHSMSignRequestChainId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
