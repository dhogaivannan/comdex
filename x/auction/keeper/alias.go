package keeper

import (
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	"github.com/comdex-official/comdex/x/collector/types"
	liquidationtypes "github.com/comdex-official/comdex/x/liquidation/types"
	vaultttypes "github.com/comdex-official/comdex/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k *Keeper) GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI {
	return k.account.GetModuleAccount(ctx, name)
}

func (k *Keeper) GetModuleAddress(ctx sdk.Context, name string) sdk.AccAddress {
	return k.account.GetModuleAddress(name)
}
func (k *Keeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	return k.bank.GetBalance(ctx, addr, denom)
}

func (k *Keeper) MintCoins(ctx sdk.Context, name string, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.bank.MintCoins(ctx, name, sdk.NewCoins(coin))
}

func (k *Keeper) BurnCoins(ctx sdk.Context, name string, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.bank.BurnCoins(ctx, name, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinsFromModuleToModule(ctx sdk.Context, senderModule string, recipientModule string, amt sdk.Coins) error {
	return k.bank.SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt)
}
func (k *Keeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	return k.bank.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}
func (k *Keeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	return k.bank.SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt)
}

func (k *Keeper) GetPriceForAsset(ctx sdk.Context, id uint64) (uint64, bool) {
	return k.market.GetPriceForAsset(ctx, id)
}

func (k *Keeper) GetLockedVaults(ctx sdk.Context) (locked_vaults []liquidationtypes.LockedVault) {
	return k.liquidation.GetLockedVaults(ctx)
}

func (k *Keeper) GetPair(ctx sdk.Context, id uint64) (assettypes.Pair, bool) {
	return k.asset.GetPair(ctx, id)
}

func (k *Keeper) GetAsset(ctx sdk.Context, id uint64) (assettypes.Asset, bool) {
	return k.asset.GetAsset(ctx, id)
}

func (k *Keeper) SetFlagIsAuctionInProgress(ctx sdk.Context, id uint64, flag bool) error {
	return k.liquidation.SetFlagIsAuctionInProgress(ctx, id, flag)
}

func (k *Keeper) SetFlagIsAuctionComplete(ctx sdk.Context, id uint64, flag bool) error {
	return k.liquidation.SetFlagIsAuctionComplete(ctx, id, flag)
}

//func (k *Keeper) UpdateAssetQuantitiesInLockedVault(ctx sdk.Context, collateral_auction auctiontypes.CollateralAuction, amountIn sdk.Int, assetIn assettypes.Asset, amountOut sdk.Int, assetOut assettypes.Asset) error {
//	return k.liquidation.UpdateAssetQuantitiesInLockedVault(ctx, collateral_auction, amountIn, assetIn, amountOut, assetOut)
//}

//func (k *Keeper) CalculateCollaterlizationRatio(
//	ctx sdk.Context,
//	amountIn sdk.Int,
//	assetIn assettypes.Asset,
//	amountOut sdk.Int,
//	assetOut assettypes.Asset,
//) (sdk.Dec, error) {
//	return k.vault.CalculateCollaterlizationRatio(ctx, amountIn, assetIn, amountOut, assetOut)
//}
//
//func (k *Keeper) BurnCAssets(ctx sdk.Context, moduleName string, collateralDenom string, denom string, amount sdk.Int) error {
//	return k.vault.BurnCAssets(ctx, moduleName, collateralDenom, denom, amount)
//}

func (k *Keeper) GetAppidToAssetCollectorMapping(ctx sdk.Context, app_id uint64) (appAssetCollectorData types.AppIdToAssetCollectorMapping, found bool) {
	return k.collector.GetAppidToAssetCollectorMapping(ctx, app_id)
}

func (k *Keeper) UpdateCollector(ctx sdk.Context, appId, asset_id uint64, CollectedStabilityFee, CollectedClosingFee, CollectedOpeningFee, LiquidationRewardsCollected sdk.Int) error {
	return k.collector.UpdateCollector(ctx, appId, asset_id, CollectedStabilityFee, CollectedClosingFee, CollectedOpeningFee, LiquidationRewardsCollected)
}

func (k *Keeper) SetCollectorLookupTable(ctx sdk.Context, records ...types.CollectorLookupTable) error {
	return k.collector.SetCollectorLookupTable(ctx, records...)
}

func (k *Keeper) GetCollectorLookupTable(ctx sdk.Context, app_id uint64) (collectorLookup types.CollectorLookup, found bool) {
	return k.collector.GetCollectorLookupTable(ctx, app_id)
}

// func (k *Keeper) SetCollectorAuctionLookupTable(ctx sdk.Context, records ...types.CollectorAuctionLookupTable) error {
// 	return k.collector.SetCollectorAuctionLookupTable(ctx, records...)
// }

// func (k *Keeper) GetCollectorAuctionLookupTable(ctx sdk.Context, app_id uint64) (appIdToAuctionData types.CollectorAuctionLookupTable, found bool) {
// 	return k.collector.GetCollectorAuctionLookupTable(ctx, app_id)
// }
func (k *Keeper) GetNetFeeCollectedData(ctx sdk.Context, app_id uint64) (netFeeData types.NetFeeCollectedData, found bool) {
	return k.collector.GetNetFeeCollectedData(ctx, app_id)
}
func (k *Keeper) GetApps(ctx sdk.Context) (apps []assettypes.AppMapping, found bool) {
	return k.asset.GetApps(ctx)
}
func (k *Keeper) GetApp(ctx sdk.Context, id uint64) (app assettypes.AppMapping, found bool) {
	return k.asset.GetApp(ctx, id)
}

func (k *Keeper) MintNewTokensForApp(ctx sdk.Context, appMappingId uint64, assetId uint64, address string, amount sdk.Int) error {
	return k.tokenmint.MintNewTokensForApp(ctx, appMappingId, assetId, address, amount)
}

func (k *Keeper) BurnTokensForApp(ctx sdk.Context, appMappingId uint64, assetId uint64, amount sdk.Int) error {
	return k.tokenmint.BurnTokensForApp(ctx, appMappingId, assetId, amount)
}

func (k *Keeper) GetAmountFromCollector(ctx sdk.Context, appId, asset_id uint64, amount sdk.Int) (sdk.Int, error) {
	return k.collector.GetAmountFromCollector(ctx, appId, asset_id, amount)
}

func (k *Keeper) SetNetFeeCollectedData(ctx sdk.Context, app_id, asset_id uint64, fee sdk.Int) error {
	return k.collector.SetNetFeeCollectedData(ctx, app_id, asset_id, fee)
}

func (k *Keeper) GetLockedVault(ctx sdk.Context, id uint64) (locked_vault liquidationtypes.LockedVault, found bool) {
	return k.liquidation.GetLockedVault(ctx, id)
}

func (k *Keeper) SetLockedVault(ctx sdk.Context, locked_vault liquidationtypes.LockedVault) {
	k.liquidation.SetLockedVault(ctx, locked_vault)
}

func (k *Keeper) GetPairsVault(ctx sdk.Context, id uint64) (pairs assettypes.ExtendedPairVault, found bool) {
	return k.asset.GetPairsVault(ctx, id)
}

func (k *Keeper) GetAppExtendedPairVaultMapping(ctx sdk.Context, appMappingId uint64) (appExtendedPairVaultData vaultttypes.AppExtendedPairVaultMapping, found bool) {
	return k.vault.GetAppExtendedPairVaultMapping(ctx, appMappingId)
}

func (k *Keeper) SetAppExtendedPairVaultMapping(ctx sdk.Context, appExtendedPairVaultData vaultttypes.AppExtendedPairVaultMapping) error {
	return k.vault.SetAppExtendedPairVaultMapping(ctx, appExtendedPairVaultData)
}

func (k *Keeper) GetAuctionMappingForApp(ctx sdk.Context, appId uint64) (collectorAuctionLookupTable types.CollectorAuctionLookupTable, found bool) {
	return k.collector.GetAuctionMappingForApp(ctx, appId)
}
func (k *Keeper) SetAuctionMappingForApp(ctx sdk.Context, records ...types.CollectorAuctionLookupTable) error {
	return k.collector.SetAuctionMappingForApp(ctx, records...)
}

func (k *Keeper) UpdateCollateralLockedAmountLockerMapping(ctx sdk.Context, valutLookupData vaultttypes.AppExtendedPairVaultMapping, extendedPairId uint64, amount sdk.Int, changeType bool) {
	k.vault.UpdateCollateralLockedAmountLockerMapping(ctx, valutLookupData, extendedPairId, amount, changeType)
}
