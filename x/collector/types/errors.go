package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/collector module sentinel errors
var (
	ErrorUnknownProposalType = sdkerrors.Register(ModuleName, 401, "unknown proposal type")
	ErrorNotFoundForAppId = sdkerrors.Register(ModuleName, 402, "Error Not Found For AppId")
	ErrorAssetDoesNotExist = sdkerrors.Register(ModuleName, 403, "asset does not exist")
	ErrorDuplicateCollectorDenomForApp = sdkerrors.Register(ModuleName, 404, "Collector Duplicate Denom For App")
	ErrorDuplicateAssetDenoms = sdkerrors.Register(ModuleName, 405, "Duplicate Asset Denoms")
)