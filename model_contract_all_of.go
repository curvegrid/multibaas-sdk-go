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

// checks if the ContractAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractAllOf{}

// ContractAllOf struct for ContractAllOf
type ContractAllOf struct {
	Abi ContractABI `json:"abi"`
	// List of the contract instances.
	Instances []ContractInstance `json:"instances,omitempty"`
}

// NewContractAllOf instantiates a new ContractAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractAllOf(abi ContractABI) *ContractAllOf {
	this := ContractAllOf{}
	this.Abi = abi
	return &this
}

// NewContractAllOfWithDefaults instantiates a new ContractAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractAllOfWithDefaults() *ContractAllOf {
	this := ContractAllOf{}
	return &this
}

// GetAbi returns the Abi field value
func (o *ContractAllOf) GetAbi() ContractABI {
	if o == nil {
		var ret ContractABI
		return ret
	}

	return o.Abi
}

// GetAbiOk returns a tuple with the Abi field value
// and a boolean to check if the value has been set.
func (o *ContractAllOf) GetAbiOk() (*ContractABI, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Abi, true
}

// SetAbi sets field value
func (o *ContractAllOf) SetAbi(v ContractABI) {
	o.Abi = v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *ContractAllOf) GetInstances() []ContractInstance {
	if o == nil || IsNil(o.Instances) {
		var ret []ContractInstance
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractAllOf) GetInstancesOk() ([]ContractInstance, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *ContractAllOf) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []ContractInstance and assigns it to the Instances field.
func (o *ContractAllOf) SetInstances(v []ContractInstance) {
	o.Instances = v
}

func (o ContractAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["abi"] = o.Abi
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	return toSerialize, nil
}

type NullableContractAllOf struct {
	value *ContractAllOf
	isSet bool
}

func (v NullableContractAllOf) Get() *ContractAllOf {
	return v.value
}

func (v *NullableContractAllOf) Set(val *ContractAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContractAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContractAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractAllOf(val *ContractAllOf) *NullableContractAllOf {
	return &NullableContractAllOf{value: val, isSet: true}
}

func (v NullableContractAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
