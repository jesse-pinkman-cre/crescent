package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jesse-pinkman-cre/crescent/x/referral/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	k.SetLastReferralID(ctx, genState.LastReferralId)
	for _, referral := range genState.Referrals {
		k.SetReferral(ctx, referral)
		k.SetReferralByAddrIndex(ctx, sdk.AccAddress(referral.Addr), referral.Id)
		k.SetReferralByCodeIndex(ctx, referral.Code, referral.Id)
	}
	for _, revenue := range genState.Revenues {
		k.SetRevenue(ctx, revenue)
	}
}

// ExportGenesis returns the module's exported genesis
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetParams(ctx),
		k.GetLastReferralID(ctx),
		k.GetAllReferrals(ctx),
		k.GetAllRevenues(ctx),
	)
}
