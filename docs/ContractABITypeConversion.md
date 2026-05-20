# ContractABITypeConversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | The type conversion mode to apply: &#x60;decimals&#x60; (for converting integer wei values to fixed-point decimal strings, e.g., wei to ether), &#x60;timestamp&#x60; (converts Unix timestamp integers to RFC 3339 formatted time strings), &#x60;stringBytes&#x60; (converts byte arrays to UTF-8 strings), &#x60;base64Bytes&#x60; (converts byte arrays to base64-encoded strings). | 
**DecimalsAbsolute** | **NullableInt64** | The fixed number of decimal places to use for &#x60;decimals&#x60; mode conversion. Must be between 0-255. | 
**DecimalsFunction** | **NullableString** | The name of a contract function that dynamically determines decimal places for &#x60;decimals&#x60; mode conversion (e.g., &#x60;decimals()&#x60; for ERC-20 tokens). The function must take no parameters and return a uint value between 0-255. Mutually exclusive with &#x60;decimalsAbsolute&#x60;. | 

## Methods

### NewContractABITypeConversion

`func NewContractABITypeConversion(mode string, decimalsAbsolute NullableInt64, decimalsFunction NullableString, ) *ContractABITypeConversion`

NewContractABITypeConversion instantiates a new ContractABITypeConversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractABITypeConversionWithDefaults

`func NewContractABITypeConversionWithDefaults() *ContractABITypeConversion`

NewContractABITypeConversionWithDefaults instantiates a new ContractABITypeConversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ContractABITypeConversion) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ContractABITypeConversion) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ContractABITypeConversion) SetMode(v string)`

SetMode sets Mode field to given value.


### GetDecimalsAbsolute

`func (o *ContractABITypeConversion) GetDecimalsAbsolute() int64`

GetDecimalsAbsolute returns the DecimalsAbsolute field if non-nil, zero value otherwise.

### GetDecimalsAbsoluteOk

`func (o *ContractABITypeConversion) GetDecimalsAbsoluteOk() (*int64, bool)`

GetDecimalsAbsoluteOk returns a tuple with the DecimalsAbsolute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalsAbsolute

`func (o *ContractABITypeConversion) SetDecimalsAbsolute(v int64)`

SetDecimalsAbsolute sets DecimalsAbsolute field to given value.


### SetDecimalsAbsoluteNil

`func (o *ContractABITypeConversion) SetDecimalsAbsoluteNil(b bool)`

 SetDecimalsAbsoluteNil sets the value for DecimalsAbsolute to be an explicit nil

### UnsetDecimalsAbsolute
`func (o *ContractABITypeConversion) UnsetDecimalsAbsolute()`

UnsetDecimalsAbsolute ensures that no value is present for DecimalsAbsolute, not even an explicit nil
### GetDecimalsFunction

`func (o *ContractABITypeConversion) GetDecimalsFunction() string`

GetDecimalsFunction returns the DecimalsFunction field if non-nil, zero value otherwise.

### GetDecimalsFunctionOk

`func (o *ContractABITypeConversion) GetDecimalsFunctionOk() (*string, bool)`

GetDecimalsFunctionOk returns a tuple with the DecimalsFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalsFunction

`func (o *ContractABITypeConversion) SetDecimalsFunction(v string)`

SetDecimalsFunction sets DecimalsFunction field to given value.


### SetDecimalsFunctionNil

`func (o *ContractABITypeConversion) SetDecimalsFunctionNil(b bool)`

 SetDecimalsFunctionNil sets the value for DecimalsFunction to be an explicit nil

### UnsetDecimalsFunction
`func (o *ContractABITypeConversion) UnsetDecimalsFunction()`

UnsetDecimalsFunction ensures that no value is present for DecimalsFunction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


