# HSMSignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The signing method to use. | 
**Address** | **string** | An ethereum address. | 
**Data** | [**EIP712TypedData**](EIP712TypedData.md) |  | 
**ChainId** | Pointer to [**HSMSignRequestPersonalSignChainId**](HSMSignRequestPersonalSignChainId.md) |  | [optional] 

## Methods

### NewHSMSignRequest

`func NewHSMSignRequest(method string, address string, data EIP712TypedData, ) *HSMSignRequest`

NewHSMSignRequest instantiates a new HSMSignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHSMSignRequestWithDefaults

`func NewHSMSignRequestWithDefaults() *HSMSignRequest`

NewHSMSignRequestWithDefaults instantiates a new HSMSignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *HSMSignRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HSMSignRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HSMSignRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetAddress

`func (o *HSMSignRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HSMSignRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HSMSignRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetData

`func (o *HSMSignRequest) GetData() EIP712TypedData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HSMSignRequest) GetDataOk() (*EIP712TypedData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HSMSignRequest) SetData(v EIP712TypedData)`

SetData sets Data field to given value.


### GetChainId

`func (o *HSMSignRequest) GetChainId() HSMSignRequestPersonalSignChainId`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *HSMSignRequest) GetChainIdOk() (*HSMSignRequestPersonalSignChainId, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *HSMSignRequest) SetChainId(v HSMSignRequestPersonalSignChainId)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *HSMSignRequest) HasChainId() bool`

HasChainId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


