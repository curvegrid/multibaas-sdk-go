# HSMData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | [**AzureAccount**](AzureAccount.md) |  | 
**Wallets** | [**[]AzureHardwareWallet**](AzureHardwareWallet.md) | An array of Azure Hardware Wallets. | 

## Methods

### NewHSMData

`func NewHSMData(configuration AzureAccount, wallets []AzureHardwareWallet, ) *HSMData`

NewHSMData instantiates a new HSMData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHSMDataWithDefaults

`func NewHSMDataWithDefaults() *HSMData`

NewHSMDataWithDefaults instantiates a new HSMData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *HSMData) GetConfiguration() AzureAccount`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *HSMData) GetConfigurationOk() (*AzureAccount, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *HSMData) SetConfiguration(v AzureAccount)`

SetConfiguration sets Configuration field to given value.


### GetWallets

`func (o *HSMData) GetWallets() []AzureHardwareWallet`

GetWallets returns the Wallets field if non-nil, zero value otherwise.

### GetWalletsOk

`func (o *HSMData) GetWalletsOk() (*[]AzureHardwareWallet, bool)`

GetWalletsOk returns a tuple with the Wallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallets

`func (o *HSMData) SetWallets(v []AzureHardwareWallet)`

SetWallets sets Wallets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


