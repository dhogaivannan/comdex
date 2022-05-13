package keeper

import (
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuftypes "github.com/gogo/protobuf/types"
	"strconv"

	"github.com/comdex-official/comdex/x/vault/types"
)

func (k *Keeper) SetID(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.IDKey
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)

	store.Set(key, value)
}

func (k *Keeper) GetID(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = types.IDKey
		value = store.Get(key)
	)

	if value == nil {
		return 0
	}

	var id protobuftypes.UInt64Value
	k.cdc.MustUnmarshal(value, &id)

	return id.GetValue()
}

func (k *Keeper) SetAppVaultTypeIdForId(ctx sdk.Context, appMapping types.AppMapping) {
	var (
		store = k.Store(ctx)
		key   = types.GetAppMappingIdPrefixKey(appMapping.AppMappingId)
	)
	bz := k.cdc.MustMarshal(&appMapping)
	store.Set(key, bz)
}

func (k *Keeper) GetAppMappingForAppMappingId(ctx sdk.Context, appMappingId uint64) (appMapping types.AppMapping, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.GetAppMappingIdPrefixKey(appMappingId)
		value = store.Get(key)
	)
	if value == nil {
		return appMapping, false
	}
	k.cdc.MustUnmarshal(value, &appMapping)
	return appMapping, true
}

func (k *Keeper) SetLookUpTableForAppMappingId(ctx sdk.Context, lookUpTable types.LookupTable, appMappingId uint64) {
	var (
		store = k.Store(ctx)
		key   = types.GetLookUpTablePrefixKey(appMappingId)
	)
	lookUpTable.Counter = lookUpTable.Counter + 1
	bz := k.cdc.MustMarshal(&lookUpTable)
	store.Set(key, bz)
}

func (k *Keeper) GetLookUpTableForAppMappingId(ctx sdk.Context, appMappingId uint64) (lookUpTable types.LookupTable) {
	var (
		store = k.Store(ctx)
		key   = types.GetLookUpTablePrefixKey(appMappingId)
		value = store.Get(key)
	)
	if value == nil {
		return types.LookupTable{AppVaultIds: make([]string, 0), Counter: 0}
	}
	k.cdc.MustUnmarshal(value, &lookUpTable)
	return lookUpTable
}

func (k *Keeper) SetUserVaultIdMapping(ctx sdk.Context, userVaultIdMapping types.UserVaultIdMapping, owner sdk.AccAddress) {
	var (
		store = k.Store(ctx)
		key   = types.GetUserVaultIdPrefixKey(owner)
	)
	bz := k.cdc.MustMarshal(&userVaultIdMapping)
	store.Set(key, bz)
}

func (k *Keeper) GetUserVaultIdMapping(ctx sdk.Context, owner sdk.AccAddress) (userVaultIdMapping types.UserVaultIdMapping, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.GetUserVaultIdPrefixKey(owner)
		value = store.Get(key)
	)
	if value == nil {
		return userVaultIdMapping, false
	}
	k.cdc.MustUnmarshal(value, &userVaultIdMapping)
	return userVaultIdMapping, true
}

func (k *Keeper) SetTokenMintStatistics(ctx sdk.Context, tokenMintStatistics types.TokenMintStatistics) {
	var (
		store = k.Store(ctx)
		key   = types.GetTokenMintPrefixKey(tokenMintStatistics.CollateralDenom)
	)
	bz := k.cdc.MustMarshal(&tokenMintStatistics)
	store.Set(key, bz)
}

func (k *Keeper) GetTokenMintStatistics(ctx sdk.Context, collateralDenom string) (tokenMintStatistics types.TokenMintStatistics, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.GetTokenMintPrefixKey(collateralDenom)
		value = store.Get(key)
	)
	if value == nil {
		return tokenMintStatistics, false
	}
	k.cdc.MustUnmarshal(value, &tokenMintStatistics)
	return tokenMintStatistics, true
}

func (k *Keeper) SetVault(ctx sdk.Context, vault types.Vault) {
	var (
		store = k.Store(ctx)
		key   = types.VaultKey(vault.AppVaultTypeId)
		value = k.cdc.MustMarshal(&vault)
	)

	store.Set(key, value)
}

func (k *Keeper) GetVault(ctx sdk.Context, appVaultTypeId string) (vault types.Vault, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.VaultKey(appVaultTypeId)
		value = store.Get(key)
	)

	if value == nil {
		return vault, false
	}

	k.cdc.MustUnmarshal(value, &vault)
	return vault, true
}

func (k *Keeper) DeleteVault(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.VaultKey(id)
	)

	store.Delete(key)
}

func (k *Keeper) GetVaults(ctx sdk.Context) (vaults []types.Vault) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.VaultKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var vault types.Vault
		k.cdc.MustUnmarshal(iter.Value(), &vault)
		vaults = append(vaults, vault)
	}

	return vaults
}

func (k *Keeper) SetVaultForAddressByPair(ctx sdk.Context, address sdk.AccAddress, appVaultTypeId string, pairID uint64, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.VaultForAddressByAppAndPair(address, appVaultTypeId, pairID)
		value = k.cdc.MustMarshal(
			&protobuftypes.UInt64Value{
				Value: id,
			},
		)
	)

	store.Set(key, value)
}

func (k *Keeper) HasVaultForAddressByPair(ctx sdk.Context, address sdk.AccAddress, appVaultTypeId string, pairID uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.VaultForAddressByAppAndPair(address, appVaultTypeId, pairID)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteVaultForAddressByPair(ctx sdk.Context, address sdk.AccAddress, appVaultTypeId string, pairID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.VaultForAddressByAppAndPair(address, appVaultTypeId, pairID)
	)

	store.Delete(key)
}

func (k *Keeper) VerifyCollaterlizationRatio(
	ctx sdk.Context,
	amountIn sdk.Int,
	assetIn assettypes.Asset,
	amountOut sdk.Int,
	assetOut assettypes.Asset,
) error {

	return nil
}

func (k *Keeper) CalculateCollaterlizationRatio(
	ctx sdk.Context,
	amountIn sdk.Int,
	assetIn assettypes.Asset,
	amountOut sdk.Int,
	assetOut assettypes.Asset,
) (sdk.Dec, error) {

	assetInPrice, found := k.GetPriceForAsset(ctx, assetIn.Id)
	if !found {
		return sdk.ZeroDec(), types.ErrorPriceDoesNotExist
	}

	assetOutPrice, found := k.GetPriceForAsset(ctx, assetOut.Id)
	if !found {
		return sdk.ZeroDec(), types.ErrorPriceDoesNotExist
	}

	totalIn := amountIn.Mul(sdk.NewIntFromUint64(assetInPrice)).ToDec()
	if totalIn.LTE(sdk.ZeroDec()) {
		return sdk.ZeroDec(), types.ErrorInvalidAmountIn
	}

	totalOut := amountOut.Mul(sdk.NewIntFromUint64(assetOutPrice)).ToDec()
	if totalOut.LTE(sdk.ZeroDec()) {
		return sdk.ZeroDec(), types.ErrorInvalidAmountOut
	}

	return totalIn.Quo(totalOut), nil
}

func (k *Keeper) UpdateUserVaultIdMapping(
	ctx sdk.Context,
	msg *types.MsgCreateRequest,
	AppVaultTypeIdWithNumber string,
	isInsert bool,
) error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil
	}
	userVaultIdMapping, found := k.GetUserVaultIdMapping(ctx, from)
	if !found && isInsert {
		userVaults := types.UserVaultIdMapping{
			Owner:        from.String(),
			UserVaultIds: nil,
		}
		vaultToAppMapping := types.VaultToAppMapping{AppMappingId: msg.AppMappingId, AppVaultTypeId: make([]string, 0)}
		vaultToAppMapping.AppVaultTypeId = append(vaultToAppMapping.AppVaultTypeId, AppVaultTypeIdWithNumber)
		userVaults.UserVaultIds = append(userVaults.UserVaultIds, &vaultToAppMapping)
	} else if !found && !isInsert {
		return types.ErrorVaultOwnerNotFound
	}
	//found owner and insert vaultid
	if found && isInsert {
		foundAppVaultType := false
		for _, element := range userVaultIdMapping.UserVaultIds {
			if element.AppMappingId == msg.AppMappingId {
				foundAppVaultType = true
				element.AppVaultTypeId = append(element.AppVaultTypeId, AppVaultTypeIdWithNumber)
				break
			}
		}
		if !foundAppVaultType {
			vaultToAppMapping := types.VaultToAppMapping{AppMappingId: msg.AppMappingId, AppVaultTypeId: make([]string, 0)}
			vaultToAppMapping.AppVaultTypeId = append(vaultToAppMapping.AppVaultTypeId, AppVaultTypeIdWithNumber)
			userVaultIdMapping.UserVaultIds = append(userVaultIdMapping.UserVaultIds, &vaultToAppMapping)
		}
	}
	//found owner and delete vaultid
	if found && !isInsert {
		for _, element1 := range userVaultIdMapping.UserVaultIds {
			if element1.AppMappingId == msg.AppMappingId {
				for index, element2 := range element1.AppVaultTypeId {
					if element2 == AppVaultTypeIdWithNumber {
						element1.AppVaultTypeId = append(element1.AppVaultTypeId[:index], element1.AppVaultTypeId[index+1:]...)
					}
				}
				break
			}
		}
	}
	k.SetUserVaultIdMapping(ctx, userVaultIdMapping, from)
	return nil
}

func (k *Keeper) UpdateCollateralVaultIdMapping(
	ctx sdk.Context,
	assetInDenom string,
	assetOutDenom string,
	vaultId uint64,
	isInsert bool,
) error {

	return nil
}

func (k *Keeper) GetCAssetTotalValueMintedForCollateral(ctx sdk.Context, collateralType assettypes.Asset) sdk.Dec {
	mintStatistics, found := k.GetCAssetMintRecords(ctx, collateralType.Denom)
	if !found {
		return sdk.NewDec(0)
	}

	availableAssets := k.GetAssets(ctx)
	cAssetDenomIdMap := make(map[string]uint64)
	for _, asset := range availableAssets {
		cAssetDenomIdMap[asset.Denom] = asset.Id
	}

	totalValueCassetMinted := sdk.NewDec(0)

	for cAssetDenom, quantity := range mintStatistics.MintedAssets {
		assetPrice, found := k.GetPriceForAsset(ctx, cAssetDenomIdMap[cAssetDenom])
		if found {
			totalValueCassetMinted = totalValueCassetMinted.Add(sdk.NewDec(int64(quantity)).Quo(sdk.NewDec(1000000)).Mul(sdk.NewDec(int64(assetPrice)).Quo(sdk.NewDec(1000000))))
		}
	}
	return totalValueCassetMinted
}

func (k *Keeper) GetCAssetMintRecords(ctx sdk.Context, collateralDenom string) (mintRecords types.TokenMintingStatisics, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.GetCAssetMintRecordsKey(collateralDenom)
		value = store.Get(key)
	)
	if value == nil {
		return mintRecords, false
	}
	k.cdc.MustUnmarshal(value, &mintRecords)

	return mintRecords, true
}

func (k *Keeper) SetCAssetMintRecords(ctx sdk.Context, mintRecords types.TokenMintingStatisics) {
	var (
		store = k.Store(ctx)
		key   = types.GetCAssetMintRecordsKey(mintRecords.TokensMinted)
		value = k.cdc.MustMarshal(&mintRecords)
	)
	store.Set(key, value)
}

func (k *Keeper) CreateNewVault(
	ctx sdk.Context,
	msg *types.MsgCreateRequest,
	AppVaultTypeId string,
) error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return err
	}
	var (
		lookUpTable              = k.GetLookUpTableForAppMappingId(ctx, msg.AppMappingId)
		AppVaultTypeIdWithNumber = AppVaultTypeId + strconv.FormatUint(lookUpTable.Counter, 10)
		vault                    = types.Vault{
			AppVaultTypeId:        AppVaultTypeIdWithNumber,
			PairID:                msg.ExtendedPairId,
			Owner:                 msg.From,
			AmountIn:              msg.AmountIn,
			AmountOut:             msg.AmountOut,
			CreatedAt:             ctx.BlockTime(),
			InterestAccumulated:   sdk.NewDec(0),
			OpeningFeeAccumulated: sdk.NewDec(0),
			ClosingFeeAccumulated: sdk.NewDec(0),
			RewardsAccumulated:    sdk.NewDec(0),
		}
	)
	k.SetVaultForAddressByPair(ctx, from, vault.AppVaultTypeId, vault.PairID, lookUpTable.Counter)
	lookUpTable.AppVaultIds = append(lookUpTable.AppVaultIds, AppVaultTypeIdWithNumber)
	k.SetLookUpTableForAppMappingId(ctx, lookUpTable, msg.AppMappingId)
	k.SetVault(ctx, vault)

	k.UpdateUserVaultIdMapping(ctx, msg, AppVaultTypeIdWithNumber, true)
	k.UpdateCollateralVaultIdMapping(ctx, assetIn.Denom, assetOut.Denom, vault.ID, true)
	return nil
}
