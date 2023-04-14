# HSMSignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | An ethereum address. | 
**IsTyped** | Pointer to **bool** | Is the data field an encapsulated EIP-712 typed message? | [optional] 
**Data** | **string** | Data to sign | 
**ChainId** | Pointer to [**HSMSignRequestChainId**](HSMSignRequestChainId.md) |  | [optional] 

## Methods

### NewHSMSignRequest

`func NewHSMSignRequest(address string, data string, ) *HSMSignRequest`

NewHSMSignRequest instantiates a new HSMSignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHSMSignRequestWithDefaults

`func NewHSMSignRequestWithDefaults() *HSMSignRequest`

NewHSMSignRequestWithDefaults instantiates a new HSMSignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetIsTyped

`func (o *HSMSignRequest) GetIsTyped() bool`

GetIsTyped returns the IsTyped field if non-nil, zero value otherwise.

### GetIsTypedOk

`func (o *HSMSignRequest) GetIsTypedOk() (*bool, bool)`

GetIsTypedOk returns a tuple with the IsTyped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTyped

`func (o *HSMSignRequest) SetIsTyped(v bool)`

SetIsTyped sets IsTyped field to given value.

### HasIsTyped

`func (o *HSMSignRequest) HasIsTyped() bool`

HasIsTyped returns a boolean if a field has been set.

### GetData

`func (o *HSMSignRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HSMSignRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HSMSignRequest) SetData(v string)`

SetData sets Data field to given value.


### GetChainId

`func (o *HSMSignRequest) GetChainId() HSMSignRequestChainId`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *HSMSignRequest) GetChainIdOk() (*HSMSignRequestChainId, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *HSMSignRequest) SetChainId(v HSMSignRequestChainId)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *HSMSignRequest) HasChainId() bool`

HasChainId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


