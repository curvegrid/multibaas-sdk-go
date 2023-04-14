# TransactionReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TransactionReceiptData**](TransactionReceiptData.md) |  | 
**Events** | Pointer to [**[]EventInformation**](EventInformation.md) |  | [optional] 

## Methods

### NewTransactionReceipt

`func NewTransactionReceipt(data TransactionReceiptData, ) *TransactionReceipt`

NewTransactionReceipt instantiates a new TransactionReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionReceiptWithDefaults

`func NewTransactionReceiptWithDefaults() *TransactionReceipt`

NewTransactionReceiptWithDefaults instantiates a new TransactionReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransactionReceipt) GetData() TransactionReceiptData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionReceipt) GetDataOk() (*TransactionReceiptData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionReceipt) SetData(v TransactionReceiptData)`

SetData sets Data field to given value.


### GetEvents

`func (o *TransactionReceipt) GetEvents() []EventInformation`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TransactionReceipt) GetEventsOk() (*[]EventInformation, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TransactionReceipt) SetEvents(v []EventInformation)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *TransactionReceipt) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


