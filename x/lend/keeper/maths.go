package keeper

import (
	"github.com/comdex-official/comdex/x/lend/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetUtilisationRatioByPoolIdAndAssetId(ctx sdk.Context, poolId, assetId uint64) (sdk.Dec, error) {
	pool, _ := k.GetPool(ctx, poolId)
	asset, _ := k.GetAsset(ctx, assetId)
	moduleBalance := k.ModuleBalance(ctx, pool.ModuleName, asset.Denom)
	assetStats, found := k.GetAssetStatsByPoolIdAndAssetId(ctx, assetId, poolId)
	if !found {
		return sdk.ZeroDec(), types.ErrAssetStatsNotFound
	}
	utilizationRatio := assetStats.TotalBorrowed.ToDec().Quo(moduleBalance.ToDec())
	return utilizationRatio, nil
}

func (k Keeper) GetBorrowAPYByAssetId(ctx sdk.Context, poolId, assetId uint64, IsStableBorrow bool) (borrowAPY sdk.Dec, err error) {
	assetRatesStats, found := k.GetAssetRatesStats(ctx, assetId)
	if !found {
		return sdk.ZeroDec(), types.ErrorAssetStatsNotFound
	}
	currentUtilisationRatio, err := k.GetUtilisationRatioByPoolIdAndAssetId(ctx, poolId, assetId)
	if err != nil {
		return sdk.ZeroDec(), err
	}
	if !IsStableBorrow {
		if currentUtilisationRatio.LT(assetRatesStats.UOptimal) {
			utilisationRatio := currentUtilisationRatio.Quo(assetRatesStats.UOptimal)
			multiplicationFactor := utilisationRatio.Mul(assetRatesStats.Slope1)
			borrowAPY = assetRatesStats.Base.Add(multiplicationFactor)
			return borrowAPY, nil
		} else {
			utilisationNumerator := currentUtilisationRatio.Sub(assetRatesStats.UOptimal)
			utilisationDenominator := sdk.OneDec().Sub(assetRatesStats.UOptimal)
			utilisationRatio := utilisationNumerator.Quo(utilisationDenominator)
			multiplicationFactor := utilisationRatio.Mul(assetRatesStats.Slope2)
			borrowAPY = assetRatesStats.Base.Add(assetRatesStats.Slope1).Add(multiplicationFactor)
			return borrowAPY, nil
		}
	} else {
		if currentUtilisationRatio.LT(assetRatesStats.UOptimal) {
			utilisationRatio := currentUtilisationRatio.Quo(assetRatesStats.UOptimal)
			multiplicationFactor := utilisationRatio.Mul(assetRatesStats.StableSlope1)
			borrowAPY = assetRatesStats.StableBase.Add(multiplicationFactor)
			return borrowAPY, nil
		} else {
			utilisationNumerator := currentUtilisationRatio.Sub(assetRatesStats.UOptimal)
			utilisationDenominator := sdk.OneDec().Sub(assetRatesStats.UOptimal)
			utilisationRatio := utilisationNumerator.Quo(utilisationDenominator)
			multiplicationFactor := utilisationRatio.Mul(assetRatesStats.StableSlope2)
			borrowAPY = assetRatesStats.StableBase.Add(assetRatesStats.StableSlope1).Add(multiplicationFactor)
			return borrowAPY, nil
		}
	}
}
