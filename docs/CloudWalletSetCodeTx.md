# CloudWalletSetCodeTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | Pointer to **int64** | Sender account nonce of the transaction | [optional] 
**GasPrice** | Pointer to **string** | Gas price of the transaction | [optional] 
**GasFeeCap** | Pointer to **string** | Fee cap per gas of the transaction | [optional] 
**GasTipCap** | Pointer to **string** | GasTipCap per gas of the transaction | [optional] 
**Gas** | **int64** | Gas limit of the transaction | 
**From** | **string** | An ethereum address. | 
**To** | **string** | An ethereum address. | 
**Value** | **string** | Amount (in wei) to send with the transaction. | 
**Data** | **string** | A hex string. | 
**Hash** | Pointer to **string** | The keccak256 hash as a hex string of 256 bits. | [optional] 
**Type** | **int64** | Transaction type | 
**AuthorizationList** | [**[]SetCodeAuthorization**](SetCodeAuthorization.md) | EIP-7702 authorization list for SetCode transactions. | 

## Methods

### NewCloudWalletSetCodeTx

`func NewCloudWalletSetCodeTx(gas int64, from string, to string, value string, data string, type_ int64, authorizationList []SetCodeAuthorization, ) *CloudWalletSetCodeTx`

NewCloudWalletSetCodeTx instantiates a new CloudWalletSetCodeTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudWalletSetCodeTxWithDefaults

`func NewCloudWalletSetCodeTxWithDefaults() *CloudWalletSetCodeTx`

NewCloudWalletSetCodeTxWithDefaults instantiates a new CloudWalletSetCodeTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *CloudWalletSetCodeTx) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CloudWalletSetCodeTx) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CloudWalletSetCodeTx) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *CloudWalletSetCodeTx) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetGasPrice

`func (o *CloudWalletSetCodeTx) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *CloudWalletSetCodeTx) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *CloudWalletSetCodeTx) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *CloudWalletSetCodeTx) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetGasFeeCap

`func (o *CloudWalletSetCodeTx) GetGasFeeCap() string`

GetGasFeeCap returns the GasFeeCap field if non-nil, zero value otherwise.

### GetGasFeeCapOk

`func (o *CloudWalletSetCodeTx) GetGasFeeCapOk() (*string, bool)`

GetGasFeeCapOk returns a tuple with the GasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeCap

`func (o *CloudWalletSetCodeTx) SetGasFeeCap(v string)`

SetGasFeeCap sets GasFeeCap field to given value.

### HasGasFeeCap

`func (o *CloudWalletSetCodeTx) HasGasFeeCap() bool`

HasGasFeeCap returns a boolean if a field has been set.

### GetGasTipCap

`func (o *CloudWalletSetCodeTx) GetGasTipCap() string`

GetGasTipCap returns the GasTipCap field if non-nil, zero value otherwise.

### GetGasTipCapOk

`func (o *CloudWalletSetCodeTx) GetGasTipCapOk() (*string, bool)`

GetGasTipCapOk returns a tuple with the GasTipCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTipCap

`func (o *CloudWalletSetCodeTx) SetGasTipCap(v string)`

SetGasTipCap sets GasTipCap field to given value.

### HasGasTipCap

`func (o *CloudWalletSetCodeTx) HasGasTipCap() bool`

HasGasTipCap returns a boolean if a field has been set.

### GetGas

`func (o *CloudWalletSetCodeTx) GetGas() int64`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *CloudWalletSetCodeTx) GetGasOk() (*int64, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *CloudWalletSetCodeTx) SetGas(v int64)`

SetGas sets Gas field to given value.


### GetFrom

`func (o *CloudWalletSetCodeTx) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CloudWalletSetCodeTx) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CloudWalletSetCodeTx) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *CloudWalletSetCodeTx) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CloudWalletSetCodeTx) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CloudWalletSetCodeTx) SetTo(v string)`

SetTo sets To field to given value.


### GetValue

`func (o *CloudWalletSetCodeTx) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CloudWalletSetCodeTx) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CloudWalletSetCodeTx) SetValue(v string)`

SetValue sets Value field to given value.


### GetData

`func (o *CloudWalletSetCodeTx) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudWalletSetCodeTx) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudWalletSetCodeTx) SetData(v string)`

SetData sets Data field to given value.


### GetHash

`func (o *CloudWalletSetCodeTx) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CloudWalletSetCodeTx) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CloudWalletSetCodeTx) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *CloudWalletSetCodeTx) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetType

`func (o *CloudWalletSetCodeTx) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudWalletSetCodeTx) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudWalletSetCodeTx) SetType(v int64)`

SetType sets Type field to given value.


### GetAuthorizationList

`func (o *CloudWalletSetCodeTx) GetAuthorizationList() []SetCodeAuthorization`

GetAuthorizationList returns the AuthorizationList field if non-nil, zero value otherwise.

### GetAuthorizationListOk

`func (o *CloudWalletSetCodeTx) GetAuthorizationListOk() (*[]SetCodeAuthorization, bool)`

GetAuthorizationListOk returns a tuple with the AuthorizationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationList

`func (o *CloudWalletSetCodeTx) SetAuthorizationList(v []SetCodeAuthorization)`

SetAuthorizationList sets AuthorizationList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


