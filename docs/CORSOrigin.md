# CORSOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Origin** | Pointer to **string** | The CORS Origin | [optional] 

## Methods

### NewCORSOrigin

`func NewCORSOrigin() *CORSOrigin`

NewCORSOrigin instantiates a new CORSOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCORSOriginWithDefaults

`func NewCORSOriginWithDefaults() *CORSOrigin`

NewCORSOriginWithDefaults instantiates a new CORSOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CORSOrigin) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CORSOrigin) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CORSOrigin) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CORSOrigin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrigin

`func (o *CORSOrigin) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *CORSOrigin) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *CORSOrigin) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *CORSOrigin) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


