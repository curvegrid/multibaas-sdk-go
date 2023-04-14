# CountEventQueryRecords200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | **int64** | The record count. | 

## Methods

### NewCountEventQueryRecords200Response

`func NewCountEventQueryRecords200Response(status int64, message string, result int64, ) *CountEventQueryRecords200Response`

NewCountEventQueryRecords200Response instantiates a new CountEventQueryRecords200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountEventQueryRecords200ResponseWithDefaults

`func NewCountEventQueryRecords200ResponseWithDefaults() *CountEventQueryRecords200Response`

NewCountEventQueryRecords200ResponseWithDefaults instantiates a new CountEventQueryRecords200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CountEventQueryRecords200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CountEventQueryRecords200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CountEventQueryRecords200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *CountEventQueryRecords200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CountEventQueryRecords200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CountEventQueryRecords200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *CountEventQueryRecords200Response) GetResult() int64`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CountEventQueryRecords200Response) GetResultOk() (*int64, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CountEventQueryRecords200Response) SetResult(v int64)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


