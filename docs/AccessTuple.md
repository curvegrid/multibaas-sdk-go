# AccessTuple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **NullableString** | An ethereum address. | 
**StorageKeys** | **[]string** |  | 

## Methods

### NewAccessTuple

`func NewAccessTuple(address NullableString, storageKeys []string, ) *AccessTuple`

NewAccessTuple instantiates a new AccessTuple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTupleWithDefaults

`func NewAccessTupleWithDefaults() *AccessTuple`

NewAccessTupleWithDefaults instantiates a new AccessTuple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AccessTuple) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AccessTuple) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AccessTuple) SetAddress(v string)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *AccessTuple) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *AccessTuple) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetStorageKeys

`func (o *AccessTuple) GetStorageKeys() []string`

GetStorageKeys returns the StorageKeys field if non-nil, zero value otherwise.

### GetStorageKeysOk

`func (o *AccessTuple) GetStorageKeysOk() (*[]string, bool)`

GetStorageKeysOk returns a tuple with the StorageKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKeys

`func (o *AccessTuple) SetStorageKeys(v []string)`

SetStorageKeys sets StorageKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


