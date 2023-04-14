# AddressAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to **string** |  | [optional] 
**Chain** | **string** |  | 
**Modules** | **[]string** |  | 
**Nonce** | Pointer to **int64** | The next transaction nonce for this address (obtained from the blockchain node). | [optional] 
**LocalNonce** | Pointer to **int64** | The next transaction nonce for this address when using the nonce management feature. | [optional] 
**CodeAt** | Pointer to **string** |  | [optional] 
**Contracts** | [**[]ContractMetadata**](ContractMetadata.md) |  | 

## Methods

### NewAddressAllOf

`func NewAddressAllOf(chain string, modules []string, contracts []ContractMetadata, ) *AddressAllOf`

NewAddressAllOf instantiates a new AddressAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressAllOfWithDefaults

`func NewAddressAllOfWithDefaults() *AddressAllOf`

NewAddressAllOfWithDefaults instantiates a new AddressAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *AddressAllOf) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AddressAllOf) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AddressAllOf) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AddressAllOf) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetChain

`func (o *AddressAllOf) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *AddressAllOf) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *AddressAllOf) SetChain(v string)`

SetChain sets Chain field to given value.


### GetModules

`func (o *AddressAllOf) GetModules() []string`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *AddressAllOf) GetModulesOk() (*[]string, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *AddressAllOf) SetModules(v []string)`

SetModules sets Modules field to given value.


### GetNonce

`func (o *AddressAllOf) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *AddressAllOf) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *AddressAllOf) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *AddressAllOf) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetLocalNonce

`func (o *AddressAllOf) GetLocalNonce() int64`

GetLocalNonce returns the LocalNonce field if non-nil, zero value otherwise.

### GetLocalNonceOk

`func (o *AddressAllOf) GetLocalNonceOk() (*int64, bool)`

GetLocalNonceOk returns a tuple with the LocalNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNonce

`func (o *AddressAllOf) SetLocalNonce(v int64)`

SetLocalNonce sets LocalNonce field to given value.

### HasLocalNonce

`func (o *AddressAllOf) HasLocalNonce() bool`

HasLocalNonce returns a boolean if a field has been set.

### GetCodeAt

`func (o *AddressAllOf) GetCodeAt() string`

GetCodeAt returns the CodeAt field if non-nil, zero value otherwise.

### GetCodeAtOk

`func (o *AddressAllOf) GetCodeAtOk() (*string, bool)`

GetCodeAtOk returns a tuple with the CodeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeAt

`func (o *AddressAllOf) SetCodeAt(v string)`

SetCodeAt sets CodeAt field to given value.

### HasCodeAt

`func (o *AddressAllOf) HasCodeAt() bool`

HasCodeAt returns a boolean if a field has been set.

### GetContracts

`func (o *AddressAllOf) GetContracts() []ContractMetadata`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *AddressAllOf) GetContractsOk() (*[]ContractMetadata, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *AddressAllOf) SetContracts(v []ContractMetadata)`

SetContracts sets Contracts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


