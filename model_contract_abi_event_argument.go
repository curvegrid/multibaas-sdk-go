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

// checks if the ContractABIEventArgument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractABIEventArgument{}

// ContractABIEventArgument A contract event argument.
type ContractABIEventArgument struct {
	Name           string                            `json:"name"`
	Type           ContractABIType                   `json:"type"`
	Indexed        bool                              `json:"indexed"`
	TypeConversion NullableContractABITypeConversion `json:"typeConversion"`
}

// NewContractABIEventArgument instantiates a new ContractABIEventArgument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractABIEventArgument(name string, type_ ContractABIType, indexed bool, typeConversion NullableContractABITypeConversion) *ContractABIEventArgument {
	this := ContractABIEventArgument{}
	this.Name = name
	this.Type = type_
	this.Indexed = indexed
	this.TypeConversion = typeConversion
	return &this
}

// NewContractABIEventArgumentWithDefaults instantiates a new ContractABIEventArgument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractABIEventArgumentWithDefaults() *ContractABIEventArgument {
	this := ContractABIEventArgument{}
	return &this
}

// GetName returns the Name field value
func (o *ContractABIEventArgument) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContractABIEventArgument) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContractABIEventArgument) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ContractABIEventArgument) GetType() ContractABIType {
	if o == nil {
		var ret ContractABIType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContractABIEventArgument) GetTypeOk() (*ContractABIType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContractABIEventArgument) SetType(v ContractABIType) {
	o.Type = v
}

// GetIndexed returns the Indexed field value
func (o *ContractABIEventArgument) GetIndexed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Indexed
}

// GetIndexedOk returns a tuple with the Indexed field value
// and a boolean to check if the value has been set.
func (o *ContractABIEventArgument) GetIndexedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indexed, true
}

// SetIndexed sets field value
func (o *ContractABIEventArgument) SetIndexed(v bool) {
	o.Indexed = v
}

// GetTypeConversion returns the TypeConversion field value
// If the value is explicit nil, the zero value for ContractABITypeConversion will be returned
func (o *ContractABIEventArgument) GetTypeConversion() ContractABITypeConversion {
	if o == nil || o.TypeConversion.Get() == nil {
		var ret ContractABITypeConversion
		return ret
	}

	return *o.TypeConversion.Get()
}

// GetTypeConversionOk returns a tuple with the TypeConversion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContractABIEventArgument) GetTypeConversionOk() (*ContractABITypeConversion, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeConversion.Get(), o.TypeConversion.IsSet()
}

// SetTypeConversion sets field value
func (o *ContractABIEventArgument) SetTypeConversion(v ContractABITypeConversion) {
	o.TypeConversion.Set(&v)
}

func (o ContractABIEventArgument) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractABIEventArgument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["indexed"] = o.Indexed
	toSerialize["typeConversion"] = o.TypeConversion.Get()
	return toSerialize, nil
}

type NullableContractABIEventArgument struct {
	value *ContractABIEventArgument
	isSet bool
}

func (v NullableContractABIEventArgument) Get() *ContractABIEventArgument {
	return v.value
}

func (v *NullableContractABIEventArgument) Set(val *ContractABIEventArgument) {
	v.value = val
	v.isSet = true
}

func (v NullableContractABIEventArgument) IsSet() bool {
	return v.isSet
}

func (v *NullableContractABIEventArgument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractABIEventArgument(val *ContractABIEventArgument) *NullableContractABIEventArgument {
	return &NullableContractABIEventArgument{value: val, isSet: true}
}

func (v NullableContractABIEventArgument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractABIEventArgument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}