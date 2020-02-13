package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName is the name of the module
	ModuleName = "zoracle"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

var (
	// GlobalStoreKeyPrefix is a prefix for global primitive state variable
	GlobalStoreKeyPrefix = []byte{0x00}

	// RequestsCountStoreKey is a key that help getting to current requests count state variable
	RequestsCountStoreKey = append(GlobalStoreKeyPrefix, []byte("RequestsCount")...)

	// UnresolvedRequestListStoreKey is a key that help getting pending request
	UnresolvedRequestListStoreKey = append(GlobalStoreKeyPrefix, []byte("PendingList")...)

	// DataSourceCountStoreKey is a key that keeps the current data source count state variable.
	DataSourceCountStoreKey = append(GlobalStoreKeyPrefix, []byte("DataSourceCount")...)

	// ========================================================================

	// RequestStoreKeyPrefix is a prefix for request store
	RequestStoreKeyPrefix = []byte{0x01}

	// ResultStoreKeyPrefix is a prefix for storing result
	ResultStoreKeyPrefix = []byte{0xff}

	// RawDataRequestStoreKeyPrefix is a prefix for storing raw data request detail.
	RawDataRequestStoreKeyPrefix = []byte{0x02}

	// RawDataReportStoreKeyPrefix is a prefix for report store
	RawDataReportStoreKeyPrefix = []byte{0x03}

	// DataSourceStoreKeyPrefix is a prefix for data source store.
	DataSourceStoreKeyPrefix = []byte{0x04}
)

// RequestStoreKey is a function to generate key for each request in store
func RequestStoreKey(requestID int64) []byte {
	return append(RequestStoreKeyPrefix, int64ToBytes(requestID)...)
}

// ResultStoreKey is a function to generate key for each result in store
func ResultStoreKey(requestID int64, codeHash []byte, params []byte) []byte {
	buf := append(ResultStoreKeyPrefix, int64ToBytes(requestID)...)
	buf = append(buf, codeHash...)
	buf = append(buf, params...)
	return buf
}

// RawDataRequestStoreKey is a function to generate key for each raw data request in store
func RawDataRequestStoreKey(requestID, externalID int64) []byte {
	buf := append(RawDataRequestStoreKeyPrefix, int64ToBytes(requestID)...)
	buf = append(buf, int64ToBytes(externalID)...)
	return buf
}

// RawDataReportStoreKey is a function to generate key for each raw data report in store.
func RawDataReportStoreKey(requestID, externalID int64, validatorAddress sdk.ValAddress) []byte {
	buf := append(RawDataReportStoreKeyPrefix, int64ToBytes(requestID)...)
	buf = append(buf, int64ToBytes(externalID)...)
	return append(buf, validatorAddress.Bytes()...)
}

// DataSourceStoreKey is a function to generate key for each data source in store.
func DataSourceStoreKey(dataSourceID int64) []byte {
	return append(DataSourceStoreKeyPrefix, int64ToBytes(dataSourceID)...)
}

// GetIteratorPrefix is a function to get specific prefix
func GetIteratorPrefix(prefix []byte, requestID int64) []byte {
	buf := int64ToBytes(requestID)
	return append(prefix, buf...)
}

// GetValidatorAddress is a function to get validator address from key
func GetValidatorAddress(key []byte, prefix []byte, requestID int64) sdk.ValAddress {
	lenRequest := len(int64ToBytes(requestID))
	return key[len(prefix)+lenRequest:]
}
