# ListContracts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**[]ContractOverview**](ContractOverview.md) |  | 

## Methods

### NewListContracts200Response

`func NewListContracts200Response(status int64, message string, result []ContractOverview, ) *ListContracts200Response`

NewListContracts200Response instantiates a new ListContracts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListContracts200ResponseWithDefaults

`func NewListContracts200ResponseWithDefaults() *ListContracts200Response`

NewListContracts200ResponseWithDefaults instantiates a new ListContracts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListContracts200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListContracts200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListContracts200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ListContracts200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListContracts200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListContracts200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *ListContracts200Response) GetResult() []ContractOverview`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListContracts200Response) GetResultOk() (*[]ContractOverview, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListContracts200Response) SetResult(v []ContractOverview)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


