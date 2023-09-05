# CreateKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientID** | **string** | The Application ID that will be accessing the Key Vault. | 
**KeyName** | **string** | The name of the key. | 
**VaultName** | **string** | The name given to the vault your key is stored in. | 
**UseHardwareModule** | **bool** |  | 

## Methods

### NewCreateKey

`func NewCreateKey(clientID string, keyName string, vaultName string, useHardwareModule bool, ) *CreateKey`

NewCreateKey instantiates a new CreateKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeyWithDefaults

`func NewCreateKeyWithDefaults() *CreateKey`

NewCreateKeyWithDefaults instantiates a new CreateKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientID

`func (o *CreateKey) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *CreateKey) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *CreateKey) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetKeyName

`func (o *CreateKey) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *CreateKey) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *CreateKey) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetVaultName

`func (o *CreateKey) GetVaultName() string`

GetVaultName returns the VaultName field if non-nil, zero value otherwise.

### GetVaultNameOk

`func (o *CreateKey) GetVaultNameOk() (*string, bool)`

GetVaultNameOk returns a tuple with the VaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultName

`func (o *CreateKey) SetVaultName(v string)`

SetVaultName sets VaultName field to given value.


### GetUseHardwareModule

`func (o *CreateKey) GetUseHardwareModule() bool`

GetUseHardwareModule returns the UseHardwareModule field if non-nil, zero value otherwise.

### GetUseHardwareModuleOk

`func (o *CreateKey) GetUseHardwareModuleOk() (*bool, bool)`

GetUseHardwareModuleOk returns a tuple with the UseHardwareModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHardwareModule

`func (o *CreateKey) SetUseHardwareModule(v bool)`

SetUseHardwareModule sets UseHardwareModule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


