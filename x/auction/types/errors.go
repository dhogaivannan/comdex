package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/auction module sentinel errors
var (
	ErrorInvalidAuctionId             = sdkerrors.Register(ModuleName, 101, "auction does not exist with given id")
	ErrorInvalidBiddingDenom          = sdkerrors.Register(ModuleName, 102, "given asset type is not accepted for bidding")
	ErrorLowBidAmount                 = sdkerrors.Register(ModuleName, 103, "bidding amount is lower than expected")
	ErrorMaxBidAmount                 = sdkerrors.Register(ModuleName, 104, "bidding amount is greater than maximum bidding amount")
	ErrorBidAlreadyExists             = sdkerrors.Register(ModuleName, 105, "bid with given amount already placed, Please try with higher bid")
	ErrorInvalidAuctioningCollateral  = sdkerrors.Register(ModuleName, 106, "collateral to be auctioned <= 0")
	ErrorInvalidAmountInAddress       = sdkerrors.Register(ModuleName, 107, "there is not sufficient balance in given address for a given denom")
	ErrorInvalidAddress               = sdkerrors.Register(ModuleName, 107, "invalid source address")
	ErrorInvalidDebtAuctionId         = sdkerrors.Register(ModuleName, 108, "debt auction does not exist with given id")
	ErrorInvalidDebtUserExpectedDenom = sdkerrors.Register(ModuleName, 109, "given asset type is not accepted for debt auction user expected token")
	ErrorDebtMoreBidAmount            = sdkerrors.Register(ModuleName, 110, "can not bid more minted amount")
	ErrorDebtExpectedUserAmount       = sdkerrors.Register(ModuleName, 111, "invalid user amount")
	ErrorInvalidDebtMintedDenom       = sdkerrors.Register(ModuleName, 112, "given asset type is not accepted for debt auction user mint token")
)

var (
	ErrorUnknownMsgType = sdkerrors.Register(ModuleName, 301, "unknown message type")
)
