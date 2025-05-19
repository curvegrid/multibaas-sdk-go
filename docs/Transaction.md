# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A hex string. | 
**ChainId** | Pointer to **NullableString** | A hex string or null. | [optional] 
**Nonce** | **string** | A hex string. | 
**To** | **NullableString** | An ethereum address. | 
**From** | Pointer to **NullableString** | An ethereum address. | [optional] 
**Gas** | **string** | A hex string. | 
**GasPrice** | Pointer to **NullableString** | A hex string or null. | [optional] 
**MaxPriorityFeePerGas** | Pointer to **NullableString** | A hex string or null. | [optional] 
**MaxFeePerGas** | Pointer to **NullableString** | A hex string or null. | [optional] 
**MaxFeePerBlobGas** | Pointer to **NullableString** | A hex string or null. | [optional] 
**Value** | **NullableString** | A hex string or null. | 
**Input** | **string** | A hex string. | 
**AccessList** | Pointer to [**[]AccessTuple**](AccessTuple.md) |  | [optional] 
**BlobVersionedHashes** | Pointer to **[]string** |  | [optional] 
**AuthorizationList** | Pointer to [**[]SetCodeAuthorization**](SetCodeAuthorization.md) |  | [optional] 
**V** | **string** | A hex string. | 
**R** | **string** | A hex string. | 
**S** | **string** | A hex string. | 
**YParity** | Pointer to **NullableString** | A hex string or null. | [optional] 
**Blobs** | Pointer to **[]string** |  | [optional] 
**Commitments** | Pointer to **[]string** |  | [optional] 
**Proofs** | Pointer to **[]string** |  | [optional] 
**Hash** | **string** | The keccak256 hash as a hex string of 256 bits. | 

## Methods

### NewTransaction

`func NewTransaction(type_ string, nonce string, to NullableString, gas string, value NullableString, input string, v string, r string, s string, hash string, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.


### GetChainId

`func (o *Transaction) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Transaction) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Transaction) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *Transaction) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### SetChainIdNil

`func (o *Transaction) SetChainIdNil(b bool)`

 SetChainIdNil sets the value for ChainId to be an explicit nil

### UnsetChainId
`func (o *Transaction) UnsetChainId()`

UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
### GetNonce

`func (o *Transaction) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *Transaction) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *Transaction) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetTo

`func (o *Transaction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Transaction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Transaction) SetTo(v string)`

SetTo sets To field to given value.


### SetToNil

`func (o *Transaction) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *Transaction) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetFrom

`func (o *Transaction) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Transaction) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Transaction) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Transaction) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *Transaction) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *Transaction) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetGas

`func (o *Transaction) GetGas() string`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *Transaction) GetGasOk() (*string, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *Transaction) SetGas(v string)`

SetGas sets Gas field to given value.


### GetGasPrice

`func (o *Transaction) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *Transaction) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *Transaction) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *Transaction) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### SetGasPriceNil

`func (o *Transaction) SetGasPriceNil(b bool)`

 SetGasPriceNil sets the value for GasPrice to be an explicit nil

### UnsetGasPrice
`func (o *Transaction) UnsetGasPrice()`

UnsetGasPrice ensures that no value is present for GasPrice, not even an explicit nil
### GetMaxPriorityFeePerGas

`func (o *Transaction) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *Transaction) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *Transaction) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.

### HasMaxPriorityFeePerGas

`func (o *Transaction) HasMaxPriorityFeePerGas() bool`

HasMaxPriorityFeePerGas returns a boolean if a field has been set.

### SetMaxPriorityFeePerGasNil

`func (o *Transaction) SetMaxPriorityFeePerGasNil(b bool)`

 SetMaxPriorityFeePerGasNil sets the value for MaxPriorityFeePerGas to be an explicit nil

### UnsetMaxPriorityFeePerGas
`func (o *Transaction) UnsetMaxPriorityFeePerGas()`

UnsetMaxPriorityFeePerGas ensures that no value is present for MaxPriorityFeePerGas, not even an explicit nil
### GetMaxFeePerGas

`func (o *Transaction) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *Transaction) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *Transaction) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *Transaction) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### SetMaxFeePerGasNil

`func (o *Transaction) SetMaxFeePerGasNil(b bool)`

 SetMaxFeePerGasNil sets the value for MaxFeePerGas to be an explicit nil

### UnsetMaxFeePerGas
`func (o *Transaction) UnsetMaxFeePerGas()`

UnsetMaxFeePerGas ensures that no value is present for MaxFeePerGas, not even an explicit nil
### GetMaxFeePerBlobGas

`func (o *Transaction) GetMaxFeePerBlobGas() string`

GetMaxFeePerBlobGas returns the MaxFeePerBlobGas field if non-nil, zero value otherwise.

### GetMaxFeePerBlobGasOk

`func (o *Transaction) GetMaxFeePerBlobGasOk() (*string, bool)`

GetMaxFeePerBlobGasOk returns a tuple with the MaxFeePerBlobGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerBlobGas

`func (o *Transaction) SetMaxFeePerBlobGas(v string)`

SetMaxFeePerBlobGas sets MaxFeePerBlobGas field to given value.

### HasMaxFeePerBlobGas

`func (o *Transaction) HasMaxFeePerBlobGas() bool`

HasMaxFeePerBlobGas returns a boolean if a field has been set.

### SetMaxFeePerBlobGasNil

`func (o *Transaction) SetMaxFeePerBlobGasNil(b bool)`

 SetMaxFeePerBlobGasNil sets the value for MaxFeePerBlobGas to be an explicit nil

### UnsetMaxFeePerBlobGas
`func (o *Transaction) UnsetMaxFeePerBlobGas()`

UnsetMaxFeePerBlobGas ensures that no value is present for MaxFeePerBlobGas, not even an explicit nil
### GetValue

`func (o *Transaction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Transaction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Transaction) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *Transaction) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Transaction) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetInput

`func (o *Transaction) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Transaction) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Transaction) SetInput(v string)`

SetInput sets Input field to given value.


### GetAccessList

`func (o *Transaction) GetAccessList() []AccessTuple`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *Transaction) GetAccessListOk() (*[]AccessTuple, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *Transaction) SetAccessList(v []AccessTuple)`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *Transaction) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### SetAccessListNil

`func (o *Transaction) SetAccessListNil(b bool)`

 SetAccessListNil sets the value for AccessList to be an explicit nil

### UnsetAccessList
`func (o *Transaction) UnsetAccessList()`

UnsetAccessList ensures that no value is present for AccessList, not even an explicit nil
### GetBlobVersionedHashes

`func (o *Transaction) GetBlobVersionedHashes() []string`

GetBlobVersionedHashes returns the BlobVersionedHashes field if non-nil, zero value otherwise.

### GetBlobVersionedHashesOk

`func (o *Transaction) GetBlobVersionedHashesOk() (*[]string, bool)`

GetBlobVersionedHashesOk returns a tuple with the BlobVersionedHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobVersionedHashes

`func (o *Transaction) SetBlobVersionedHashes(v []string)`

SetBlobVersionedHashes sets BlobVersionedHashes field to given value.

### HasBlobVersionedHashes

`func (o *Transaction) HasBlobVersionedHashes() bool`

HasBlobVersionedHashes returns a boolean if a field has been set.

### SetBlobVersionedHashesNil

`func (o *Transaction) SetBlobVersionedHashesNil(b bool)`

 SetBlobVersionedHashesNil sets the value for BlobVersionedHashes to be an explicit nil

### UnsetBlobVersionedHashes
`func (o *Transaction) UnsetBlobVersionedHashes()`

UnsetBlobVersionedHashes ensures that no value is present for BlobVersionedHashes, not even an explicit nil
### GetAuthorizationList

`func (o *Transaction) GetAuthorizationList() []SetCodeAuthorization`

GetAuthorizationList returns the AuthorizationList field if non-nil, zero value otherwise.

### GetAuthorizationListOk

`func (o *Transaction) GetAuthorizationListOk() (*[]SetCodeAuthorization, bool)`

GetAuthorizationListOk returns a tuple with the AuthorizationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationList

`func (o *Transaction) SetAuthorizationList(v []SetCodeAuthorization)`

SetAuthorizationList sets AuthorizationList field to given value.

### HasAuthorizationList

`func (o *Transaction) HasAuthorizationList() bool`

HasAuthorizationList returns a boolean if a field has been set.

### SetAuthorizationListNil

`func (o *Transaction) SetAuthorizationListNil(b bool)`

 SetAuthorizationListNil sets the value for AuthorizationList to be an explicit nil

### UnsetAuthorizationList
`func (o *Transaction) UnsetAuthorizationList()`

UnsetAuthorizationList ensures that no value is present for AuthorizationList, not even an explicit nil
### GetV

`func (o *Transaction) GetV() string`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *Transaction) GetVOk() (*string, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *Transaction) SetV(v string)`

SetV sets V field to given value.


### GetR

`func (o *Transaction) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *Transaction) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *Transaction) SetR(v string)`

SetR sets R field to given value.


### GetS

`func (o *Transaction) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *Transaction) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *Transaction) SetS(v string)`

SetS sets S field to given value.


### GetYParity

`func (o *Transaction) GetYParity() string`

GetYParity returns the YParity field if non-nil, zero value otherwise.

### GetYParityOk

`func (o *Transaction) GetYParityOk() (*string, bool)`

GetYParityOk returns a tuple with the YParity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYParity

`func (o *Transaction) SetYParity(v string)`

SetYParity sets YParity field to given value.

### HasYParity

`func (o *Transaction) HasYParity() bool`

HasYParity returns a boolean if a field has been set.

### SetYParityNil

`func (o *Transaction) SetYParityNil(b bool)`

 SetYParityNil sets the value for YParity to be an explicit nil

### UnsetYParity
`func (o *Transaction) UnsetYParity()`

UnsetYParity ensures that no value is present for YParity, not even an explicit nil
### GetBlobs

`func (o *Transaction) GetBlobs() []string`

GetBlobs returns the Blobs field if non-nil, zero value otherwise.

### GetBlobsOk

`func (o *Transaction) GetBlobsOk() (*[]string, bool)`

GetBlobsOk returns a tuple with the Blobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobs

`func (o *Transaction) SetBlobs(v []string)`

SetBlobs sets Blobs field to given value.

### HasBlobs

`func (o *Transaction) HasBlobs() bool`

HasBlobs returns a boolean if a field has been set.

### SetBlobsNil

`func (o *Transaction) SetBlobsNil(b bool)`

 SetBlobsNil sets the value for Blobs to be an explicit nil

### UnsetBlobs
`func (o *Transaction) UnsetBlobs()`

UnsetBlobs ensures that no value is present for Blobs, not even an explicit nil
### GetCommitments

`func (o *Transaction) GetCommitments() []string`

GetCommitments returns the Commitments field if non-nil, zero value otherwise.

### GetCommitmentsOk

`func (o *Transaction) GetCommitmentsOk() (*[]string, bool)`

GetCommitmentsOk returns a tuple with the Commitments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitments

`func (o *Transaction) SetCommitments(v []string)`

SetCommitments sets Commitments field to given value.

### HasCommitments

`func (o *Transaction) HasCommitments() bool`

HasCommitments returns a boolean if a field has been set.

### SetCommitmentsNil

`func (o *Transaction) SetCommitmentsNil(b bool)`

 SetCommitmentsNil sets the value for Commitments to be an explicit nil

### UnsetCommitments
`func (o *Transaction) UnsetCommitments()`

UnsetCommitments ensures that no value is present for Commitments, not even an explicit nil
### GetProofs

`func (o *Transaction) GetProofs() []string`

GetProofs returns the Proofs field if non-nil, zero value otherwise.

### GetProofsOk

`func (o *Transaction) GetProofsOk() (*[]string, bool)`

GetProofsOk returns a tuple with the Proofs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofs

`func (o *Transaction) SetProofs(v []string)`

SetProofs sets Proofs field to given value.

### HasProofs

`func (o *Transaction) HasProofs() bool`

HasProofs returns a boolean if a field has been set.

### SetProofsNil

`func (o *Transaction) SetProofsNil(b bool)`

 SetProofsNil sets the value for Proofs to be an explicit nil

### UnsetProofs
`func (o *Transaction) UnsetProofs()`

UnsetProofs ensures that no value is present for Proofs, not even an explicit nil
### GetHash

`func (o *Transaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Transaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Transaction) SetHash(v string)`

SetHash sets Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


