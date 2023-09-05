# ListCorsOrigins200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | The status code. | 
**Message** | **string** | The human-readable status message. | 
**Result** | [**[]CORSOrigin**](CORSOrigin.md) |  | 

## Methods

### NewListCorsOrigins200Response

`func NewListCorsOrigins200Response(status int64, message string, result []CORSOrigin, ) *ListCorsOrigins200Response`

NewListCorsOrigins200Response instantiates a new ListCorsOrigins200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCorsOrigins200ResponseWithDefaults

`func NewListCorsOrigins200ResponseWithDefaults() *ListCorsOrigins200Response`

NewListCorsOrigins200ResponseWithDefaults instantiates a new ListCorsOrigins200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListCorsOrigins200Response) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListCorsOrigins200Response) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListCorsOrigins200Response) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ListCorsOrigins200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListCorsOrigins200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListCorsOrigins200Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *ListCorsOrigins200Response) GetResult() []CORSOrigin`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListCorsOrigins200Response) GetResultOk() (*[]CORSOrigin, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListCorsOrigins200Response) SetResult(v []CORSOrigin)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


