# AzureAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | A label. | 
**ClientID** | **string** | The Application ID that will be accessing the Key Vault. | 
**ClientSecret** | **string** | The application’s secret key that you generate when you first register the application in Azure. | 
**TenantID** | **string** | Also known as Directory ID. | 
**SubscriptionID** | **string** | The ID linked to your subscription to Azure services. | 
**BaseGroupName** | **string** | The Resource Group Name for the resource being accessed. | 
**Id** | **int64** |  | 

## Methods

### NewAzureAccount

`func NewAzureAccount(label string, clientID string, clientSecret string, tenantID string, subscriptionID string, baseGroupName string, id int64, ) *AzureAccount`

NewAzureAccount instantiates a new AzureAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureAccountWithDefaults

`func NewAzureAccountWithDefaults() *AzureAccount`

NewAzureAccountWithDefaults instantiates a new AzureAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *AzureAccount) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AzureAccount) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AzureAccount) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetClientID

`func (o *AzureAccount) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *AzureAccount) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *AzureAccount) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetClientSecret

`func (o *AzureAccount) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AzureAccount) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AzureAccount) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetTenantID

`func (o *AzureAccount) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *AzureAccount) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *AzureAccount) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.


### GetSubscriptionID

`func (o *AzureAccount) GetSubscriptionID() string`

GetSubscriptionID returns the SubscriptionID field if non-nil, zero value otherwise.

### GetSubscriptionIDOk

`func (o *AzureAccount) GetSubscriptionIDOk() (*string, bool)`

GetSubscriptionIDOk returns a tuple with the SubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionID

`func (o *AzureAccount) SetSubscriptionID(v string)`

SetSubscriptionID sets SubscriptionID field to given value.


### GetBaseGroupName

`func (o *AzureAccount) GetBaseGroupName() string`

GetBaseGroupName returns the BaseGroupName field if non-nil, zero value otherwise.

### GetBaseGroupNameOk

`func (o *AzureAccount) GetBaseGroupNameOk() (*string, bool)`

GetBaseGroupNameOk returns a tuple with the BaseGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseGroupName

`func (o *AzureAccount) SetBaseGroupName(v string)`

SetBaseGroupName sets BaseGroupName field to given value.


### GetId

`func (o *AzureAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureAccount) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


