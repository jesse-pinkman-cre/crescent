package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/jesse-pinkman-cre/crescent/x/referral/types"
)

func (k Keeper) GetLastReferralID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LastReferralIDKey)
	if bz == nil {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetLastReferralID(ctx sdk.Context, lastReferralID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastReferralIDKey, sdk.Uint64ToBigEndian(lastReferralID))
}

func (k Keeper) GetReferral(ctx sdk.Context, referralID uint64) (referral types.Referral, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetReferralKey(referralID))
	if bz == nil {
		return
	}
	k.cdc.MustUnmarshal(bz, &referral)
	return referral, true
}

func (k Keeper) SetReferral(ctx sdk.Context, referral types.Referral) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&referral)
	store.Set(types.GetReferralKey(referral.Id), bz)
}

func (k Keeper) DeleteReferral(ctx sdk.Context, referral types.Referral) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetReferralKey(referral.Id))
}

func (k Keeper) IterateAllReferrals(ctx sdk.Context, cb func(referral types.Referral) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ReferralKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var referral types.Referral
		k.cdc.MustUnmarshal(iter.Value(), &referral)
		if cb(referral) {
			break
		}
	}
}

func (k Keeper) GetAllReferrals(ctx sdk.Context) (referrals []types.Referral) {
	k.IterateAllReferrals(ctx, func(referral types.Referral) (stop bool) {
		referrals = append(referrals, referral)
		return false
	})
	return
}

func (k Keeper) GetReferralByAddr(ctx sdk.Context, addr sdk.AccAddress) (referral types.Referral, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetReferralByAddrIndexKey(addr))
	if bz == nil {
		return
	}
	//return k.GetReferral(ctx, sdk.BigEndianToUint64(bz))
	var val gogotypes.UInt64Value
	k.cdc.MustUnmarshal(bz, &val)
	referralID := val.GetValue()
	return k.GetReferral(ctx, referralID)
}

func (k Keeper) GetReferralByCode(ctx sdk.Context, code string) (referral types.Referral, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetReferralByCodeIndexKey(code))
	if bz == nil {
		return
	}
	//return k.GetReferral(ctx, sdk.BigEndianToUint64(bz))
	var val gogotypes.UInt64Value
	k.cdc.MustUnmarshal(bz, &val)
	referralID := val.GetValue()
	return k.GetReferral(ctx, referralID)
}

func (k Keeper) SetReferralByAddrIndex(ctx sdk.Context, addr sdk.AccAddress, referralID uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: referralID})
	store.Set(types.GetReferralByAddrIndexKey(addr), bz)
}

func (k Keeper) SetReferralByCodeIndex(ctx sdk.Context, code string, referralID uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: referralID})
	store.Set(types.GetReferralByCodeIndexKey(code), bz)
}

func (k Keeper) DeleteReferralByAddrIndex(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetReferralByAddrIndexKey(addr))
}

func (k Keeper) DeleteReferralByCodeIndex(ctx sdk.Context, code string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetReferralByCodeIndexKey(code))
}
