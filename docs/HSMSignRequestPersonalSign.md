# HSMSignRequestPersonalSign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The signing method to use. | 
**Address** | **string** | An ethereum address. | 
**Data** | **string** | A hex string. | 
**ChainId** | Pointer to [**HSMSignRequestPersonalSignChainId**](HSMSignRequestPersonalSignChainId.md) |  | [optional] 

## Methods

### NewHSMSignRequestPersonalSign

`func NewHSMSignRequestPersonalSign(method string, address string, data string, ) *HSMSignRequestPersonalSign`

NewHSMSignRequestPersonalSign instantiates a new HSMSignRequestPersonalSign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHSMSignRequestPersonalSignWithDefaults

`func NewHSMSignRequestPersonalSignWithDefaults() *HSMSignRequestPersonalSign`

NewHSMSignRequestPersonalSignWithDefaults instantiates a new HSMSignRequestPersonalSign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *HSMSignRequestPersonalSign) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HSMSignRequestPersonalSign) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HSMSignRequestPersonalSign) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetAddress

`func (o *HSMSignRequestPersonalSign) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HSMSignRequestPersonalSign) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HSMSignRequestPersonalSign) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetData

`func (o *HSMSignRequestPersonalSign) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HSMSignRequestPersonalSign) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HSMSignRequestPersonalSign) SetData(v string)`

SetData sets Data field to given value.


### GetChainId

`func (o *HSMSignRequestPersonalSign) GetChainId() HSMSignRequestPersonalSignChainId`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *HSMSignRequestPersonalSign) GetChainIdOk() (*HSMSignRequestPersonalSignChainId, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *HSMSignRequestPersonalSign) SetChainId(v HSMSignRequestPersonalSignChainId)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *HSMSignRequestPersonalSign) HasChainId() bool`

HasChainId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


