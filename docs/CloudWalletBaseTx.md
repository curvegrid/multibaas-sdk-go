# CloudWalletBaseTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | Pointer to **int64** | Sender account nonce of the transaction | [optional] 
**GasPrice** | Pointer to **string** | Gas price of the transaction | [optional] 
**GasFeeCap** | Pointer to **string** | Fee cap per gas of the transaction | [optional] 
**GasTipCap** | Pointer to **string** | GasTipCap per gas of the transaction | [optional] 
**Gas** | **int64** | Gas limit of the transaction | 
**From** | **string** | An ethereum address. | 
**To** | Pointer to **NullableString** | An ethereum address. | [optional] 
**Value** | **string** | Amount (in wei) to send with the transaction. | 
**Data** | **string** | A hex string. | 
**Hash** | Pointer to **string** | The keccak256 hash as a hex string of 256 bits. | [optional] 
**Type** | **int64** | Transaction type | 

## Methods

### NewCloudWalletBaseTx

`func NewCloudWalletBaseTx(gas int64, from string, value string, data string, type_ int64, ) *CloudWalletBaseTx`

NewCloudWalletBaseTx instantiates a new CloudWalletBaseTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudWalletBaseTxWithDefaults

`func NewCloudWalletBaseTxWithDefaults() *CloudWalletBaseTx`

NewCloudWalletBaseTxWithDefaults instantiates a new CloudWalletBaseTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *CloudWalletBaseTx) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CloudWalletBaseTx) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CloudWalletBaseTx) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *CloudWalletBaseTx) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetGasPrice

`func (o *CloudWalletBaseTx) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *CloudWalletBaseTx) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *CloudWalletBaseTx) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *CloudWalletBaseTx) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetGasFeeCap

`func (o *CloudWalletBaseTx) GetGasFeeCap() string`

GetGasFeeCap returns the GasFeeCap field if non-nil, zero value otherwise.

### GetGasFeeCapOk

`func (o *CloudWalletBaseTx) GetGasFeeCapOk() (*string, bool)`

GetGasFeeCapOk returns a tuple with the GasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeCap

`func (o *CloudWalletBaseTx) SetGasFeeCap(v string)`

SetGasFeeCap sets GasFeeCap field to given value.

### HasGasFeeCap

`func (o *CloudWalletBaseTx) HasGasFeeCap() bool`

HasGasFeeCap returns a boolean if a field has been set.

### GetGasTipCap

`func (o *CloudWalletBaseTx) GetGasTipCap() string`

GetGasTipCap returns the GasTipCap field if non-nil, zero value otherwise.

### GetGasTipCapOk

`func (o *CloudWalletBaseTx) GetGasTipCapOk() (*string, bool)`

GetGasTipCapOk returns a tuple with the GasTipCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTipCap

`func (o *CloudWalletBaseTx) SetGasTipCap(v string)`

SetGasTipCap sets GasTipCap field to given value.

### HasGasTipCap

`func (o *CloudWalletBaseTx) HasGasTipCap() bool`

HasGasTipCap returns a boolean if a field has been set.

### GetGas

`func (o *CloudWalletBaseTx) GetGas() int64`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *CloudWalletBaseTx) GetGasOk() (*int64, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *CloudWalletBaseTx) SetGas(v int64)`

SetGas sets Gas field to given value.


### GetFrom

`func (o *CloudWalletBaseTx) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CloudWalletBaseTx) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CloudWalletBaseTx) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *CloudWalletBaseTx) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CloudWalletBaseTx) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CloudWalletBaseTx) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CloudWalletBaseTx) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *CloudWalletBaseTx) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *CloudWalletBaseTx) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetValue

`func (o *CloudWalletBaseTx) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CloudWalletBaseTx) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CloudWalletBaseTx) SetValue(v string)`

SetValue sets Value field to given value.


### GetData

`func (o *CloudWalletBaseTx) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudWalletBaseTx) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudWalletBaseTx) SetData(v string)`

SetData sets Data field to given value.


### GetHash

`func (o *CloudWalletBaseTx) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CloudWalletBaseTx) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CloudWalletBaseTx) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *CloudWalletBaseTx) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetType

`func (o *CloudWalletBaseTx) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudWalletBaseTx) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudWalletBaseTx) SetType(v int64)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


