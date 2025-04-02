# DeployContractTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tx** | [**TransactionToSignTx**](TransactionToSignTx.md) |  | 
**Submitted** | **bool** |  | 
**DeployAt** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | An alias to easily identify and reference the entity in subsequent requests. | [optional] 

## Methods

### NewDeployContractTransaction

`func NewDeployContractTransaction(tx TransactionToSignTx, submitted bool, ) *DeployContractTransaction`

NewDeployContractTransaction instantiates a new DeployContractTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployContractTransactionWithDefaults

`func NewDeployContractTransactionWithDefaults() *DeployContractTransaction`

NewDeployContractTransactionWithDefaults instantiates a new DeployContractTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTx

`func (o *DeployContractTransaction) GetTx() TransactionToSignTx`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *DeployContractTransaction) GetTxOk() (*TransactionToSignTx, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *DeployContractTransaction) SetTx(v TransactionToSignTx)`

SetTx sets Tx field to given value.


### GetSubmitted

`func (o *DeployContractTransaction) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *DeployContractTransaction) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *DeployContractTransaction) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.


### GetDeployAt

`func (o *DeployContractTransaction) GetDeployAt() string`

GetDeployAt returns the DeployAt field if non-nil, zero value otherwise.

### GetDeployAtOk

`func (o *DeployContractTransaction) GetDeployAtOk() (*string, bool)`

GetDeployAtOk returns a tuple with the DeployAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployAt

`func (o *DeployContractTransaction) SetDeployAt(v string)`

SetDeployAt sets DeployAt field to given value.

### HasDeployAt

`func (o *DeployContractTransaction) HasDeployAt() bool`

HasDeployAt returns a boolean if a field has been set.

### GetLabel

`func (o *DeployContractTransaction) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DeployContractTransaction) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DeployContractTransaction) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DeployContractTransaction) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


