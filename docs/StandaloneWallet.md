# StandaloneWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientID** | Pointer to **string** | The Application ID that accesses the Key Vault. | [optional] 
**BaseGroupName** | **string** | The Resource Group Name for the resource being accessed. | 
**VaultName** | Pointer to **string** | The name given to the vault your key is stored in. | [optional] 
**KeyName** | **string** | The name of the key. | 
**KeyVersion** | Pointer to **string** | The version of the key. | [optional] 
**PublicAddress** | **string** | An ethereum address. | 

## Methods

### NewStandaloneWallet

`func NewStandaloneWallet(baseGroupName string, keyName string, publicAddress string, ) *StandaloneWallet`

NewStandaloneWallet instantiates a new StandaloneWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandaloneWalletWithDefaults

`func NewStandaloneWalletWithDefaults() *StandaloneWallet`

NewStandaloneWalletWithDefaults instantiates a new StandaloneWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientID

`func (o *StandaloneWallet) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *StandaloneWallet) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *StandaloneWallet) SetClientID(v string)`

SetClientID sets ClientID field to given value.

### HasClientID

`func (o *StandaloneWallet) HasClientID() bool`

HasClientID returns a boolean if a field has been set.

### GetBaseGroupName

`func (o *StandaloneWallet) GetBaseGroupName() string`

GetBaseGroupName returns the BaseGroupName field if non-nil, zero value otherwise.

### GetBaseGroupNameOk

`func (o *StandaloneWallet) GetBaseGroupNameOk() (*string, bool)`

GetBaseGroupNameOk returns a tuple with the BaseGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseGroupName

`func (o *StandaloneWallet) SetBaseGroupName(v string)`

SetBaseGroupName sets BaseGroupName field to given value.


### GetVaultName

`func (o *StandaloneWallet) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *StandaloneWallet) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *StandaloneWallet) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.

### HasVaultName

`func (o *StandaloneWallet) HasVaultName() bool`

HasVaultName returns a boolean if a field has been set.

### GetKeyName

`func (o *StandaloneWallet) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *StandaloneWallet) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *StandaloneWallet) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetKeyVersion

`func (o *StandaloneWallet) GetKeyVersion() string`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *StandaloneWallet) GetKeyVersionOk() (*string, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *StandaloneWallet) SetKeyVersion(v string)`

SetKeyVersion sets KeyVersion field to given value.

### HasKeyVersion

`func (o *StandaloneWallet) HasKeyVersion() bool`

HasKeyVersion returns a boolean if a field has been set.

### GetPublicAddress

`func (o *StandaloneWallet) GetPublicAddress() string`

GetPublicAddress returns the PublicAddress field if non-nil, zero value otherwise.

### GetPublicAddressOk

`func (o *StandaloneWallet) GetPublicAddressOk() (*string, bool)`

GetPublicAddressOk returns a tuple with the PublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAddress

`func (o *StandaloneWallet) SetPublicAddress(v string)`

SetPublicAddress sets PublicAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


