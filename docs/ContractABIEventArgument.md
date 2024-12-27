# ContractABIEventArgument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**ContractABIType**](ContractABIType.md) |  | 
**TypeName** | **string** |  | 
**Indexed** | **bool** |  | 
**TypeConversion** | [**NullableContractABITypeConversion**](ContractABITypeConversion.md) |  | 
**Notes** | **string** | The developer documentation. | 

## Methods

### NewContractABIEventArgument

`func NewContractABIEventArgument(name string, type_ ContractABIType, typeName string, indexed bool, typeConversion NullableContractABITypeConversion, notes string, ) *ContractABIEventArgument`

NewContractABIEventArgument instantiates a new ContractABIEventArgument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractABIEventArgumentWithDefaults

`func NewContractABIEventArgumentWithDefaults() *ContractABIEventArgument`

NewContractABIEventArgumentWithDefaults instantiates a new ContractABIEventArgument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContractABIEventArgument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractABIEventArgument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractABIEventArgument) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ContractABIEventArgument) GetType() ContractABIType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContractABIEventArgument) GetTypeOk() (*ContractABIType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContractABIEventArgument) SetType(v ContractABIType)`

SetType sets Type field to given value.


### GetTypeName

`func (o *ContractABIEventArgument) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *ContractABIEventArgument) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *ContractABIEventArgument) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetIndexed

`func (o *ContractABIEventArgument) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *ContractABIEventArgument) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *ContractABIEventArgument) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.


### GetTypeConversion

`func (o *ContractABIEventArgument) GetTypeConversion() ContractABITypeConversion`

GetTypeConversion returns the TypeConversion field if non-nil, zero value otherwise.

### GetTypeConversionOk

`func (o *ContractABIEventArgument) GetTypeConversionOk() (*ContractABITypeConversion, bool)`

GetTypeConversionOk returns a tuple with the TypeConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeConversion

`func (o *ContractABIEventArgument) SetTypeConversion(v ContractABITypeConversion)`

SetTypeConversion sets TypeConversion field to given value.


### SetTypeConversionNil

`func (o *ContractABIEventArgument) SetTypeConversionNil(b bool)`

 SetTypeConversionNil sets the value for TypeConversion to be an explicit nil

### UnsetTypeConversion
`func (o *ContractABIEventArgument) UnsetTypeConversion()`

UnsetTypeConversion ensures that no value is present for TypeConversion, not even an explicit nil
### GetNotes

`func (o *ContractABIEventArgument) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ContractABIEventArgument) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ContractABIEventArgument) SetNotes(v string)`

SetNotes sets Notes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


