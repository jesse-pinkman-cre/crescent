package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jesse-pinkman-cre/crescent/x/referral/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	fmt.Println("===================================================")
	fmt.Println(k.bankKeeper.GetSupply(ctx, "ucre"))
	fmt.Println("===================================================")

	k.paramSpace.GetParamSet(ctx, &params)
	return
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetMaxParentDepth returns the maximum parent depth for bonus distribution parameter.
func (k Keeper) GetMaxParentDepth(ctx sdk.Context) (maxParentDepth sdk.Int) {
	k.paramSpace.Get(ctx, types.KeyMaxParentDepth, &maxParentDepth)
	return
}

// SetMaxParentDepth sets the maximum parent depth for bonus distribution.
func (k Keeper) SetMaxParentDepth(ctx sdk.Context, i sdk.Int) {
	k.paramSpace.Set(ctx, types.KeyMaxParentDepth, i)
}

// GetMaxParentDepth returns the maximum parent depth for bonus distribution parameter.
func (k Keeper) GetSwapFeeReferralRate(ctx sdk.Context) (swapFeeReferralRate sdk.Dec) {
	k.paramSpace.Get(ctx, types.KeySwapFeeReferralRate, &swapFeeReferralRate)
	return
}
