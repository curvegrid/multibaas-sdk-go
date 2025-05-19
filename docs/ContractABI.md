# ContractABI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constructor** | [**ContractABIMethod1**](ContractABIMethod1.md) |  | 
**Methods** | [**map[string]ContractABIMethod**](ContractABIMethod.md) |  | 
**Events** | [**map[string]ContractABIEvent**](ContractABIEvent.md) |  | 
**Errors** | Pointer to [**map[string]ContractABIError**](ContractABIError.md) |  | [optional] 
**Fallback** | [**ContractABIMethod**](ContractABIMethod.md) |  | 
**Receive** | [**ContractABIMethod**](ContractABIMethod.md) |  | 

## Methods

### NewContractABI

`func NewContractABI(constructor ContractABIMethod1, methods map[string]ContractABIMethod, events map[string]ContractABIEvent, fallback ContractABIMethod, receive ContractABIMethod, ) *ContractABI`

NewContractABI instantiates a new ContractABI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractABIWithDefaults

`func NewContractABIWithDefaults() *ContractABI`

NewContractABIWithDefaults instantiates a new ContractABI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstructor

`func (o *ContractABI) GetConstructor() ContractABIMethod1`

GetConstructor returns the Constructor field if non-nil, zero value otherwise.

### GetConstructorOk

`func (o *ContractABI) GetConstructorOk() (*ContractABIMethod1, bool)`

GetConstructorOk returns a tuple with the Constructor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstructor

`func (o *ContractABI) SetConstructor(v ContractABIMethod1)`

SetConstructor sets Constructor field to given value.


### GetMethods

`func (o *ContractABI) GetMethods() map[string]ContractABIMethod`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *ContractABI) GetMethodsOk() (*map[string]ContractABIMethod, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *ContractABI) SetMethods(v map[string]ContractABIMethod)`

SetMethods sets Methods field to given value.


### GetEvents

`func (o *ContractABI) GetEvents() map[string]ContractABIEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ContractABI) GetEventsOk() (*map[string]ContractABIEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ContractABI) SetEvents(v map[string]ContractABIEvent)`

SetEvents sets Events field to given value.


### GetErrors

`func (o *ContractABI) GetErrors() map[string]ContractABIError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ContractABI) GetErrorsOk() (*map[string]ContractABIError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ContractABI) SetErrors(v map[string]ContractABIError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ContractABI) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetFallback

`func (o *ContractABI) GetFallback() ContractABIMethod`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *ContractABI) GetFallbackOk() (*ContractABIMethod, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *ContractABI) SetFallback(v ContractABIMethod)`

SetFallback sets Fallback field to given value.


### GetReceive

`func (o *ContractABI) GetReceive() ContractABIMethod`

GetReceive returns the Receive field if non-nil, zero value otherwise.

### GetReceiveOk

`func (o *ContractABI) GetReceiveOk() (*ContractABIMethod, bool)`

GetReceiveOk returns a tuple with the Receive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceive

`func (o *ContractABI) SetReceive(v ContractABIMethod)`

SetReceive sets Receive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


