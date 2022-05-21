package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/types/time"

	"github.com/comdex-official/comdex/x/vault/types"
)

var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	Keeper
}

func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &msgServer{
		Keeper: keeper,
	}
}

//Creating a new CDP
func (k *msgServer) MsgCreate(c context.Context, msg *types.MsgCreateRequest) (*types.MsgCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	//Checking if extended pair exists
	extended_pair_vault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultID)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extended_pair_vault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	assetInData, found := k.GetAsset(ctx, pairData.AssetIn)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	assetOutData, found := k.GetAsset(ctx, pairData.AssetOut)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}

	//Checking if app_mapping_id exists
	app_mapping, found := k.GetApp(ctx, msg.AppMappingId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}

	//Checking if the app_mapping_id in the msg_create & extended_pair_vault_are same or not
	if app_mapping.Id != extended_pair_vault.AppMappingId {
		return nil, types.ErrorAppMappingIdMismatch
	}

	//Converting user address for bank transaction
	depositor_address, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Checking if this is a stableMint pair or not  -- stableMintPair == psmPair
	if extended_pair_vault.IsPsmPair {
		return nil, types.ErrorCannotCreateStableMintVault
	}
	//Checking
	if !extended_pair_vault.IsVaultActive {
		return nil, types.ErrorVaultCreationInactive

	}
	//Checking UserMapping Details
	//Checking if user mapping exists
	//if does then check app to extendedPair mapping has any vault key
	//if it does throw error
	user_vault_extendedPair_mapping, user_exists := k.GetUserVaultExtendedPairMapping(ctx, msg.From)

	if user_exists {
		_, already_exists := k.CheckUserAppToExtendedPairMapping(ctx, user_vault_extendedPair_mapping, extended_pair_vault.Id, app_mapping.Id)
		if already_exists {
			return nil, types.ErrorUserVaultAlreadyExists

		}

	}
	//Call CheckAppExtendedPairVaultMapping function to get counter - it also initialised the kv store if app_mapping_id does not exists, or extended_pair_vault_id does not exists.

	counter_val, token_minted_statistics := k.CheckAppExtendedPairVaultMapping(ctx, app_mapping.Id, extended_pair_vault.Id)

	//Check debt Floor
	if !msg.AmountOut.GTE(extended_pair_vault.DebtFloor) {

		return nil, types.ErrorAmountOutLessThanDebtFloor
	}
	//Check Debt Ceil
	current_minted_statistics := token_minted_statistics.Add(msg.AmountOut)

	if current_minted_statistics.GTE(extended_pair_vault.DebtCeiling) {
		return nil, types.ErrorAmountOutGreaterThanDebtCeiling
	}

	//Calculate CR - make necessary changes to the calculate collateralization function
	if err := k.VerifyCollaterlizationRatio(ctx, extended_pair_vault.Id, msg.AmountIn, msg.AmountOut, extended_pair_vault.MinCr); err != nil {
		return nil, err
	}

	//Take amount from user
	if err := k.SendCoinFromAccountToModule(ctx, depositor_address, types.ModuleName, sdk.NewCoin(assetInData.Denom, msg.AmountIn)); err != nil {
		return nil, err
	}
	//Mint Tokens for user

	if err := k.MintCoin(ctx, types.ModuleName, sdk.NewCoin(assetOutData.Denom, msg.AmountOut)); err != nil {
		return nil, err
	}

	//Calculating Closing Fee
	//----Done inside the vault-----//

	//Send Fees to Accumulator
	//Deducting Opening Fee if 0 opening fee then act accordingly
	if extended_pair_vault.CreationFee.IsZero() {

		//Send Rest to user
		if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositor_address, sdk.NewCoin(assetOutData.Denom, msg.AmountOut)); err != nil {
			return nil, err
		}

	} else {
		//If not zero deduct send to collector//////////
		//
		//			COLLECTOR FUNCTION
		//
		//
		/////////////////////////////////////////////////

		collectorShare := (msg.AmountOut.Mul(sdk.Int(extended_pair_vault.CreationFee))).Quo(sdk.NewInt(100))

		// and send the rest to the user
		amountToUser := msg.AmountOut.Sub(collectorShare)
		if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositor_address, sdk.NewCoin(assetOutData.Denom, amountToUser)); err != nil {
			return nil, err
		}

	}

	//If all correct  create vault
	zero_val := sdk.ZeroInt()
	var new_vault types.Vault
	updated_counter := counter_val + 1
	new_vault.Id = app_mapping.ShortName + strconv.FormatUint(updated_counter, 10)
	new_vault.AmountIn = msg.AmountIn

	closingFeeVal := (sdk.Dec(msg.AmountOut).Mul((extended_pair_vault.ClosingFee))).Quo(sdk.NewDec(100))

	new_vault.ClosingFeeAccumulated = (*sdk.Int)(&closingFeeVal)
	new_vault.AmountOut = msg.AmountOut
	new_vault.AppMappingId = app_mapping.Id
	new_vault.InterestAccumulated = &zero_val
	new_vault.Owner = msg.From
	new_vault.CreatedAt = time.Now()
	new_vault.ExtendedPairVaultID = extended_pair_vault.Id

	k.SetVault(ctx, new_vault)

	//Update mapping data - take proper approach
	// lookup table already exists
	//only need to update counter and token statistics value
	k.UpdateAppExtendedPairVaultMappingDataOnMsgCreate(ctx, updated_counter, new_vault)

	//update user data
	//Check and update - similar fashion as Locker module
	user_vault_extendedPair_mapping_data, user_exists := k.GetUserVaultExtendedPairMapping(ctx, msg.From)
	if !user_exists {
		//UserData does not exists
		//Create a new instance
		var user_mapping_data types.UserVaultAssetMapping
		var user_app_data types.VaultToAppMapping
		var user_extendedPair_data types.ExtendedPairToVaultMapping

		user_extendedPair_data.ExtendedPairId = new_vault.ExtendedPairVaultID
		user_extendedPair_data.VaultId = new_vault.Id
		user_app_data.AppMappingId = app_mapping.Id
		user_app_data.UserExtendedPairVault = append(user_app_data.UserExtendedPairVault, &user_extendedPair_data)
		user_mapping_data.Owner = msg.From
		user_mapping_data.UserVaultApp = append(user_mapping_data.UserVaultApp, &user_app_data)

		k.SetUserVaultExtendedPairMapping(ctx, user_mapping_data)
	} else {
		///Check if user app_mapping data exits

		app_exists := k.CheckUserToAppMapping(ctx, user_vault_extendedPair_mapping_data, app_mapping.Id)
		if app_exists {

			//User has the app_mapping added
			//So only need to add the locker id with asset
			var user_extendedPair_data types.ExtendedPairToVaultMapping
			user_extendedPair_data.VaultId = new_vault.Id
			user_extendedPair_data.ExtendedPairId = new_vault.ExtendedPairVaultID

			for _, appData := range user_vault_extendedPair_mapping_data.UserVaultApp {
				if appData.AppMappingId == app_mapping.Id {

					appData.UserExtendedPairVault = append(appData.UserExtendedPairVault, &user_extendedPair_data)
				}

			}
			k.SetUserVaultExtendedPairMapping(ctx, user_vault_extendedPair_mapping_data)

		} else {
			//Will need to create new app and add it to the user
			var user_app_data types.VaultToAppMapping
			var user_extendedPair_data types.ExtendedPairToVaultMapping

			user_extendedPair_data.ExtendedPairId = new_vault.ExtendedPairVaultID
			user_extendedPair_data.VaultId = new_vault.Id
			user_app_data.AppMappingId = app_mapping.Id
			user_app_data.UserExtendedPairVault = append(user_app_data.UserExtendedPairVault, &user_extendedPair_data)
			user_vault_extendedPair_mapping_data.UserVaultApp = append(user_vault_extendedPair_mapping_data.UserVaultApp, &user_app_data)
			k.SetUserVaultExtendedPairMapping(ctx, user_vault_extendedPair_mapping_data)

		}

	}

	return &types.MsgCreateResponse{}, nil
}

//Only for depositing new collateral
func (k *msgServer) MsgDeposit(c context.Context, msg *types.MsgDepositRequest) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	depositor, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	//checks if extended pair exists
	extended_pair_vault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultID)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extended_pair_vault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	assetInData, found := k.GetAsset(ctx, pairData.AssetIn)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}

	//Checking if app_mapping_id exists
	app_mapping, found := k.GetApp(ctx, msg.AppMappingId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	//Checking if vault acccess disabled
	if !extended_pair_vault.IsVaultActive {
		return nil, types.ErrorVaultInactive

	}

	//Checking if the app_mapping_id in the msg_create & extended_pair_vault_are same or not
	if app_mapping.Id != extended_pair_vault.AppMappingId {
		return nil, types.ErrorAppMappingIdMismatch
	}

	userVault, found := k.GetVault(ctx, msg.UserVaultID)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}
	if userVault.Owner != msg.From {

		return nil, types.ErrVaultAccessUnauthorised
	}

	if app_mapping.Id != userVault.AppMappingId {
		return nil, types.ErrorInvalidAppMappingData
	}
	if extended_pair_vault.Id != userVault.ExtendedPairVaultID {
		return nil, types.ErrorInvalidExtendedPairMappingData
	}

	userVault.AmountIn = userVault.AmountIn.Add(msg.Amount)
	if !userVault.AmountIn.IsPositive() {
		return nil, types.ErrorInvalidAmount
	}

	if err := k.SendCoinFromAccountToModule(ctx, depositor, types.ModuleName, sdk.NewCoin(assetInData.Denom, msg.Amount)); err != nil {
		return nil, err
	}

	k.SetVault(ctx, userVault)
	//Updating appExtendedPairvaultMappingData data -
	appExtendedPairVaultData, _ := k.GetAppExtendedPairVaultMapping(ctx, app_mapping.Id)
	k.UpdateCollateralLockedAmountLockerMapping(ctx, appExtendedPairVaultData, extended_pair_vault.Id, msg.Amount, true)

	return &types.MsgDepositResponse{}, nil

}

//Withdrawing collateral
func (k *msgServer) MsgWithdraw(c context.Context, msg *types.MsgWithdrawRequest) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	depositor, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	//checks if extended pair exists
	extended_pair_vault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultID)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extended_pair_vault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	assetInData, found := k.GetAsset(ctx, pairData.AssetIn)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	// assetOutData, found := k.GetAsset(ctx, pairData.AssetOut)
	// if !found {
	// 	return nil, types.ErrorAssetDoesNotExist
	// }

	//Checking if app_mapping_id exists
	app_mapping, found := k.GetApp(ctx, msg.AppMappingId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	//Checking if vault acccess disabled
	if !extended_pair_vault.IsVaultActive {
		return nil, types.ErrorVaultInactive

	}
	//Checking if the app_mapping_id in the msg_create & extended_pair_vault_are same or not
	if app_mapping.Id != extended_pair_vault.AppMappingId {
		return nil, types.ErrorAppMappingIdMismatch
	}

	userVault, found := k.GetVault(ctx, msg.UserVaultID)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}
	if userVault.Owner != msg.From {

		return nil, types.ErrVaultAccessUnauthorised
	}

	if app_mapping.Id != userVault.AppMappingId {

		return nil, types.ErrorInvalidAppMappingData

	}
	if extended_pair_vault.Id != userVault.ExtendedPairVaultID {

		return nil, types.ErrorInvalidExtendedPairMappingData
	}

	userVault.AmountIn = userVault.AmountIn.Sub(msg.Amount)
	if !userVault.AmountIn.IsPositive() {
		return nil, types.ErrorInvalidAmount
	}

	totalDebtCalculation := userVault.AmountOut.Add(*userVault.InterestAccumulated)
	totalDebtCalculation = totalDebtCalculation.Add(*userVault.ClosingFeeAccumulated)

	//Calculate CR - make necessary changes to the calculate collateralization function
	if err := k.VerifyCollaterlizationRatio(ctx, extended_pair_vault.Id, userVault.AmountIn, totalDebtCalculation, extended_pair_vault.MinCr); err != nil {
		return nil, err
	}

	if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositor, sdk.NewCoin(assetInData.Denom, msg.Amount)); err != nil {
		return nil, err
	}

	k.SetVault(ctx, userVault)

	//Updating appExtendedPairVaultMappingData
	appExtendedPairVaultData, _ := k.GetAppExtendedPairVaultMapping(ctx, app_mapping.Id)
	k.UpdateCollateralLockedAmountLockerMapping(ctx, appExtendedPairVaultData, extended_pair_vault.Id, msg.Amount, false)

	return &types.MsgWithdrawResponse{}, nil
}

//To borrow more amount
func (k *msgServer) MsgDraw(c context.Context, msg *types.MsgDrawRequest) (*types.MsgDrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	depositor, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	//checks if extended pair exists
	extended_pair_vault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultID)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extended_pair_vault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	// assetInData, found := k.GetAsset(ctx, pairData.AssetIn)
	// if !found {
	// 	return nil, types.ErrorAssetDoesNotExist
	// }
	assetOutData, found := k.GetAsset(ctx, pairData.AssetOut)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}

	//Checking if app_mapping_id exists
	app_mapping, found := k.GetApp(ctx, msg.AppMappingId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	//Checking if vault acccess disabled
	if !extended_pair_vault.IsVaultActive {
		return nil, types.ErrorVaultInactive

	}
	//Checking if the app_mapping_id in the msg_create & extended_pair_vault_are same or not
	if app_mapping.Id != extended_pair_vault.AppMappingId {
		return nil, types.ErrorAppMappingIdMismatch
	}

	userVault, found := k.GetVault(ctx, msg.UserVaultID)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}
	if userVault.Owner != msg.From {

		return nil, types.ErrVaultAccessUnauthorised
	}

	if app_mapping.Id != userVault.AppMappingId {

		return nil, types.ErrorInvalidAppMappingData

	}
	if extended_pair_vault.Id != userVault.ExtendedPairVaultID {

		return nil, types.ErrorInvalidExtendedPairMappingData
	}
	if msg.Amount.LTE(sdk.NewInt(0)) {
		return nil, types.ErrorInvalidAmount
	}

	userVault.AmountOut = userVault.AmountOut.Add(msg.Amount)
	totaldebt := userVault.AmountOut.Add(*userVault.InterestAccumulated)
	totaldebt = totaldebt.Add(*userVault.ClosingFeeAccumulated)

	_, token_minted_statistics := k.CheckAppExtendedPairVaultMapping(ctx, app_mapping.Id, extended_pair_vault.Id)

	//Check Debt Ceil
	current_minted_statistics := token_minted_statistics.Add(msg.Amount)

	if current_minted_statistics.GTE(extended_pair_vault.DebtCeiling) {
		return nil, types.ErrorAmountOutGreaterThanDebtCeiling
	}

	if err := k.VerifyCollaterlizationRatio(ctx, extended_pair_vault.Id, userVault.AmountIn, totaldebt, extended_pair_vault.MinCr); err != nil {
		return nil, err
	}

	if err := k.MintCoin(ctx, types.ModuleName, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
		return nil, err
	}
	if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositor, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
		return nil, err
	}

	k.SetVault(ctx, userVault)

	//Updating appExtendedPairVaultMappingData
	appExtendedPairVaultData, _ := k.GetAppExtendedPairVaultMapping(ctx, app_mapping.Id)
	k.UpdateTokenMintedAmountLockerMapping(ctx, appExtendedPairVaultData, extended_pair_vault.Id, msg.Amount, true)

	return &types.MsgDrawResponse{}, nil
}

func (k *msgServer) MsgRepay(c context.Context, msg *types.MsgRepayRequest) (*types.MsgRepayResponse, error) {

	ctx := sdk.UnwrapSDKContext(c)

	depositor, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	//checks if extended pair exists
	extended_pair_vault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultID)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extended_pair_vault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	// assetInData, found := k.GetAsset(ctx, pairData.AssetIn)
	// if !found {
	// 	return nil, types.ErrorAssetDoesNotExist
	// }
	assetOutData, found := k.GetAsset(ctx, pairData.AssetOut)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}

	//Checking if app_mapping_id exists
	app_mapping, found := k.GetApp(ctx, msg.AppMappingId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	//Checking if vault acccess disabled
	if !extended_pair_vault.IsVaultActive {
		return nil, types.ErrorVaultInactive

	}
	//Checking if the app_mapping_id in the msg_create & extended_pair_vault_are same or not
	if app_mapping.Id != extended_pair_vault.AppMappingId {
		return nil, types.ErrorAppMappingIdMismatch
	}

	userVault, found := k.GetVault(ctx, msg.UserVaultID)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}
	if userVault.Owner != msg.From {

		return nil, types.ErrVaultAccessUnauthorised
	}

	if app_mapping.Id != userVault.AppMappingId {

		return nil, types.ErrorInvalidAppMappingData

	}
	if extended_pair_vault.Id != userVault.ExtendedPairVaultID {

		return nil, types.ErrorInvalidExtendedPairMappingData
	}
	if msg.Amount.LTE(sdk.NewInt(0)) {
		return nil, types.ErrorInvalidAmount
	}

	if msg.Amount.LTE(*userVault.InterestAccumulated) {
		//Amount is less than equal to the interest acccumulated
		//substract that as interest
		reducedFees := userVault.InterestAccumulated.Sub(msg.Amount)
		userVault.InterestAccumulated = &reducedFees
		//and send it to the collector module
		if err := k.SendCoinFromAccountToModule(ctx, depositor, types.ModuleName, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
			return nil, err
		}
		////////////////////////////////////////
		//
		//			SEND TO COLLECTOR- msg.Amount
		//
		//////////////////////////////////////////
		k.SetVault(ctx, userVault)

	} else {
		updatedUserSentAmountAfterFeesDeduction := msg.Amount.Sub(*userVault.InterestAccumulated)
		updatedUserDebt := userVault.AmountOut.Sub(updatedUserSentAmountAfterFeesDeduction)

		// //If user's closing fees is a bigger amount than the debt floor, user will not close the debt floor
		// totalUpdatedDebt:=updatedUserDebt.Add(*userVault.ClosingFeeAccumulated)
		// if err := k.VerifyCollaterlizationRatio(ctx, extended_pair_vault.Id, userVault.AmountIn, totalUpdatedDebt, extended_pair_vault.MinCr); err != nil {
		// 	return nil, err
		// }

		if !updatedUserDebt.GTE(extended_pair_vault.DebtFloor) {
			return nil, types.ErrorAmountOutLessThanDebtFloor
		}
		if err := k.SendCoinFromAccountToModule(ctx, depositor, types.ModuleName, sdk.NewCoin(assetOutData.Denom, msg.Amount)); err != nil {
			return nil, err
		}

		if err := k.BurnCoin(ctx, types.ModuleName, sdk.NewCoin(assetOutData.Denom, updatedUserSentAmountAfterFeesDeduction)); err != nil {
			return nil, err
		}
		////////////////////////////////////////
		//
		//			SEND TO COLLECTOR----userVault.InterestAccumulated
		//
		//////////////////////////////////////////
		userVault.AmountOut = updatedUserDebt
		zeroval := sdk.ZeroInt()
		userVault.InterestAccumulated = &zeroval
		k.SetVault(ctx, userVault)
		appExtendedPairVaultData, _ := k.GetAppExtendedPairVaultMapping(ctx, app_mapping.Id)
		k.UpdateTokenMintedAmountLockerMapping(ctx, appExtendedPairVaultData, extended_pair_vault.Id, updatedUserSentAmountAfterFeesDeduction, false)

	}

	return &types.MsgRepayResponse{}, nil
}

func (k *msgServer) MsgClose(c context.Context, msg *types.MsgCloseRequest) (*types.MsgCloseResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	depositor, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	//checks if extended pair exists
	extended_pair_vault, found := k.GetPairsVault(ctx, msg.ExtendedPairVaultID)
	if !found {
		return nil, types.ErrorExtendedPairVaultDoesNotExists
	}
	pairData, found := k.GetPair(ctx, extended_pair_vault.PairId)
	if !found {
		return nil, types.ErrorPairDoesNotExist
	}
	assetInData, found := k.GetAsset(ctx, pairData.AssetIn)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}
	assetOutData, found := k.GetAsset(ctx, pairData.AssetOut)
	if !found {
		return nil, types.ErrorAssetDoesNotExist
	}

	//Checking if app_mapping_id exists
	app_mapping, found := k.GetApp(ctx, msg.AppMappingId)
	if !found {
		return nil, types.ErrorAppMappingDoesNotExist
	}
	//Checking if vault acccess disabled
	if !extended_pair_vault.IsVaultActive {
		return nil, types.ErrorVaultInactive

	}
	//Checking if the app_mapping_id in the msg_create & extended_pair_vault_are same or not
	if app_mapping.Id != extended_pair_vault.AppMappingId {
		return nil, types.ErrorAppMappingIdMismatch
	}

	userVault, found := k.GetVault(ctx, msg.UserVaultID)
	if !found {
		return nil, types.ErrorVaultDoesNotExist
	}
	if userVault.Owner != msg.From {

		return nil, types.ErrVaultAccessUnauthorised
	}

	if app_mapping.Id != userVault.AppMappingId {

		return nil, types.ErrorInvalidAppMappingData

	}
	if extended_pair_vault.Id != userVault.ExtendedPairVaultID {

		return nil, types.ErrorInvalidExtendedPairMappingData
	}

	totalUserDebt := userVault.AmountOut.Add(*userVault.InterestAccumulated)
	totalUserDebt = totalUserDebt.Add(*userVault.ClosingFeeAccumulated)
	if err := k.SendCoinFromAccountToModule(ctx, depositor, types.ModuleName, sdk.NewCoin(assetOutData.Denom, totalUserDebt)); err != nil {
		return nil, err
	}
	////////////////////////////////////////
	//
	//			SEND TO COLLECTOR----userVault.InterestAccumulated & userVault.ClosingFees
	//
	//////////////////////////////////////////

	if err := k.BurnCoin(ctx, types.ModuleName, sdk.NewCoin(assetOutData.Denom, userVault.AmountOut)); err != nil {
		return nil, err
	}

	if err := k.SendCoinFromModuleToAccount(ctx, types.ModuleName, depositor, sdk.NewCoin(assetInData.Denom, userVault.AmountIn)); err != nil {
		return nil, err
	}

	//Update LookupTable minting Status
	appExtendedPairVaultData, _ := k.GetAppExtendedPairVaultMapping(ctx, app_mapping.Id)

	k.UpdateCollateralLockedAmountLockerMapping(ctx, appExtendedPairVaultData, extended_pair_vault.Id, userVault.AmountIn, false)
	k.UpdateTokenMintedAmountLockerMapping(ctx, appExtendedPairVaultData, extended_pair_vault.Id, userVault.AmountOut, false)

	//Remove address from lookup table
	k.DeleteAddressFromAppExtendedPairVaultMapping(ctx, extended_pair_vault.Id, userVault.Id, app_mapping.Id)

	//Remove user extendedPair to address field in UserLookupStruct
	k.UpdateUserVaultExtendedPairMapping(ctx, extended_pair_vault.Id, msg.From, app_mapping.Id)

	//Delete Vault
	k.DeleteVault(ctx, userVault.Id)

	return &types.MsgCloseResponse{}, nil
}
