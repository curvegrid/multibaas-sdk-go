# ContractOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | An alias to easily identify and reference the entity in subsequent requests. | 
**ContractName** | **string** | The name of the contract. | 
**Version** | **string** | The contract version. | 
**IsFavorite** | Pointer to **bool** |  | [optional] 
**Deployable** | **bool** |  | 
**Instances** | [**[]ContractInstance**](ContractInstance.md) | List of contract instances. | 

## Methods

### NewContractOverview

`func NewContractOverview(label string, contractName string, version string, deployable bool, instances []ContractInstance, ) *ContractOverview`

NewContractOverview instantiates a new ContractOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractOverviewWithDefaults

`func NewContractOverviewWithDefaults() *ContractOverview`

NewContractOverviewWithDefaults instantiates a new ContractOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ContractOverview) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContractOverview) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContractOverview) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetContractName

`func (o *ContractOverview) GetContractName() string`

GetContractName returns the ContractName field if non-nil, zero value otherwise.

### GetContractNameOk

`func (o *ContractOverview) GetContractNameOk() (*string, bool)`

GetContractNameOk returns a tuple with the ContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractName

`func (o *ContractOverview) SetContractName(v string)`

SetContractName sets ContractName field to given value.


### GetVersion

`func (o *ContractOverview) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContractOverview) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContractOverview) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetIsFavorite

`func (o *ContractOverview) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *ContractOverview) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *ContractOverview) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *ContractOverview) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetDeployable

`func (o *ContractOverview) GetDeployable() bool`

GetDeployable returns the Deployable field if non-nil, zero value otherwise.

### GetDeployableOk

`func (o *ContractOverview) GetDeployableOk() (*bool, bool)`

GetDeployableOk returns a tuple with the Deployable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployable

`func (o *ContractOverview) SetDeployable(v bool)`

SetDeployable sets Deployable field to given value.


### GetInstances

`func (o *ContractOverview) GetInstances() []ContractInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ContractOverview) GetInstancesOk() (*[]ContractInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ContractOverview) SetInstances(v []ContractInstance)`

SetInstances sets Instances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


