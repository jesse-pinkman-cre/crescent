package keeper

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/jesse-pinkman-cre/crescent/x/referral/types"
)

func (k Keeper) InputOutputCoinsByReferral(ctx sdk.Context, inputs []banktypes.Input, outputs []banktypes.Output) error {
	// 1. check whether output address has referral & parent
	// 	- if not, continue. if yes, continue below
	// 2. check params until MaxParentDepth reached
	// 3. distribute output by two address (owner & parent)
	// 4. recursive until
	// 5. store state for tracking
	fmt.Printf("==============InputOutputCoinsByReferral...================\n")
	var newOutputs []banktypes.Output
	for _, out := range outputs {
		// assume that len of inputs and outputs are same
		fmt.Println(out.Address)

		cntDepth := sdk.NewInt(1)
		for {
			// check whether parent exists
			referral, found := k.GetReferralByAddr(ctx, sdk.MustAccAddressFromBech32(out.Address))
			if !found || referral.Parent == "" {
				fmt.Println("not found referral " + out.Address)
				newOutputs = append(newOutputs, out)
				//store state
				fmt.Println(cntDepth.String() + " " + out.Address)
				if cntDepth.GT(sdk.NewInt(1)) {
					fmt.Println("jihon 18")
					_, err := k.AddRevenue(ctx, referral, out.Coins)
					if err != nil {
						return err
					}
				}
				break
			}
			// query parent's address
			parentReferral, found := k.GetReferralByCode(ctx, referral.Parent)
			if !found {
				fmt.Println("referral parent not found by GetReferralByCode")
				newOutputs = append(newOutputs, out)
				//store state
				if cntDepth.GT(sdk.NewInt(1)) {
					_, err := k.AddRevenue(ctx, referral, out.Coins)
					if err != nil {
						return err
					}
				}
				break
			}
			// check whether GetMaxParentDepth reached
			if cntDepth.GT(k.GetMaxParentDepth(ctx)) {
				fmt.Println("reached to MaxParentDepth")
				newOutputs = append(newOutputs, out)
				//store state
				if cntDepth.GT(sdk.NewInt(1)) {
					_, err := k.AddRevenue(ctx, referral, out.Coins)
					if err != nil {
						return err
					}
				}
				break
			}

			// distribute Coins = []Coin = []struct{Denom, Amount}
			var (
				newCoinsForOwner  sdk.Coins
				newCoinsForParent sdk.Coins
			)
			for _, coin := range out.Coins {
				newAmtForParent := coin.Amount.ToDec().MulTruncate(k.GetSwapFeeReferralRate(ctx)).TruncateInt()
				newAmtForOwner := coin.Amount.Sub(newAmtForParent)
				newCoinsForOwner = append(newCoinsForOwner, sdk.NewCoin(coin.Denom, newAmtForOwner))
				if newAmtForParent.GT(sdk.NewInt(0)) {
					newCoinsForParent = append(newCoinsForParent, sdk.NewCoin(coin.Denom, newAmtForParent))
				}
			}
			outForOwner := banktypes.Output{
				Address: out.Address,
				Coins:   newCoinsForOwner,
			}
			outForParent := banktypes.Output{
				Address: parentReferral.Addr,
				Coins:   newCoinsForParent,
			}
			newOutputs = append(newOutputs, outForOwner)

			//store state
			if cntDepth.GT(sdk.NewInt(1)) {
				_, err := k.AddRevenue(ctx, referral, outForOwner.Coins)
				if err != nil {
					return err
				}
			}

			//recursive distribute
			out = outForParent
			cntDepth = cntDepth.Add(sdk.NewInt(1))
			// ctx.EventManager().EmitEvents(sdk.Events{
			// 	sdk.NewEvent(
			// 		types.EventTypePoolOrderMatched,
			// 		sdk.NewAttribute(types.AttributeKeyOrderDirection, r.OrderDirection.String()),
			// 		sdk.NewAttribute(types.AttributeKeyPairId, strconv.FormatUint(pair.Id, 10)),
			// 		sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(r.PoolId, 10)),
			// 		sdk.NewAttribute(types.AttributeKeyMatchedAmount, r.MatchedAmount.String()),
			// 		sdk.NewAttribute(types.AttributeKeyPaidCoin, r.PaidCoin.String()),
			// 		sdk.NewAttribute(types.AttributeKeyReceivedCoin, r.ReceivedCoin.String()),
			// 	),
			// })
		}
	}
	fmt.Println(newOutputs)
	return k.bankKeeper.InputOutputCoins(ctx, inputs, newOutputs)
}

func (k Keeper) CreateRevenue(ctx sdk.Context, referral types.Referral) (revenue types.Revenue, err error) {
	fmt.Println("==================== CreateRevenue =======================")
	revenue, found := k.GetRevenue(ctx, referral.Id)

	if !found {
		revenue = types.NewRevenue(referral.Id)
		k.SetRevenue(ctx, revenue)
	} else {
		err = errors.New("revenue already exists")
		return
	}

	return revenue, nil
}

func (k Keeper) AddRevenue(ctx sdk.Context, referral types.Referral, coins sdk.Coins) (revenue types.Revenue, err error) {
	fmt.Println("==================== AddRevenue =======================")
	revenue, found := k.GetRevenue(ctx, referral.Id)
	if !found {
		revenue, err = k.CreateRevenue(ctx, referral)
		if err != nil {
			return
		}
	}

	// append revenue to existing revenues
	for _, coin := range coins {
		revenue.Coins = revenue.Coins.Add(coin)
	}
	k.SetRevenue(ctx, revenue)
	return revenue, nil
}
