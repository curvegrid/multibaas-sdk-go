# Invite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The invitee&#39;s email address. | 
**GroupIDs** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewInvite

`func NewInvite(email string, ) *Invite`

NewInvite instantiates a new Invite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteWithDefaults

`func NewInviteWithDefaults() *Invite`

NewInviteWithDefaults instantiates a new Invite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Invite) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Invite) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Invite) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetGroupIDs

`func (o *Invite) GetGroupIDs() []int64`

GetGroupIDs returns the GroupIDs field if non-nil, zero value otherwise.

### GetGroupIDsOk

`func (o *Invite) GetGroupIDsOk() (*[]int64, bool)`

GetGroupIDsOk returns a tuple with the GroupIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIDs

`func (o *Invite) SetGroupIDs(v []int64)`

SetGroupIDs sets GroupIDs field to given value.

### HasGroupIDs

`func (o *Invite) HasGroupIDs() bool`

HasGroupIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


