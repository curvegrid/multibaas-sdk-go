# ContractInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | An ethereum address. | 
**AddressAlias** | **string** | An alias to easily identify and reference addresses. | 
**Name** | **string** | The name of the contract. | 
**Label** | **string** | An alias to easily identify and reference the entity in subsequent requests. | 

## Methods

### NewContractInformation

`func NewContractInformation(address string, addressAlias string, name string, label string, ) *ContractInformation`

NewContractInformation instantiates a new ContractInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractInformationWithDefaults

`func NewContractInformationWithDefaults() *ContractInformation`

NewContractInformationWithDefaults instantiates a new ContractInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ContractInformation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractInformation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractInformation) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressAlias

`func (o *ContractInformation) GetAddressAlias() string`

GetAddressAlias returns the AddressAlias field if non-nil, zero value otherwise.

### GetAddressAliasOk

`func (o *ContractInformation) GetAddressAliasOk() (*string, bool)`

GetAddressAliasOk returns a tuple with the AddressAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAlias

`func (o *ContractInformation) SetAddressAlias(v string)`

SetAddressAlias sets AddressAlias field to given value.


### GetName

`func (o *ContractInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractInformation) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ContractInformation) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContractInformation) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContractInformation) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


