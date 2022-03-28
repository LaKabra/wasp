// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:gocritic
package fairroulette

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type FairRouletteEvents struct {
}

func (e FairRouletteEvents) Bet(address wasmtypes.ScAddress, amount uint64, number uint16) {
	evt := wasmlib.NewEventEncoder("fairroulette.bet")
	evt.Encode(wasmtypes.AddressToString(address))
	evt.Encode(wasmtypes.Uint64ToString(amount))
	evt.Encode(wasmtypes.Uint16ToString(number))
	evt.Emit()
}

func (e FairRouletteEvents) Payout(address wasmtypes.ScAddress, amount uint64) {
	evt := wasmlib.NewEventEncoder("fairroulette.payout")
	evt.Encode(wasmtypes.AddressToString(address))
	evt.Encode(wasmtypes.Uint64ToString(amount))
	evt.Emit()
}

func (e FairRouletteEvents) Round(number uint32) {
	evt := wasmlib.NewEventEncoder("fairroulette.round")
	evt.Encode(wasmtypes.Uint32ToString(number))
	evt.Emit()
}

func (e FairRouletteEvents) Start() {
	evt := wasmlib.NewEventEncoder("fairroulette.start")
	evt.Emit()
}

func (e FairRouletteEvents) Stop() {
	evt := wasmlib.NewEventEncoder("fairroulette.stop")
	evt.Emit()
}

func (e FairRouletteEvents) Winner(number uint16) {
	evt := wasmlib.NewEventEncoder("fairroulette.winner")
	evt.Encode(wasmtypes.Uint16ToString(number))
	evt.Emit()
}
