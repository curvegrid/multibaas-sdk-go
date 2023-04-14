# ContractABIEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Signature** | **string** |  | 
**Anonymous** | **bool** |  | 
**Inputs** | [**[]ContractABIEventArgument**](ContractABIEventArgument.md) | List of contract event&#39;s input arguments. | 

## Methods

### NewContractABIEvent

`func NewContractABIEvent(name string, signature string, anonymous bool, inputs []ContractABIEventArgument, ) *ContractABIEvent`

NewContractABIEvent instantiates a new ContractABIEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractABIEventWithDefaults

`func NewContractABIEventWithDefaults() *ContractABIEvent`

NewContractABIEventWithDefaults instantiates a new ContractABIEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContractABIEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractABIEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractABIEvent) SetName(v string)`

SetName sets Name field to given value.


### GetSignature

`func (o *ContractABIEvent) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ContractABIEvent) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ContractABIEvent) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetAnonymous

`func (o *ContractABIEvent) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *ContractABIEvent) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *ContractABIEvent) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.


### GetInputs

`func (o *ContractABIEvent) GetInputs() []ContractABIEventArgument`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ContractABIEvent) GetInputsOk() (*[]ContractABIEventArgument, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ContractABIEvent) SetInputs(v []ContractABIEventArgument)`

SetInputs sets Inputs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


