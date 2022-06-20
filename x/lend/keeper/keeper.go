package keeper

import (
	"fmt"
	"github.com/comdex-official/comdex/x/lend/expected"
	"github.com/comdex-official/comdex/x/lend/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
		bank       expected.BankKeeper
		account    expected.AccountKeeper
		asset      expected.AssetKeeper
		market     expected.MarketKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	bank expected.BankKeeper,
	account expected.AccountKeeper,
	asset expected.AssetKeeper,
	market expected.MarketKeeper,

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
		bank:       bank,
		account:    account,
		asset:      asset,
		market:     market,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) ModuleBalance(ctx sdk.Context, moduleName string, denom string) sdk.Int {
	return k.bank.GetBalance(ctx, authtypes.NewModuleAddress(moduleName), denom).Amount
}

func (k Keeper) LendAsset(ctx sdk.Context, lenderAddr string, AssetId uint64, Amount sdk.Coin, PoolId uint64) error {

	asset, _ := k.GetAsset(ctx, AssetId)
	pool, _ := k.GetPool(ctx, PoolId)

	if Amount.Denom != asset.Denom {
		return sdkerrors.Wrap(types.ErrBadOfferCoinAmount, Amount.Denom)
	}

	addr, _ := sdk.AccAddressFromBech32(lenderAddr)

	if k.HasLendForAddressByAsset(ctx, addr, AssetId) {
		return types.ErrorDuplicateLend
	}

	loanTokens := sdk.NewCoins(Amount)

	cToken, err := k.ExchangeToken(ctx, Amount, asset.Name)
	if err != nil {
		return err
	}

	if err := k.bank.SendCoinsFromAccountToModule(ctx, addr, pool.ModuleName, loanTokens); err != nil {
		return err
	}
	// mint c/Token and set new total cToken supply

	cTokens := sdk.NewCoins(cToken)
	if err = k.bank.MintCoins(ctx, pool.ModuleName, cTokens); err != nil {
		return err
	}
	// adding to the total supply
	if err = k.setCTokenSupply(ctx, k.GetCTokenSupply(ctx, cToken.Denom).Add(cToken)); err != nil {
		return err
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, pool.ModuleName, addr, cTokens)
	if err != nil {
		return err
	}

	//TODO:
	// create lend position
	// update that position from next block
	// set initial interest 0
	// loan amt, pool, check if bridged asset,

	//////////////////////////////////////////////
	lendId := k.GetUserLendIDHistory(ctx)

	lendPos := types.LendAsset{
		ID:                 lendId + 1,
		AssetId:            AssetId,
		PoolId:             PoolId,
		Owner:              lenderAddr,
		AmountIn:           Amount,
		LendingTime:        ctx.BlockTime(),
		UpdatedAmountIn:    Amount.Amount,
		Reward_Accumulated: sdk.ZeroInt(),
	}

	k.SetUserLendIDHistory(ctx, lendPos.ID)
	k.SetLend(ctx, lendPos)
	k.SetLendForAddressByAsset(ctx, addr, lendPos.AssetId, lendPos.ID)
	err = k.UpdateUserLendIdMapping(ctx, lenderAddr, lendPos.ID, true)
	if err != nil {
		return err
	}
	err = k.UpdateLendIdsMapping(ctx, lendPos.ID, true)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) WithdrawAsset(ctx sdk.Context, addr string, lendId uint64, withdrawal sdk.Coin) error {

	//TODO:
	// check if lend position exists
	// check borrow for that lend position
	// get current CR
	// calculate available to withdraw
	// take balance in c/Token and return Token
	// if borrow available update CR
	// Ensure module account has sufficient unreserved tokens to withdraw

	lenderAddr, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return err
	}

	lendPos, found := k.GetLend(ctx, lendId)
	if !found {
		return types.ErrLendNotFound
	}
	getAsset, _ := k.GetAsset(ctx, lendPos.AssetId)
	pool, _ := k.GetPool(ctx, lendPos.PoolId)

	if lendPos.Owner != addr {
		return types.ErrLendAccessUnauthorised
	}

	if withdrawal.Denom != getAsset.Denom {
		return sdkerrors.Wrap(types.ErrBadOfferCoinAmount, withdrawal.Denom)
	}

	reservedAmount := k.GetReserveFunds(ctx, withdrawal.Denom)
	availableAmount := k.ModuleBalance(ctx, pool.ModuleName, withdrawal.Denom)

	if withdrawal.Amount.GT(lendPos.AmountIn.Amount) {
		return sdkerrors.Wrap(types.ErrWithdrawlAmountExceeds, withdrawal.String())
	}

	if withdrawal.Amount.GT(availableAmount.Sub(reservedAmount)) {
		return sdkerrors.Wrap(types.ErrLendingPoolInsufficient, withdrawal.String())
	}

	tokens := sdk.NewCoins(withdrawal)
	cToken, err := k.ExchangeToken(ctx, withdrawal, getAsset.Name)

	if err != nil {
		return err
	}
	_, found = k.GetLendIdToBorrowIdMapping(ctx, lendId)
	if !found {

		if withdrawal.Amount.LT(lendPos.UpdatedAmountIn) {

			//TODO:
			// update lend & Updated amount in position
			// create a lend to borrow mapping
			// borrow calculations for CR

			if err := k.SendCoinFromAccountToModule(ctx, lenderAddr, pool.ModuleName, cToken); err != nil {
				return err
			}

			//burn c/Token
			cTokens := sdk.NewCoins(cToken)
			err = k.bank.BurnCoins(ctx, pool.ModuleName, cTokens)
			if err != nil {
				return err
			}

			if err := k.bank.SendCoinsFromModuleToAccount(ctx, pool.ModuleName, lenderAddr, tokens); err != nil {
				return err
			}

			lendPos.AmountIn = lendPos.AmountIn.Sub(withdrawal)
			lendPos.UpdatedAmountIn = lendPos.UpdatedAmountIn.Sub(withdrawal.Amount)
			k.SetLend(ctx, lendPos)

		} else {
			return nil
		}
	} else {
		if withdrawal.Amount.LT(lendPos.UpdatedAmountIn) {
			// add CR validation
			// lend to borrow mapping

			if err := k.SendCoinFromAccountToModule(ctx, lenderAddr, pool.ModuleName, cToken); err != nil {
				return err
			}

			//burn c/Token
			cTokens := sdk.NewCoins(cToken)
			err = k.bank.BurnCoins(ctx, pool.ModuleName, cTokens)
			if err != nil {
				return err
			}

			if err := k.bank.SendCoinsFromModuleToAccount(ctx, pool.ModuleName, lenderAddr, tokens); err != nil {
				return err
			}

			lendPos.AmountIn = lendPos.AmountIn.Sub(withdrawal)
			lendPos.UpdatedAmountIn = lendPos.UpdatedAmountIn.Sub(withdrawal.Amount)
			k.SetLend(ctx, lendPos)

		} else {
			return nil
		}
	}

	return nil
}

func (k Keeper) DepositAsset(ctx sdk.Context, addr string, lendId uint64, deposit sdk.Coin) error {

	//TODO:
	// check lend position
	// mint additional c/Token
	// send c/token to user
	// update LendPos

	lenderAddr, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return err
	}

	lendPos, found := k.GetLend(ctx, lendId)
	if !found {
		return types.ErrLendNotFound
	}

	getAsset, _ := k.GetAsset(ctx, lendPos.AssetId)
	pool, _ := k.GetPool(ctx, lendPos.PoolId)

	if deposit.Denom != getAsset.Denom {
		return sdkerrors.Wrap(types.ErrBadOfferCoinAmount, deposit.Denom)
	}

	cToken, err := k.ExchangeToken(ctx, deposit, getAsset.Name)
	if err != nil {
		return err
	}
	cTokens := sdk.NewCoins(cToken)

	if err = k.bank.MintCoins(ctx, pool.ModuleName, cTokens); err != nil {
		return err
	}

	if err = k.setCTokenSupply(ctx, k.GetCTokenSupply(ctx, cToken.Denom).Add(cToken)); err != nil {
		return err
	}

	if err := k.bank.SendCoinsFromAccountToModule(ctx, lenderAddr, pool.ModuleName, sdk.NewCoins(deposit)); err != nil {
		return err
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, pool.ModuleName, lenderAddr, cTokens)
	if err != nil {
		return err
	}

	lendPos.AmountIn = lendPos.AmountIn.Add(deposit)
	lendPos.UpdatedAmountIn = lendPos.UpdatedAmountIn.Add(deposit.Amount)
	k.SetLend(ctx, lendPos)

	return nil
}

func (k Keeper) CloseLend(ctx sdk.Context, addr string, lendId uint64) error {

	lenderAddr, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return err
	}

	lendPos, found := k.GetLend(ctx, lendId)
	if !found {
		return types.ErrLendNotFound
	}
	getAsset, _ := k.GetAsset(ctx, lendPos.AssetId)
	pool, _ := k.GetPool(ctx, lendPos.PoolId)

	if lendPos.Owner != addr {
		return types.ErrLendAccessUnauthorised
	}

	_, found = k.GetLendIdToBorrowIdMapping(ctx, lendId)
	if found {
		return types.ErrBorrowingPositionOpen
	}
	reservedAmount := k.GetReserveFunds(ctx, lendPos.AmountIn.Denom)
	availableAmount := k.ModuleBalance(ctx, pool.ModuleName, lendPos.AmountIn.Denom)

	if lendPos.UpdatedAmountIn.GT(availableAmount.Sub(reservedAmount)) {
		return sdkerrors.Wrap(types.ErrLendingPoolInsufficient, lendPos.AmountIn.String())
	}

	tokens := sdk.NewCoins(sdk.NewCoin(lendPos.AmountIn.Denom, lendPos.UpdatedAmountIn))

	cToken, err := k.ExchangeToken(ctx, sdk.NewCoin(lendPos.AmountIn.Denom, lendPos.UpdatedAmountIn), getAsset.Name)
	if err != nil {
		return err
	}

	if err := k.SendCoinFromAccountToModule(ctx, lenderAddr, pool.ModuleName, cToken); err != nil {
		return err
	}

	cTokens := sdk.NewCoins(cToken)
	err = k.bank.BurnCoins(ctx, pool.ModuleName, cTokens)
	if err != nil {
		return err
	}

	if err := k.bank.SendCoinsFromModuleToAccount(ctx, pool.ModuleName, lenderAddr, tokens); err != nil {
		return err
	}

	k.DeleteLendForAddressByAsset(ctx, lenderAddr, lendPos.AssetId)

	err = k.UpdateUserLendIdMapping(ctx, addr, lendPos.ID, false)
	if err != nil {
		return err
	}
	err = k.UpdateLendIdsMapping(ctx, lendPos.ID, false)
	if err != nil {
		return err
	}
	k.DeleteLend(ctx, lendPos.ID)
	return nil
}

func uint64InSlice(a uint64, list []uint64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (k Keeper) BorrowAsset(ctx sdk.Context, addr string, lendId, pairId uint64, IsStableBorrow bool, AmountIn, loan sdk.Coin) error {

	//TODO:
	// take Lending Id and check if exists
	// take the pair id, check if that pair id exists for lend AssetId (Asset to pair mapping)
	// Amount In is the u/Token user provided (updated Amount In) At the time of creation
	// option to deposit additional u/Token to inc CR will be added in DepositBorrow function
	// From Amount Out calculate the CR, if below LR fail
	// update kv stores accordingly

	lenderAddr, _ := sdk.AccAddressFromBech32(addr)

	lendPos, found := k.GetLend(ctx, lendId)
	if !found {
		return types.ErrLendNotFound
	}
	getAsset, _ := k.GetAsset(ctx, lendPos.AssetId)
	if lendPos.Owner != addr {
		return types.ErrLendAccessUnauthorised
	}

	if k.HasBorrowForAddressByPair(ctx, lenderAddr, pairId) {
		return types.ErrorDuplicateBorrow
	}

	pairMapping, _ := k.GetAssetToPair(ctx, lendPos.AssetId)
	found = uint64InSlice(pairId, pairMapping.PairId)
	if !found {
		return types.ErrorPairNotFound
	}
	pair, found := k.GetLendPair(ctx, pairId)
	if !found {
		return types.ErrorPairNotFound
	}
	AssetInPool, _ := k.GetPool(ctx, lendPos.PoolId)
	AssetOutPool, _ := k.GetPool(ctx, pair.AssetOutPoolId)

	//check cr ratio
	assetIn, _ := k.GetAsset(ctx, lendPos.AssetId)
	assetOut, _ := k.GetAsset(ctx, pair.AssetOut)

	assetRatesStats, _ := k.GetAssetRatesStats(ctx, pair.AssetOut)
	err := k.VerifyCollaterlizationRatio(ctx, AmountIn.Amount, assetIn, loan.Amount, assetOut, assetRatesStats.LiquidationThreshold)
	if err != nil {
		return err
	}
	borrowId := k.GetUserBorrowIDHistory(ctx)

	if !pair.IsInterPool {
		// check sufficient amt in pool to borrow
		reservedAmount := k.GetReserveFunds(ctx, loan.Denom)
		availableAmount := k.ModuleBalance(ctx, AssetOutPool.ModuleName, loan.Denom)

		if loan.Amount.GT(availableAmount.Sub(reservedAmount)) {
			return sdkerrors.Wrap(types.ErrBorrowingPoolInsufficient, loan.String())
		}

		AmountIn := sdk.NewCoin(lendPos.AmountIn.Denom, AmountIn.Amount)
		AmountOut := loan
		// take c/Tokens from the user
		cToken, err := k.ExchangeToken(ctx, sdk.NewCoin(lendPos.AmountIn.Denom, AmountIn.Amount), getAsset.Name)
		if err != nil {
			return err
		}

		if err := k.SendCoinFromAccountToModule(ctx, lenderAddr, AssetInPool.ModuleName, cToken); err != nil {
			return err
		}

		if err := k.SendCoinFromModuleToAccount(ctx, AssetOutPool.ModuleName, lenderAddr, loan); err != nil {
			return err
		}

		var StableBorrowRate sdk.Dec
		if assetRatesStats.EnableStableBorrow {
			if IsStableBorrow {
				StableBorrowRate, err = k.GetBorrowAPYByAssetId(ctx, AssetOutPool.PoolId, pair.AssetOut, IsStableBorrow)
				if err != nil {
					return err
				}
			} else {
				StableBorrowRate = sdk.ZeroDec()
			}
		} else {
			IsStableBorrow = false
			StableBorrowRate = sdk.ZeroDec()
		}

		borrowPos := types.BorrowAsset{
			ID:                   borrowId + 1,
			LendingID:            lendId,
			PairID:               pairId,
			AmountIn:             AmountIn,
			AmountOut:            AmountOut,
			IsStableBorrow:       IsStableBorrow,
			StableBorrowRate:     StableBorrowRate,
			BorrowingTime:        ctx.BlockTime(),
			UpdatedAmountOut:     AmountOut.Amount,
			Interest_Accumulated: sdk.ZeroInt(),
		}
		k.SetUserBorrowIDHistory(ctx, borrowPos.ID)
		k.SetBorrow(ctx, borrowPos)
		k.SetBorrowForAddressByPair(ctx, lenderAddr, pairId, borrowPos.ID)
		err = k.UpdateUserBorrowIdMapping(ctx, lendPos.Owner, borrowPos.ID, true)
		if err != nil {
			return err
		}
		err = k.UpdateLendIdToBorrowIdMapping(ctx, borrowPos.LendingID, borrowPos.ID, true)
		if err != nil {
			return err
		}

	} else {
		reservedAmount := k.GetReserveFunds(ctx, loan.Denom)
		availableAmount := k.ModuleBalance(ctx, AssetOutPool.ModuleName, loan.Denom)

		if loan.Amount.GT(availableAmount.Sub(reservedAmount)) {
			return sdkerrors.Wrap(types.ErrBorrowingPoolInsufficient, loan.String())
		}
		assetIn := lendPos.UpdatedAmountIn
		priceAssetIn, _ := k.GetPriceForAsset(ctx, pair.AssetIn)
		amtIn := assetIn.Mul(sdk.NewIntFromUint64(priceAssetIn))

		priceFirstBridgedAsset, _ := k.GetPriceForAsset(ctx, AssetInPool.FirstBridgedAssetId)
		priceSecondBridgedAsset, _ := k.GetPriceForAsset(ctx, AssetInPool.SecondBridgedAssetId)
		firstBridgedAsset, _ := k.GetAsset(ctx, AssetInPool.FirstBridgedAssetId)
		secondBridgedAsset, _ := k.GetAsset(ctx, AssetInPool.SecondBridgedAssetId)

		// qty of first and second bridged asset to be sent over different pool according to the borrow Pool

		firstBridgedAssetQty := amtIn.Quo(sdk.NewIntFromUint64(priceFirstBridgedAsset))
		firstBridgedAssetBal := k.ModuleBalance(ctx, AssetInPool.ModuleName, firstBridgedAsset.Denom)
		secondBridgedAssetQty := amtIn.Quo(sdk.NewIntFromUint64(priceSecondBridgedAsset))
		secondBridgedAssetBal := k.ModuleBalance(ctx, AssetInPool.ModuleName, secondBridgedAsset.Denom)

		if firstBridgedAssetQty.LT(firstBridgedAssetBal) {

			// take c/Tokens from the user
			cToken, err := k.ExchangeToken(ctx, sdk.NewCoin(lendPos.AmountIn.Denom, AmountIn.Amount), getAsset.Name)
			if err != nil {
				return err
			}

			if err := k.SendCoinFromAccountToModule(ctx, lenderAddr, AssetInPool.ModuleName, cToken); err != nil {
				return err
			}

			if err := k.SendCoinFromModuleToModule(ctx, AssetInPool.ModuleName, AssetOutPool.ModuleName, sdk.NewCoins(sdk.NewCoin(firstBridgedAsset.Denom, firstBridgedAssetQty))); err != nil {
				return err
			}

			if err := k.SendCoinFromModuleToAccount(ctx, AssetOutPool.ModuleName, lenderAddr, loan); err != nil {
				return err
			}

			AmountIn := sdk.NewCoin(lendPos.AmountIn.Denom, lendPos.UpdatedAmountIn)
			AmountOut := loan

			var StableBorrowRate sdk.Dec
			if assetRatesStats.EnableStableBorrow {
				if IsStableBorrow {
					StableBorrowRate, err = k.GetBorrowAPYByAssetId(ctx, AssetOutPool.PoolId, pair.AssetOut, IsStableBorrow)
					if err != nil {
						return err
					}
				} else {
					StableBorrowRate = sdk.ZeroDec()
				}
			} else {
				IsStableBorrow = false
				StableBorrowRate = sdk.ZeroDec()
			}

			borrowPos := types.BorrowAsset{
				ID:                   borrowId + 1,
				LendingID:            lendId,
				PairID:               pairId,
				AmountIn:             AmountIn,
				AmountOut:            AmountOut,
				IsStableBorrow:       IsStableBorrow,
				StableBorrowRate:     StableBorrowRate,
				BorrowingTime:        ctx.BlockTime(),
				UpdatedAmountOut:     AmountOut.Amount,
				Interest_Accumulated: sdk.ZeroInt(),
			}
			k.SetUserBorrowIDHistory(ctx, borrowPos.ID)
			k.SetBorrow(ctx, borrowPos)
			k.SetBorrowForAddressByPair(ctx, lenderAddr, pairId, borrowPos.ID)
			err = k.UpdateUserBorrowIdMapping(ctx, lendPos.Owner, borrowPos.ID, true)
			if err != nil {
				return err
			}
			err = k.UpdateLendIdToBorrowIdMapping(ctx, borrowPos.LendingID, borrowPos.ID, true)
			if err != nil {
				return err
			}

		} else if secondBridgedAssetQty.LT(secondBridgedAssetBal) {
			// take c/Tokens from the user
			cToken, err := k.ExchangeToken(ctx, sdk.NewCoin(lendPos.AmountIn.Denom, AmountIn.Amount), getAsset.Name)
			if err != nil {
				return err
			}

			if err := k.SendCoinFromAccountToModule(ctx, lenderAddr, AssetInPool.ModuleName, cToken); err != nil {
				return err
			}

			if err := k.SendCoinFromModuleToModule(ctx, AssetInPool.ModuleName, AssetOutPool.ModuleName, sdk.NewCoins(sdk.NewCoin(secondBridgedAsset.Denom, secondBridgedAssetQty))); err != nil {
				return err
			}

			if err := k.SendCoinFromModuleToAccount(ctx, AssetOutPool.ModuleName, lenderAddr, loan); err != nil {
				return err
			}

			AmountIn := sdk.NewCoin(lendPos.AmountIn.Denom, lendPos.UpdatedAmountIn)
			AmountOut := loan

			var StableBorrowRate sdk.Dec
			if assetRatesStats.EnableStableBorrow {
				if IsStableBorrow {
					StableBorrowRate, err = k.GetBorrowAPYByAssetId(ctx, AssetOutPool.PoolId, pair.AssetOut, IsStableBorrow)
					if err != nil {
						return err
					}
				} else {
					StableBorrowRate = sdk.ZeroDec()
				}
			} else {
				IsStableBorrow = false
				StableBorrowRate = sdk.ZeroDec()
			}

			borrowPos := types.BorrowAsset{
				ID:                   borrowId + 1,
				LendingID:            lendId,
				PairID:               pairId,
				AmountIn:             AmountIn,
				AmountOut:            AmountOut,
				IsStableBorrow:       IsStableBorrow,
				StableBorrowRate:     StableBorrowRate,
				BorrowingTime:        ctx.BlockTime(),
				UpdatedAmountOut:     AmountOut.Amount,
				Interest_Accumulated: sdk.ZeroInt(),
			}
			k.SetUserBorrowIDHistory(ctx, borrowPos.ID)
			k.SetBorrow(ctx, borrowPos)
			k.SetBorrowForAddressByPair(ctx, lenderAddr, pairId, borrowPos.ID)
			err = k.UpdateUserBorrowIdMapping(ctx, lendPos.Owner, borrowPos.ID, true)
			if err != nil {
				return err
			}

			err = k.UpdateLendIdToBorrowIdMapping(ctx, borrowPos.LendingID, borrowPos.ID, true)
			if err != nil {
				return err
			}

		} else {
			return types.ErrBorrowingPoolInsufficient
		}

	}
	return nil
}

func (k Keeper) RepayAsset(ctx sdk.Context, borrowId uint64, borrowerAddr string, payment sdk.Coin) error {
	//TODO:
	// check borrow position
	// payment is LT UpdatedAmountOut
	// take the payment from borrower and reduce UpdatedAmountOut
	// update KV store accordingly

	borrowPos, found := k.GetBorrow(ctx, borrowId)
	if !found {
		return types.ErrBorrowNotFound
	}
	addr, _ := sdk.AccAddressFromBech32(borrowerAddr)
	pair, _ := k.GetLendPair(ctx, borrowPos.PairID)
	pool, _ := k.GetPool(ctx, pair.AssetOutPoolId)
	lendPos, _ := k.GetLend(ctx, borrowPos.LendingID)
	if lendPos.Owner != borrowerAddr {
		return types.ErrLendAccessUnauthorised
	}
	if borrowPos.AmountOut.Denom != payment.Denom {
		return types.ErrBadOfferCoinAmount
	}

	if payment.Amount.LT(borrowPos.UpdatedAmountOut) {
		// sending repayment to moduleAcc from borrower
		if err := k.bank.SendCoinsFromAccountToModule(ctx, addr, pool.ModuleName, sdk.NewCoins(payment)); err != nil {
			return err
		}
		borrowPos.UpdatedAmountOut = borrowPos.UpdatedAmountOut.Sub(payment.Amount)
		borrowPos.AmountOut = borrowPos.AmountOut.Sub(payment)
		k.SetBorrow(ctx, borrowPos)
	} else {
		return types.ErrInvalidRepayment
	}
	return nil
}

func (k Keeper) DepositBorrowAsset(ctx sdk.Context, borrowId uint64, addr string, AmountIn sdk.Coin) error {

	borrowPos, found := k.GetBorrow(ctx, borrowId)
	if !found {
		return types.ErrBorrowNotFound
	}
	lendId := borrowPos.LendingID
	pairId := borrowPos.PairID
	lenderAddr, _ := sdk.AccAddressFromBech32(addr)

	lendPos, found := k.GetLend(ctx, lendId)
	if !found {
		return types.ErrLendNotFound
	}
	getAsset, _ := k.GetAsset(ctx, lendPos.AssetId)
	if lendPos.Owner != addr {
		return types.ErrLendAccessUnauthorised
	}

	if k.HasBorrowForAddressByPair(ctx, lenderAddr, pairId) {
		return types.ErrorDuplicateBorrow
	}

	pairMapping, _ := k.GetAssetToPair(ctx, lendPos.AssetId)
	found = uint64InSlice(pairId, pairMapping.PairId)
	if !found {
		return types.ErrorPairNotFound
	}
	pair, found := k.GetLendPair(ctx, pairId)
	if !found {
		return types.ErrorPairNotFound
	}
	AssetInPool, _ := k.GetPool(ctx, lendPos.PoolId)
	AssetOutPool, _ := k.GetPool(ctx, pair.AssetOutPoolId)

	if !pair.IsInterPool {

		AmountIn := sdk.NewCoin(lendPos.AmountIn.Denom, AmountIn.Amount)
		// take c/Tokens from the user
		cToken, err := k.ExchangeToken(ctx, sdk.NewCoin(lendPos.AmountIn.Denom, AmountIn.Amount), getAsset.Name)
		if err != nil {
			return err
		}

		if err := k.SendCoinFromAccountToModule(ctx, lenderAddr, AssetInPool.ModuleName, cToken); err != nil {
			return err
		}

		borrowPos.AmountIn = borrowPos.AmountIn.Add(AmountIn)
		k.SetUserBorrowIDHistory(ctx, borrowPos.ID)
		k.SetBorrow(ctx, borrowPos)
		k.SetBorrowForAddressByPair(ctx, lenderAddr, pairId, borrowPos.ID)
		err = k.UpdateUserBorrowIdMapping(ctx, lendPos.Owner, borrowPos.ID, true)
		if err != nil {
			return err
		}
		err = k.UpdateLendIdToBorrowIdMapping(ctx, borrowPos.LendingID, borrowPos.ID, true)
		if err != nil {
			return err
		}

	} else {
		assetIn := lendPos.UpdatedAmountIn
		priceAssetIn, _ := k.GetPriceForAsset(ctx, pair.AssetIn)
		amtIn := assetIn.Mul(sdk.NewIntFromUint64(priceAssetIn))

		priceFirstBridgedAsset, _ := k.GetPriceForAsset(ctx, AssetInPool.FirstBridgedAssetId)
		priceSecondBridgedAsset, _ := k.GetPriceForAsset(ctx, AssetInPool.SecondBridgedAssetId)
		firstBridgedAsset, _ := k.GetAsset(ctx, AssetInPool.FirstBridgedAssetId)
		secondBridgedAsset, _ := k.GetAsset(ctx, AssetInPool.SecondBridgedAssetId)

		// qty of first and second bridged asset to be sent over different pool according to the borrow Pool

		firstBridgedAssetQty := amtIn.Quo(sdk.NewIntFromUint64(priceFirstBridgedAsset))
		firstBridgedAssetBal := k.ModuleBalance(ctx, AssetInPool.ModuleName, firstBridgedAsset.Denom)
		secondBridgedAssetQty := amtIn.Quo(sdk.NewIntFromUint64(priceSecondBridgedAsset))
		secondBridgedAssetBal := k.ModuleBalance(ctx, AssetInPool.ModuleName, secondBridgedAsset.Denom)

		if firstBridgedAssetQty.LT(firstBridgedAssetBal) {

			// take c/Tokens from the user
			cToken, err := k.ExchangeToken(ctx, sdk.NewCoin(lendPos.AmountIn.Denom, AmountIn.Amount), getAsset.Name)
			if err != nil {
				return err
			}

			if err := k.SendCoinFromAccountToModule(ctx, lenderAddr, AssetInPool.ModuleName, cToken); err != nil {
				return err
			}

			if err := k.SendCoinFromModuleToModule(ctx, AssetInPool.ModuleName, AssetOutPool.ModuleName, sdk.NewCoins(sdk.NewCoin(firstBridgedAsset.Denom, firstBridgedAssetQty))); err != nil {
				return err
			}

			borrowPos.AmountIn = borrowPos.AmountIn.Add(AmountIn)

			k.SetUserBorrowIDHistory(ctx, borrowPos.ID)
			k.SetBorrow(ctx, borrowPos)
			k.SetBorrowForAddressByPair(ctx, lenderAddr, pairId, borrowPos.ID)
			err = k.UpdateUserBorrowIdMapping(ctx, lendPos.Owner, borrowPos.ID, true)
			if err != nil {
				return err
			}
			err = k.UpdateLendIdToBorrowIdMapping(ctx, borrowPos.LendingID, borrowPos.ID, true)
			if err != nil {
				return err
			}

		} else if secondBridgedAssetQty.LT(secondBridgedAssetBal) {
			// take c/Tokens from the user
			cToken, err := k.ExchangeToken(ctx, sdk.NewCoin(lendPos.AmountIn.Denom, AmountIn.Amount), getAsset.Name)
			if err != nil {
				return err
			}

			if err := k.SendCoinFromAccountToModule(ctx, lenderAddr, AssetInPool.ModuleName, cToken); err != nil {
				return err
			}

			if err := k.SendCoinFromModuleToModule(ctx, AssetInPool.ModuleName, AssetOutPool.ModuleName, sdk.NewCoins(sdk.NewCoin(secondBridgedAsset.Denom, secondBridgedAssetQty))); err != nil {
				return err
			}

			borrowPos.AmountIn = borrowPos.AmountIn.Add(AmountIn)

			k.SetUserBorrowIDHistory(ctx, borrowPos.ID)
			k.SetBorrow(ctx, borrowPos)
			k.SetBorrowForAddressByPair(ctx, lenderAddr, pairId, borrowPos.ID)
			err = k.UpdateUserBorrowIdMapping(ctx, lendPos.Owner, borrowPos.ID, true)
			if err != nil {
				return err
			}

			err = k.UpdateLendIdToBorrowIdMapping(ctx, borrowPos.LendingID, borrowPos.ID, true)
			if err != nil {
				return err
			}

		} else {
			return types.ErrBorrowingPoolInsufficient
		}

	}
	return nil
}

func (k Keeper) DrawAsset(ctx sdk.Context, borrowId uint64, borrowerAddr string, payment sdk.Coin) error {

	//TODO:
	// check borrow position
	// check current CR and verify

	borrowPos, found := k.GetBorrow(ctx, borrowId)
	if !found {
		return types.ErrBorrowNotFound
	}
	addr, _ := sdk.AccAddressFromBech32(borrowerAddr)
	pair, _ := k.GetLendPair(ctx, borrowPos.PairID)
	pool, _ := k.GetPool(ctx, pair.AssetOutPoolId)
	lendPos, _ := k.GetLend(ctx, borrowPos.LendingID)
	if lendPos.Owner != borrowerAddr {
		return types.ErrLendAccessUnauthorised
	}
	if borrowPos.AmountOut.Denom != payment.Denom {
		return types.ErrBadOfferCoinAmount
	}
	if borrowPos.UpdatedAmountOut.GT(payment.Amount) {
		borrowPos.UpdatedAmountOut = borrowPos.UpdatedAmountOut.Add(payment.Amount)
		assetIn, _ := k.GetAsset(ctx, lendPos.AssetId)
		assetOut, _ := k.GetAsset(ctx, pair.AssetOut)

		assetRatesStats, _ := k.GetAssetRatesStats(ctx, pair.AssetOut)
		err := k.VerifyCollaterlizationRatio(ctx, borrowPos.AmountIn.Amount, assetIn, borrowPos.UpdatedAmountOut, assetOut, assetRatesStats.LiquidationThreshold)
		if err != nil {
			return err
		}
		if err := k.SendCoinFromModuleToAccount(ctx, pool.ModuleName, addr, payment); err != nil {
			return err
		}
		borrowPos.AmountOut.Add(payment)
		k.DeleteLendIdToBorrowIdMapping(ctx, borrowPos.LendingID)
		k.SetBorrow(ctx, borrowPos)
	}
	return nil
}

func (k Keeper) CloseBorrow(ctx sdk.Context, borrowerAddr string, borrowId uint64) error {

	//TODO:
	// take borrow position
	// match user
	// Delete BorrowPos and mappings

	borrowPos, found := k.GetBorrow(ctx, borrowId)
	if !found {
		return types.ErrBorrowNotFound
	}
	addr, _ := sdk.AccAddressFromBech32(borrowerAddr)
	pair, _ := k.GetLendPair(ctx, borrowPos.PairID)
	pool, _ := k.GetPool(ctx, pair.AssetOutPoolId)
	lendPos, _ := k.GetLend(ctx, borrowPos.LendingID)
	if lendPos.Owner != borrowerAddr {
		return types.ErrLendAccessUnauthorised
	}
	assetOut, _ := k.GetAsset(ctx, pair.AssetOut)

	amt := sdk.NewCoins(sdk.NewCoin(assetOut.Denom, borrowPos.UpdatedAmountOut))

	if err := k.bank.SendCoinsFromAccountToModule(ctx, addr, pool.ModuleName, amt); err != nil {
		return err
	}
	err := k.UpdateUserBorrowIdMapping(ctx, lendPos.Owner, borrowPos.ID, false)
	if err != nil {
		return err
	}
	k.DeleteBorrowForAddressByPair(ctx, addr, borrowPos.PairID)
	err = k.UpdateLendIdToBorrowIdMapping(ctx, borrowPos.LendingID, borrowPos.ID, false)
	if err != nil {
		return err
	}
	k.DeleteBorrow(ctx, borrowId)

	return nil
}

func (k Keeper) FundModAcc(ctx sdk.Context, moduleName string, assetId uint64, lenderAddr sdk.AccAddress, payment sdk.Coin) error {

	loanTokens := sdk.NewCoins(payment)
	if err := k.bank.SendCoinsFromAccountToModule(ctx, lenderAddr, moduleName, loanTokens); err != nil {
		return err
	}

	getAsset, found := k.GetAsset(ctx, assetId)
	if !found {
		return types.ErrLendNotFound
	}

	cToken, err := k.ExchangeToken(ctx, payment, getAsset.Name)
	if err != nil {
		return err
	}
	err = k.MintCoin(ctx, moduleName, cToken)
	if err != nil {
		return err
	}

	currentCollateral := k.GetCollateralAmount(ctx, lenderAddr, payment.Denom)
	if err := k.setCollateralAmount(ctx, lenderAddr, currentCollateral.Add(payment)); err != nil {
		return err
	}

	return nil
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}
