package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgLend(lender sdk.AccAddress, pairID uint64, amount sdk.Coin) *MsgLend {
	return &MsgLend{
		Lender: lender.String(),
		PairId: pairID,
		Amount: amount,
	}
}

func (msg MsgLend) Route() string { return ModuleName }
func (msg MsgLend) Type() string  { return EventTypeLoanAsset }

func (msg *MsgLend) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetLender())
	if err != nil {
		return err
	}

	if asset := msg.GetAmount(); !asset.IsValid() {
		return sdkerrors.Wrap(ErrInvalidAsset, asset.String())
	}

	return nil
}

func (msg *MsgLend) GetSigners() []sdk.AccAddress {
	lender, _ := sdk.AccAddressFromBech32(msg.GetLender())
	return []sdk.AccAddress{lender}
}

// GetSignBytes get the bytes for the message signer to sign on
func (msg *MsgLend) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func NewMsgWithdraw(lender sdk.AccAddress, lendID uint64, amount sdk.Coin) *MsgWithdraw {
	return &MsgWithdraw{
		Lender: lender.String(),
		LendId: lendID,
		Amount: amount,
	}
}

func (msg MsgWithdraw) Route() string { return ModuleName }
func (msg MsgWithdraw) Type() string  { return EventTypeWithdrawLoanedAsset }

func (msg *MsgWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetLender())
	if err != nil {
		return err
	}

	if asset := msg.GetAmount(); !asset.IsValid() {
		return sdkerrors.Wrap(ErrInvalidAsset, asset.String())
	}

	return nil
}

func (msg *MsgWithdraw) GetSigners() []sdk.AccAddress {
	lender, _ := sdk.AccAddressFromBech32(msg.GetLender())
	return []sdk.AccAddress{lender}
}

// GetSignBytes get the bytes for the message signer to sign on
func (msg *MsgWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func NewMsgDeposit(lender sdk.AccAddress, lendID uint64, amount sdk.Coin) *MsgDeposit {
	return &MsgDeposit{
		From:   lender.String(),
		LendId: lendID,
		Amount: amount,
	}
}

func (msg MsgDeposit) Route() string { return ModuleName }
func (msg MsgDeposit) Type() string  { return EventTypeLoanAsset }

func (msg *MsgDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetFrom())
	if err != nil {
		return err
	}

	if asset := msg.GetAmount(); !asset.IsValid() {
		return sdkerrors.Wrap(ErrInvalidAsset, asset.String())
	}

	return nil
}

func (msg *MsgDeposit) GetSigners() []sdk.AccAddress {
	lender, _ := sdk.AccAddressFromBech32(msg.GetFrom())
	return []sdk.AccAddress{lender}
}

// GetSignBytes get the bytes for the message signer to sign on
func (msg *MsgDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func NewMsgBorrow(borrower sdk.AccAddress, pairID uint64, amountIn, amountOut sdk.Coin) *MsgBorrow {
	return &MsgBorrow{
		Borrower:  borrower.String(),
		AmountIn:  amountIn,
		AmountOut: amountOut,
		LendId:    pairID,
	}
}

func (msg MsgBorrow) Route() string { return ModuleName }
func (msg MsgBorrow) Type() string  { return EventTypeLoanAsset }

func (msg *MsgBorrow) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetBorrower())
	if err != nil {
		return err
	}

	if asset := msg.GetAmountIn(); !asset.IsValid() {
		return sdkerrors.Wrap(ErrInvalidAsset, asset.String())
	}
	return nil
}

func (msg *MsgBorrow) GetSigners() []sdk.AccAddress {
	lender, _ := sdk.AccAddressFromBech32(msg.GetBorrower())
	return []sdk.AccAddress{lender}
}

// GetSignBytes get the bytes for the message signer to sign on
func (msg *MsgBorrow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func NewMsgDraw(borrower sdk.AccAddress, borrowID uint64, amount sdk.Coin) *MsgDraw {
	return &MsgDraw{
		Borrower: borrower.String(),
		BorrowId: borrowID,
		Amount:   amount,
	}
}

func (msg MsgDraw) Route() string { return ModuleName }
func (msg MsgDraw) Type() string  { return EventTypeLoanAsset }

func (msg *MsgDraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetBorrower())
	if err != nil {
		return err
	}

	if asset := msg.GetAmount(); !asset.IsValid() {
		return sdkerrors.Wrap(ErrInvalidAsset, asset.String())
	}

	return nil
}

func (msg *MsgDraw) GetSigners() []sdk.AccAddress {
	lender, _ := sdk.AccAddressFromBech32(msg.GetBorrower())
	return []sdk.AccAddress{lender}
}

// GetSignBytes get the bytes for the message signer to sign on
func (msg *MsgDraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func NewMsgRepay(borrower sdk.AccAddress, borrowID uint64, amount sdk.Coin) *MsgRepay {
	return &MsgRepay{
		Borrower: borrower.String(),
		BorrowId: borrowID,
		Amount:   amount,
	}
}

func (msg MsgRepay) Route() string { return ModuleName }
func (msg MsgRepay) Type() string  { return EventTypeLoanAsset }

func (msg *MsgRepay) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetBorrower())
	if err != nil {
		return err
	}

	if asset := msg.GetAmount(); !asset.IsValid() {
		return sdkerrors.Wrap(ErrInvalidAsset, asset.String())
	}

	return nil
}

func (msg *MsgRepay) GetSigners() []sdk.AccAddress {
	lender, _ := sdk.AccAddressFromBech32(msg.GetBorrower())
	return []sdk.AccAddress{lender}
}

// GetSignBytes get the bytes for the message signer to sign on
func (msg *MsgRepay) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func NewMsgFundModuleAccounts(moduleName string, lender sdk.AccAddress, amount sdk.Coin) *MsgFundModuleAccounts {
	return &MsgFundModuleAccounts{
		ModuleName: moduleName,
		Lender:     lender.String(),
		Amount:     amount,
	}
}

func (msg MsgFundModuleAccounts) Route() string { return ModuleName }
func (msg MsgFundModuleAccounts) Type() string  { return EventTypeLoanAsset }

func (msg *MsgFundModuleAccounts) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetLender())
	if err != nil {
		return err
	}

	if asset := msg.GetAmount(); !asset.IsValid() {
		return sdkerrors.Wrap(ErrInvalidAsset, asset.String())
	}

	return nil
}

func (msg *MsgFundModuleAccounts) GetSigners() []sdk.AccAddress {
	lender, _ := sdk.AccAddressFromBech32(msg.GetLender())
	return []sdk.AccAddress{lender}
}

// GetSignBytes get the bytes for the message signer to sign on
func (msg *MsgFundModuleAccounts) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
