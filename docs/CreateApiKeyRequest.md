# CreateApiKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | A label. | 
**GroupIDs** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewCreateApiKeyRequest

`func NewCreateApiKeyRequest(label string, ) *CreateApiKeyRequest`

NewCreateApiKeyRequest instantiates a new CreateApiKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKeyRequestWithDefaults

`func NewCreateApiKeyRequestWithDefaults() *CreateApiKeyRequest`

NewCreateApiKeyRequestWithDefaults instantiates a new CreateApiKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateApiKeyRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateApiKeyRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateApiKeyRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetGroupIDs

`func (o *CreateApiKeyRequest) GetGroupIDs() []int64`

GetGroupIDs returns the GroupIDs field if non-nil, zero value otherwise.

### GetGroupIDsOk

`func (o *CreateApiKeyRequest) GetGroupIDsOk() (*[]int64, bool)`

GetGroupIDsOk returns a tuple with the GroupIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIDs

`func (o *CreateApiKeyRequest) SetGroupIDs(v []int64)`

SetGroupIDs sets GroupIDs field to given value.

### HasGroupIDs

`func (o *CreateApiKeyRequest) HasGroupIDs() bool`

HasGroupIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


