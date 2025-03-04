package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/curating module sentinel errors
var (
	ErrStaketNotFound     = sdkerrors.Register(ModuleName, 1, "stake not found")
	ErrAmountTooLarge     = sdkerrors.Register(ModuleName, 2, "unstake amount too large")
	ErrCurationNotExpired = sdkerrors.Register(ModuleName, 3, "post is still being curated")
)
