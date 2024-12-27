# HSMSignRequestTypedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The signing method to use. | 
**Address** | **string** | An ethereum address. | 
**Data** | [**EIP712TypedData**](EIP712TypedData.md) |  | 

## Methods

### NewHSMSignRequestTypedData

`func NewHSMSignRequestTypedData(method string, address string, data EIP712TypedData, ) *HSMSignRequestTypedData`

NewHSMSignRequestTypedData instantiates a new HSMSignRequestTypedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHSMSignRequestTypedDataWithDefaults

`func NewHSMSignRequestTypedDataWithDefaults() *HSMSignRequestTypedData`

NewHSMSignRequestTypedDataWithDefaults instantiates a new HSMSignRequestTypedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *HSMSignRequestTypedData) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HSMSignRequestTypedData) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HSMSignRequestTypedData) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetAddress

`func (o *HSMSignRequestTypedData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HSMSignRequestTypedData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HSMSignRequestTypedData) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetData

`func (o *HSMSignRequestTypedData) GetData() EIP712TypedData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HSMSignRequestTypedData) GetDataOk() (*EIP712TypedData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HSMSignRequestTypedData) SetData(v EIP712TypedData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


