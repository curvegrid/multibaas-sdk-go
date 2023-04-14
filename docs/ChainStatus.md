# ChainStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockNumber** | **int64** |  | 
**Version** | **string** | The client version for this chain node. | 
**ChainID** | **int64** |  | 
**NetworkID** | **int64** |  | 
**BaseFee** | Pointer to **string** | The current base fee (only available for chains that support EIP-1559). | [optional] 

## Methods

### NewChainStatus

`func NewChainStatus(blockNumber int64, version string, chainID int64, networkID int64, ) *ChainStatus`

NewChainStatus instantiates a new ChainStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainStatusWithDefaults

`func NewChainStatusWithDefaults() *ChainStatus`

NewChainStatusWithDefaults instantiates a new ChainStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockNumber

`func (o *ChainStatus) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *ChainStatus) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *ChainStatus) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.


### GetVersion

`func (o *ChainStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ChainStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ChainStatus) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetChainID

`func (o *ChainStatus) GetChainID() int64`

GetChainID returns the ChainID field if non-nil, zero value otherwise.

### GetChainIDOk

`func (o *ChainStatus) GetChainIDOk() (*int64, bool)`

GetChainIDOk returns a tuple with the ChainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainID

`func (o *ChainStatus) SetChainID(v int64)`

SetChainID sets ChainID field to given value.


### GetNetworkID

`func (o *ChainStatus) GetNetworkID() int64`

GetNetworkID returns the NetworkID field if non-nil, zero value otherwise.

### GetNetworkIDOk

`func (o *ChainStatus) GetNetworkIDOk() (*int64, bool)`

GetNetworkIDOk returns a tuple with the NetworkID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkID

`func (o *ChainStatus) SetNetworkID(v int64)`

SetNetworkID sets NetworkID field to given value.


### GetBaseFee

`func (o *ChainStatus) GetBaseFee() string`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *ChainStatus) GetBaseFeeOk() (*string, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *ChainStatus) SetBaseFee(v string)`

SetBaseFee sets BaseFee field to given value.

### HasBaseFee

`func (o *ChainStatus) HasBaseFee() bool`

HasBaseFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


