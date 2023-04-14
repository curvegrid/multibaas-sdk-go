# ContractABIMethodArgument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ContractABIType**](ContractABIType.md) |  | [optional] 
**TypeConversion** | [**NullableContractABITypeConversion**](ContractABITypeConversion.md) |  | 
**Notes** | Pointer to **string** |  | [optional] 

## Methods

### NewContractABIMethodArgument

`func NewContractABIMethodArgument(typeConversion NullableContractABITypeConversion, ) *ContractABIMethodArgument`

NewContractABIMethodArgument instantiates a new ContractABIMethodArgument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractABIMethodArgumentWithDefaults

`func NewContractABIMethodArgumentWithDefaults() *ContractABIMethodArgument`

NewContractABIMethodArgumentWithDefaults instantiates a new ContractABIMethodArgument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContractABIMethodArgument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractABIMethodArgument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractABIMethodArgument) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContractABIMethodArgument) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ContractABIMethodArgument) GetType() ContractABIType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContractABIMethodArgument) GetTypeOk() (*ContractABIType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContractABIMethodArgument) SetType(v ContractABIType)`

SetType sets Type field to given value.

### HasType

`func (o *ContractABIMethodArgument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeConversion

`func (o *ContractABIMethodArgument) GetTypeConversion() ContractABITypeConversion`

GetTypeConversion returns the TypeConversion field if non-nil, zero value otherwise.

### GetTypeConversionOk

`func (o *ContractABIMethodArgument) GetTypeConversionOk() (*ContractABITypeConversion, bool)`

GetTypeConversionOk returns a tuple with the TypeConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeConversion

`func (o *ContractABIMethodArgument) SetTypeConversion(v ContractABITypeConversion)`

SetTypeConversion sets TypeConversion field to given value.


### SetTypeConversionNil

`func (o *ContractABIMethodArgument) SetTypeConversionNil(b bool)`

 SetTypeConversionNil sets the value for TypeConversion to be an explicit nil

### UnsetTypeConversion
`func (o *ContractABIMethodArgument) UnsetTypeConversion()`

UnsetTypeConversion ensures that no value is present for TypeConversion, not even an explicit nil
### GetNotes

`func (o *ContractABIMethodArgument) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ContractABIMethodArgument) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ContractABIMethodArgument) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ContractABIMethodArgument) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


