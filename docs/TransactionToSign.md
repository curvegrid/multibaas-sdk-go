# TransactionToSign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tx** | **map[string]interface{}** |  | 
**Submitted** | **bool** |  | 

## Methods

### NewTransactionToSign

`func NewTransactionToSign(tx map[string]interface{}, submitted bool, ) *TransactionToSign`

NewTransactionToSign instantiates a new TransactionToSign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionToSignWithDefaults

`func NewTransactionToSignWithDefaults() *TransactionToSign`

NewTransactionToSignWithDefaults instantiates a new TransactionToSign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTx

`func (o *TransactionToSign) GetTx() map[string]interface{}`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *TransactionToSign) GetTxOk() (*map[string]interface{}, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *TransactionToSign) SetTx(v map[string]interface{})`

SetTx sets Tx field to given value.


### GetSubmitted

`func (o *TransactionToSign) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *TransactionToSign) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *TransactionToSign) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


