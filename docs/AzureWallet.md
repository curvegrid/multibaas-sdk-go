# AzureWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyName** | **string** |  | 
**KeyVersion** | **string** | The current version of your key. | 
**PublicAddress** | **string** | An ethereum address. | 

## Methods

### NewAzureWallet

`func NewAzureWallet(keyName string, keyVersion string, publicAddress string, ) *AzureWallet`

NewAzureWallet instantiates a new AzureWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureWalletWithDefaults

`func NewAzureWalletWithDefaults() *AzureWallet`

NewAzureWalletWithDefaults instantiates a new AzureWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyName

`func (o *AzureWallet) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *AzureWallet) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *AzureWallet) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetKeyVersion

`func (o *AzureWallet) GetKeyVersion() string`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *AzureWallet) GetKeyVersionOk() (*string, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *AzureWallet) SetKeyVersion(v string)`

SetKeyVersion sets KeyVersion field to given value.


### GetPublicAddress

`func (o *AzureWallet) GetPublicAddress() string`

GetPublicAddress returns the PublicAddress field if non-nil, zero value otherwise.

### GetPublicAddressOk

`func (o *AzureWallet) GetPublicAddressOk() (*string, bool)`

GetPublicAddressOk returns a tuple with the PublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAddress

`func (o *AzureWallet) SetPublicAddress(v string)`

SetPublicAddress sets PublicAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


