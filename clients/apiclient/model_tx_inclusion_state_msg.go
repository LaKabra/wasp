/*
Wasp API

REST API for the Wasp node

API version: 0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the TxInclusionStateMsg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TxInclusionStateMsg{}

// TxInclusionStateMsg struct for TxInclusionStateMsg
type TxInclusionStateMsg struct {
	// The inclusion state
	State string `json:"state"`
	// The transaction ID
	TxId string `json:"txId"`
}

// NewTxInclusionStateMsg instantiates a new TxInclusionStateMsg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxInclusionStateMsg(state string, txId string) *TxInclusionStateMsg {
	this := TxInclusionStateMsg{}
	this.State = state
	this.TxId = txId
	return &this
}

// NewTxInclusionStateMsgWithDefaults instantiates a new TxInclusionStateMsg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxInclusionStateMsgWithDefaults() *TxInclusionStateMsg {
	this := TxInclusionStateMsg{}
	return &this
}

// GetState returns the State field value
func (o *TxInclusionStateMsg) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *TxInclusionStateMsg) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *TxInclusionStateMsg) SetState(v string) {
	o.State = v
}

// GetTxId returns the TxId field value
func (o *TxInclusionStateMsg) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *TxInclusionStateMsg) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *TxInclusionStateMsg) SetTxId(v string) {
	o.TxId = v
}

func (o TxInclusionStateMsg) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TxInclusionStateMsg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["txId"] = o.TxId
	return toSerialize, nil
}

type NullableTxInclusionStateMsg struct {
	value *TxInclusionStateMsg
	isSet bool
}

func (v NullableTxInclusionStateMsg) Get() *TxInclusionStateMsg {
	return v.value
}

func (v *NullableTxInclusionStateMsg) Set(val *TxInclusionStateMsg) {
	v.value = val
	v.isSet = true
}

func (v NullableTxInclusionStateMsg) IsSet() bool {
	return v.isSet
}

func (v *NullableTxInclusionStateMsg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxInclusionStateMsg(val *TxInclusionStateMsg) *NullableTxInclusionStateMsg {
	return &NullableTxInclusionStateMsg{value: val, isSet: true}
}

func (v NullableTxInclusionStateMsg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxInclusionStateMsg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


