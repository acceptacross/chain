package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/tss module sentinel errors
var (
	ErrInvalidAccAddressFormat = sdkerrors.Register(ModuleName, 2, "account address format is invalid")
	ErrGroupNotFound           = sdkerrors.Register(ModuleName, 3, "group not found")
	ErrMemberNotFound          = sdkerrors.Register(ModuleName, 4, "member not found")
	ErrRound1DataNotFound      = sdkerrors.Register(ModuleName, 5, "round 1 data not found")
	ErrDKGContextNotFound      = sdkerrors.Register(ModuleName, 6, "dkg context not found")
	ErrMemberNotAuthorized     = sdkerrors.Register(ModuleName, 7, "member is not authorized for this group")
	ErrRoundExpired            = sdkerrors.Register(ModuleName, 8, "round expired")
	ErrAlreadySubmitRound1     = sdkerrors.Register(ModuleName, 9, "already submit round 1 message")
	ErrVerifyOneTimeSigFailed  = sdkerrors.Register(ModuleName, 10, "fail to verify one time sign")
	ErrVerifyA0SigFailed       = sdkerrors.Register(ModuleName, 11, "fail to verify a0 sign")
)
