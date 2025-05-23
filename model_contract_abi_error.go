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

// checks if the ContractABIError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractABIError{}

// ContractABIError A contract error.
type ContractABIError struct {
	// The keccak256 hash as a hex string of 256 bits.
	Id        string `json:"id" validate:"regexp=^(0x[0-9a-f]{64}|0X[0-9A-F]{64})$"`
	Name      string `json:"name"`
	Signature string `json:"signature"`
	// List of contract event's input arguments.
	Inputs []ContractABIErrorArgument `json:"inputs"`
	// The developer documentation.
	Notes *string `json:"notes,omitempty"`
	// The user documentation.
	Description *string `json:"description,omitempty"`
}

type _ContractABIError ContractABIError

// NewContractABIError instantiates a new ContractABIError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractABIError(id string, name string, signature string, inputs []ContractABIErrorArgument) *ContractABIError {
	this := ContractABIError{}
	this.Id = id
	this.Name = name
	this.Signature = signature
	this.Inputs = inputs
	return &this
}

// NewContractABIErrorWithDefaults instantiates a new ContractABIError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractABIErrorWithDefaults() *ContractABIError {
	this := ContractABIError{}
	return &this
}

// GetId returns the Id field value
func (o *ContractABIError) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContractABIError) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContractABIError) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ContractABIError) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContractABIError) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContractABIError) SetName(v string) {
	o.Name = v
}

// GetSignature returns the Signature field value
func (o *ContractABIError) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *ContractABIError) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *ContractABIError) SetSignature(v string) {
	o.Signature = v
}

// GetInputs returns the Inputs field value
func (o *ContractABIError) GetInputs() []ContractABIErrorArgument {
	if o == nil {
		var ret []ContractABIErrorArgument
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ContractABIError) GetInputsOk() ([]ContractABIErrorArgument, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *ContractABIError) SetInputs(v []ContractABIErrorArgument) {
	o.Inputs = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ContractABIError) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractABIError) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ContractABIError) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ContractABIError) SetNotes(v string) {
	o.Notes = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ContractABIError) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractABIError) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ContractABIError) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ContractABIError) SetDescription(v string) {
	o.Description = &v
}

func (o ContractABIError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["signature"] = o.Signature
	toSerialize["inputs"] = o.Inputs
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableContractABIError struct {
	value *ContractABIError
	isSet bool
}

func (v NullableContractABIError) Get() *ContractABIError {
	return v.value
}

func (v *NullableContractABIError) Set(val *ContractABIError) {
	v.value = val
	v.isSet = true
}

func (v NullableContractABIError) IsSet() bool {
	return v.isSet
}

func (v *NullableContractABIError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractABIError(val *ContractABIError) *NullableContractABIError {
	return &NullableContractABIError{value: val, isSet: true}
}

func (v NullableContractABIError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractABIError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
