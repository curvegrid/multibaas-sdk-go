# LinkAddressContractRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | An alias to easily identify and reference the entity in subsequent requests. | 
**Version** | Pointer to **string** | The contract version. | [optional] 
**StartingBlock** | Pointer to **string** | The block number from which to start syncing events. The value can be &#x60;latest&#x60; for the latest block number, an absolute block number (e.g. &#x60;123&#x60; for the block number 123), or a relative block number (e.g. &#x60;-100&#x60; for 100 blocks before the latest block). If absent, the event monitor will be disabled for this contract and events won&#39;t be synced. | [optional] 

## Methods

### NewLinkAddressContractRequest

`func NewLinkAddressContractRequest(label string, ) *LinkAddressContractRequest`

NewLinkAddressContractRequest instantiates a new LinkAddressContractRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkAddressContractRequestWithDefaults

`func NewLinkAddressContractRequestWithDefaults() *LinkAddressContractRequest`

NewLinkAddressContractRequestWithDefaults instantiates a new LinkAddressContractRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *LinkAddressContractRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LinkAddressContractRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LinkAddressContractRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetVersion

`func (o *LinkAddressContractRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LinkAddressContractRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LinkAddressContractRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LinkAddressContractRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStartingBlock

`func (o *LinkAddressContractRequest) GetStartingBlock() string`

GetStartingBlock returns the StartingBlock field if non-nil, zero value otherwise.

### GetStartingBlockOk

`func (o *LinkAddressContractRequest) GetStartingBlockOk() (*string, bool)`

GetStartingBlockOk returns a tuple with the StartingBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingBlock

`func (o *LinkAddressContractRequest) SetStartingBlock(v string)`

SetStartingBlock sets StartingBlock field to given value.

### HasStartingBlock

`func (o *LinkAddressContractRequest) HasStartingBlock() bool`

HasStartingBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


