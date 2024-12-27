# EIP712Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human-readable name of the signing domain. | [optional] 
**Version** | Pointer to **string** | Current major version of the signing domain. | [optional] 
**ChainId** | Pointer to [**EIP712DomainChainId**](EIP712DomainChainId.md) |  | [optional] 
**VerifyingContract** | Pointer to **string** | An ethereum address. | [optional] 
**Salt** | Pointer to **string** | A hex string. | [optional] 

## Methods

### NewEIP712Domain

`func NewEIP712Domain() *EIP712Domain`

NewEIP712Domain instantiates a new EIP712Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEIP712DomainWithDefaults

`func NewEIP712DomainWithDefaults() *EIP712Domain`

NewEIP712DomainWithDefaults instantiates a new EIP712Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EIP712Domain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EIP712Domain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EIP712Domain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EIP712Domain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *EIP712Domain) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EIP712Domain) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EIP712Domain) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EIP712Domain) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetChainId

`func (o *EIP712Domain) GetChainId() EIP712DomainChainId`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *EIP712Domain) GetChainIdOk() (*EIP712DomainChainId, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *EIP712Domain) SetChainId(v EIP712DomainChainId)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *EIP712Domain) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetVerifyingContract

`func (o *EIP712Domain) GetVerifyingContract() string`

GetVerifyingContract returns the VerifyingContract field if non-nil, zero value otherwise.

### GetVerifyingContractOk

`func (o *EIP712Domain) GetVerifyingContractOk() (*string, bool)`

GetVerifyingContractOk returns a tuple with the VerifyingContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyingContract

`func (o *EIP712Domain) SetVerifyingContract(v string)`

SetVerifyingContract sets VerifyingContract field to given value.

### HasVerifyingContract

`func (o *EIP712Domain) HasVerifyingContract() bool`

HasVerifyingContract returns a boolean if a field has been set.

### GetSalt

`func (o *EIP712Domain) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *EIP712Domain) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *EIP712Domain) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *EIP712Domain) HasSalt() bool`

HasSalt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


