package test

import (
	"testing"

	"github.com/iotaledger/wasp/contracts/wasm/testcore/go/testcore"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmhost"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/coreaccounts"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func Test2Chains(t *testing.T) {
	run2(t, func(t *testing.T, w bool) {
		chain1 := wasmsolo.StartChain(t, "chain1")
		chain1.CheckAccountLedger()

		chain2 := wasmsolo.StartChain(t, "chain2", chain1.Env)
		chain2.CheckAccountLedger()

		user := wasmsolo.NewSoloAgent(chain1.Env, "user")
		userL1 := user.Balance()

		ctx1 := deployTestCoreOnChain(t, w, chain1, nil)
		require.NoError(t, ctx1.Err)
		ctx2 := deployTestCoreOnChain(t, w, chain2, nil)
		require.NoError(t, ctx2.Err)

		ctxAcc1 := ctx1.SoloContextForCore(t, coreaccounts.ScName, coreaccounts.OnDispatch)
		ctxAcc2 := ctx2.SoloContextForCore(t, coreaccounts.ScName, coreaccounts.OnDispatch)

		testcore1 := ctx1.Account()
		testcore2 := ctx2.Account()
		accounts1 := ctxAcc1.Account()
		accounts2 := ctxAcc2.Account()
		bal1 := ctx1.Balances(user, testcore2, accounts2)
		bal2 := ctx2.Balances(user, testcore1, accounts1)

		ctx1.WaitForPendingRequestsMark()
		ctx2.WaitForPendingRequestsMark()

		const baseTokens = 7 * isc.Million
		const gasBuffer = 10_000

		xfer := coreaccounts.ScFuncs.TransferAllowanceTo(ctxAcc1.Sign(user))
		xfer.Params.AgentID().SetValue(testcore2.ScAgentID())
		xfer.Func.TransferBaseTokens(baseTokens + gasBuffer).
			AllowanceBaseTokens(baseTokens).Post()
		require.NoError(t, ctxAcc1.Err)

		userL1 -= baseTokens + gasBuffer
		require.Equal(t, userL1, user.Balance())
		bal1.Chain += ctxAcc1.GasFee
		bal1.Add(user, gasBuffer-ctxAcc1.GasFee)
		bal1.Add(testcore2, baseTokens)
		bal1.VerifyBalances(t)
		bal2.VerifyBalances(t)

		f := testcore.ScFuncs.WithdrawFromChain(ctx2)
		f.Params.ChainID().SetValue(ctx1.CurrentChainID())
		f.Params.BaseTokensWithdrawal().SetValue(baseTokens)
		const allowanceForWithdraw = 33_333
		f.Func.AllowanceBaseTokens(allowanceForWithdraw).Post()
		require.NoError(t, ctx2.Err)

		require.True(t, ctx1.WaitForPendingRequests(2))
		require.True(t, ctx2.WaitForPendingRequests(2))
		ctx1.UpdateGasFees()

		bal1.Add(testcore2, allowanceForWithdraw-ctx1.GasFee-baseTokens)
		bal1.Chain += ctx1.GasFee
		bal1.VerifyBalances(t)

		bal2.Add(accounts1, accounts.ConstDepositFeeTmp-ctx2.GasFee)
		bal2.Account += baseTokens - accounts.ConstDepositFeeTmp
		bal2.Originator += wasmhost.WasmStorageDeposit - ctx2.GasFee - allowanceForWithdraw
		bal2.Chain += 2 * ctx2.GasFee
		bal2.VerifyBalances(t)
	})
}
