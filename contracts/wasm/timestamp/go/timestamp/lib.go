// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:dupl
package timestamp

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

var exportMap = wasmlib.ScExportMap{
	Names: []string{
    	FuncNow,
    	ViewGetTimestamp,
	},
	Funcs: []wasmlib.ScFuncContextFunction{
    	funcNowThunk,
	},
	Views: []wasmlib.ScViewContextFunction{
    	viewGetTimestampThunk,
	},
}

func OnLoad(index int32) {
	if index >= 0 {
		wasmlib.ScExportsCall(index, &exportMap)
		return
	}

	wasmlib.ScExportsExport(&exportMap)
}

type NowContext struct {
	State   MutabletimestampState
}

func funcNowThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("timestamp.funcNow")
	f := &NowContext{
		State: MutabletimestampState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcNow(ctx, f)
	ctx.Log("timestamp.funcNow ok")
}

type GetTimestampContext struct {
	Results MutableGetTimestampResults
	State   ImmutabletimestampState
}

func viewGetTimestampThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("timestamp.viewGetTimestamp")
	results := wasmlib.NewScDict()
	f := &GetTimestampContext{
		Results: MutableGetTimestampResults{
			proxy: results.AsProxy(),
		},
		State: ImmutabletimestampState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	viewGetTimestamp(ctx, f)
	ctx.Results(results)
	ctx.Log("timestamp.viewGetTimestamp ok")
}
