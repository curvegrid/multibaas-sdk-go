# EIP712TypedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Types** | [**EIP712Types**](EIP712Types.md) |  | 
**PrimaryType** | **string** | The root type of the message object. Must correspond to a key in the &#x60;types&#x60; object. | 
**Domain** | [**EIP712Domain**](EIP712Domain.md) |  | 
**Message** | **map[string]interface{}** | The actual data, conforming to the &#x60;primaryType&#x60; definition in &#x60;types&#x60;. | 

## Methods

### NewEIP712TypedData

`func NewEIP712TypedData(types EIP712Types, primaryType string, domain EIP712Domain, message map[string]interface{}, ) *EIP712TypedData`

NewEIP712TypedData instantiates a new EIP712TypedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEIP712TypedDataWithDefaults

`func NewEIP712TypedDataWithDefaults() *EIP712TypedData`

NewEIP712TypedDataWithDefaults instantiates a new EIP712TypedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypes

`func (o *EIP712TypedData) GetTypes() EIP712Types`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *EIP712TypedData) GetTypesOk() (*EIP712Types, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *EIP712TypedData) SetTypes(v EIP712Types)`

SetTypes sets Types field to given value.


### GetPrimaryType

`func (o *EIP712TypedData) GetPrimaryType() string`

GetPrimaryType returns the PrimaryType field if non-nil, zero value otherwise.

### GetPrimaryTypeOk

`func (o *EIP712TypedData) GetPrimaryTypeOk() (*string, bool)`

GetPrimaryTypeOk returns a tuple with the PrimaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryType

`func (o *EIP712TypedData) SetPrimaryType(v string)`

SetPrimaryType sets PrimaryType field to given value.


### GetDomain

`func (o *EIP712TypedData) GetDomain() EIP712Domain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EIP712TypedData) GetDomainOk() (*EIP712Domain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EIP712TypedData) SetDomain(v EIP712Domain)`

SetDomain sets Domain field to given value.


### GetMessage

`func (o *EIP712TypedData) GetMessage() map[string]interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EIP712TypedData) GetMessageOk() (*map[string]interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EIP712TypedData) SetMessage(v map[string]interface{})`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


