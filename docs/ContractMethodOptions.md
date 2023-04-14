# ContractMethodOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | Pointer to **string** | The function signature. | [optional] 
**Inputs** | [**[]ContractParameter**](ContractParameter.md) | List of function input parameters. | 
**Outputs** | Pointer to [**[]ContractParameter**](ContractParameter.md) | List of function output parameters. | [optional] 

## Methods

### NewContractMethodOptions

`func NewContractMethodOptions(inputs []ContractParameter, ) *ContractMethodOptions`

NewContractMethodOptions instantiates a new ContractMethodOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractMethodOptionsWithDefaults

`func NewContractMethodOptionsWithDefaults() *ContractMethodOptions`

NewContractMethodOptionsWithDefaults instantiates a new ContractMethodOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *ContractMethodOptions) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ContractMethodOptions) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ContractMethodOptions) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ContractMethodOptions) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetInputs

`func (o *ContractMethodOptions) GetInputs() []ContractParameter`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ContractMethodOptions) GetInputsOk() (*[]ContractParameter, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ContractMethodOptions) SetInputs(v []ContractParameter)`

SetInputs sets Inputs field to given value.


### GetOutputs

`func (o *ContractMethodOptions) GetOutputs() []ContractParameter`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ContractMethodOptions) GetOutputsOk() (*[]ContractParameter, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ContractMethodOptions) SetOutputs(v []ContractParameter)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *ContractMethodOptions) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


