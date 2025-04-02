# CloudWalletTXToSignTx

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

### NewCloudWalletTXToSignTx

`func NewCloudWalletTXToSignTx(gas int64, from string, value string, data string, type_ int64, ) *CloudWalletTXToSignTx`

NewCloudWalletTXToSignTx instantiates a new CloudWalletTXToSignTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudWalletTXToSignTxWithDefaults

`func NewCloudWalletTXToSignTxWithDefaults() *CloudWalletTXToSignTx`

NewCloudWalletTXToSignTxWithDefaults instantiates a new CloudWalletTXToSignTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *CloudWalletTXToSignTx) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CloudWalletTXToSignTx) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CloudWalletTXToSignTx) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *CloudWalletTXToSignTx) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetGasPrice

`func (o *CloudWalletTXToSignTx) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *CloudWalletTXToSignTx) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *CloudWalletTXToSignTx) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *CloudWalletTXToSignTx) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetGasFeeCap

`func (o *CloudWalletTXToSignTx) GetGasFeeCap() string`

GetGasFeeCap returns the GasFeeCap field if non-nil, zero value otherwise.

### GetGasFeeCapOk

`func (o *CloudWalletTXToSignTx) GetGasFeeCapOk() (*string, bool)`

GetGasFeeCapOk returns a tuple with the GasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeCap

`func (o *CloudWalletTXToSignTx) SetGasFeeCap(v string)`

SetGasFeeCap sets GasFeeCap field to given value.

### HasGasFeeCap

`func (o *CloudWalletTXToSignTx) HasGasFeeCap() bool`

HasGasFeeCap returns a boolean if a field has been set.

### GetGasTipCap

`func (o *CloudWalletTXToSignTx) GetGasTipCap() string`

GetGasTipCap returns the GasTipCap field if non-nil, zero value otherwise.

### GetGasTipCapOk

`func (o *CloudWalletTXToSignTx) GetGasTipCapOk() (*string, bool)`

GetGasTipCapOk returns a tuple with the GasTipCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTipCap

`func (o *CloudWalletTXToSignTx) SetGasTipCap(v string)`

SetGasTipCap sets GasTipCap field to given value.

### HasGasTipCap

`func (o *CloudWalletTXToSignTx) HasGasTipCap() bool`

HasGasTipCap returns a boolean if a field has been set.

### GetGas

`func (o *CloudWalletTXToSignTx) GetGas() int64`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *CloudWalletTXToSignTx) GetGasOk() (*int64, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *CloudWalletTXToSignTx) SetGas(v int64)`

SetGas sets Gas field to given value.


### GetFrom

`func (o *CloudWalletTXToSignTx) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CloudWalletTXToSignTx) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CloudWalletTXToSignTx) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *CloudWalletTXToSignTx) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CloudWalletTXToSignTx) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CloudWalletTXToSignTx) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CloudWalletTXToSignTx) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *CloudWalletTXToSignTx) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *CloudWalletTXToSignTx) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetValue

`func (o *CloudWalletTXToSignTx) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CloudWalletTXToSignTx) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CloudWalletTXToSignTx) SetValue(v string)`

SetValue sets Value field to given value.


### GetData

`func (o *CloudWalletTXToSignTx) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudWalletTXToSignTx) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudWalletTXToSignTx) SetData(v string)`

SetData sets Data field to given value.


### GetHash

`func (o *CloudWalletTXToSignTx) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CloudWalletTXToSignTx) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CloudWalletTXToSignTx) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *CloudWalletTXToSignTx) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetType

`func (o *CloudWalletTXToSignTx) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudWalletTXToSignTx) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudWalletTXToSignTx) SetType(v int64)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


