# ContractABIType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Size** | Pointer to **int64** |  | [optional] 
**TupleElems** | Pointer to [**[]ContractABIType**](ContractABIType.md) |  | [optional] 
**TupleRawNames** | Pointer to **[]string** |  | [optional] 
**Elem** | Pointer to [**ContractABIType**](ContractABIType.md) |  | [optional] 

## Methods

### NewContractABIType

`func NewContractABIType(type_ string, ) *ContractABIType`

NewContractABIType instantiates a new ContractABIType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractABITypeWithDefaults

`func NewContractABITypeWithDefaults() *ContractABIType`

NewContractABITypeWithDefaults instantiates a new ContractABIType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContractABIType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContractABIType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContractABIType) SetType(v string)`

SetType sets Type field to given value.


### GetSize

`func (o *ContractABIType) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContractABIType) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContractABIType) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ContractABIType) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTupleElems

`func (o *ContractABIType) GetTupleElems() []ContractABIType`

GetTupleElems returns the TupleElems field if non-nil, zero value otherwise.

### GetTupleElemsOk

`func (o *ContractABIType) GetTupleElemsOk() (*[]ContractABIType, bool)`

GetTupleElemsOk returns a tuple with the TupleElems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleElems

`func (o *ContractABIType) SetTupleElems(v []ContractABIType)`

SetTupleElems sets TupleElems field to given value.

### HasTupleElems

`func (o *ContractABIType) HasTupleElems() bool`

HasTupleElems returns a boolean if a field has been set.

### GetTupleRawNames

`func (o *ContractABIType) GetTupleRawNames() []string`

GetTupleRawNames returns the TupleRawNames field if non-nil, zero value otherwise.

### GetTupleRawNamesOk

`func (o *ContractABIType) GetTupleRawNamesOk() (*[]string, bool)`

GetTupleRawNamesOk returns a tuple with the TupleRawNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleRawNames

`func (o *ContractABIType) SetTupleRawNames(v []string)`

SetTupleRawNames sets TupleRawNames field to given value.

### HasTupleRawNames

`func (o *ContractABIType) HasTupleRawNames() bool`

HasTupleRawNames returns a boolean if a field has been set.

### GetElem

`func (o *ContractABIType) GetElem() ContractABIType`

GetElem returns the Elem field if non-nil, zero value otherwise.

### GetElemOk

`func (o *ContractABIType) GetElemOk() (*ContractABIType, bool)`

GetElemOk returns a tuple with the Elem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElem

`func (o *ContractABIType) SetElem(v ContractABIType)`

SetElem sets Elem field to given value.

### HasElem

`func (o *ContractABIType) HasElem() bool`

HasElem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


