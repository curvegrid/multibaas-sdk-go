# SignedTransactionSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignedTx** | **string** | The pre-signed raw transaction. | 

## Methods

### NewSignedTransactionSubmission

`func NewSignedTransactionSubmission(signedTx string, ) *SignedTransactionSubmission`

NewSignedTransactionSubmission instantiates a new SignedTransactionSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignedTransactionSubmissionWithDefaults

`func NewSignedTransactionSubmissionWithDefaults() *SignedTransactionSubmission`

NewSignedTransactionSubmissionWithDefaults instantiates a new SignedTransactionSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignedTx

`func (o *SignedTransactionSubmission) GetSignedTx() string`

GetSignedTx returns the SignedTx field if non-nil, zero value otherwise.

### GetSignedTxOk

`func (o *SignedTransactionSubmission) GetSignedTxOk() (*string, bool)`

GetSignedTxOk returns a tuple with the SignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTx

`func (o *SignedTransactionSubmission) SetSignedTx(v string)`

SetSignedTx sets SignedTx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


