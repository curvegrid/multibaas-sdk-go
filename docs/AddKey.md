# AddKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientID** | **string** | The Application ID that will be accessing the Key Vault. | 
**KeyName** | **string** | The name of the key. | 
**KeyVersion** | **string** | The version of the key. | 
**VaultName** | **string** | The name given to the vault your key is stored in. | 

## Methods

### NewAddKey

`func NewAddKey(clientID string, keyName string, keyVersion string, vaultName string, ) *AddKey`

NewAddKey instantiates a new AddKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddKeyWithDefaults

`func NewAddKeyWithDefaults() *AddKey`

NewAddKeyWithDefaults instantiates a new AddKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientID

`func (o *AddKey) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *AddKey) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *AddKey) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetKeyName

`func (o *AddKey) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *AddKey) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *AddKey) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetKeyVersion

`func (o *AddKey) GetKeyVersion() string`

GetKeyVersion returns the KeyVersion field if non-nil, zero value otherwise.

### GetKeyVersionOk

`func (o *AddKey) GetKeyVersionOk() (*string, bool)`

GetKeyVersionOk returns a tuple with the KeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVersion

`func (o *AddKey) SetKeyVersion(v string)`

SetKeyVersion sets KeyVersion field to given value.


### GetVaultName

`func (o *AddKey) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *AddKey) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *AddKey) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


