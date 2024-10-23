# ContractMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | An alias to easily identify and reference the entity in subsequent requests. | 
**Name** | **string** | The name of the contract. | 
**Version** | **string** | The contract version. | 

## Methods

### NewContractMetadata

`func NewContractMetadata(label string, name string, version string, ) *ContractMetadata`

NewContractMetadata instantiates a new ContractMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractMetadataWithDefaults

`func NewContractMetadataWithDefaults() *ContractMetadata`

NewContractMetadataWithDefaults instantiates a new ContractMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ContractMetadata) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContractMetadata) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContractMetadata) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *ContractMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *ContractMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContractMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContractMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


