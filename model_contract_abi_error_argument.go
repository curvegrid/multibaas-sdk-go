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

// checks if the ContractABIErrorArgument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractABIErrorArgument{}

// ContractABIErrorArgument A contract error argument.
type ContractABIErrorArgument struct {
	Name     string          `json:"name"`
	Type     ContractABIType `json:"type"`
	TypeName string          `json:"typeName"`
	Indexed  bool            `json:"indexed"`
	// The developer documentation.
	Notes string `json:"notes"`
}

type _ContractABIErrorArgument ContractABIErrorArgument

// NewContractABIErrorArgument instantiates a new ContractABIErrorArgument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractABIErrorArgument(name string, type_ ContractABIType, typeName string, indexed bool, notes string) *ContractABIErrorArgument {
	this := ContractABIErrorArgument{}
	this.Name = name
	this.Type = type_
	this.TypeName = typeName
	this.Indexed = indexed
	this.Notes = notes
	return &this
}

// NewContractABIErrorArgumentWithDefaults instantiates a new ContractABIErrorArgument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractABIErrorArgumentWithDefaults() *ContractABIErrorArgument {
	this := ContractABIErrorArgument{}
	return &this
}

// GetName returns the Name field value
func (o *ContractABIErrorArgument) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContractABIErrorArgument) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContractABIErrorArgument) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ContractABIErrorArgument) GetType() ContractABIType {
	if o == nil {
		var ret ContractABIType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContractABIErrorArgument) GetTypeOk() (*ContractABIType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContractABIErrorArgument) SetType(v ContractABIType) {
	o.Type = v
}

// GetTypeName returns the TypeName field value
func (o *ContractABIErrorArgument) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *ContractABIErrorArgument) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *ContractABIErrorArgument) SetTypeName(v string) {
	o.TypeName = v
}

// GetIndexed returns the Indexed field value
func (o *ContractABIErrorArgument) GetIndexed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Indexed
}

// GetIndexedOk returns a tuple with the Indexed field value
// and a boolean to check if the value has been set.
func (o *ContractABIErrorArgument) GetIndexedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indexed, true
}

// SetIndexed sets field value
func (o *ContractABIErrorArgument) SetIndexed(v bool) {
	o.Indexed = v
}

// GetNotes returns the Notes field value
func (o *ContractABIErrorArgument) GetNotes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *ContractABIErrorArgument) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notes, true
}

// SetNotes sets field value
func (o *ContractABIErrorArgument) SetNotes(v string) {
	o.Notes = v
}

func (o ContractABIErrorArgument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["typeName"] = o.TypeName
	toSerialize["indexed"] = o.Indexed
	toSerialize["notes"] = o.Notes
	return toSerialize, nil
}

type NullableContractABIErrorArgument struct {
	value *ContractABIErrorArgument
	isSet bool
}

func (v NullableContractABIErrorArgument) Get() *ContractABIErrorArgument {
	return v.value
}

func (v *NullableContractABIErrorArgument) Set(val *ContractABIErrorArgument) {
	v.value = val
	v.isSet = true
}

func (v NullableContractABIErrorArgument) IsSet() bool {
	return v.isSet
}

func (v *NullableContractABIErrorArgument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractABIErrorArgument(val *ContractABIErrorArgument) *NullableContractABIErrorArgument {
	return &NullableContractABIErrorArgument{value: val, isSet: true}
}

func (v NullableContractABIErrorArgument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractABIErrorArgument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
