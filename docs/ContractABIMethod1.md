# ContractABIMethod1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A hex string. | 
**Name** | **string** | Name of the function. | 
**Signature** | **string** | The function signature. | 
**Const** | **bool** |  | 
**Payable** | **bool** |  | 
**Inputs** | [**[]ContractABIMethodArgument**](ContractABIMethodArgument.md) | List of function arguments. | 
**Outputs** | [**[]ContractABIMethodArgument**](ContractABIMethodArgument.md) | List of function outputs. | 
**Author** | **string** |  | 
**Notes** | **string** |  | 
**Description** | **string** | The function description. | 

## Methods

### NewContractABIMethod1

`func NewContractABIMethod1(id string, name string, signature string, const_ bool, payable bool, inputs []ContractABIMethodArgument, outputs []ContractABIMethodArgument, author string, notes string, description string, ) *ContractABIMethod1`

NewContractABIMethod1 instantiates a new ContractABIMethod1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractABIMethod1WithDefaults

`func NewContractABIMethod1WithDefaults() *ContractABIMethod1`

NewContractABIMethod1WithDefaults instantiates a new ContractABIMethod1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContractABIMethod1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractABIMethod1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractABIMethod1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ContractABIMethod1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractABIMethod1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractABIMethod1) SetName(v string)`

SetName sets Name field to given value.


### GetSignature

`func (o *ContractABIMethod1) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ContractABIMethod1) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ContractABIMethod1) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetConst

`func (o *ContractABIMethod1) GetConst() bool`

GetConst returns the Const field if non-nil, zero value otherwise.

### GetConstOk

`func (o *ContractABIMethod1) GetConstOk() (*bool, bool)`

GetConstOk returns a tuple with the Const field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConst

`func (o *ContractABIMethod1) SetConst(v bool)`

SetConst sets Const field to given value.


### GetPayable

`func (o *ContractABIMethod1) GetPayable() bool`

GetPayable returns the Payable field if non-nil, zero value otherwise.

### GetPayableOk

`func (o *ContractABIMethod1) GetPayableOk() (*bool, bool)`

GetPayableOk returns a tuple with the Payable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayable

`func (o *ContractABIMethod1) SetPayable(v bool)`

SetPayable sets Payable field to given value.


### GetInputs

`func (o *ContractABIMethod1) GetInputs() []ContractABIMethodArgument`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ContractABIMethod1) GetInputsOk() (*[]ContractABIMethodArgument, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ContractABIMethod1) SetInputs(v []ContractABIMethodArgument)`

SetInputs sets Inputs field to given value.


### GetOutputs

`func (o *ContractABIMethod1) GetOutputs() []ContractABIMethodArgument`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ContractABIMethod1) GetOutputsOk() (*[]ContractABIMethodArgument, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ContractABIMethod1) SetOutputs(v []ContractABIMethodArgument)`

SetOutputs sets Outputs field to given value.


### GetAuthor

`func (o *ContractABIMethod1) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ContractABIMethod1) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ContractABIMethod1) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetNotes

`func (o *ContractABIMethod1) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ContractABIMethod1) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ContractABIMethod1) SetNotes(v string)`

SetNotes sets Notes field to given value.


### GetDescription

`func (o *ContractABIMethod1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContractABIMethod1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContractABIMethod1) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


