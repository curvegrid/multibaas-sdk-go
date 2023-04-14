# AzureHardwareWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**AzureAccountID** | **int64** |  | 
**VaultName** | **string** | The name given to the vault your key is stored in. | 
**KeyName** | **string** |  | 
**KeyVersion** | **string** | The current version of your key. | 
**PublicAddress** | **string** | An ethereum address. | 

## Methods

### NewAzureHardwareWallet

`func NewAzureHardwareWallet(id int64, azureAccountID int64, vaultName string, keyName string, keyVersion string, publicAddress string, ) *AzureHardwareWallet`

NewAzureHardwareWallet instantiates a new AzureHardwareWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureHardwareWalletWithDefaults

`func NewAzureHardwareWalletWithDefaults() *AzureHardwareWallet`

NewAzureHardwareWalletWithDefaults instantiates a new AzureHardwareWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureHardwareWallet) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureHardwareWallet) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureHardwareWallet) SetId(v int64)`

SetId sets Id field to given value.


### GetAzureAccountID

`func (o *AzureHardwareWallet) GetAzureAccountID() int64`

GetAzureAccountID returns the AzureAccountID field if non-nil, zero value otherwise.

### GetAzureAccountIDOk

`func (o *AzureHardwareWallet) GetAzureAccountIDOk() (*int64, bool)`

GetAzureAccountIDOk returns a tuple with the AzureAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAccountID

`func (o *AzureHardwareWallet) SetAzureAccountID(v int64)`

SetAzureAccountID sets AzureAccountID field to given value.


### GetVaultName

`func (o *AzureHardwareWallet) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *AzureHardwareWallet) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *AzureHardwareWallet) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.


### GetKeyName

`func (o *AzureHardwareWallet) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *AzureHardwareWallet) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *AzureHardwareWallet) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetKeyVersion

`func (o *AzureHardwareWallet) GetKeyVersion() string`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *AzureHardwareWallet) GetKeyVersionOk() (*string, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *AzureHardwareWallet) SetKeyVersion(v string)`

SetKeyVersion sets KeyVersion field to given value.


### GetPublicAddress

`func (o *AzureHardwareWallet) GetPublicAddress() string`

GetPublicAddress returns the PublicAddress field if non-nil, zero value otherwise.

### GetPublicAddressOk

`func (o *AzureHardwareWallet) GetPublicAddressOk() (*string, bool)`

GetPublicAddressOk returns a tuple with the PublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAddress

`func (o *AzureHardwareWallet) SetPublicAddress(v string)`

SetPublicAddress sets PublicAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


