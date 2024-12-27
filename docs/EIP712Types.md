# EIP712Types

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EIP712Domain** | [**[]EIP712TypeEntry**](EIP712TypeEntry.md) |  | 

## Methods

### NewEIP712Types

`func NewEIP712Types(eIP712Domain []EIP712TypeEntry, ) *EIP712Types`

NewEIP712Types instantiates a new EIP712Types object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEIP712TypesWithDefaults

`func NewEIP712TypesWithDefaults() *EIP712Types`

NewEIP712TypesWithDefaults instantiates a new EIP712Types object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEIP712Domain

`func (o *EIP712Types) GetEIP712Domain() []EIP712TypeEntry`

GetEIP712Domain returns the EIP712Domain field if non-nil, zero value otherwise.

### GetEIP712DomainOk

`func (o *EIP712Types) GetEIP712DomainOk() (*[]EIP712TypeEntry, bool)`

GetEIP712DomainOk returns a tuple with the EIP712Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEIP712Domain

`func (o *EIP712Types) SetEIP712Domain(v []EIP712TypeEntry)`

SetEIP712Domain sets EIP712Domain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


