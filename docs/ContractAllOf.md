# ContractAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abi** | [**ContractABI**](ContractABI.md) |  | 
**Instances** | Pointer to [**[]ContractInstance**](ContractInstance.md) | List of the contract instances. | [optional] 

## Methods

### NewContractAllOf

`func NewContractAllOf(abi ContractABI, ) *ContractAllOf`

NewContractAllOf instantiates a new ContractAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractAllOfWithDefaults

`func NewContractAllOfWithDefaults() *ContractAllOf`

NewContractAllOfWithDefaults instantiates a new ContractAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbi

`func (o *ContractAllOf) GetAbi() ContractABI`

GetAbi returns the Abi field if non-nil, zero value otherwise.

### GetAbiOk

`func (o *ContractAllOf) GetAbiOk() (*ContractABI, bool)`

GetAbiOk returns a tuple with the Abi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbi

`func (o *ContractAllOf) SetAbi(v ContractABI)`

SetAbi sets Abi field to given value.


### GetInstances

`func (o *ContractAllOf) GetInstances() []ContractInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ContractAllOf) GetInstancesOk() (*[]ContractInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ContractAllOf) SetInstances(v []ContractInstance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ContractAllOf) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


