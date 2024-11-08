# SignedTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tx** | [**Transaction**](Transaction.md) |  | 

## Methods

### NewSignedTransactionResponse

`func NewSignedTransactionResponse(tx Transaction, ) *SignedTransactionResponse`

NewSignedTransactionResponse instantiates a new SignedTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignedTransactionResponseWithDefaults

`func NewSignedTransactionResponseWithDefaults() *SignedTransactionResponse`

NewSignedTransactionResponseWithDefaults instantiates a new SignedTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTx

`func (o *SignedTransactionResponse) GetTx() Transaction`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *SignedTransactionResponse) GetTxOk() (*Transaction, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *SignedTransactionResponse) SetTx(v Transaction)`

SetTx sets Tx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


