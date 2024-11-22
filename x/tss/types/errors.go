package types

import errorsmod "cosmossdk.io/errors"

// x/tss module sentinel errors
var (
	ErrGroupCreationFailed                = errorsmod.Register(ModuleName, 2, "fail to create group")
	ErrGroupNotFound                      = errorsmod.Register(ModuleName, 3, "group not found")
	ErrMemberNotFound                     = errorsmod.Register(ModuleName, 4, "member not found")
	ErrMemberAlreadySubmit                = errorsmod.Register(ModuleName, 5, "member is already submit message")
	ErrRound1InfoNotFound                 = errorsmod.Register(ModuleName, 6, "round 1 info not found")
	ErrDKGContextNotFound                 = errorsmod.Register(ModuleName, 7, "dkg context not found")
	ErrInvalidGroupStatus                 = errorsmod.Register(ModuleName, 8, "invalid group status")
	ErrVerifyOneTimeSignatureFailed       = errorsmod.Register(ModuleName, 9, "failed to verify one time signature")
	ErrVerifyA0SignatureFailed            = errorsmod.Register(ModuleName, 10, "failed to verify a0 signature")
	ErrAddCoeffCommit                     = errorsmod.Register(ModuleName, 11, "failed to add coefficient commit")
	ErrInvalidLengthCoeffCommits          = errorsmod.Register(ModuleName, 12, "coefficients commit length is invalid")
	ErrRound2InfoNotFound                 = errorsmod.Register(ModuleName, 13, "round 2 info not found")
	ErrInvalidLengthEncryptedSecretShares = errorsmod.Register(
		ModuleName,
		14,
		"encrypted secret shares length is invalid",
	)
	ErrComputeOwnPubKeyFailed       = errorsmod.Register(ModuleName, 15, "failed to compute own public key")
	ErrComplainFailed               = errorsmod.Register(ModuleName, 16, "failed to complain")
	ErrConfirmFailed                = errorsmod.Register(ModuleName, 17, "failed to confirm")
	ErrConfirmNotFound              = errorsmod.Register(ModuleName, 18, "confirm not found")
	ErrComplaintsWithStatusNotFound = errorsmod.Register(ModuleName, 19, "complaints with status not found")
	ErrDENotFound                   = errorsmod.Register(ModuleName, 20, "DE not found")
	ErrGroupIsNotActive             = errorsmod.Register(ModuleName, 21, "group is not active")
	ErrInsufficientSigners          = errorsmod.Register(ModuleName, 22, "insufficient members for signing message")
	ErrBadDrbgInitialization        = errorsmod.Register(ModuleName, 23, "bad drbg initialization")
	ErrPartialSignatureNotFound     = errorsmod.Register(ModuleName, 24, "partial signature not found")
	ErrSubmitSignatureFailed        = errorsmod.Register(ModuleName, 25, "fail to submit signature")
	ErrSigningNotFound              = errorsmod.Register(ModuleName, 26, "signing not found")
	ErrAlreadySigned                = errorsmod.Register(ModuleName, 27, "already signed")
	ErrSigningAlreadySuccess        = errorsmod.Register(ModuleName, 28, "signing already success")
	ErrMemberNotAssigned            = errorsmod.Register(ModuleName, 29, "member is not assigned")
	ErrSubmitSigningSignatureFailed = errorsmod.Register(ModuleName, 30, "failed to submit signing signature")
	ErrDELimitExceeded              = errorsmod.Register(ModuleName, 31, "the number of DEs exceeds the limit")
	ErrHandleSignatureOrderFailed   = errorsmod.Register(ModuleName, 32, "failed to handle signature order")
	ErrHandlerNotFound              = errorsmod.Register(ModuleName, 33, "handler not found")
	ErrInvalidCoefficientCommit     = errorsmod.Register(ModuleName, 34, "invalid coefficient commit")
	ErrInvalidPublicKey             = errorsmod.Register(ModuleName, 35, "invalid public key")
	ErrInvalidSignature             = errorsmod.Register(ModuleName, 36, "invalid signature")
	ErrInvalidSecretShare           = errorsmod.Register(ModuleName, 37, "invalid secret share")
	ErrInvalidComplaint             = errorsmod.Register(ModuleName, 38, "invalid complaint")
	ErrInvalidDE                    = errorsmod.Register(ModuleName, 39, "invalid DE")
	ErrMaxSigningAttemptExceeded    = errorsmod.Register(ModuleName, 40, "signing attempt exceeds the limit")
	ErrEncodeOriginatorFailed       = errorsmod.Register(ModuleName, 41, "failed to encode originator")
	ErrInvalidOriginator            = errorsmod.Register(ModuleName, 42, "invalid originator")
	ErrInvalidMessage               = errorsmod.Register(ModuleName, 43, "invalid message")
	ErrSigningAttemptNotFound       = errorsmod.Register(ModuleName, 44, "signing attempt not found")
	ErrInvalidMember                = errorsmod.Register(ModuleName, 45, "invalid member")
	ErrInvalidGroup                 = errorsmod.Register(ModuleName, 46, "invalid group")
	ErrInvalidSigning               = errorsmod.Register(ModuleName, 47, "invalid signing")
	ErrCreateSigningFailed          = errorsmod.Register(ModuleName, 48, "failed to create signing")
)