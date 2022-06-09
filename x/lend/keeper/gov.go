package keeper

import (
	"github.com/comdex-official/comdex/x/lend/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) HandleAddWhitelistedPairsRecords(ctx sdk.Context, p *types.LendPairsProposal) error {
	return k.AddLendPairsRecords(ctx, p.Pairs...)
}

func (k Keeper) HandleUpdateWhitelistedPairRecords(ctx sdk.Context, p *types.UpdatePairProposal) error {
	return k.UpdateLendPairRecords(ctx, p.Pair)
}

func (k Keeper) HandleAddPoolRecords(ctx sdk.Context, p *types.AddPoolsProposal) error {
	return k.AddPoolRecords(ctx, p.Pool)
}

func (k Keeper) HandleAddAssetToPairRecords(ctx sdk.Context, p *types.AddAssetToPairProposal) error {
	return k.AddAssetToPair(ctx, p.AssetToPairMapping)
}
