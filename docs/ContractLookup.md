# ContractLookup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | An ethereum address. | 
**Name** | Pointer to **string** | The name of the contract. | [optional] 
**Abi** | **string** | The contract ABI JSON string. | 
**Bytecode** | Pointer to **string** | The smart-contract bytecode. | [optional] 
**Source** | Pointer to **string** | The contract&#39;s source code. | [optional] 
**Userdoc** | Pointer to **string** | The user documentation JSON string. | [optional] 
**Devdoc** | Pointer to **string** | The developer documentation JSON string. | [optional] 
**Verified** | **bool** | Indicates whether the contract has been verified. | 
**VerifiedSource** | Pointer to **string** | The name of the service that provided the contract verification. | [optional] 
**VerifiedLink** | Pointer to **string** | The URL to the contract&#39;s verification details on the verification service. | [optional] 
**Proxy** | **bool** | Indicates whether the contract is a proxy contract. | 

## Methods

### NewContractLookup

`func NewContractLookup(address string, abi string, verified bool, proxy bool, ) *ContractLookup`

NewContractLookup instantiates a new ContractLookup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractLookupWithDefaults

`func NewContractLookupWithDefaults() *ContractLookup`

NewContractLookupWithDefaults instantiates a new ContractLookup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ContractLookup) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractLookup) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractLookup) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetName

`func (o *ContractLookup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractLookup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractLookup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContractLookup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAbi

`func (o *ContractLookup) GetAbi() string`

GetAbi returns the Abi field if non-nil, zero value otherwise.

### GetAbiOk

`func (o *ContractLookup) GetAbiOk() (*string, bool)`

GetAbiOk returns a tuple with the Abi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbi

`func (o *ContractLookup) SetAbi(v string)`

SetAbi sets Abi field to given value.


### GetBytecode

`func (o *ContractLookup) GetBytecode() string`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *ContractLookup) GetBytecodeOk() (*string, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *ContractLookup) SetBytecode(v string)`

SetBytecode sets Bytecode field to given value.

### HasBytecode

`func (o *ContractLookup) HasBytecode() bool`

HasBytecode returns a boolean if a field has been set.

### GetSource

`func (o *ContractLookup) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ContractLookup) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ContractLookup) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ContractLookup) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUserdoc

`func (o *ContractLookup) GetUserdoc() string`

GetUserdoc returns the Userdoc field if non-nil, zero value otherwise.

### GetUserdocOk

`func (o *ContractLookup) GetUserdocOk() (*string, bool)`

GetUserdocOk returns a tuple with the Userdoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdoc

`func (o *ContractLookup) SetUserdoc(v string)`

SetUserdoc sets Userdoc field to given value.

### HasUserdoc

`func (o *ContractLookup) HasUserdoc() bool`

HasUserdoc returns a boolean if a field has been set.

### GetDevdoc

`func (o *ContractLookup) GetDevdoc() string`

GetDevdoc returns the Devdoc field if non-nil, zero value otherwise.

### GetDevdocOk

`func (o *ContractLookup) GetDevdocOk() (*string, bool)`

GetDevdocOk returns a tuple with the Devdoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevdoc

`func (o *ContractLookup) SetDevdoc(v string)`

SetDevdoc sets Devdoc field to given value.

### HasDevdoc

`func (o *ContractLookup) HasDevdoc() bool`

HasDevdoc returns a boolean if a field has been set.

### GetVerified

`func (o *ContractLookup) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ContractLookup) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ContractLookup) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetVerifiedSource

`func (o *ContractLookup) GetVerifiedSource() string`

GetVerifiedSource returns the VerifiedSource field if non-nil, zero value otherwise.

### GetVerifiedSourceOk

`func (o *ContractLookup) GetVerifiedSourceOk() (*string, bool)`

GetVerifiedSourceOk returns a tuple with the VerifiedSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedSource

`func (o *ContractLookup) SetVerifiedSource(v string)`

SetVerifiedSource sets VerifiedSource field to given value.

### HasVerifiedSource

`func (o *ContractLookup) HasVerifiedSource() bool`

HasVerifiedSource returns a boolean if a field has been set.

### GetVerifiedLink

`func (o *ContractLookup) GetVerifiedLink() string`

GetVerifiedLink returns the VerifiedLink field if non-nil, zero value otherwise.

### GetVerifiedLinkOk

`func (o *ContractLookup) GetVerifiedLinkOk() (*string, bool)`

GetVerifiedLinkOk returns a tuple with the VerifiedLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedLink

`func (o *ContractLookup) SetVerifiedLink(v string)`

SetVerifiedLink sets VerifiedLink field to given value.

### HasVerifiedLink

`func (o *ContractLookup) HasVerifiedLink() bool`

HasVerifiedLink returns a boolean if a field has been set.

### GetProxy

`func (o *ContractLookup) GetProxy() bool`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ContractLookup) GetProxyOk() (*bool, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ContractLookup) SetProxy(v bool)`

SetProxy sets Proxy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


