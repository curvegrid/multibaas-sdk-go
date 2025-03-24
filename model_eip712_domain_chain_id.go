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

// EIP712DomainChainId - The EIP-155 chain ID of the application using the typed data.
type EIP712DomainChainId struct {
	Int64  *int64
	String *string
}

// int64AsEIP712DomainChainId is a convenience function that returns int64 wrapped in EIP712DomainChainId
func Int64AsEIP712DomainChainId(v *int64) EIP712DomainChainId {
	return EIP712DomainChainId{
		Int64: v,
	}
}

// stringAsEIP712DomainChainId is a convenience function that returns string wrapped in EIP712DomainChainId
func StringAsEIP712DomainChainId(v *string) EIP712DomainChainId {
	return EIP712DomainChainId{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EIP712DomainChainId) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(EIP712DomainChainId)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EIP712DomainChainId)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EIP712DomainChainId) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EIP712DomainChainId) GetActualInstance() interface{} {
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

// Get the actual instance value
func (obj EIP712DomainChainId) GetActualInstanceValue() interface{} {
	if obj.Int64 != nil {
		return *obj.Int64
	}

	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableEIP712DomainChainId struct {
	value *EIP712DomainChainId
	isSet bool
}

func (v NullableEIP712DomainChainId) Get() *EIP712DomainChainId {
	return v.value
}

func (v *NullableEIP712DomainChainId) Set(val *EIP712DomainChainId) {
	v.value = val
	v.isSet = true
}

func (v NullableEIP712DomainChainId) IsSet() bool {
	return v.isSet
}

func (v *NullableEIP712DomainChainId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEIP712DomainChainId(val *EIP712DomainChainId) *NullableEIP712DomainChainId {
	return &NullableEIP712DomainChainId{value: val, isSet: true}
}

func (v NullableEIP712DomainChainId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEIP712DomainChainId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
