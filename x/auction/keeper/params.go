package keeper

import (
	"github.com/comdex-official/comdex/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k *Keeper) LiquidationPenaltyPercent(ctx sdk.Context) (s string) {
	k.paramstore.Get(ctx, types.KeyLiquidationPenaltyPercent, &s)
	return
}

func (k *Keeper) AuctionDiscountPercent(ctx sdk.Context) (s string) {
	k.paramstore.Get(ctx, types.KeyAuctionDiscountPercent, &s)
	return
}

func (k *Keeper) AuctionDurationHours(ctx sdk.Context) (s uint64) {
	k.paramstore.Get(ctx, types.KeyAuctionDurationSeconds, &s)
	return
}

func (k *Keeper) AuctionDecreasePercentage(ctx sdk.Context) (s sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyDebtMintTokenDecreasePercentage, &s)
	return
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.LiquidationPenaltyPercent(ctx),
		k.AuctionDiscountPercent(ctx),
		k.AuctionDurationHours(ctx),
		k.AuctionDecreasePercentage(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
