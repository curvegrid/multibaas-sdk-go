# HSMSignResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | **string** |  | 
**Signature** | **string** |  | 

## Methods

### NewHSMSignResponse

`func NewHSMSignResponse(publicKey string, signature string, ) *HSMSignResponse`

NewHSMSignResponse instantiates a new HSMSignResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHSMSignResponseWithDefaults

`func NewHSMSignResponseWithDefaults() *HSMSignResponse`

NewHSMSignResponseWithDefaults instantiates a new HSMSignResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *HSMSignResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *HSMSignResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *HSMSignResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetSignature

`func (o *HSMSignResponse) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *HSMSignResponse) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *HSMSignResponse) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


