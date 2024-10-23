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

// checks if the ContractABIMethod1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractABIMethod1{}

// ContractABIMethod1 A contract function.
type ContractABIMethod1 struct {
	// A hex string.
	Id string `json:"id" validate:"regexp=^(0x[0-9a-f]*|0X[0-9A-F]*)$"`
	// Name of the function.
	Name string `json:"name"`
	// The function signature.
	Signature string `json:"signature"`
	Const     bool   `json:"const"`
	Payable   bool   `json:"payable"`
	// List of function arguments.
	Inputs []ContractABIMethodArgument `json:"inputs"`
	// List of function outputs.
	Outputs []ContractABIMethodArgument `json:"outputs"`
	Author  string                      `json:"author"`
	Notes   string                      `json:"notes"`
	Returns string                      `json:"returns"`
	// The function description.
	Description string `json:"description"`
}

type _ContractABIMethod1 ContractABIMethod1

// NewContractABIMethod1 instantiates a new ContractABIMethod1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractABIMethod1(id string, name string, signature string, const_ bool, payable bool, inputs []ContractABIMethodArgument, outputs []ContractABIMethodArgument, author string, notes string, returns string, description string) *ContractABIMethod1 {
	this := ContractABIMethod1{}
	this.Id = id
	this.Name = name
	this.Signature = signature
	this.Const = const_
	this.Payable = payable
	this.Inputs = inputs
	this.Outputs = outputs
	this.Author = author
	this.Notes = notes
	this.Returns = returns
	this.Description = description
	return &this
}

// NewContractABIMethod1WithDefaults instantiates a new ContractABIMethod1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractABIMethod1WithDefaults() *ContractABIMethod1 {
	this := ContractABIMethod1{}
	return &this
}

// GetId returns the Id field value
func (o *ContractABIMethod1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContractABIMethod1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContractABIMethod1) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ContractABIMethod1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContractABIMethod1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContractABIMethod1) SetName(v string) {
	o.Name = v
}

// GetSignature returns the Signature field value
func (o *ContractABIMethod1) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *ContractABIMethod1) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *ContractABIMethod1) SetSignature(v string) {
	o.Signature = v
}

// GetConst returns the Const field value
func (o *ContractABIMethod1) GetConst() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Const
}

// GetConstOk returns a tuple with the Const field value
// and a boolean to check if the value has been set.
func (o *ContractABIMethod1) GetConstOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Const, true
}

// SetConst sets field value
func (o *ContractABIMethod1) SetConst(v bool) {
	o.Const = v
}

// GetPayable returns the Payable field value
func (o *ContractABIMethod1) GetPayable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Payable
}

// GetPayableOk returns a tuple with the Payable field value
// and a boolean to check if the value has been set.
func (o *ContractABIMethod1) GetPayableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payable, true
}

// SetPayable sets field value
func (o *ContractABIMethod1) SetPayable(v bool) {
	o.Payable = v
}

// GetInputs returns the Inputs field value
func (o *ContractABIMethod1) GetInputs() []ContractABIMethodArgument {
	if o == nil {
		var ret []ContractABIMethodArgument
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ContractABIMethod1) GetInputsOk() ([]ContractABIMethodArgument, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *ContractABIMethod1) SetInputs(v []ContractABIMethodArgument) {
	o.Inputs = v
}

// GetOutputs returns the Outputs field value
func (o *ContractABIMethod1) GetOutputs() []ContractABIMethodArgument {
	if o == nil {
		var ret []ContractABIMethodArgument
		return ret
	}

	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value
// and a boolean to check if the value has been set.
func (o *ContractABIMethod1) GetOutputsOk() ([]ContractABIMethodArgument, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outputs, true
}

// SetOutputs sets field value
func (o *ContractABIMethod1) SetOutputs(v []ContractABIMethodArgument) {
	o.Outputs = v
}

// GetAuthor returns the Author field value
func (o *ContractABIMethod1) GetAuthor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *ContractABIMethod1) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *ContractABIMethod1) SetAuthor(v string) {
	o.Author = v
}

// GetNotes returns the Notes field value
func (o *ContractABIMethod1) GetNotes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *ContractABIMethod1) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notes, true
}

// SetNotes sets field value
func (o *ContractABIMethod1) SetNotes(v string) {
	o.Notes = v
}

// GetReturns returns the Returns field value
func (o *ContractABIMethod1) GetReturns() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Returns
}

// GetReturnsOk returns a tuple with the Returns field value
// and a boolean to check if the value has been set.
func (o *ContractABIMethod1) GetReturnsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Returns, true
}

// SetReturns sets field value
func (o *ContractABIMethod1) SetReturns(v string) {
	o.Returns = v
}

// GetDescription returns the Description field value
func (o *ContractABIMethod1) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ContractABIMethod1) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ContractABIMethod1) SetDescription(v string) {
	o.Description = v
}

func (o ContractABIMethod1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractABIMethod1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["signature"] = o.Signature
	toSerialize["const"] = o.Const
	toSerialize["payable"] = o.Payable
	toSerialize["inputs"] = o.Inputs
	toSerialize["outputs"] = o.Outputs
	toSerialize["author"] = o.Author
	toSerialize["notes"] = o.Notes
	toSerialize["returns"] = o.Returns
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *ContractABIMethod1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"signature",
		"const",
		"payable",
		"inputs",
		"outputs",
		"author",
		"notes",
		"returns",
		"description",
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

	varContractABIMethod1 := _ContractABIMethod1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContractABIMethod1)

	if err != nil {
		return err
	}

	*o = ContractABIMethod1(varContractABIMethod1)

	return err
}

type NullableContractABIMethod1 struct {
	value *ContractABIMethod1
	isSet bool
}

func (v NullableContractABIMethod1) Get() *ContractABIMethod1 {
	return v.value
}

func (v *NullableContractABIMethod1) Set(val *ContractABIMethod1) {
	v.value = val
	v.isSet = true
}

func (v NullableContractABIMethod1) IsSet() bool {
	return v.isSet
}

func (v *NullableContractABIMethod1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractABIMethod1(val *ContractABIMethod1) *NullableContractABIMethod1 {
	return &NullableContractABIMethod1{value: val, isSet: true}
}

func (v NullableContractABIMethod1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractABIMethod1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
