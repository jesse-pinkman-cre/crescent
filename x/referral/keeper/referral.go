package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jesse-pinkman-cre/crescent/x/referral/types"
)

func (k Keeper) AddReferral(ctx sdk.Context, fromAddr sdk.AccAddress, code string) (referral types.Referral, err error) {
	referral, found := k.GetReferralByAddr(ctx, fromAddr)

	// empty liquidity pool
	if !found {
		// check whether code already exists
		_, found := k.GetReferralByCode(ctx, code)
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
	} else {
		err = errors.New("referral already exists")
		return
	}

	// mintedShare = sdk.NewCoin(shareDenom, share)
	// if err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(mintedShare)); err != nil {
	// 	return
	// }
	// if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, fromAddr, sdk.NewCoins(mintedShare)); err != nil {
	// 	return
	// }
	return referral, nil
}
