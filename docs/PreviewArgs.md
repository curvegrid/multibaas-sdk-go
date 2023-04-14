# PreviewArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputsOnly** | **bool** | Only preview the effect of a Type Conversion on the inputs. Only applicable for write function calls, where the output is an unsigned transaction. | 
**Inputs** | [**[]ContractABIMethodArgument**](ContractABIMethodArgument.md) | Type Conversion information for the function inputs. The number of inputs must match that of the actual function inputs. The parameter is a contract function argument where only the type conversion information is used. | 
**Outputs** | [**[]ContractABIMethodArgument**](ContractABIMethodArgument.md) | Type Conversion information for the function outputs. The number of outputs must match that of the actual function outputs. The parameter is a contract function argument where only the type conversion information is used. | 

## Methods

### NewPreviewArgs

`func NewPreviewArgs(inputsOnly bool, inputs []ContractABIMethodArgument, outputs []ContractABIMethodArgument, ) *PreviewArgs`

NewPreviewArgs instantiates a new PreviewArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewArgsWithDefaults

`func NewPreviewArgsWithDefaults() *PreviewArgs`

NewPreviewArgsWithDefaults instantiates a new PreviewArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputsOnly

`func (o *PreviewArgs) GetInputsOnly() bool`

GetInputsOnly returns the InputsOnly field if non-nil, zero value otherwise.

### GetInputsOnlyOk

`func (o *PreviewArgs) GetInputsOnlyOk() (*bool, bool)`

GetInputsOnlyOk returns a tuple with the InputsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputsOnly

`func (o *PreviewArgs) SetInputsOnly(v bool)`

SetInputsOnly sets InputsOnly field to given value.


### GetInputs

`func (o *PreviewArgs) GetInputs() []ContractABIMethodArgument`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *PreviewArgs) GetInputsOk() (*[]ContractABIMethodArgument, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *PreviewArgs) SetInputs(v []ContractABIMethodArgument)`

SetInputs sets Inputs field to given value.


### GetOutputs

`func (o *PreviewArgs) GetOutputs() []ContractABIMethodArgument`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *PreviewArgs) GetOutputsOk() (*[]ContractABIMethodArgument, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *PreviewArgs) SetOutputs(v []ContractABIMethodArgument)`

SetOutputs sets Outputs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


