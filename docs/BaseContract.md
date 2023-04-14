# BaseContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | A label. | 
**ContractName** | **string** | The name of the contract. | 
**Version** | **string** | The contract version. | 
**Bin** | Pointer to **string** | The smart-contract bytecode. | [optional] 
**RawAbi** | **string** | The contract raw ABI JSON string. | 
**UserDoc** | **string** | The user documentation JSON string. | 
**DeveloperDoc** | **string** | The developer documentation JSON string. | 
**Metadata** | Pointer to **string** | The contract metadata JSON string. | [optional] 
**IsFavorite** | Pointer to **bool** |  | [optional] 

## Methods

### NewBaseContract

`func NewBaseContract(label string, contractName string, version string, rawAbi string, userDoc string, developerDoc string, ) *BaseContract`

NewBaseContract instantiates a new BaseContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseContractWithDefaults

`func NewBaseContractWithDefaults() *BaseContract`

NewBaseContractWithDefaults instantiates a new BaseContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *BaseContract) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BaseContract) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BaseContract) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetContractName

`func (o *BaseContract) GetContractName() string`

GetContractName returns the ContractName field if non-nil, zero value otherwise.

### GetContractNameOk

`func (o *BaseContract) GetContractNameOk() (*string, bool)`

GetContractNameOk returns a tuple with the ContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractName

`func (o *BaseContract) SetContractName(v string)`

SetContractName sets ContractName field to given value.


### GetVersion

`func (o *BaseContract) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BaseContract) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BaseContract) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetBin

`func (o *BaseContract) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *BaseContract) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *BaseContract) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *BaseContract) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetRawAbi

`func (o *BaseContract) GetRawAbi() string`

GetRawAbi returns the RawAbi field if non-nil, zero value otherwise.

### GetRawAbiOk

`func (o *BaseContract) GetRawAbiOk() (*string, bool)`

GetRawAbiOk returns a tuple with the RawAbi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAbi

`func (o *BaseContract) SetRawAbi(v string)`

SetRawAbi sets RawAbi field to given value.


### GetUserDoc

`func (o *BaseContract) GetUserDoc() string`

GetUserDoc returns the UserDoc field if non-nil, zero value otherwise.

### GetUserDocOk

`func (o *BaseContract) GetUserDocOk() (*string, bool)`

GetUserDocOk returns a tuple with the UserDoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDoc

`func (o *BaseContract) SetUserDoc(v string)`

SetUserDoc sets UserDoc field to given value.


### GetDeveloperDoc

`func (o *BaseContract) GetDeveloperDoc() string`

GetDeveloperDoc returns the DeveloperDoc field if non-nil, zero value otherwise.

### GetDeveloperDocOk

`func (o *BaseContract) GetDeveloperDocOk() (*string, bool)`

GetDeveloperDocOk returns a tuple with the DeveloperDoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperDoc

`func (o *BaseContract) SetDeveloperDoc(v string)`

SetDeveloperDoc sets DeveloperDoc field to given value.


### GetMetadata

`func (o *BaseContract) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BaseContract) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BaseContract) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BaseContract) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetIsFavorite

`func (o *BaseContract) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *BaseContract) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *BaseContract) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *BaseContract) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


