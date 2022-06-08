package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName   = "auction"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	MemStoreKey  = ModuleName

	ActiveAuctionStatus = "active"
	ClosedAuctionStatus = "inactive"

	PlacedBiddingStatus   = "placed"
	RejectedBiddingStatus = "rejected"
	SuccessBiddingStatus  = "success"
)

const AuctionStartNoBids uint64 = 0
const AuctionGoingOn uint64 = 1
const AuctionEnded uint64 = 2

const NoAuction uint64 = 0
const StartedSurplusAuction uint64 = 1
const StartedDebtAuction uint64 = 2
const SurplusString = "surplus"
const DebtString = "debt"
const DutchString = "dutch"
const TestFlag = 0

var (
	AuctionKeyPrefix            = []byte{0x11}
	UserKeyPrefix               = []byte{0x12}
	AuctionIdKey                = []byte{0x13}
	UserBiddingsIdKey           = []byte{0x14}
	HistoryAuctionKeyPrefix     = []byte{0x15}
	HistoryUserKeyPrefix        = []byte{0x16}
	ProtocolStatisticsPrefixKey = []byte{0x17}
	AuctionParamsKeyPrefix      = []byte{0x18}
)

func AuctionKey(appId uint64, auctionType string, auctionId uint64) []byte {
	return append(append(append(AuctionKeyPrefix, sdk.Uint64ToBigEndian(appId)...), auctionType...), sdk.Uint64ToBigEndian(auctionId)...)
}

func UserKey(bidder string, appId uint64, auctionType string, bidId uint64) []byte {
	return append(append(append(append(UserKeyPrefix, bidder...), sdk.Uint64ToBigEndian(appId)...), auctionType...), sdk.Uint64ToBigEndian(bidId)...)
}

func UserAuctionTypeKey(bidder string, appId uint64, auctionType string) []byte {
	return append(append(append(UserKeyPrefix, bidder...), sdk.Uint64ToBigEndian(appId)...), auctionType...)
}

func AuctionTypeKey(appId uint64, auctionType string) []byte {
	return append(append(AuctionKeyPrefix, sdk.Uint64ToBigEndian(appId)...), auctionType...)
}

func HistoryAuctionKey(appId uint64, auctionType string, auctionId uint64) []byte {
	return append(append(append(HistoryAuctionKeyPrefix, sdk.Uint64ToBigEndian(appId)...), auctionType...), sdk.Uint64ToBigEndian(auctionId)...)
}

func HistoryUserKey(bidder string, appId uint64, auctionType string, bidId uint64) []byte {
	return append(append(append(append(HistoryUserKeyPrefix, bidder...), sdk.Uint64ToBigEndian(appId)...), auctionType...), sdk.Uint64ToBigEndian(bidId)...)
}

func HistoryUserAuctionTypeKey(bidder string, appId uint64, auctionType string) []byte {
	return append(append(append(HistoryUserKeyPrefix, bidder...), sdk.Uint64ToBigEndian(appId)...), auctionType...)
}

func HistoryAuctionTypeKey(appId uint64, auctionType string) []byte {
	return append(append(HistoryAuctionKeyPrefix, sdk.Uint64ToBigEndian(appId)...), auctionType...)
}

func ProtocolStatisticsKey(appId, assetId uint64) []byte {
	return append(append(ProtocolStatisticsPrefixKey, sdk.Uint64ToBigEndian(appId)...), sdk.Uint64ToBigEndian(assetId)...)
}

func ProtocolStatisticsAppIdKey(appId uint64) []byte {
	return append(ProtocolStatisticsPrefixKey, sdk.Uint64ToBigEndian(appId)...)
}

func AuctionParamsKey(id uint64) []byte {
	return append(AuctionParamsKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}
