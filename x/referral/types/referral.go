package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
// ReserveAddressPrefix = "referral/reserve/"
// ShareDenomPrefix     = "amm/share/"
)

func NewReferral(referralID uint64, addr, code string) Referral {
	return Referral{
		Id:     referralID,
		Addr:   addr,
		Code:   code,
		Parent: "",
		Tier:   1,
	}
}

func SetParent(referral Referral, parent string) Referral {
	return Referral{
		Id:     referral.Id,
		Addr:   referral.Addr,
		Code:   referral.Code,
		Parent: parent,
		Tier:   referral.Tier,
	}
}

func NewRevenue(revenueID uint64) Revenue {
	return Revenue{
		ReferralId: revenueID,
		Coins:      sdk.NewCoins(),
	}
}
