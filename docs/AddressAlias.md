# AddressAlias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | An alias to easily identify and reference addresses. | 
**Address** | **string** | An ethereum address. | 

## Methods

### NewAddressAlias

`func NewAddressAlias(alias string, address string, ) *AddressAlias`

NewAddressAlias instantiates a new AddressAlias object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressAliasWithDefaults

`func NewAddressAliasWithDefaults() *AddressAlias`

NewAddressAliasWithDefaults instantiates a new AddressAlias object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *AddressAlias) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AddressAlias) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AddressAlias) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetAddress

`func (o *AddressAlias) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressAlias) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressAlias) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


