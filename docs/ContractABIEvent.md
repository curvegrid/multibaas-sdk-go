# ContractABIEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The keccak256 hash as a hex string of 256 bits. | 
**Name** | **string** |  | 
**Signature** | **string** |  | 
**Anonymous** | **bool** |  | 
**Inputs** | [**[]ContractABIEventArgument**](ContractABIEventArgument.md) | List of contract event&#39;s input arguments. | 
**Notes** | **string** | The developer documentation. | 
**Description** | **string** | The user documentation. | 

## Methods

### NewContractABIEvent

`func NewContractABIEvent(id string, name string, signature string, anonymous bool, inputs []ContractABIEventArgument, notes string, description string, ) *ContractABIEvent`

NewContractABIEvent instantiates a new ContractABIEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractABIEventWithDefaults

`func NewContractABIEventWithDefaults() *ContractABIEvent`

NewContractABIEventWithDefaults instantiates a new ContractABIEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContractABIEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContractABIEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContractABIEvent) SetId(v string)`

SetId sets Id field to given value.


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


### GetNotes

`func (o *ContractABIEvent) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ContractABIEvent) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ContractABIEvent) SetNotes(v string)`

SetNotes sets Notes field to given value.


### GetDescription

`func (o *ContractABIEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContractABIEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContractABIEvent) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


