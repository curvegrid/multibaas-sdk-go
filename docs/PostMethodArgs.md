# PostMethodArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | Pointer to **string** |  | [optional] 
**Args** | Pointer to **[]interface{}** | List of the function arguments. | [optional] 
**From** | Pointer to **string** | An ethereum address. | [optional] 
**Nonce** | Pointer to **int64** | Nonce to use for the transaction execution. | [optional] 
**GasPrice** | Pointer to **int64** | Gas price to use for the transaction execution. | [optional] 
**GasFeeCap** | Pointer to **int64** | Gas fee cap to use for the 1559 transaction execution. | [optional] 
**GasTipCap** | Pointer to **int64** | Gas priority fee cap to use for the 1559 transaction execution. | [optional] 
**Gas** | Pointer to **int64** | Gas limit to set for the transaction execution. | [optional] 
**To** | Pointer to **string** | An ethereum address. | [optional] 
**Value** | Pointer to **string** | Amount (in wei) to send with the transaction. | [optional] 
**SignAndSubmit** | Pointer to **bool** | If the &#x60;from&#x60; address is an HSM address and this flag is set to &#x60;true&#x60;, the transaction will be automatically signed and submitted to the blockchain. | [optional] [default to false]
**NonceManagement** | Pointer to **bool** | If the &#x60;from&#x60; address is an HSM address and this flag is set to &#x60;true&#x60;, MultiBaas will keep track of the nonce and set it accordingly. This is particularly useful when submitting multiple transactions concurrently or in a very short period of time. If this flag is set to &#x60;true&#x60; and a &#x60;nonce&#x60; is provided, it will reset the nonce tracker to the given nonce (useful if the nonce tracker is out of sync). | [optional] [default to false]
**PreEIP1559** | Pointer to **bool** | If set to &#x60;true&#x60;, forces a legacy type transaction. Otherwise an EIP-1559 transaction will created if the network supports it. | [optional] [default to false]
**Signer** | Pointer to **string** | An ethereum address. | [optional] 
**FormatInts** | Pointer to **string** | Mode to format integer outputs in the function call&#39;s responses. There are 3 possible modes:   - &#x60;auto&#x60; (the default option), where number format is decided by its type:     - If the type has size at most 32 bits, then the number is returned verbatim.     - If the type has size larger than 32 bits, then the number is returned as a string.   - &#x60;as-numbers&#x60;, where all numbers are returned as strings.   - &#x60;as-strings&#x60;, where all numbers are returned verbatim.  | [optional] [default to "auto"]
**Timestamp** | Pointer to **string** | Call the function at a specific timestamp. Only available for read functions calls and if the &#x60;historical_blocks_feature&#x60; is enabled (see the plan endpoint). Mutually exclusive with the &#x60;blockNumber&#x60; parameter. | [optional] 
**BlockNumber** | Pointer to **string** | Call the function at a specific block. Only available for read functions calls and if the &#x60;historical_blocks_feature&#x60; is enabled (see the plan endpoint). Mutually exclusive with the &#x60;timestamp&#x60; parameter. | [optional] 
**ContractOverride** | Pointer to **bool** | If set to true the given address and contract don&#39;t need to be linked for the function to be called. | [optional] 
**Preview** | Pointer to [**PreviewArgs**](PreviewArgs.md) |  | [optional] 

## Methods

### NewPostMethodArgs

`func NewPostMethodArgs() *PostMethodArgs`

NewPostMethodArgs instantiates a new PostMethodArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostMethodArgsWithDefaults

`func NewPostMethodArgsWithDefaults() *PostMethodArgs`

NewPostMethodArgsWithDefaults instantiates a new PostMethodArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *PostMethodArgs) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PostMethodArgs) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PostMethodArgs) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *PostMethodArgs) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetArgs

`func (o *PostMethodArgs) GetArgs() []interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *PostMethodArgs) GetArgsOk() (*[]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *PostMethodArgs) SetArgs(v []interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *PostMethodArgs) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetFrom

`func (o *PostMethodArgs) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PostMethodArgs) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PostMethodArgs) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PostMethodArgs) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetNonce

`func (o *PostMethodArgs) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *PostMethodArgs) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *PostMethodArgs) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *PostMethodArgs) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetGasPrice

`func (o *PostMethodArgs) GetGasPrice() int64`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *PostMethodArgs) GetGasPriceOk() (*int64, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *PostMethodArgs) SetGasPrice(v int64)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *PostMethodArgs) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetGasFeeCap

`func (o *PostMethodArgs) GetGasFeeCap() int64`

GetGasFeeCap returns the GasFeeCap field if non-nil, zero value otherwise.

### GetGasFeeCapOk

`func (o *PostMethodArgs) GetGasFeeCapOk() (*int64, bool)`

GetGasFeeCapOk returns a tuple with the GasFeeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeCap

`func (o *PostMethodArgs) SetGasFeeCap(v int64)`

SetGasFeeCap sets GasFeeCap field to given value.

### HasGasFeeCap

`func (o *PostMethodArgs) HasGasFeeCap() bool`

HasGasFeeCap returns a boolean if a field has been set.

### GetGasTipCap

`func (o *PostMethodArgs) GetGasTipCap() int64`

GetGasTipCap returns the GasTipCap field if non-nil, zero value otherwise.

### GetGasTipCapOk

`func (o *PostMethodArgs) GetGasTipCapOk() (*int64, bool)`

GetGasTipCapOk returns a tuple with the GasTipCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTipCap

`func (o *PostMethodArgs) SetGasTipCap(v int64)`

SetGasTipCap sets GasTipCap field to given value.

### HasGasTipCap

`func (o *PostMethodArgs) HasGasTipCap() bool`

HasGasTipCap returns a boolean if a field has been set.

### GetGas

`func (o *PostMethodArgs) GetGas() int64`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *PostMethodArgs) GetGasOk() (*int64, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *PostMethodArgs) SetGas(v int64)`

SetGas sets Gas field to given value.

### HasGas

`func (o *PostMethodArgs) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetTo

`func (o *PostMethodArgs) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PostMethodArgs) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PostMethodArgs) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *PostMethodArgs) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetValue

`func (o *PostMethodArgs) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PostMethodArgs) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PostMethodArgs) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PostMethodArgs) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSignAndSubmit

`func (o *PostMethodArgs) GetSignAndSubmit() bool`

GetSignAndSubmit returns the SignAndSubmit field if non-nil, zero value otherwise.

### GetSignAndSubmitOk

`func (o *PostMethodArgs) GetSignAndSubmitOk() (*bool, bool)`

GetSignAndSubmitOk returns a tuple with the SignAndSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAndSubmit

`func (o *PostMethodArgs) SetSignAndSubmit(v bool)`

SetSignAndSubmit sets SignAndSubmit field to given value.

### HasSignAndSubmit

`func (o *PostMethodArgs) HasSignAndSubmit() bool`

HasSignAndSubmit returns a boolean if a field has been set.

### GetNonceManagement

`func (o *PostMethodArgs) GetNonceManagement() bool`

GetNonceManagement returns the NonceManagement field if non-nil, zero value otherwise.

### GetNonceManagementOk

`func (o *PostMethodArgs) GetNonceManagementOk() (*bool, bool)`

GetNonceManagementOk returns a tuple with the NonceManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonceManagement

`func (o *PostMethodArgs) SetNonceManagement(v bool)`

SetNonceManagement sets NonceManagement field to given value.

### HasNonceManagement

`func (o *PostMethodArgs) HasNonceManagement() bool`

HasNonceManagement returns a boolean if a field has been set.

### GetPreEIP1559

`func (o *PostMethodArgs) GetPreEIP1559() bool`

GetPreEIP1559 returns the PreEIP1559 field if non-nil, zero value otherwise.

### GetPreEIP1559Ok

`func (o *PostMethodArgs) GetPreEIP1559Ok() (*bool, bool)`

GetPreEIP1559Ok returns a tuple with the PreEIP1559 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreEIP1559

`func (o *PostMethodArgs) SetPreEIP1559(v bool)`

SetPreEIP1559 sets PreEIP1559 field to given value.

### HasPreEIP1559

`func (o *PostMethodArgs) HasPreEIP1559() bool`

HasPreEIP1559 returns a boolean if a field has been set.

### GetSigner

`func (o *PostMethodArgs) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *PostMethodArgs) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *PostMethodArgs) SetSigner(v string)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *PostMethodArgs) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetFormatInts

`func (o *PostMethodArgs) GetFormatInts() string`

GetFormatInts returns the FormatInts field if non-nil, zero value otherwise.

### GetFormatIntsOk

`func (o *PostMethodArgs) GetFormatIntsOk() (*string, bool)`

GetFormatIntsOk returns a tuple with the FormatInts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatInts

`func (o *PostMethodArgs) SetFormatInts(v string)`

SetFormatInts sets FormatInts field to given value.

### HasFormatInts

`func (o *PostMethodArgs) HasFormatInts() bool`

HasFormatInts returns a boolean if a field has been set.

### GetTimestamp

`func (o *PostMethodArgs) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PostMethodArgs) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PostMethodArgs) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PostMethodArgs) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBlockNumber

`func (o *PostMethodArgs) GetBlockNumber() string`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *PostMethodArgs) GetBlockNumberOk() (*string, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *PostMethodArgs) SetBlockNumber(v string)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *PostMethodArgs) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetContractOverride

`func (o *PostMethodArgs) GetContractOverride() bool`

GetContractOverride returns the ContractOverride field if non-nil, zero value otherwise.

### GetContractOverrideOk

`func (o *PostMethodArgs) GetContractOverrideOk() (*bool, bool)`

GetContractOverrideOk returns a tuple with the ContractOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractOverride

`func (o *PostMethodArgs) SetContractOverride(v bool)`

SetContractOverride sets ContractOverride field to given value.

### HasContractOverride

`func (o *PostMethodArgs) HasContractOverride() bool`

HasContractOverride returns a boolean if a field has been set.

### GetPreview

`func (o *PostMethodArgs) GetPreview() PreviewArgs`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *PostMethodArgs) GetPreviewOk() (*PreviewArgs, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *PostMethodArgs) SetPreview(v PreviewArgs)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *PostMethodArgs) HasPreview() bool`

HasPreview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


