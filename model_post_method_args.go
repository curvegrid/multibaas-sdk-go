/*
MultiBaas API

MultiBaas's REST APIv0.

API version: 0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package multibaas

import (
	"encoding/json"
)

// checks if the PostMethodArgs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostMethodArgs{}

// PostMethodArgs Arguments to be passed into a contract function.
type PostMethodArgs struct {
	Signature *string `json:"signature,omitempty"`
	// List of the function arguments.
	Args []interface{} `json:"args,omitempty"`
	// An ethereum address.
	From *string `json:"from,omitempty" validate:"regexp=^0[xX][a-fA-F0-9]{40}$"`
	// Nonce to use for the transaction execution.
	Nonce *int64 `json:"nonce,omitempty"`
	// Gas price to use for the transaction execution.
	GasPrice *int64 `json:"gasPrice,omitempty"`
	// Gas fee cap to use for the 1559 transaction execution.
	GasFeeCap *int64 `json:"gasFeeCap,omitempty"`
	// Gas priority fee cap to use for the 1559 transaction execution.
	GasTipCap *int64 `json:"gasTipCap,omitempty"`
	// Gas limit to set for the transaction execution.
	Gas *int64 `json:"gas,omitempty"`
	// An ethereum address.
	To *string `json:"to,omitempty" validate:"regexp=^0[xX][a-fA-F0-9]{40}$"`
	// Amount (in wei) to send with the transaction.
	Value *string `json:"value,omitempty"`
	// If the `from` address is an HSM address and this flag is set to `true`, the transaction will be automatically signed and submitted to the blockchain.
	SignAndSubmit *bool `json:"signAndSubmit,omitempty"`
	// If the `from` address is an HSM address and this flag is set to `true`, MultiBaas will keep track of the nonce and set it accordingly. This is particularly useful when submitting multiple transactions concurrently or in a very short period of time. If this flag is set to `true` and a `nonce` is provided, it will reset the nonce tracker to the given nonce (useful if the nonce tracker is out of sync).
	NonceManagement *bool `json:"nonceManagement,omitempty"`
	// If set to `true`, forces a legacy type transaction. Otherwise an EIP-1559 transaction will created if the network supports it.
	PreEIP1559 *bool `json:"preEIP1559,omitempty"`
	// An ethereum address.
	Signer *string `json:"signer,omitempty" validate:"regexp=^0[xX][a-fA-F0-9]{40}$"`
	// Mode to format integer outputs in the function call's responses. There are 3 possible modes:   - `auto` (the default option), where number format is decided by its type:     - If the type has size at most 32 bits, then the number is returned verbatim.     - If the type has size larger than 32 bits, then the number is returned as a string.   - `as_numbers`, where all numbers are returned verbatim.   - `as_strings`, where all numbers are returned as strings.
	FormatInts *string `json:"formatInts,omitempty"`
	// Call the function at a specific timestamp. Only available for read functions calls and if the `historical_blocks_feature` is enabled (see the plan endpoint). Mutually exclusive with the `blockNumber` parameter.
	Timestamp *string `json:"timestamp,omitempty"`
	// Call the function at a specific block. Only available for read functions calls and if the `historical_blocks_feature` is enabled (see the plan endpoint). Mutually exclusive with the `timestamp` parameter.
	BlockNumber *string `json:"blockNumber,omitempty"`
	// If set to true the given address and contract don't need to be linked for the function to be called.
	ContractOverride *bool        `json:"contractOverride,omitempty"`
	Preview          *PreviewArgs `json:"preview,omitempty"`
}

// NewPostMethodArgs instantiates a new PostMethodArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostMethodArgs() *PostMethodArgs {
	this := PostMethodArgs{}
	var signAndSubmit bool = false
	this.SignAndSubmit = &signAndSubmit
	var nonceManagement bool = false
	this.NonceManagement = &nonceManagement
	var preEIP1559 bool = false
	this.PreEIP1559 = &preEIP1559
	var formatInts string = "auto"
	this.FormatInts = &formatInts
	return &this
}

// NewPostMethodArgsWithDefaults instantiates a new PostMethodArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostMethodArgsWithDefaults() *PostMethodArgs {
	this := PostMethodArgs{}
	var signAndSubmit bool = false
	this.SignAndSubmit = &signAndSubmit
	var nonceManagement bool = false
	this.NonceManagement = &nonceManagement
	var preEIP1559 bool = false
	this.PreEIP1559 = &preEIP1559
	var formatInts string = "auto"
	this.FormatInts = &formatInts
	return &this
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *PostMethodArgs) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *PostMethodArgs) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *PostMethodArgs) SetSignature(v string) {
	o.Signature = &v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *PostMethodArgs) GetArgs() []interface{} {
	if o == nil || IsNil(o.Args) {
		var ret []interface{}
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetArgsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Args) {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *PostMethodArgs) HasArgs() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []interface{} and assigns it to the Args field.
func (o *PostMethodArgs) SetArgs(v []interface{}) {
	o.Args = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *PostMethodArgs) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *PostMethodArgs) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *PostMethodArgs) SetFrom(v string) {
	o.From = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *PostMethodArgs) GetNonce() int64 {
	if o == nil || IsNil(o.Nonce) {
		var ret int64
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetNonceOk() (*int64, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *PostMethodArgs) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given int64 and assigns it to the Nonce field.
func (o *PostMethodArgs) SetNonce(v int64) {
	o.Nonce = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *PostMethodArgs) GetGasPrice() int64 {
	if o == nil || IsNil(o.GasPrice) {
		var ret int64
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetGasPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.GasPrice) {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *PostMethodArgs) HasGasPrice() bool {
	if o != nil && !IsNil(o.GasPrice) {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given int64 and assigns it to the GasPrice field.
func (o *PostMethodArgs) SetGasPrice(v int64) {
	o.GasPrice = &v
}

// GetGasFeeCap returns the GasFeeCap field value if set, zero value otherwise.
func (o *PostMethodArgs) GetGasFeeCap() int64 {
	if o == nil || IsNil(o.GasFeeCap) {
		var ret int64
		return ret
	}
	return *o.GasFeeCap
}

// GetGasFeeCapOk returns a tuple with the GasFeeCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetGasFeeCapOk() (*int64, bool) {
	if o == nil || IsNil(o.GasFeeCap) {
		return nil, false
	}
	return o.GasFeeCap, true
}

// HasGasFeeCap returns a boolean if a field has been set.
func (o *PostMethodArgs) HasGasFeeCap() bool {
	if o != nil && !IsNil(o.GasFeeCap) {
		return true
	}

	return false
}

// SetGasFeeCap gets a reference to the given int64 and assigns it to the GasFeeCap field.
func (o *PostMethodArgs) SetGasFeeCap(v int64) {
	o.GasFeeCap = &v
}

// GetGasTipCap returns the GasTipCap field value if set, zero value otherwise.
func (o *PostMethodArgs) GetGasTipCap() int64 {
	if o == nil || IsNil(o.GasTipCap) {
		var ret int64
		return ret
	}
	return *o.GasTipCap
}

// GetGasTipCapOk returns a tuple with the GasTipCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetGasTipCapOk() (*int64, bool) {
	if o == nil || IsNil(o.GasTipCap) {
		return nil, false
	}
	return o.GasTipCap, true
}

// HasGasTipCap returns a boolean if a field has been set.
func (o *PostMethodArgs) HasGasTipCap() bool {
	if o != nil && !IsNil(o.GasTipCap) {
		return true
	}

	return false
}

// SetGasTipCap gets a reference to the given int64 and assigns it to the GasTipCap field.
func (o *PostMethodArgs) SetGasTipCap(v int64) {
	o.GasTipCap = &v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *PostMethodArgs) GetGas() int64 {
	if o == nil || IsNil(o.Gas) {
		var ret int64
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetGasOk() (*int64, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *PostMethodArgs) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given int64 and assigns it to the Gas field.
func (o *PostMethodArgs) SetGas(v int64) {
	o.Gas = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *PostMethodArgs) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *PostMethodArgs) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *PostMethodArgs) SetTo(v string) {
	o.To = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PostMethodArgs) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PostMethodArgs) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PostMethodArgs) SetValue(v string) {
	o.Value = &v
}

// GetSignAndSubmit returns the SignAndSubmit field value if set, zero value otherwise.
func (o *PostMethodArgs) GetSignAndSubmit() bool {
	if o == nil || IsNil(o.SignAndSubmit) {
		var ret bool
		return ret
	}
	return *o.SignAndSubmit
}

// GetSignAndSubmitOk returns a tuple with the SignAndSubmit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetSignAndSubmitOk() (*bool, bool) {
	if o == nil || IsNil(o.SignAndSubmit) {
		return nil, false
	}
	return o.SignAndSubmit, true
}

// HasSignAndSubmit returns a boolean if a field has been set.
func (o *PostMethodArgs) HasSignAndSubmit() bool {
	if o != nil && !IsNil(o.SignAndSubmit) {
		return true
	}

	return false
}

// SetSignAndSubmit gets a reference to the given bool and assigns it to the SignAndSubmit field.
func (o *PostMethodArgs) SetSignAndSubmit(v bool) {
	o.SignAndSubmit = &v
}

// GetNonceManagement returns the NonceManagement field value if set, zero value otherwise.
func (o *PostMethodArgs) GetNonceManagement() bool {
	if o == nil || IsNil(o.NonceManagement) {
		var ret bool
		return ret
	}
	return *o.NonceManagement
}

// GetNonceManagementOk returns a tuple with the NonceManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetNonceManagementOk() (*bool, bool) {
	if o == nil || IsNil(o.NonceManagement) {
		return nil, false
	}
	return o.NonceManagement, true
}

// HasNonceManagement returns a boolean if a field has been set.
func (o *PostMethodArgs) HasNonceManagement() bool {
	if o != nil && !IsNil(o.NonceManagement) {
		return true
	}

	return false
}

// SetNonceManagement gets a reference to the given bool and assigns it to the NonceManagement field.
func (o *PostMethodArgs) SetNonceManagement(v bool) {
	o.NonceManagement = &v
}

// GetPreEIP1559 returns the PreEIP1559 field value if set, zero value otherwise.
func (o *PostMethodArgs) GetPreEIP1559() bool {
	if o == nil || IsNil(o.PreEIP1559) {
		var ret bool
		return ret
	}
	return *o.PreEIP1559
}

// GetPreEIP1559Ok returns a tuple with the PreEIP1559 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetPreEIP1559Ok() (*bool, bool) {
	if o == nil || IsNil(o.PreEIP1559) {
		return nil, false
	}
	return o.PreEIP1559, true
}

// HasPreEIP1559 returns a boolean if a field has been set.
func (o *PostMethodArgs) HasPreEIP1559() bool {
	if o != nil && !IsNil(o.PreEIP1559) {
		return true
	}

	return false
}

// SetPreEIP1559 gets a reference to the given bool and assigns it to the PreEIP1559 field.
func (o *PostMethodArgs) SetPreEIP1559(v bool) {
	o.PreEIP1559 = &v
}

// GetSigner returns the Signer field value if set, zero value otherwise.
func (o *PostMethodArgs) GetSigner() string {
	if o == nil || IsNil(o.Signer) {
		var ret string
		return ret
	}
	return *o.Signer
}

// GetSignerOk returns a tuple with the Signer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetSignerOk() (*string, bool) {
	if o == nil || IsNil(o.Signer) {
		return nil, false
	}
	return o.Signer, true
}

// HasSigner returns a boolean if a field has been set.
func (o *PostMethodArgs) HasSigner() bool {
	if o != nil && !IsNil(o.Signer) {
		return true
	}

	return false
}

// SetSigner gets a reference to the given string and assigns it to the Signer field.
func (o *PostMethodArgs) SetSigner(v string) {
	o.Signer = &v
}

// GetFormatInts returns the FormatInts field value if set, zero value otherwise.
func (o *PostMethodArgs) GetFormatInts() string {
	if o == nil || IsNil(o.FormatInts) {
		var ret string
		return ret
	}
	return *o.FormatInts
}

// GetFormatIntsOk returns a tuple with the FormatInts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetFormatIntsOk() (*string, bool) {
	if o == nil || IsNil(o.FormatInts) {
		return nil, false
	}
	return o.FormatInts, true
}

// HasFormatInts returns a boolean if a field has been set.
func (o *PostMethodArgs) HasFormatInts() bool {
	if o != nil && !IsNil(o.FormatInts) {
		return true
	}

	return false
}

// SetFormatInts gets a reference to the given string and assigns it to the FormatInts field.
func (o *PostMethodArgs) SetFormatInts(v string) {
	o.FormatInts = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *PostMethodArgs) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *PostMethodArgs) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *PostMethodArgs) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetBlockNumber returns the BlockNumber field value if set, zero value otherwise.
func (o *PostMethodArgs) GetBlockNumber() string {
	if o == nil || IsNil(o.BlockNumber) {
		var ret string
		return ret
	}
	return *o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetBlockNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BlockNumber) {
		return nil, false
	}
	return o.BlockNumber, true
}

// HasBlockNumber returns a boolean if a field has been set.
func (o *PostMethodArgs) HasBlockNumber() bool {
	if o != nil && !IsNil(o.BlockNumber) {
		return true
	}

	return false
}

// SetBlockNumber gets a reference to the given string and assigns it to the BlockNumber field.
func (o *PostMethodArgs) SetBlockNumber(v string) {
	o.BlockNumber = &v
}

// GetContractOverride returns the ContractOverride field value if set, zero value otherwise.
func (o *PostMethodArgs) GetContractOverride() bool {
	if o == nil || IsNil(o.ContractOverride) {
		var ret bool
		return ret
	}
	return *o.ContractOverride
}

// GetContractOverrideOk returns a tuple with the ContractOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetContractOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.ContractOverride) {
		return nil, false
	}
	return o.ContractOverride, true
}

// HasContractOverride returns a boolean if a field has been set.
func (o *PostMethodArgs) HasContractOverride() bool {
	if o != nil && !IsNil(o.ContractOverride) {
		return true
	}

	return false
}

// SetContractOverride gets a reference to the given bool and assigns it to the ContractOverride field.
func (o *PostMethodArgs) SetContractOverride(v bool) {
	o.ContractOverride = &v
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *PostMethodArgs) GetPreview() PreviewArgs {
	if o == nil || IsNil(o.Preview) {
		var ret PreviewArgs
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostMethodArgs) GetPreviewOk() (*PreviewArgs, bool) {
	if o == nil || IsNil(o.Preview) {
		return nil, false
	}
	return o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *PostMethodArgs) HasPreview() bool {
	if o != nil && !IsNil(o.Preview) {
		return true
	}

	return false
}

// SetPreview gets a reference to the given PreviewArgs and assigns it to the Preview field.
func (o *PostMethodArgs) SetPreview(v PreviewArgs) {
	o.Preview = &v
}

func (o PostMethodArgs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostMethodArgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !IsNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.GasPrice) {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if !IsNil(o.GasFeeCap) {
		toSerialize["gasFeeCap"] = o.GasFeeCap
	}
	if !IsNil(o.GasTipCap) {
		toSerialize["gasTipCap"] = o.GasTipCap
	}
	if !IsNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.SignAndSubmit) {
		toSerialize["signAndSubmit"] = o.SignAndSubmit
	}
	if !IsNil(o.NonceManagement) {
		toSerialize["nonceManagement"] = o.NonceManagement
	}
	if !IsNil(o.PreEIP1559) {
		toSerialize["preEIP1559"] = o.PreEIP1559
	}
	if !IsNil(o.Signer) {
		toSerialize["signer"] = o.Signer
	}
	if !IsNil(o.FormatInts) {
		toSerialize["formatInts"] = o.FormatInts
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.BlockNumber) {
		toSerialize["blockNumber"] = o.BlockNumber
	}
	if !IsNil(o.ContractOverride) {
		toSerialize["contractOverride"] = o.ContractOverride
	}
	if !IsNil(o.Preview) {
		toSerialize["preview"] = o.Preview
	}
	return toSerialize, nil
}

type NullablePostMethodArgs struct {
	value *PostMethodArgs
	isSet bool
}

func (v NullablePostMethodArgs) Get() *PostMethodArgs {
	return v.value
}

func (v *NullablePostMethodArgs) Set(val *PostMethodArgs) {
	v.value = val
	v.isSet = true
}

func (v NullablePostMethodArgs) IsSet() bool {
	return v.isSet
}

func (v *NullablePostMethodArgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostMethodArgs(val *PostMethodArgs) *NullablePostMethodArgs {
	return &NullablePostMethodArgs{value: val, isSet: true}
}

func (v NullablePostMethodArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostMethodArgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
