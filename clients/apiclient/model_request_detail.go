/*
Wasp API

REST API for the Wasp node

API version: 0.4.0-alpha.2-403-g907bc6c54
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the RequestDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestDetail{}

// RequestDetail struct for RequestDetail
type RequestDetail struct {
	Allowance Allowance `json:"allowance"`
	CallTarget CallTarget `json:"callTarget"`
	FungibleTokens FungibleTokens `json:"fungibleTokens"`
	GasGudget int64 `json:"gasGudget"`
	IsEVM bool `json:"isEVM"`
	IsOffLedger bool `json:"isOffLedger"`
	Nft NFTDataResponse `json:"nft"`
	Params JSONDict `json:"params"`
	RequestId string `json:"requestId"`
	SenderAccount string `json:"senderAccount"`
	TargetAddress string `json:"targetAddress"`
}

// NewRequestDetail instantiates a new RequestDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestDetail(allowance Allowance, callTarget CallTarget, fungibleTokens FungibleTokens, gasGudget int64, isEVM bool, isOffLedger bool, nft NFTDataResponse, params JSONDict, requestId string, senderAccount string, targetAddress string) *RequestDetail {
	this := RequestDetail{}
	this.Allowance = allowance
	this.CallTarget = callTarget
	this.FungibleTokens = fungibleTokens
	this.GasGudget = gasGudget
	this.IsEVM = isEVM
	this.IsOffLedger = isOffLedger
	this.Nft = nft
	this.Params = params
	this.RequestId = requestId
	this.SenderAccount = senderAccount
	this.TargetAddress = targetAddress
	return &this
}

// NewRequestDetailWithDefaults instantiates a new RequestDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestDetailWithDefaults() *RequestDetail {
	this := RequestDetail{}
	return &this
}

// GetAllowance returns the Allowance field value
func (o *RequestDetail) GetAllowance() Allowance {
	if o == nil {
		var ret Allowance
		return ret
	}

	return o.Allowance
}

// GetAllowanceOk returns a tuple with the Allowance field value
// and a boolean to check if the value has been set.
func (o *RequestDetail) GetAllowanceOk() (*Allowance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allowance, true
}

// SetAllowance sets field value
func (o *RequestDetail) SetAllowance(v Allowance) {
	o.Allowance = v
}

// GetCallTarget returns the CallTarget field value
func (o *RequestDetail) GetCallTarget() CallTarget {
	if o == nil {
		var ret CallTarget
		return ret
	}

	return o.CallTarget
}

// GetCallTargetOk returns a tuple with the CallTarget field value
// and a boolean to check if the value has been set.
func (o *RequestDetail) GetCallTargetOk() (*CallTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallTarget, true
}

// SetCallTarget sets field value
func (o *RequestDetail) SetCallTarget(v CallTarget) {
	o.CallTarget = v
}

// GetFungibleTokens returns the FungibleTokens field value
func (o *RequestDetail) GetFungibleTokens() FungibleTokens {
	if o == nil {
		var ret FungibleTokens
		return ret
	}

	return o.FungibleTokens
}

// GetFungibleTokensOk returns a tuple with the FungibleTokens field value
// and a boolean to check if the value has been set.
func (o *RequestDetail) GetFungibleTokensOk() (*FungibleTokens, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FungibleTokens, true
}

// SetFungibleTokens sets field value
func (o *RequestDetail) SetFungibleTokens(v FungibleTokens) {
	o.FungibleTokens = v
}

// GetGasGudget returns the GasGudget field value
func (o *RequestDetail) GetGasGudget() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GasGudget
}

// GetGasGudgetOk returns a tuple with the GasGudget field value
// and a boolean to check if the value has been set.
func (o *RequestDetail) GetGasGudgetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasGudget, true
}

// SetGasGudget sets field value
func (o *RequestDetail) SetGasGudget(v int64) {
	o.GasGudget = v
}

// GetIsEVM returns the IsEVM field value
func (o *RequestDetail) GetIsEVM() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEVM
}

// GetIsEVMOk returns a tuple with the IsEVM field value
// and a boolean to check if the value has been set.
func (o *RequestDetail) GetIsEVMOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEVM, true
}

// SetIsEVM sets field value
func (o *RequestDetail) SetIsEVM(v bool) {
	o.IsEVM = v
}

// GetIsOffLedger returns the IsOffLedger field value
func (o *RequestDetail) GetIsOffLedger() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsOffLedger
}

// GetIsOffLedgerOk returns a tuple with the IsOffLedger field value
// and a boolean to check if the value has been set.
func (o *RequestDetail) GetIsOffLedgerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsOffLedger, true
}

// SetIsOffLedger sets field value
func (o *RequestDetail) SetIsOffLedger(v bool) {
	o.IsOffLedger = v
}

// GetNft returns the Nft field value
func (o *RequestDetail) GetNft() NFTDataResponse {
	if o == nil {
		var ret NFTDataResponse
		return ret
	}

	return o.Nft
}

// GetNftOk returns a tuple with the Nft field value
// and a boolean to check if the value has been set.
func (o *RequestDetail) GetNftOk() (*NFTDataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nft, true
}

// SetNft sets field value
func (o *RequestDetail) SetNft(v NFTDataResponse) {
	o.Nft = v
}

// GetParams returns the Params field value
func (o *RequestDetail) GetParams() JSONDict {
	if o == nil {
		var ret JSONDict
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *RequestDetail) GetParamsOk() (*JSONDict, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *RequestDetail) SetParams(v JSONDict) {
	o.Params = v
}

// GetRequestId returns the RequestId field value
func (o *RequestDetail) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *RequestDetail) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *RequestDetail) SetRequestId(v string) {
	o.RequestId = v
}

// GetSenderAccount returns the SenderAccount field value
func (o *RequestDetail) GetSenderAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderAccount
}

// GetSenderAccountOk returns a tuple with the SenderAccount field value
// and a boolean to check if the value has been set.
func (o *RequestDetail) GetSenderAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAccount, true
}

// SetSenderAccount sets field value
func (o *RequestDetail) SetSenderAccount(v string) {
	o.SenderAccount = v
}

// GetTargetAddress returns the TargetAddress field value
func (o *RequestDetail) GetTargetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAddress
}

// GetTargetAddressOk returns a tuple with the TargetAddress field value
// and a boolean to check if the value has been set.
func (o *RequestDetail) GetTargetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetAddress, true
}

// SetTargetAddress sets field value
func (o *RequestDetail) SetTargetAddress(v string) {
	o.TargetAddress = v
}

func (o RequestDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowance"] = o.Allowance
	toSerialize["callTarget"] = o.CallTarget
	toSerialize["fungibleTokens"] = o.FungibleTokens
	toSerialize["gasGudget"] = o.GasGudget
	toSerialize["isEVM"] = o.IsEVM
	toSerialize["isOffLedger"] = o.IsOffLedger
	toSerialize["nft"] = o.Nft
	toSerialize["params"] = o.Params
	toSerialize["requestId"] = o.RequestId
	toSerialize["senderAccount"] = o.SenderAccount
	toSerialize["targetAddress"] = o.TargetAddress
	return toSerialize, nil
}

type NullableRequestDetail struct {
	value *RequestDetail
	isSet bool
}

func (v NullableRequestDetail) Get() *RequestDetail {
	return v.value
}

func (v *NullableRequestDetail) Set(val *RequestDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestDetail(val *RequestDetail) *NullableRequestDetail {
	return &NullableRequestDetail{value: val, isSet: true}
}

func (v NullableRequestDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


