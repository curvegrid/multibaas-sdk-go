# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | A label. | 
**Address** | **string** | An ethereum address. | 
**Balance** | Pointer to **string** |  | [optional] 
**Chain** | **string** |  | 
**Modules** | **[]string** |  | 
**Nonce** | Pointer to **int64** | The next transaction nonce for this address (obtained from the blockchain node). | [optional] 
**LocalNonce** | Pointer to **int64** | The next transaction nonce for this address when using the nonce management feature. | [optional] 
**CodeAt** | Pointer to **string** |  | [optional] 
**Contracts** | [**[]ContractMetadata**](ContractMetadata.md) |  | 

## Methods

### NewAddress

`func NewAddress(label string, address string, chain string, modules []string, contracts []ContractMetadata, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *Address) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Address) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Address) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetAddress

`func (o *Address) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Address) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Address) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBalance

`func (o *Address) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Address) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Address) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *Address) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetChain

`func (o *Address) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *Address) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *Address) SetChain(v string)`

SetChain sets Chain field to given value.


### GetModules

`func (o *Address) GetModules() []string`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *Address) GetModulesOk() (*[]string, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *Address) SetModules(v []string)`

SetModules sets Modules field to given value.


### GetNonce

`func (o *Address) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *Address) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *Address) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *Address) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetLocalNonce

`func (o *Address) GetLocalNonce() int64`

GetLocalNonce returns the LocalNonce field if non-nil, zero value otherwise.

### GetLocalNonceOk

`func (o *Address) GetLocalNonceOk() (*int64, bool)`

GetLocalNonceOk returns a tuple with the LocalNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNonce

`func (o *Address) SetLocalNonce(v int64)`

SetLocalNonce sets LocalNonce field to given value.

### HasLocalNonce

`func (o *Address) HasLocalNonce() bool`

HasLocalNonce returns a boolean if a field has been set.

### GetCodeAt

`func (o *Address) GetCodeAt() string`

GetCodeAt returns the CodeAt field if non-nil, zero value otherwise.

### GetCodeAtOk

`func (o *Address) GetCodeAtOk() (*string, bool)`

GetCodeAtOk returns a tuple with the CodeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeAt

`func (o *Address) SetCodeAt(v string)`

SetCodeAt sets CodeAt field to given value.

### HasCodeAt

`func (o *Address) HasCodeAt() bool`

HasCodeAt returns a boolean if a field has been set.

### GetContracts

`func (o *Address) GetContracts() []ContractMetadata`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *Address) GetContractsOk() (*[]ContractMetadata, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *Address) SetContracts(v []ContractMetadata)`

SetContracts sets Contracts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


