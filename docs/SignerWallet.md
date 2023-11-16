# SignerWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the signer. | 
**Wallet** | **string** | An ethereum address. | 
**Signer** | **string** | An ethereum address. | 
**Label** | **string** | The label of the signer. | 

## Methods

### NewSignerWallet

`func NewSignerWallet(type_ string, wallet string, signer string, label string, ) *SignerWallet`

NewSignerWallet instantiates a new SignerWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignerWalletWithDefaults

`func NewSignerWalletWithDefaults() *SignerWallet`

NewSignerWalletWithDefaults instantiates a new SignerWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SignerWallet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignerWallet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignerWallet) SetType(v string)`

SetType sets Type field to given value.


### GetWallet

`func (o *SignerWallet) GetWallet() string`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *SignerWallet) GetWalletOk() (*string, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *SignerWallet) SetWallet(v string)`

SetWallet sets Wallet field to given value.


### GetSigner

`func (o *SignerWallet) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *SignerWallet) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *SignerWallet) SetSigner(v string)`

SetSigner sets Signer field to given value.


### GetLabel

`func (o *SignerWallet) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SignerWallet) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SignerWallet) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


