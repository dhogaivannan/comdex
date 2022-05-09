package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bonding module sentinel errors.
var (
	ErrNotLockOwner                      = sdkerrors.Register(ModuleName, 1, "msg sender is not the owner of specified lock")
	ErrSyntheticBondingAlreadyExists     = sdkerrors.Register(ModuleName, 2, "synthetic bonding already exists for same lock and suffix")
	ErrSyntheticDurationLongerThanNative = sdkerrors.Register(ModuleName, 3, "synthetic bonding duration should be shorter than native bonding duration")
	ErrBondingNotFound                   = sdkerrors.Register(ModuleName, 4, "bonding not found")
)