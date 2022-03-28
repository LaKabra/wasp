// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";

export const ScName        = "governance";
export const ScDescription = "Core governance contract";
export const HScName       = new wasmtypes.ScHname(0x17cf909f);

export const ParamChainOwner             = "oi";
export const ParamHname                  = "hn";
export const ParamMaxBlobSize            = "bs";
export const ParamMaxEventSize           = "es";
export const ParamMaxEventsPerReq        = "ne";
export const ParamOwnerFee               = "of";
export const ParamStateControllerAddress = "S";
export const ParamValidatorFee           = "vf";

export const ResultAllowedStateControllerAddresses = "a";
export const ResultChainID                         = "c";
export const ResultChainOwnerID                    = "o";
export const ResultDefaultOwnerFee                 = "do";
export const ResultDefaultValidatorFee             = "dv";
export const ResultDescription                     = "d";
export const ResultMaxBlobSize                     = "mb";
export const ResultMaxEventSize                    = "me";
export const ResultMaxEventsPerReq                 = "mr";
export const ResultOwnerFee                        = "of";
export const ResultValidatorFee                    = "vf";

export const FuncAddAllowedStateControllerAddress    = "addAllowedStateControllerAddress";
export const FuncClaimChainOwnership                 = "claimChainOwnership";
export const FuncDelegateChainOwnership              = "delegateChainOwnership";
export const FuncRemoveAllowedStateControllerAddress = "removeAllowedStateControllerAddress";
export const FuncRotateStateController               = "rotateStateController";
export const FuncSetChainInfo                        = "setChainInfo";
export const FuncSetContractFee                      = "setContractFee";
export const FuncSetDefaultFee                       = "setDefaultFee";
export const ViewGetAllowedStateControllerAddresses  = "getAllowedStateControllerAddresses";
export const ViewGetChainInfo                        = "getChainInfo";
export const ViewGetFeeInfo                          = "getFeeInfo";
export const ViewGetMaxBlobSize                      = "getMaxBlobSize";

export const HFuncAddAllowedStateControllerAddress    = new wasmtypes.ScHname(0x9469d567);
export const HFuncClaimChainOwnership                 = new wasmtypes.ScHname(0x03ff0fc0);
export const HFuncDelegateChainOwnership              = new wasmtypes.ScHname(0x93ecb6ad);
export const HFuncRemoveAllowedStateControllerAddress = new wasmtypes.ScHname(0x31f69447);
export const HFuncRotateStateController               = new wasmtypes.ScHname(0x244d1038);
export const HFuncSetChainInfo                        = new wasmtypes.ScHname(0x702f5d2b);
export const HFuncSetContractFee                      = new wasmtypes.ScHname(0x8421a42b);
export const HFuncSetDefaultFee                       = new wasmtypes.ScHname(0x3310ecd0);
export const HViewGetAllowedStateControllerAddresses  = new wasmtypes.ScHname(0xf3505183);
export const HViewGetChainInfo                        = new wasmtypes.ScHname(0x434477e2);
export const HViewGetFeeInfo                          = new wasmtypes.ScHname(0x9fe54b48);
export const HViewGetMaxBlobSize                      = new wasmtypes.ScHname(0xe1db3d28);
