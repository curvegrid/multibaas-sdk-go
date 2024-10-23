# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | An alias to easily identify and reference the entity in subsequent requests. | 
**ContractName** | **string** | The name of the contract. | 
**Version** | **string** | The contract version. | 
**Bin** | Pointer to **string** | The smart-contract bytecode. | [optional] 
**RawAbi** | **string** | The contract raw ABI JSON string. | 
**UserDoc** | **string** | The user documentation JSON string. | 
**DeveloperDoc** | **string** | The developer documentation JSON string. | 
**Metadata** | Pointer to **string** | The contract metadata JSON string. | [optional] 
**IsFavorite** | Pointer to **bool** |  | [optional] 
**Abi** | [**ContractABI**](ContractABI.md) |  | 
**Instances** | Pointer to [**[]ContractInstance**](ContractInstance.md) | List of the contract instances. | [optional] 

## Methods

### NewContract

`func NewContract(label string, contractName string, version string, rawAbi string, userDoc string, developerDoc string, abi ContractABI, ) *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *Contract) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Contract) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Contract) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetContractName

`func (o *Contract) GetContractName() string`

GetContractName returns the ContractName field if non-nil, zero value otherwise.

### GetContractNameOk

`func (o *Contract) GetContractNameOk() (*string, bool)`

GetContractNameOk returns a tuple with the ContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractName

`func (o *Contract) SetContractName(v string)`

SetContractName sets ContractName field to given value.


### GetVersion

`func (o *Contract) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Contract) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Contract) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetBin

`func (o *Contract) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *Contract) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *Contract) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *Contract) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetRawAbi

`func (o *Contract) GetRawAbi() string`

GetRawAbi returns the RawAbi field if non-nil, zero value otherwise.

### GetRawAbiOk

`func (o *Contract) GetRawAbiOk() (*string, bool)`

GetRawAbiOk returns a tuple with the RawAbi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAbi

`func (o *Contract) SetRawAbi(v string)`

SetRawAbi sets RawAbi field to given value.


### GetUserDoc

`func (o *Contract) GetUserDoc() string`

GetUserDoc returns the UserDoc field if non-nil, zero value otherwise.

### GetUserDocOk

`func (o *Contract) GetUserDocOk() (*string, bool)`

GetUserDocOk returns a tuple with the UserDoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDoc

`func (o *Contract) SetUserDoc(v string)`

SetUserDoc sets UserDoc field to given value.


### GetDeveloperDoc

`func (o *Contract) GetDeveloperDoc() string`

GetDeveloperDoc returns the DeveloperDoc field if non-nil, zero value otherwise.

### GetDeveloperDocOk

`func (o *Contract) GetDeveloperDocOk() (*string, bool)`

GetDeveloperDocOk returns a tuple with the DeveloperDoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperDoc

`func (o *Contract) SetDeveloperDoc(v string)`

SetDeveloperDoc sets DeveloperDoc field to given value.


### GetMetadata

`func (o *Contract) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Contract) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Contract) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Contract) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetIsFavorite

`func (o *Contract) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *Contract) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *Contract) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *Contract) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetAbi

`func (o *Contract) GetAbi() ContractABI`

GetAbi returns the Abi field if non-nil, zero value otherwise.

### GetAbiOk

`func (o *Contract) GetAbiOk() (*ContractABI, bool)`

GetAbiOk returns a tuple with the Abi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbi

`func (o *Contract) SetAbi(v ContractABI)`

SetAbi sets Abi field to given value.


### GetInstances

`func (o *Contract) GetInstances() []ContractInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *Contract) GetInstancesOk() (*[]ContractInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *Contract) SetInstances(v []ContractInstance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *Contract) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


