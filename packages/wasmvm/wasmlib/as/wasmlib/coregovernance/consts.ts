// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

import * as wasmtypes from '../wasmtypes';

export const ScName        = 'governance';
export const ScDescription = 'Governance contract';
export const HScName       = new wasmtypes.ScHname(0x17cf909f);

export const ParamAccessNodeInfoAccessAPI    = 'ia';
export const ParamAccessNodeInfoCertificate  = 'ic';
export const ParamAccessNodeInfoForCommittee = 'i';
export const ParamAccessNodeInfoPubKey       = 'ip';
export const ParamChainOwner                 = 'o';
export const ParamChangeAccessNodesActions   = 'n';
export const ParamFeePolicyBytes             = 'g';
export const ParamStateControllerAddress     = 'S';

export const ResultAccessNodeCandidates            = 'ac';
export const ResultAccessNodes                     = 'an';
export const ResultAllowedStateControllerAddresses = 'a';
export const ResultChainID                         = 'c';
export const ResultChainOwner                      = 'o';
export const ResultChainOwnerID                    = 'o';
export const ResultCustomMetadata                  = 'x';
export const ResultFeePolicyBytes                  = 'g';
export const ResultGasFeePolicyBytes               = 'g';

export const FuncAddAllowedStateControllerAddress    = 'addAllowedStateControllerAddress';
export const FuncAddCandidateNode                    = 'addCandidateNode';
export const FuncChangeAccessNodes                   = 'changeAccessNodes';
export const FuncClaimChainOwnership                 = 'claimChainOwnership';
export const FuncDelegateChainOwnership              = 'delegateChainOwnership';
export const FuncRemoveAllowedStateControllerAddress = 'removeAllowedStateControllerAddress';
export const FuncRevokeAccessNode                    = 'revokeAccessNode';
export const FuncRotateStateController               = 'rotateStateController';
export const FuncSetFeePolicy                        = 'setFeePolicy';
export const ViewGetAllowedStateControllerAddresses  = 'getAllowedStateControllerAddresses';
export const ViewGetChainInfo                        = 'getChainInfo';
export const ViewGetChainNodes                       = 'getChainNodes';
export const ViewGetChainOwner                       = 'getChainOwner';
export const ViewGetFeePolicy                        = 'getFeePolicy';

export const HFuncAddAllowedStateControllerAddress    = new wasmtypes.ScHname(0x9469d567);
export const HFuncAddCandidateNode                    = new wasmtypes.ScHname(0xb745b382);
export const HFuncChangeAccessNodes                   = new wasmtypes.ScHname(0x7bca3700);
export const HFuncClaimChainOwnership                 = new wasmtypes.ScHname(0x03ff0fc0);
export const HFuncDelegateChainOwnership              = new wasmtypes.ScHname(0x93ecb6ad);
export const HFuncRemoveAllowedStateControllerAddress = new wasmtypes.ScHname(0x31f69447);
export const HFuncRevokeAccessNode                    = new wasmtypes.ScHname(0x5459512d);
export const HFuncRotateStateController               = new wasmtypes.ScHname(0x244d1038);
export const HFuncSetFeePolicy                        = new wasmtypes.ScHname(0x5b791c9f);
export const HViewGetAllowedStateControllerAddresses  = new wasmtypes.ScHname(0xf3505183);
export const HViewGetChainInfo                        = new wasmtypes.ScHname(0x434477e2);
export const HViewGetChainNodes                       = new wasmtypes.ScHname(0xe1832289);
export const HViewGetChainOwner                       = new wasmtypes.ScHname(0x9b2ef0ac);
export const HViewGetFeePolicy                        = new wasmtypes.ScHname(0xf8c89790);
