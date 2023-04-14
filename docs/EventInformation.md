# EventInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the event. | 
**Signature** | **string** | The event signature. | 
**Inputs** | [**[]EventField**](EventField.md) | The list of input arguments for the event. | 
**RawFields** | Pointer to **string** | The raw output from an event. Omitted when returned as part of a transaction receipt. | [optional] 
**Contract** | [**ContractInformation**](ContractInformation.md) |  | 
**IndexInLog** | **int64** | The event&#39;s index in the log. | 

## Methods

### NewEventInformation

`func NewEventInformation(name string, signature string, inputs []EventField, contract ContractInformation, indexInLog int64, ) *EventInformation`

NewEventInformation instantiates a new EventInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventInformationWithDefaults

`func NewEventInformationWithDefaults() *EventInformation`

NewEventInformationWithDefaults instantiates a new EventInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventInformation) SetName(v string)`

SetName sets Name field to given value.


### GetSignature

`func (o *EventInformation) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *EventInformation) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *EventInformation) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetInputs

`func (o *EventInformation) GetInputs() []EventField`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *EventInformation) GetInputsOk() (*[]EventField, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *EventInformation) SetInputs(v []EventField)`

SetInputs sets Inputs field to given value.


### GetRawFields

`func (o *EventInformation) GetRawFields() string`

GetRawFields returns the RawFields field if non-nil, zero value otherwise.

### GetRawFieldsOk

`func (o *EventInformation) GetRawFieldsOk() (*string, bool)`

GetRawFieldsOk returns a tuple with the RawFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawFields

`func (o *EventInformation) SetRawFields(v string)`

SetRawFields sets RawFields field to given value.

### HasRawFields

`func (o *EventInformation) HasRawFields() bool`

HasRawFields returns a boolean if a field has been set.

### GetContract

`func (o *EventInformation) GetContract() ContractInformation`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *EventInformation) GetContractOk() (*ContractInformation, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *EventInformation) SetContract(v ContractInformation)`

SetContract sets Contract field to given value.


### GetIndexInLog

`func (o *EventInformation) GetIndexInLog() int64`

GetIndexInLog returns the IndexInLog field if non-nil, zero value otherwise.

### GetIndexInLogOk

`func (o *EventInformation) GetIndexInLogOk() (*int64, bool)`

GetIndexInLogOk returns a tuple with the IndexInLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexInLog

`func (o *EventInformation) SetIndexInLog(v int64)`

SetIndexInLog sets IndexInLog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


