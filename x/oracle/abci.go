package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v3/x/oracle/keeper"
	"github.com/bandprotocol/chain/v3/x/oracle/types"
)

// BeginBlocker re-calculates and saves the rolling seed value based on block hashes.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) error {
	// Update rolling seed used for pseudorandom oracle provider selection.
	hash := ctx.HeaderInfo().Hash
	// On the first block in the test. it's possible to have empty hash.
	if len(hash) > 0 {
		rollingSeed := k.GetRollingSeed(ctx)
		k.SetRollingSeed(ctx, append(rollingSeed[1:], hash[0]))
	}
	// Reward a portion of block rewards (inflation + tx fee) to active oracle validators.
	k.AllocateTokens(ctx, ctx.VoteInfos())
	return nil
}

// handleEndBlock cleans up the state during end block. See comment in the implementation!
func EndBlocker(ctx sdk.Context, k keeper.Keeper) error {
	// Loops through all requests in the resolvable list to resolve all of them!
	for _, reqID := range k.GetPendingResolveList(ctx) {
		k.ResolveRequest(ctx, reqID)
	}

	// Once all the requests are resolved, we can clear the list.
	k.SetPendingResolveList(ctx, []types.RequestID{})
	// Lastly, we clean up data requests that are supposed to be expired.
	k.ProcessExpiredRequests(ctx)
	// NOTE: We can remove old requests from state to optimize space, using `k.DeleteRequest`
	// and `k.DeleteReports`. We don't do that now as it is premature optimization at this state.
	return nil
}
