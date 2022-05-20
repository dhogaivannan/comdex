package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidAmount    = errors.Register(ModuleName, 101, "invalid amount")
	ErrorInvalidAmountIn  = errors.Register(ModuleName, 102, "invalid amount_in")
	ErrorInvalidAmountOut = errors.Register(ModuleName, 103, "invalid amount_out")
	ErrorInvalidFrom      = errors.Register(ModuleName, 104, "invalid from")
	ErrorInvalidID        = errors.Register(ModuleName, 105, "invalid id")
	ErrorAppIstoExtendedAppId        = errors.Register(ModuleName, 106, "app id does not match with extended pair app id")
)

var (
	
	// ErrorUnauthorized                  = errors.Register(ModuleName, 203, "unauthorized")
	// ErrorDuplicateVault                = errors.Register(ModuleName, 204, "duplicate vault")
	
	ErrorExtendedPairVaultDoesNotExists= errors.Register(ModuleName, 201, "Extended pair vault does not exists for the given id")
	ErrorAppMappingDoesNotExist= errors.Register(ModuleName, 202, "App Mapping Id does not exists")
	ErrorAppMappingIdMismatch= errors.Register(ModuleName, 203, "App Mapping Id mismatch, use the correct App Mapping ID in request")
	ErrorVaultCreationInactive=errors.Register(ModuleName, 204, "Vault Creation Inactive")
	ErrorUserVaultAlreadyExists=errors.Register(ModuleName, 205, "User vault already exists for teh given extended pair vault id ")
	ErrorAmountOutLessThanDebtFloor=errors.Register(ModuleName, 206, "Amount Out is less than Debt Floor")
	ErrorAmountOutGreaterThanDebtCeiling=errors.Register(ModuleName, 207, "Amount Out is greater than Debt Ceiling")
	ErrorPairDoesNotExist              = errors.Register(ModuleName, 208, "Pair does not exists")
    ErrorAssetDoesNotExist             = errors.Register(ModuleName, 210, "Asset does not exists")
	ErrorPriceDoesNotExist             = errors.Register(ModuleName, 211, "Price does not exist")
	ErrorInvalidCollateralizationRatio = errors.Register(ModuleName, 212, "Invalid collateralization ratio")
	ErrorVaultDoesNotExist             = errors.Register(ModuleName, 213, "Vault does not exist")
	ErrVaultAccessUnauthorised		=errors.Register(ModuleName, 214, "Unauthorized user for the tx")
	ErrorInvalidAppMappingData		=errors.Register(ModuleName, 215, "Invalid App Mapping data sent as compared to data exists in vault")
	ErrorInvalidExtendedPairMappingData = errors.Register(ModuleName, 215, "Invalid Extended Pair Vault Mapping data sent as compared to data exists in vault")
	ErrorVaultInactive=errors.Register(ModuleName, 216, "Vault tx Inactive")



)

var (
	ErrorUnknownMsgType = errors.Register(ModuleName, 301, "unknown message type")
)

var (
	ErrorCannotCreateStableMintVault   = errors.Register(ModuleName, 401, "Cannot Create Stable Mint Vault, StableMint tx command")
	ErrorAmtGreaterDebt  = errors.Register(ModuleName, 402, "amt should be > debt floor")
	ErrorIdnotFound  = errors.Register(ModuleName, 403, "not found")
)