package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "restake"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the restake module
	QuerierRoute = ModuleName
)

var (
	GlobalStoreKeyPrefix = []byte{0x00}

	RemainderStoreKey = append(GlobalStoreKeyPrefix, []byte("Remainder")...)

	KeyStoreKeyPrefix    = []byte{0x01}
	StakeStoreKeyPrefix  = []byte{0x02}
	RewardStoreKeyPrefix = []byte{0x03}

	StakesByAmountIndexKeyPrefix = []byte{0x10}
)

func KeyStoreKey(keyName string) []byte {
	return append(KeyStoreKeyPrefix, []byte(keyName)...)
}

func StakesStoreKey(address sdk.AccAddress) []byte {
	return append(StakeStoreKeyPrefix, address...)
}

func StakeStoreKey(address sdk.AccAddress, keyName string) []byte {
	return append(StakesStoreKey(address), []byte(keyName)...)
}

func RewardsStoreKey(address sdk.AccAddress) []byte {
	return append(RewardStoreKeyPrefix, address...)
}

func RewardStoreKey(address sdk.AccAddress, keyName string) []byte {
	return append(RewardsStoreKey(address), []byte(keyName)...)
}

func StakesByAmountIndexKey(address sdk.AccAddress) []byte {
	return append(StakesByAmountIndexKeyPrefix, address...)
}

func StakeByAmountIndexKey(stake Stake) []byte {
	address := sdk.MustAccAddressFromBech32(stake.Address)

	amountBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(amountBytes, stake.Amount.Uint64())

	// key is of format prefix || address || amountBytes || keyBytes
	bz := append(StakesByAmountIndexKey(address), amountBytes...)
	return append(bz, []byte(stake.Key)...)
}
