package keeper

import (
	"github.com/comdex-official/comdex/x/lend/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetUtilisationRatioByPoolIdAndAssetId(ctx sdk.Context, poolId, assetId uint64) (sdk.Int, error) {
	pool, _ := k.GetPool(ctx, poolId)
	asset, _ := k.GetAsset(ctx, assetId)
	moduleBalance := k.ModuleBalance(ctx, pool.ModuleName, asset.Denom)
	assetStats, found := k.GetAssetStatsByPoolIdAndAssetId(ctx, assetId, poolId)
	if !found {
		return sdk.ZeroInt(), types.ErrAssetStatsNotFound
	}
	utilizationRatio := assetStats.TotalBorrowed.Quo(moduleBalance)
	return utilizationRatio, nil
}
