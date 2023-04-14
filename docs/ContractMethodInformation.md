# ContractMethodInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the method. | 
**Signature** | **string** | The method signature. | 
**Inputs** | [**[]MethodArg**](MethodArg.md) |  | 

## Methods

### NewContractMethodInformation

`func NewContractMethodInformation(name string, signature string, inputs []MethodArg, ) *ContractMethodInformation`

NewContractMethodInformation instantiates a new ContractMethodInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractMethodInformationWithDefaults

`func NewContractMethodInformationWithDefaults() *ContractMethodInformation`

NewContractMethodInformationWithDefaults instantiates a new ContractMethodInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContractMethodInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractMethodInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractMethodInformation) SetName(v string)`

SetName sets Name field to given value.


### GetSignature

`func (o *ContractMethodInformation) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ContractMethodInformation) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ContractMethodInformation) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetInputs

`func (o *ContractMethodInformation) GetInputs() []MethodArg`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ContractMethodInformation) GetInputsOk() (*[]MethodArg, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ContractMethodInformation) SetInputs(v []MethodArg)`

SetInputs sets Inputs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


