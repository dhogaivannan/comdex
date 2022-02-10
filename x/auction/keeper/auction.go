package keeper

import (
	"time"

	assettypes "github.com/comdex-official/comdex/x/asset/types"
	auctiontypes "github.com/comdex-official/comdex/x/auction/types"
	liquidationtypes "github.com/comdex-official/comdex/x/liquidation/types"
	vaulttypes "github.com/comdex-official/comdex/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuftypes "github.com/gogo/protobuf/types"
)

func (k Keeper) CreateNewAuctions(ctx sdk.Context) {
	locked_vaults := k.GetLockedVaults(ctx)
	for _, locked_vault := range locked_vaults {
		pair, found := k.GetPair(ctx, locked_vault.PairId)
		if !found {
			continue
		}
		assetIn, found := k.GetAsset(ctx, pair.AssetIn)
		if !found {
			continue
		}

		assetOut, found := k.GetAsset(ctx, pair.AssetOut)
		if !found {
			continue
		}
		collateralizationRatio, err := k.CalculateCollaterlizationRatio(ctx, locked_vault.AmountIn, assetIn, locked_vault.AmountOut, assetOut)
		if err != nil {
			continue
		}
		if sdk.Dec.LT(collateralizationRatio, pair.LiquidationRatio) && !locked_vault.IsAuctionInProgress {
			k.StartCollateralAuction(ctx, locked_vault, pair, assetIn, assetOut)
		}
	}
}

func (k Keeper) CloseAuctions(ctx sdk.Context) {
	collateral_auctions := k.GetCollateralAuctions(ctx)
	for _, collateral_auction := range collateral_auctions {
		if ctx.BlockTime().After(collateral_auction.EndTime) {
			k.CloseCollateralAuction(ctx, collateral_auction)
		}
	}
}

func (k Keeper) StartCollateralAuction(
	ctx sdk.Context,
	locked_vault liquidationtypes.LockedVault,
	pair assettypes.Pair,
	assetIn assettypes.Asset,
	assetOut assettypes.Asset,
) error {

	assetInPrice, found := k.GetPriceForAsset(ctx, pair.AssetIn)
	if !found {
		return assettypes.ErrorAssetDoesNotExist
	}

	assetOutPrice, found := k.GetPriceForAsset(ctx, pair.AssetOut)
	if !found {
		return assettypes.ErrorAssetDoesNotExist
	}

	liquidatedQuantity := sdk.NewDec(locked_vault.CollateralToBeAuctioned.Quo(sdk.NewDec(int64(assetInPrice))).RoundInt64())
	penaltyQuantity := liquidatedQuantity.Sub(liquidatedQuantity.Mul(sdk.NewDec(85).Quo(sdk.NewDec(100)))).RoundInt64()
	DiscountedQuantity := liquidatedQuantity.Sub(liquidatedQuantity.Mul(sdk.NewDec(95).Quo(sdk.NewDec(100)))).RoundInt64()
	AuctioningQuantity := liquidatedQuantity.Sub(sdk.NewDec(int64(penaltyQuantity + DiscountedQuantity))).RoundInt64()

	minBid := sdk.NewDec(AuctioningQuantity * int64(assetInPrice)).Quo(sdk.NewDec(int64(assetOutPrice))).Ceil().RoundInt()
	maxBid := sdk.NewDec((AuctioningQuantity + DiscountedQuantity) * int64(assetInPrice)).Quo(sdk.NewDec(int64(assetOutPrice))).Ceil().RoundInt()

	auction := auctiontypes.CollateralAuction{
		LockedVaultId:       locked_vault.LockedVaultId,
		AuctionedCollateral: sdk.NewCoin(assetIn.Denom, sdk.NewInt(AuctioningQuantity)),
		DiscountQuantity:    sdk.NewCoin(assetIn.Denom, sdk.NewInt(DiscountedQuantity)),
		Bidder:              nil,
		Bid:                 sdk.NewCoin(assetOut.Denom, sdk.NewInt(0)),
		MinBid:              sdk.NewCoin(assetOut.Denom, minBid),
		MaxBid:              sdk.NewCoin(assetOut.Denom, maxBid),
		EndTime:             ctx.BlockTime().Add(time.Hour * 6),
		Pair:                auctiontypes.Pair(pair),
	}
	auction.Id = k.GetCollateralAuctionID(ctx) + 1
	k.SetCollateralAuctionID(ctx, auction.Id)
	k.SetCollateralAuction(ctx, auction)
	k.SetFlagIsAuctionInProgress(ctx, locked_vault.LockedVaultId, true)
	return nil
}

func (k Keeper) CloseCollateralAuction(
	ctx sdk.Context,
	collateral_auction auctiontypes.CollateralAuction,
) error {

	if collateral_auction.Bidder != nil {

		assetIn, found := k.GetAsset(ctx, collateral_auction.Pair.AssetIn)
		if !found {
			return assettypes.ErrorAssetDoesNotExist
		}
		assetOut, found := k.GetAsset(ctx, collateral_auction.Pair.AssetOut)
		if !found {
			return assettypes.ErrorAssetDoesNotExist
		}

		assetInPrice, found := k.GetPriceForAsset(ctx, collateral_auction.Pair.AssetIn)
		if !found {
			return assettypes.ErrorAssetDoesNotExist
		}

		assetOutPrice, found := k.GetPriceForAsset(ctx, collateral_auction.Pair.AssetOut)
		if !found {
			return assettypes.ErrorAssetDoesNotExist
		}

		highestBidReceived := collateral_auction.Bid
		collateralQuantity := sdk.NewDec(highestBidReceived.Amount.Int64()).Mul(sdk.NewDec(int64(assetOutPrice))).Quo(sdk.NewDec(int64(assetInPrice))).RoundInt64()

		err := k.bank.SendCoinsFromModuleToAccount(ctx, vaulttypes.ModuleName, collateral_auction.Bidder, sdk.NewCoins(sdk.NewCoin(assetIn.Denom, sdk.NewInt(collateralQuantity))))
		if err != nil {
			return err
		}
		k.BurnCoin(ctx, vaulttypes.ModuleName, highestBidReceived)
		k.UpdateAssetQuantitiesInLockedVault(ctx, collateral_auction, sdk.NewInt(collateralQuantity), assetIn, highestBidReceived.Amount, assetOut)
	}
	k.SetFlagIsAuctionComplete(ctx, collateral_auction.LockedVaultId, true)
	k.SetFlagIsAuctionInProgress(ctx, collateral_auction.LockedVaultId, false)
	k.DeleteCollateralAuction(ctx, collateral_auction.Id)
	return nil
}

func (k *Keeper) GetCollateralAuctionID(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = auctiontypes.CollateralAuctionIdKey
		value = store.Get(key)
	)
	if value == nil {
		return 0
	}
	var id protobuftypes.UInt64Value
	k.cdc.MustUnmarshal(value, &id)

	return id.GetValue()
}

func (k *Keeper) SetCollateralAuctionID(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = auctiontypes.CollateralAuctionIdKey
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)

	store.Set(key, value)
}

func (k *Keeper) SetCollateralAuction(ctx sdk.Context, auction auctiontypes.CollateralAuction) {
	var (
		store = k.Store(ctx)
		key   = auctiontypes.CollateralAuctionKey(auction.Id)
		value = k.cdc.MustMarshal(&auction)
	)
	store.Set(key, value)
}

func (k *Keeper) DeleteCollateralAuction(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = auctiontypes.CollateralAuctionKey(id)
	)
	store.Delete(key)
}

func (k *Keeper) GetCollateralAuction(ctx sdk.Context, id uint64) (auction auctiontypes.CollateralAuction, found bool) {
	var (
		store = k.Store(ctx)
		key   = auctiontypes.CollateralAuctionKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return auction, false
	}

	k.cdc.MustUnmarshal(value, &auction)
	return auction, true
}

func (k *Keeper) GetCollateralAuctions(ctx sdk.Context) (auctions []auctiontypes.CollateralAuction) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, auctiontypes.CollateralAuctionKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var auction auctiontypes.CollateralAuction
		k.cdc.MustUnmarshal(iter.Value(), &auction)
		auctions = append(auctions, auction)
	}

	return auctions
}

func (k Keeper) PlaceBid(ctx sdk.Context, auctionId uint64, bidder sdk.AccAddress, bid sdk.Coin) error {
	auction, found := k.GetCollateralAuction(ctx, auctionId)
	if !found {
		return auctiontypes.ErrorInvalidAuctionId
	}
	if bid.Denom != auction.MinBid.Denom {
		return auctiontypes.ErrorInvalidBiddingDenom
	}
	if bid.Amount.LT(auction.MinBid.Amount) {
		return auctiontypes.ErrorLowBidAmount
	}
	if bid.Amount.GT(auction.MaxBid.Amount) {
		return auctiontypes.ErrorMaxBidAmount
	}
	if bid.Amount.LT(auction.Bid.Amount.Add(sdk.NewInt(1))) {
		return auctiontypes.ErrorBidAlreadyExists
	}
	err := k.SendCoinsFromAccountToModule(ctx, bidder, liquidationtypes.ModuleName, sdk.NewCoins(bid))
	if err != nil {
		return err
	}
	err = k.bank.SendCoinsFromModuleToAccount(ctx, liquidationtypes.ModuleName, auction.Bidder, sdk.NewCoins(auction.Bid))
	if err != nil {
		return err
	}
	auction.Bidder = bidder
	auction.Bid = bid
	k.SetCollateralAuction(ctx, auction)
	return nil
}