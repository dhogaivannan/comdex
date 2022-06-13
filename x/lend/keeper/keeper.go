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
	fmt.Println("ModuleBalance...ModuleBalance", moduleName)
	fmt.Println("ModuleBalance...denom", denom)

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

	if withdrawal.Amount.LT(lendPos.UpdatedAmountIn) {

		//TODO:
		// update lend & Updated amount in position
		// create a lend to borrow mapping
		// borrow calculations for CR

		if err := k.SendCoinFromAccountToModule(ctx, lenderAddr, pool.ModuleName, cToken); err != nil {
			return err
		}

		//burn c/Token

		if err := k.bank.SendCoinsFromModuleToAccount(ctx, pool.ModuleName, lenderAddr, tokens); err != nil {
			return err
		}

		lendPos.AmountIn = lendPos.AmountIn.Sub(withdrawal)
		lendPos.UpdatedAmountIn = lendPos.UpdatedAmountIn.Sub(withdrawal.Amount)
		k.SetLend(ctx, lendPos)

	} else {
		return nil
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

	reservedAmount := k.GetReserveFunds(ctx, lendPos.AmountIn.Denom)
	availableAmount := k.ModuleBalance(ctx, pool.ModuleName, lendPos.AmountIn.Denom)

	if lendPos.AmountIn.Amount.GT(availableAmount.Sub(reservedAmount)) {
		return sdkerrors.Wrap(types.ErrLendingPoolInsufficient, lendPos.AmountIn.String())
	}

	tokens := sdk.NewCoins(lendPos.AmountIn)

	cToken, err := k.ExchangeToken(ctx, lendPos.AmountIn, getAsset.Name)
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

	lendPos.AmountIn = lendPos.AmountIn.Sub(lendPos.AmountIn)
	lendPos.UpdatedAmountIn = lendPos.UpdatedAmountIn.Sub(lendPos.AmountIn.Amount)

	k.DeleteLendForAddressByAsset(ctx, lenderAddr, lendPos.AssetId)

	err = k.UpdateUserLendIdMapping(ctx, addr, lendPos.ID, false)
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

func (k Keeper) BorrowAsset(ctx sdk.Context, addr string, lendId, pairId uint64, loan sdk.Coin) error {

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

	err := k.VerifyCollaterlizationRatio(ctx, lendPos.UpdatedAmountIn, assetIn, loan.Amount, assetOut, pair.LiquidationRatio)
	if err != nil {
		return err
	}
	borrowId := k.GetUserBorrowIDHistory(ctx)

	if !pair.IsInterPool {
		// check sufficient amt in pool to borrow
		reservedAmount := k.GetReserveFunds(ctx, loan.Denom)
		availableAmount := k.ModuleBalance(ctx, AssetOutPool.ModuleName, loan.Denom)

		fmt.Println("reservedAmount..", reservedAmount)
		fmt.Println("availableAmount..", availableAmount)
		fmt.Println("loan.Amount..", loan.Amount)

		if loan.Amount.GT(availableAmount.Sub(reservedAmount)) {
			return sdkerrors.Wrap(types.ErrLendingPoolInsufficient, loan.String())
		}

		AmountIn := sdk.NewCoin(lendPos.AmountIn.Denom, lendPos.UpdatedAmountIn)
		AmountOut := loan
		// take u/Tokens from the user
		cToken, err := k.ExchangeToken(ctx, sdk.NewCoin(lendPos.AmountIn.Denom, lendPos.UpdatedAmountIn), getAsset.Name)
		if err != nil {
			return err
		}

		fmt.Println("cToken...", cToken)

		if err := k.SendCoinFromAccountToModule(ctx, lenderAddr, AssetInPool.ModuleName, cToken); err != nil {
			return err
		}

		if err := k.SendCoinFromModuleToAccount(ctx, AssetOutPool.ModuleName, lenderAddr, loan); err != nil {
			return err
		}

		borrowPos := types.BorrowAsset{
			ID:                   borrowId + 1,
			LendingID:            lendId,
			PairID:               pairId,
			AmountIn:             AmountIn,
			AmountOut:            AmountOut,
			BorrowingTime:        ctx.BlockTime(),
			UpdatedAmountOut:     AmountOut.Amount,
			Interest_Accumulated: sdk.ZeroInt(),
		}
		k.SetUserBorrowIDHistory(ctx, borrowPos.ID)
		k.SetBorrow(ctx, borrowPos)
		k.SetBorrowForAddressByPair(ctx, lenderAddr, pair.AssetOut, borrowPos.ID)
		err = k.UpdateUserBorrowIdMapping(ctx, lendPos.Owner, borrowPos.ID, true)
		if err != nil {
			return err
		}

	}

	return nil
}

func (k Keeper) RepayAsset(ctx sdk.Context, borrowerAddr sdk.AccAddress, payment sdk.Coin) error {
	if !payment.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidAsset, payment.String())
	}

	return nil
}

func (k Keeper) DepositBorrowAsset(ctx sdk.Context, borrowerAddr sdk.AccAddress, payment sdk.Coin) error {
	if !payment.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidAsset, payment.String())
	}

	return nil
}

func (k Keeper) DrawAsset(ctx sdk.Context, borrowerAddr sdk.AccAddress, payment sdk.Coin) error {
	if !payment.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidAsset, payment.String())
	}

	return nil
}

func (k Keeper) CloseBorrow(ctx sdk.Context, borrowerAddr string, borrowId uint64) error {

	return nil
}

func (k Keeper) FundModAcc(ctx sdk.Context, moduleName string, lenderAddr sdk.AccAddress, payment sdk.Coin) error {

	loanTokens := sdk.NewCoins(payment)
	if err := k.bank.SendCoinsFromAccountToModule(ctx, lenderAddr, moduleName, loanTokens); err != nil {
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
