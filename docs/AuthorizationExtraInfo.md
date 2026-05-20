# AuthorizationExtraInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | **string** | An ethereum address. | 
**FormatValid** | **bool** | Whether the format of the authorization is valid. | 
**Notes** | **string** | Additional notes about the validity of the authorization. | 

## Methods

### NewAuthorizationExtraInfo

`func NewAuthorizationExtraInfo(authority string, formatValid bool, notes string, ) *AuthorizationExtraInfo`

NewAuthorizationExtraInfo instantiates a new AuthorizationExtraInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationExtraInfoWithDefaults

`func NewAuthorizationExtraInfoWithDefaults() *AuthorizationExtraInfo`

NewAuthorizationExtraInfoWithDefaults instantiates a new AuthorizationExtraInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *AuthorizationExtraInfo) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *AuthorizationExtraInfo) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *AuthorizationExtraInfo) SetAuthority(v string)`

SetAuthority sets Authority field to given value.


### GetFormatValid

`func (o *AuthorizationExtraInfo) GetFormatValid() bool`

GetFormatValid returns the FormatValid field if non-nil, zero value otherwise.

### GetFormatValidOk

`func (o *AuthorizationExtraInfo) GetFormatValidOk() (*bool, bool)`

GetFormatValidOk returns a tuple with the FormatValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatValid

`func (o *AuthorizationExtraInfo) SetFormatValid(v bool)`

SetFormatValid sets FormatValid field to given value.


### GetNotes

`func (o *AuthorizationExtraInfo) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AuthorizationExtraInfo) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AuthorizationExtraInfo) SetNotes(v string)`

SetNotes sets Notes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


