package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jesse-pinkman-cre/crescent/x/referral/types"
)

func (k Keeper) AddReferral(ctx sdk.Context, fromAddr sdk.AccAddress, code string) (referral types.Referral, err error) {
	_, found := k.GetReferralByAddr(ctx, fromAddr)]
	if found {
		err = errors.New("referral already exists")
		return
	}

	_, found = k.GetReferralByCode(ctx, code)
	if found {
		err = errors.New("duplicate code name")
		return
	}

	referralID := k.GetLastReferralID(ctx)
	referralID++
	k.SetLastReferralID(ctx, referralID)

	referral = types.NewReferral(referralID, fromAddr.String(), code)
	k.SetReferral(ctx, referral)
	k.SetReferralByAddrIndex(ctx, fromAddr, referralID)
	k.SetReferralByCodeIndex(ctx, code, referralID)

	return
}

func (k Keeper) SetParent(ctx sdk.Context, fromAddr sdk.AccAddress, parent string) (referral types.Referral, err error) {
	referral, found := k.GetReferralByAddr(ctx, fromAddr)
	if !found { // Sanity check
		err = sdkerrors.Wrapf(sdkerrors.ErrNotFound, "referral %s not found", fromAddr)
		return
	}
	if len(referral.Parent) > 0 {
		err = errors.New("parent already exists")
		return
	}
	// check whether parent code exists
	_, found = k.GetReferralByCode(ctx, parent)
	if !found { // Sanity check
		err = sdkerrors.Wrapf(sdkerrors.ErrNotFound, "parent code %s not found", parent)
		return
	}

	referral.Parent = parent
	k.SetReferral(ctx, referral)
	return referral, nil
}
