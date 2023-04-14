# ContractABIMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the function. | 
**Signature** | **string** | The function signature. | 
**Const** | **bool** |  | 
**Payable** | **bool** |  | 
**Inputs** | [**[]ContractABIMethodArgument**](ContractABIMethodArgument.md) | List of function arguments. | 
**Outputs** | [**[]ContractABIMethodArgument**](ContractABIMethodArgument.md) | List of function outputs. | 
**Author** | **string** |  | 
**Notes** | **string** |  | 
**Returns** | **string** |  | 
**Description** | **string** | The function description. | 

## Methods

### NewContractABIMethod

`func NewContractABIMethod(name string, signature string, const_ bool, payable bool, inputs []ContractABIMethodArgument, outputs []ContractABIMethodArgument, author string, notes string, returns string, description string, ) *ContractABIMethod`

NewContractABIMethod instantiates a new ContractABIMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractABIMethodWithDefaults

`func NewContractABIMethodWithDefaults() *ContractABIMethod`

NewContractABIMethodWithDefaults instantiates a new ContractABIMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContractABIMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractABIMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractABIMethod) SetName(v string)`

SetName sets Name field to given value.


### GetSignature

`func (o *ContractABIMethod) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ContractABIMethod) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ContractABIMethod) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetConst

`func (o *ContractABIMethod) GetConst() bool`

GetConst returns the Const field if non-nil, zero value otherwise.

### GetConstOk

`func (o *ContractABIMethod) GetConstOk() (*bool, bool)`

GetConstOk returns a tuple with the Const field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConst

`func (o *ContractABIMethod) SetConst(v bool)`

SetConst sets Const field to given value.


### GetPayable

`func (o *ContractABIMethod) GetPayable() bool`

GetPayable returns the Payable field if non-nil, zero value otherwise.

### GetPayableOk

`func (o *ContractABIMethod) GetPayableOk() (*bool, bool)`

GetPayableOk returns a tuple with the Payable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayable

`func (o *ContractABIMethod) SetPayable(v bool)`

SetPayable sets Payable field to given value.


### GetInputs

`func (o *ContractABIMethod) GetInputs() []ContractABIMethodArgument`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ContractABIMethod) GetInputsOk() (*[]ContractABIMethodArgument, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ContractABIMethod) SetInputs(v []ContractABIMethodArgument)`

SetInputs sets Inputs field to given value.


### GetOutputs

`func (o *ContractABIMethod) GetOutputs() []ContractABIMethodArgument`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ContractABIMethod) GetOutputsOk() (*[]ContractABIMethodArgument, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ContractABIMethod) SetOutputs(v []ContractABIMethodArgument)`

SetOutputs sets Outputs field to given value.


### GetAuthor

`func (o *ContractABIMethod) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ContractABIMethod) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ContractABIMethod) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetNotes

`func (o *ContractABIMethod) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ContractABIMethod) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ContractABIMethod) SetNotes(v string)`

SetNotes sets Notes field to given value.


### GetReturns

`func (o *ContractABIMethod) GetReturns() string`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *ContractABIMethod) GetReturnsOk() (*string, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *ContractABIMethod) SetReturns(v string)`

SetReturns sets Returns field to given value.


### GetDescription

`func (o *ContractABIMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContractABIMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContractABIMethod) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


