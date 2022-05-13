package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName   = "vault"
	QuerierRoute = ModuleName
	RouterKey    = ModuleName
	StoreKey     = ModuleName
)

var (
	TypeMsgCreateRequest    = ModuleName + ":create"
	TypeMsgDepositRequest   = ModuleName + ":deposit"
	TypeMsgWithdrawRequest  = ModuleName + ":withdraw"
	TypeMsgDrawRequest      = ModuleName + ":draw"
	TypeMsgRepayRequest     = ModuleName + ":repay"
	TypeMsgLiquidateRequest = ModuleName + ":liquidate"
)

var (
	IDKey                          = []byte{0x00}
	VaultKeyPrefix                 = []byte{0x10}
	VaultForAddressByPairKeyPrefix = []byte{0x20}
	AppMappingPrefixKey            = []byte{0x30}
	LookUpTablePrefixKey           = []byte{0x40}
	UserVaultIdPrefixKey           = []byte{0x50}
	TokenMintPrefixKey             = []byte{0x60}
	CAssetMintStatisticsKeyPrefix  = []byte{0x70}
)

func VaultKey(appVaultTypeId string) []byte {
	return append(VaultKeyPrefix, []byte(appVaultTypeId)...)
}

func VaultForAddressByAppAndPair(address sdk.AccAddress, appVaultTypeId string, pairID uint64) []byte {
	if len(address.Bytes()) != 20 {
		panic(fmt.Errorf("invalid address length %d; expected %d", len(address.Bytes()), 20))
	}
	appVaultTypeIdBytes := []byte(appVaultTypeId)
	v := append(append(VaultForAddressByPairKeyPrefix, appVaultTypeIdBytes...), address.Bytes()...)
	return append(v, sdk.Uint64ToBigEndian(pairID)...)
}

func GetAppMappingIdPrefixKey(id uint64) []byte {
	return append(AppMappingPrefixKey, sdk.Uint64ToBigEndian(id)...)
}

func GetLookUpTablePrefixKey(id uint64) []byte {
	return append(LookUpTablePrefixKey, sdk.Uint64ToBigEndian(id)...)
}

func GetUserVaultIdPrefixKey(address sdk.AccAddress) []byte {
	return append(UserVaultIdPrefixKey, address.Bytes()...)
}

func GetTokenMintPrefixKey(collateralDenom string) []byte {
	return append(UserVaultIdPrefixKey, []byte(collateralDenom)...)
}

func GetCAssetMintRecordsKey(collateralDenom string) []byte {

	return append(CAssetMintStatisticsKeyPrefix, collateralDenom...)
}
