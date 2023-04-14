# SetNonceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | Pointer to **int64** | If nonce is specified the provided value is set, otherwise the value is read from the blockchain. | [optional] 

## Methods

### NewSetNonceRequest

`func NewSetNonceRequest() *SetNonceRequest`

NewSetNonceRequest instantiates a new SetNonceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetNonceRequestWithDefaults

`func NewSetNonceRequestWithDefaults() *SetNonceRequest`

NewSetNonceRequestWithDefaults instantiates a new SetNonceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *SetNonceRequest) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *SetNonceRequest) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *SetNonceRequest) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *SetNonceRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


