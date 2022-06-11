package keeper

import (
	"github.com/comdex-official/comdex/x/lend/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	protobuftypes "github.com/gogo/protobuf/types"
)

func (k Keeper) GetCollateralAmount(ctx sdk.Context, borrowerAddr sdk.AccAddress, denom string) sdk.Coin {
	store := ctx.KVStore(k.storeKey)
	collateral := sdk.NewCoin(denom, sdk.ZeroInt())
	key := types.CreateCollateralAmountKey(borrowerAddr, denom)

	if bz := store.Get(key); bz != nil {
		err := collateral.Amount.Unmarshal(bz)
		if err != nil {
			panic(err)
		}
	}

	return collateral
}

func (k Keeper) setCollateralAmount(ctx sdk.Context, borrowerAddr sdk.AccAddress, collateral sdk.Coin) error {
	if !collateral.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidAsset, collateral.String())
	}

	if borrowerAddr.Empty() {
		return types.ErrEmptyAddress
	}

	bz, err := collateral.Amount.Marshal()
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.CreateCollateralAmountKey(borrowerAddr, collateral.Denom)

	if collateral.Amount.IsZero() {
		store.Delete(key)
	} else {
		store.Set(key, bz)
	}
	return nil
}

func (k *Keeper) SetUserLendIDHistory(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LendHistoryIdPrefix
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)
	store.Set(key, value)
}

func (k *Keeper) GetUserLendIDHistory(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = types.LendHistoryIdPrefix
		value = store.Get(key)
	)

	if value == nil {
		return 0
	}

	var id protobuftypes.UInt64Value
	k.cdc.MustUnmarshal(value, &id)

	return id.GetValue()
}

func (k *Keeper) SetLend(ctx sdk.Context, lend types.LendAsset) {
	var (
		store = k.Store(ctx)
		key   = types.LendUserKey(lend.ID)
		value = k.cdc.MustMarshal(&lend)
	)

	store.Set(key, value)
}

func (k *Keeper) GetLend(ctx sdk.Context, id uint64) (lend types.LendAsset, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.LendUserKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return lend, false
	}

	k.cdc.MustUnmarshal(value, &lend)
	return lend, true
}

func (k *Keeper) SetPoolId(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PoolIdPrefix
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)
	store.Set(key, value)
}

func (k *Keeper) GetPoolId(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = types.PoolIdPrefix
		value = store.Get(key)
	)

	if value == nil {
		return 0
	}

	var id protobuftypes.UInt64Value
	k.cdc.MustUnmarshal(value, &id)

	return id.GetValue()
}

func (k *Keeper) SetPool(ctx sdk.Context, pool types.Pool) {
	var (
		store = k.Store(ctx)
		key   = types.PoolKey(pool.PoolId)
		value = k.cdc.MustMarshal(&pool)
	)

	store.Set(key, value)
}

func (k *Keeper) GetPool(ctx sdk.Context, id uint64) (pool types.Pool, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.PoolKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return pool, false
	}

	k.cdc.MustUnmarshal(value, &pool)
	return pool, true
}

func (k *Keeper) SetAssetToPair(ctx sdk.Context, assetToPair types.AssetToPairMapping) {
	var (
		store = k.Store(ctx)
		key   = types.AssetToPairMappingKey(assetToPair.AssetId)
		value = k.cdc.MustMarshal(&assetToPair)
	)

	store.Set(key, value)
}

func (k *Keeper) GetAssetToPair(ctx sdk.Context, id uint64) (assetToPair types.AssetToPairMapping, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.AssetToPairMappingKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return assetToPair, false
	}

	k.cdc.MustUnmarshal(value, &assetToPair)
	return assetToPair, true
}

func (k *Keeper) SetLendForAddressByAsset(ctx sdk.Context, address sdk.AccAddress, assetID, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LendForAddressByAsset(address, assetID)
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)

	store.Set(key, value)
}

func (k *Keeper) HasLendForAddressByAsset(ctx sdk.Context, address sdk.AccAddress, assetID uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.LendForAddressByAsset(address, assetID)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteLendForAddressByAsset(ctx sdk.Context, address sdk.AccAddress, assetID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LendForAddressByAsset(address, assetID)
	)

	store.Delete(key)
}

func (k *Keeper) UpdateUserLendIdMapping(
	ctx sdk.Context,
	lendOwner string,
	vaultId uint64,
	isInsert bool,
) error {

	userVaults, found := k.GetUserLends(ctx, lendOwner)

	if !found && isInsert {
		userVaults = types.UserLendIdMapping{
			Owner:   lendOwner,
			LendIds: nil,
		}
	} else if !found && !isInsert {
		return types.ErrorLendOwnerNotFound
	}

	if isInsert {
		userVaults.LendIds = append(userVaults.LendIds, vaultId)
	} else {
		for index, id := range userVaults.LendIds {
			if id == vaultId {
				userVaults.LendIds = append(userVaults.LendIds[:index], userVaults.LendIds[index+1:]...)
				break
			}
		}
	}

	k.SetUserLends(ctx, userVaults)
	return nil
}

func (k *Keeper) GetUserLends(ctx sdk.Context, address string) (userVaults types.UserLendIdMapping, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.UserLendsForAddressKey(address)
		value = store.Get(key)
	)
	if value == nil {
		return userVaults, false
	}
	k.cdc.MustUnmarshal(value, &userVaults)

	return userVaults, true
}

func (k *Keeper) SetUserLends(ctx sdk.Context, userVaults types.UserLendIdMapping) {
	var (
		store = k.Store(ctx)
		key   = types.UserLendsForAddressKey(userVaults.Owner)
		value = k.cdc.MustMarshal(&userVaults)
	)
	store.Set(key, value)
}
