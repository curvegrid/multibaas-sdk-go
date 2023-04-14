# ContractEventOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | Pointer to **string** |  | [optional] 
**Inputs** | [**[]ContractParameter**](ContractParameter.md) |  | 

## Methods

### NewContractEventOptions

`func NewContractEventOptions(inputs []ContractParameter, ) *ContractEventOptions`

NewContractEventOptions instantiates a new ContractEventOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractEventOptionsWithDefaults

`func NewContractEventOptionsWithDefaults() *ContractEventOptions`

NewContractEventOptionsWithDefaults instantiates a new ContractEventOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *ContractEventOptions) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ContractEventOptions) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ContractEventOptions) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ContractEventOptions) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetInputs

`func (o *ContractEventOptions) GetInputs() []ContractParameter`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ContractEventOptions) GetInputsOk() (*[]ContractParameter, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ContractEventOptions) SetInputs(v []ContractParameter)`

SetInputs sets Inputs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


