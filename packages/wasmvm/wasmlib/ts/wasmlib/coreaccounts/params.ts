// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ImmutableDepositParams extends wasmtypes.ScProxy {
	agentID(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamAgentID));
	}
}

export class MutableDepositParams extends wasmtypes.ScProxy {
	agentID(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamAgentID));
	}
}

export class ImmutableHarvestParams extends wasmtypes.ScProxy {
	withdrawAmount(): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.root(sc.ParamWithdrawAmount));
	}
}

export class MutableHarvestParams extends wasmtypes.ScProxy {
	withdrawAmount(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.root(sc.ParamWithdrawAmount));
	}
}

export class ImmutableGetAccountNonceParams extends wasmtypes.ScProxy {
	agentID(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamAgentID));
	}
}

export class MutableGetAccountNonceParams extends wasmtypes.ScProxy {
	agentID(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamAgentID));
	}
}
