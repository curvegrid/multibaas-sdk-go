# AzureAccountConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientID** | **string** | The Application ID that will be accessing the Key Vault. | 
**ClientSecret** | **string** | The application’s secret key that you generate when you first register the application in Azure. | 
**TenantID** | **string** | Also known as Directory ID. | 
**SubscriptionID** | **string** | The ID linked to your subscription to Azure services. | 
**BaseGroupName** | **string** | The Resource Group Name for the resource being accessed. | 

## Methods

### NewAzureAccountConfig

`func NewAzureAccountConfig(clientID string, clientSecret string, tenantID string, subscriptionID string, baseGroupName string, ) *AzureAccountConfig`

NewAzureAccountConfig instantiates a new AzureAccountConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureAccountConfigWithDefaults

`func NewAzureAccountConfigWithDefaults() *AzureAccountConfig`

NewAzureAccountConfigWithDefaults instantiates a new AzureAccountConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientID

`func (o *AzureAccountConfig) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *AzureAccountConfig) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *AzureAccountConfig) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetClientSecret

`func (o *AzureAccountConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AzureAccountConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AzureAccountConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetTenantID

`func (o *AzureAccountConfig) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *AzureAccountConfig) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *AzureAccountConfig) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.


### GetSubscriptionID

`func (o *AzureAccountConfig) GetSubscriptionID() string`

GetSubscriptionID returns the SubscriptionID field if non-nil, zero value otherwise.

### GetSubscriptionIDOk

`func (o *AzureAccountConfig) GetSubscriptionIDOk() (*string, bool)`

GetSubscriptionIDOk returns a tuple with the SubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionID

`func (o *AzureAccountConfig) SetSubscriptionID(v string)`

SetSubscriptionID sets SubscriptionID field to given value.


### GetBaseGroupName

`func (o *AzureAccountConfig) GetBaseGroupName() string`

GetBaseGroupName returns the BaseGroupName field if non-nil, zero value otherwise.

### GetBaseGroupNameOk

`func (o *AzureAccountConfig) GetBaseGroupNameOk() (*string, bool)`

GetBaseGroupNameOk returns a tuple with the BaseGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseGroupName

`func (o *AzureAccountConfig) SetBaseGroupName(v string)`

SetBaseGroupName sets BaseGroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


