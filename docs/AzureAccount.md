# AzureAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**ClientID** | **string** | The Application ID accesses the Key Vault. | 
**ClientSecret** | **string** | The application’s secret key that you generate when you first register the application in Azure. | 
**TenantID** | **string** | Also known as Directory ID. | 
**SubscriptionID** | **string** | The ID linked to your subscription to Azure services. | 
**BaseGroupName** | **string** | The Resource Group Name for the resource being accessed. | 
**DeletedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAzureAccount

`func NewAzureAccount(id int64, clientID string, clientSecret string, tenantID string, subscriptionID string, baseGroupName string, ) *AzureAccount`

NewAzureAccount instantiates a new AzureAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureAccountWithDefaults

`func NewAzureAccountWithDefaults() *AzureAccount`

NewAzureAccountWithDefaults instantiates a new AzureAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetDeletedAt

`func (o *AzureAccount) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AzureAccount) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AzureAccount) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AzureAccount) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *AzureAccount) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *AzureAccount) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


