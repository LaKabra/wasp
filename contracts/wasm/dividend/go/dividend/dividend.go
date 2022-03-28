// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// This example implements 'dividend', a simple smart contract that will
// automatically disperse iota tokens which are sent to the contract to a group
// of member addresses according to predefined division factors. The intent is
// to showcase basic functionality of WasmLib through a minimal implementation
// and not to come up with a complete robust real-world solution.
// Note that we have drawn sometimes out constructs that could have been done
// in a single line over multiple statements to be able to properly document
// step by step what is happening in the code. We also unnecessarily annotate
// all 'var' statements with their assignment type to improve understanding.

//nolint:revive
package dividend

import (
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
)

// 'init' is used as a way to initialize a smart contract. It is an optional
// function that will automatically be called upon contract deployment. In this
// case we use it to initialize the 'owner' state variable so that we can later
// use this information to prevent non-owners from calling certain functions.
// The 'init' function takes a single optional parameter:
// - 'owner', which is the agent id of the entity owning the contract.
// When this parameter is omitted the owner will default to the contract creator.
func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
	// The schema tool has already created a proper InitContext for this function that
	// allows us to access call parameters and state storage in a type-safe manner.

	// First we set up a default value for the owner in case the optional
	// 'owner' parameter was omitted.
	var owner wasmtypes.ScAgentID = ctx.ContractCreator()

	// Now we check if the optional 'owner' parameter is present in the params map.
	if f.Params.Owner().Exists() {
		// Yes, it was present, so now we overwrite the default owner with
		// the one specified by the 'owner' parameter.
		owner = f.Params.Owner().Value()
	}

	// Now that we have sorted out which agent will be the owner of this contract
	// we will save this value in the 'owner' variable in state storage on the host.
	// Read the documentation on schema.json to understand why this state variable is
	// supported at compile-time by code generated from schema.json by the schema tool.
	f.State.Owner().SetValue(owner)
}

// 'member' is a function that can be used only by the entity that owns the
// 'dividend' smart contract. It can be used to define the group of member
// addresses and dispersal factors one by one prior to sending tokens to the
// smart contract's 'divide' function. The 'member' function takes 2 parameters,
// which are both required:
// - 'address', which is an Address to use as member in the group, and
// - 'factor',  which is an Int64 relative dispersal factor associated with
//              that address
// The 'member' function will save the address/factor combination in state storage
// and also calculate and store a running sum of all factors so that the 'divide'
// function can simply start using these precalculated values when called.
func funcMember(ctx wasmlib.ScFuncContext, f *MemberContext) {
	// Note that the schema tool has already dealt with making sure that this function
	// can only be called by the owner and that the required parameters are present.
	// So once we get to this point in the code we can take that as a given.

	// Since we are sure that the 'factor' parameter actually exists we can
	// retrieve its actual value into an uint64. Note that we use Go's built-in
	// data types when manipulating Int64, String, or Bytes value objects.
	var factor uint64 = f.Params.Factor().Value()

	// Since we are sure that the 'address' parameter actually exists we can
	// retrieve its actual value into an ScAddress value type.
	var address wasmtypes.ScAddress = f.Params.Address().Value()

	// We will store the address/factor combinations in a key/value sub-map of the
	// state storage named 'members'. The schema tool has generated an appropriately
	// type-checked proxy map for us from the schema.json state storage definition.
	// If there is no 'members' map present yet in state storage an empty map will
	// automatically be created on the host.
	var members MapAddressToMutableUint64 = f.State.Members()

	// Now we create an ScMutableUint64 proxy for the value stored in the 'members'
	// map under the key defined by the 'address' parameter we retrieved earlier.
	var currentFactor wasmtypes.ScMutableUint64 = members.GetUint64(address)

	// We check to see if this key/value combination exists in the 'members' map.
	if !currentFactor.Exists() {
		// If it does not exist yet then we have to add this new address to the
		// 'memberList' array that keeps track of all address keys used in the
		// 'members' map. The schema tool has again created the appropriate type
		// for us already. Here too, if the address array was not present yet it
		// will automatically be created on the host.
		var memberList ArrayOfMutableAddress = f.State.MemberList()

		// Now we will append the new address to the memberList array.
		// We create an ScMutableAddress proxy to an address value that lives
		// at the end of the memberList array (no value yet, since we're appending).
		var newAddress wasmtypes.ScMutableAddress = memberList.AppendAddress()

		// And finally we append the new address to the array by telling the proxy
		// to update the value it refers to with the 'address' parameter.
		newAddress.SetValue(address)

		// Note that we could have achieved the last 3 lines of code in a single line:
		// memberList.GetAddress(memberList.Length()).SetValue(&address)
	}

	// Create an ScMutableUint64 proxy named 'totalFactor' for an Uint64 value in
	// state storage. Note that we don't care whether this value exists or not,
	// because WasmLib will treat it as if it has the default value of zero.
	var totalFactor wasmtypes.ScMutableUint64 = f.State.TotalFactor()

	// Now we calculate the new running total sum of factors by first getting the
	// current value of 'totalFactor' from the state storage, then subtracting the
	// current value of the factor associated with the 'address' parameter, if any
	// exists. Again, if the associated value doesn't exist, WasmLib will assume it
	// to be zero. Finally we add the factor retrieved from the parameters,
	// resulting in the new totalFactor.
	var newTotalFactor uint64 = totalFactor.Value() - currentFactor.Value() + factor

	// Now we store the new totalFactor in the state storage.
	totalFactor.SetValue(newTotalFactor)

	// And we also store the factor from the parameters under the address from the
	// parameters in the state storage that the proxy refers to.
	currentFactor.SetValue(factor)
}

// 'divide' is a function that will take any iotas it receives and properly
// disperse them to the addresses in the member list according to the dispersion
// factors associated with these addresses.
// Anyone can send iota tokens to this function and they will automatically be
// divided over the member list. Note that this function does not deal with
// fractions. It simply truncates the calculated amount to the nearest lower
// integer and keeps any remaining iotas in the sender account.
func funcDivide(ctx wasmlib.ScFuncContext, f *DivideContext) {
	// Create an ScBalances map proxy to the account balances for this
	// smart contract. Note that ScBalances wraps an ScImmutableMap of
	// token color/amount combinations in a simpler to use interface.
	var allowance wasmlib.ScBalances = ctx.Allowance()

	// Retrieve the amount of plain iota tokens from the account balance
	var amount uint64 = allowance.Balance(wasmtypes.IOTA)

	// Retrieve the pre-calculated totalFactor value from the state storage.
	var totalFactor uint64 = f.State.TotalFactor().Value()

	// Get the proxy to the 'members' map in the state storage.
	var members MapAddressToMutableUint64 = f.State.Members()

	// Get the proxy to the 'memberList' array in the state storage.
	var memberList ArrayOfMutableAddress = f.State.MemberList()

	// Determine the current length of the memberList array.
	var size uint32 = memberList.Length()

	// Loop through all indexes of the memberList array.
	for i := uint32(0); i < size; i++ {
		// Retrieve the next indexed address from the memberList array.
		var address wasmtypes.ScAddress = memberList.GetAddress(i).Value()

		// Retrieve the factor associated with the address from the members map.
		var factor uint64 = members.GetUint64(address).Value()

		// Calculate the fair share of iotas to disperse to this member based on the
		// factor we just retrieved. Note that the result will been truncated.
		var share uint64 = amount * factor / totalFactor

		// Is there anything to disperse to this member?
		if share > 0 {
			// Yes, so let's set up an ScTransfer map proxy that transfers the
			// calculated amount of iotas. Note that ScTransfer wraps an
			// ScMutableMap of token color/amount combinations in a simpler to use
			// interface. The constructor we use here creates and initializes a
			// single token color transfer in a single statement. The actual color
			// and amount values passed in will be stored in a new map on the host.
			var transfers wasmlib.ScTransfer = wasmlib.NewScTransferIotas(share)

			// Perform the actual transfer of tokens from the smart contract to the
			// member address. The transfer_to_address() method receives the address
			// value and the proxy to the new transfers map on the host, and will
			// call the corresponding host sandbox function with these values.
			ctx.TransferAllowed(address.AsAgentID(), transfers, true)
		}
	}
}

// 'setOwner' is used to change the owner of the smart contract.
// It updates the 'owner' state variable with the provided agent id.
// The 'setOwner' function takes a single mandatory parameter:
// - 'owner', which is the agent id of the entity that will own the contract.
// Only the current owner can change the owner.
func funcSetOwner(ctx wasmlib.ScFuncContext, f *SetOwnerContext) {
	// Note that the schema tool has already dealt with making sure that this function
	// can only be called by the owner and that the required parameter is present.
	// So once we get to this point in the code we can take that as a given.

	// Save the new owner parameter value in the 'owner' variable in state storage.
	f.State.Owner().SetValue(f.Params.Owner().Value())
}

// 'getFactor' is a simple View function. It will retrieve the factor
// associated with the (mandatory) address parameter it was provided with.
func viewGetFactor(ctx wasmlib.ScViewContext, f *GetFactorContext) {
	// Since we are sure that the 'address' parameter actually exists we can
	// retrieve its actual value into an ScAddress value type.
	var address wasmtypes.ScAddress = f.Params.Address().Value()

	// Create an ScImmutableMap proxy to the 'members' map in the state storage.
	// Note that for views this is an *immutable* map as opposed to the *mutable*
	// map we can access from the *mutable* state that gets passed to funcs.
	var members MapAddressToImmutableUint64 = f.State.Members()

	// Retrieve the factor associated with the address parameter.
	var factor uint64 = members.GetUint64(address).Value()

	// Set the factor in the results map of the function context.
	// The contents of this results map is returned to the caller of the function.
	f.Results.Factor().SetValue(factor)
}

// 'getOwner' can be used to retrieve the current owner of the dividend contract
func viewGetOwner(ctx wasmlib.ScViewContext, f *GetOwnerContext) {
	f.Results.Owner().SetValue(f.State.Owner().Value())
}
