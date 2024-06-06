# ContractABIError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**Name** | **string** |  | 
**Signature** | **string** |  | 
**Inputs** | [**[]ContractABIErrorArgument**](ContractABIErrorArgument.md) | List of contract event&#39;s input arguments. | 

## Methods

### NewContractABIError

`func NewContractABIError(id string, name string, signature string, inputs []ContractABIErrorArgument, ) *ContractABIError`

NewContractABIError instantiates a new ContractABIError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractABIErrorWithDefaults

`func NewContractABIErrorWithDefaults() *ContractABIError`

NewContractABIErrorWithDefaults instantiates a new ContractABIError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContractABIError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractABIError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractABIError) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ContractABIError) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractABIError) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractABIError) SetName(v string)`

SetName sets Name field to given value.


### GetSignature

`func (o *ContractABIError) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ContractABIError) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ContractABIError) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetInputs

`func (o *ContractABIError) GetInputs() []ContractABIErrorArgument`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ContractABIError) GetInputsOk() (*[]ContractABIErrorArgument, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ContractABIError) SetInputs(v []ContractABIErrorArgument)`

SetInputs sets Inputs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


