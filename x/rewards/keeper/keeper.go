package keeper

import (
	"fmt"
	"github.com/comdex-official/comdex/x/rewards/expected"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/comdex-official/comdex/x/rewards/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
		locker     expected.LockerKeeper
		collector  expected.CollectorKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	locker expected.LockerKeeper,
	collector expected.CollectorKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		locker:     locker,
		collector:  collector,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func uint64InSlice(a uint64, list []uint64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (k Keeper) WhitelistAsset(ctx sdk.Context, appMappingId uint64, assetId []uint64) error {
	lockerAssets, _ := k.locker.GetLockerProductAssetMapping(ctx, appMappingId)
	for i := range assetId {
		found := uint64InSlice(assetId[i], lockerAssets.AssetIds)
		if !found {
			return types.ErrAssetIdDoesNotExist
		}
	}

	internalRewards := types.InternalRewards{
		App_mapping_ID: appMappingId,
		Asset_ID:       assetId,
	}

	k.SetReward(ctx, internalRewards)
	return nil
}

func (k Keeper) RemoveWhitelistAsset(ctx sdk.Context, appMappingId uint64, assetId uint64) error {

	rewards, found := k.GetReward(ctx, appMappingId)
	if found != true {
		return nil
	}
	var newAssetIds []uint64
	fmt.Println(rewards.Asset_ID)
	for i := range rewards.Asset_ID {
		if assetId != rewards.Asset_ID[i] {
			newAssetId := rewards.Asset_ID[i]
			newAssetIds = append(newAssetIds, newAssetId)
		}

	}
	newRewards := types.InternalRewards{
		App_mapping_ID: appMappingId,
		Asset_ID:       newAssetIds,
	}
	k.SetReward(ctx, newRewards)
	return nil
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}
