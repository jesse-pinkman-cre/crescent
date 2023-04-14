package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "referral"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_referral"
)

var (
	LastReferralIDKey = []byte{0xdf} // key for the latest id

	ReferralKeyPrefix            = []byte{0xde}
	ReferralByAddrIndexKeyPrefix = []byte{0xdd}
	ReferralByCodeIndexKeyPrefix = []byte{0xdc}
	RevenueKeyPrefix             = []byte{0xdb}
)

func GetReferralKey(referralID uint64) []byte {
	return append(ReferralKeyPrefix, sdk.Uint64ToBigEndian(referralID)...)
}

func GetReferralByAddrIndexKey(addr sdk.AccAddress) []byte {
	return append(ReferralByAddrIndexKeyPrefix, address.MustLengthPrefix(addr)...)
	//addr.String()
}

func GetReferralByCodeIndexKey(code string) []byte {
	return append(ReferralByCodeIndexKeyPrefix, LengthPrefixString(code)...)
}

func GetRevenueKey(revenueID uint64) []byte {
	return append(RevenueKeyPrefix, sdk.Uint64ToBigEndian(revenueID)...)
}

// LengthPrefixString returns length-prefixed bytes representation
// of a string.
func LengthPrefixString(s string) []byte {
	bz := []byte(s)
	bzLen := len(bz)
	if bzLen == 0 {
		return bz
	}
	return append([]byte{byte(bzLen)}, bz...)
}
