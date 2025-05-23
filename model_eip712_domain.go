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

// checks if the EIP712Domain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EIP712Domain{}

// EIP712Domain The domain fields for EIP-712. All fields are optional per the specification.
type EIP712Domain struct {
	// Human-readable name of the signing domain.
	Name *string `json:"name,omitempty"`
	// Current major version of the signing domain.
	Version *string              `json:"version,omitempty"`
	ChainId *EIP712DomainChainId `json:"chainId,omitempty"`
	// An ethereum address.
	VerifyingContract *string `json:"verifyingContract,omitempty" validate:"regexp=^0[xX][a-fA-F0-9]{40}$"`
	// A hex string.
	Salt *string `json:"salt,omitempty" validate:"regexp=^(0x[0-9a-f]*|0X[0-9A-F]*)$"`
}

// NewEIP712Domain instantiates a new EIP712Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEIP712Domain() *EIP712Domain {
	this := EIP712Domain{}
	return &this
}

// NewEIP712DomainWithDefaults instantiates a new EIP712Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEIP712DomainWithDefaults() *EIP712Domain {
	this := EIP712Domain{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EIP712Domain) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EIP712Domain) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EIP712Domain) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EIP712Domain) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EIP712Domain) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EIP712Domain) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EIP712Domain) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EIP712Domain) SetVersion(v string) {
	o.Version = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *EIP712Domain) GetChainId() EIP712DomainChainId {
	if o == nil || IsNil(o.ChainId) {
		var ret EIP712DomainChainId
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EIP712Domain) GetChainIdOk() (*EIP712DomainChainId, bool) {
	if o == nil || IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *EIP712Domain) HasChainId() bool {
	if o != nil && !IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given EIP712DomainChainId and assigns it to the ChainId field.
func (o *EIP712Domain) SetChainId(v EIP712DomainChainId) {
	o.ChainId = &v
}

// GetVerifyingContract returns the VerifyingContract field value if set, zero value otherwise.
func (o *EIP712Domain) GetVerifyingContract() string {
	if o == nil || IsNil(o.VerifyingContract) {
		var ret string
		return ret
	}
	return *o.VerifyingContract
}

// GetVerifyingContractOk returns a tuple with the VerifyingContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EIP712Domain) GetVerifyingContractOk() (*string, bool) {
	if o == nil || IsNil(o.VerifyingContract) {
		return nil, false
	}
	return o.VerifyingContract, true
}

// HasVerifyingContract returns a boolean if a field has been set.
func (o *EIP712Domain) HasVerifyingContract() bool {
	if o != nil && !IsNil(o.VerifyingContract) {
		return true
	}

	return false
}

// SetVerifyingContract gets a reference to the given string and assigns it to the VerifyingContract field.
func (o *EIP712Domain) SetVerifyingContract(v string) {
	o.VerifyingContract = &v
}

// GetSalt returns the Salt field value if set, zero value otherwise.
func (o *EIP712Domain) GetSalt() string {
	if o == nil || IsNil(o.Salt) {
		var ret string
		return ret
	}
	return *o.Salt
}

// GetSaltOk returns a tuple with the Salt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EIP712Domain) GetSaltOk() (*string, bool) {
	if o == nil || IsNil(o.Salt) {
		return nil, false
	}
	return o.Salt, true
}

// HasSalt returns a boolean if a field has been set.
func (o *EIP712Domain) HasSalt() bool {
	if o != nil && !IsNil(o.Salt) {
		return true
	}

	return false
}

// SetSalt gets a reference to the given string and assigns it to the Salt field.
func (o *EIP712Domain) SetSalt(v string) {
	o.Salt = &v
}

func (o EIP712Domain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.ChainId) {
		toSerialize["chainId"] = o.ChainId
	}
	if !IsNil(o.VerifyingContract) {
		toSerialize["verifyingContract"] = o.VerifyingContract
	}
	if !IsNil(o.Salt) {
		toSerialize["salt"] = o.Salt
	}
	return toSerialize, nil
}

type NullableEIP712Domain struct {
	value *EIP712Domain
	isSet bool
}

func (v NullableEIP712Domain) Get() *EIP712Domain {
	return v.value
}

func (v *NullableEIP712Domain) Set(val *EIP712Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableEIP712Domain) IsSet() bool {
	return v.isSet
}

func (v *NullableEIP712Domain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEIP712Domain(val *EIP712Domain) *NullableEIP712Domain {
	return &NullableEIP712Domain{value: val, isSet: true}
}

func (v NullableEIP712Domain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEIP712Domain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
