# APIKeyWithSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | An alias to easily identify and reference the entity in subsequent requests. | 
**Id** | **int64** |  | 
**CreatedAt** | **time.Time** | The time the API key was created. | 
**LastUsedAt** | Pointer to **time.Time** | The time the API key was last used. | [optional] 
**CreatedBy** | **int64** | The ID of the user that created the API key. | 
**Signature** | **string** | The signature of the API key. | 
**Key** | **string** | The secret key of the API key. | 

## Methods

### NewAPIKeyWithSecret

`func NewAPIKeyWithSecret(label string, id int64, createdAt time.Time, createdBy int64, signature string, key string, ) *APIKeyWithSecret`

NewAPIKeyWithSecret instantiates a new APIKeyWithSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyWithSecretWithDefaults

`func NewAPIKeyWithSecretWithDefaults() *APIKeyWithSecret`

NewAPIKeyWithSecretWithDefaults instantiates a new APIKeyWithSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *APIKeyWithSecret) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *APIKeyWithSecret) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *APIKeyWithSecret) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetId

`func (o *APIKeyWithSecret) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIKeyWithSecret) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIKeyWithSecret) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *APIKeyWithSecret) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APIKeyWithSecret) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APIKeyWithSecret) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastUsedAt

`func (o *APIKeyWithSecret) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *APIKeyWithSecret) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *APIKeyWithSecret) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *APIKeyWithSecret) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *APIKeyWithSecret) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *APIKeyWithSecret) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *APIKeyWithSecret) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetSignature

`func (o *APIKeyWithSecret) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *APIKeyWithSecret) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *APIKeyWithSecret) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetKey

`func (o *APIKeyWithSecret) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *APIKeyWithSecret) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *APIKeyWithSecret) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


