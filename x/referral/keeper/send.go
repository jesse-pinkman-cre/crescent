package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k Keeper) InputOutputCoinsByReferral(ctx sdk.Context, inputs []banktypes.Input, outputs []banktypes.Output) error {
	// 1. outputs 이 referral 인지 확인
	// 1-1. 대상이 아니라면 바로 전송
	// 2. 수수료율 확인(param)
	// 3. 수수료율에 맞게 코인 전송
	// 4. 대상이라면 recursive 하게 전송
	// 5. 보낸 내역을 store 에 저장
	fmt.Printf("==============InputOutputCoinsByReferral...================\n")
	var newOutputs []banktypes.Output
	// for _, tx := range op.txs {
	// 	inputs = append(inputs, banktypes.NewInput(tx.from, tx.amt))
	// 	outputs = append(outputs, banktypes.NewOutput(tx.to, tx.amt))
	// }
	for _, out := range outputs {
		// assume that len of inputs and outputs are same
		fmt.Println(out.Address)
		//ctx := sdk.UnwrapSDKContext(ctx)
		// check whether parent exists
		referral, found := k.GetReferralByAddr(ctx, sdk.MustAccAddressFromBech32(out.Address))
		if !found || referral.Parent == "" {
			fmt.Println("not found referral")
			newOutputs = append(newOutputs, out)
			continue
		}
		// query parent's address
		parentReferral, found := k.GetReferralByCode(ctx, referral.Parent)
		if !found {
			fmt.Println("referral parent not found by GetReferralByCode")
			newOutputs = append(newOutputs, out)
			continue
		}

		// distribute Coins = []Coin = []struct{Denom, Amount}
		var (
			newCoinsForOwner  sdk.Coins
			newCoinsForParent sdk.Coins
		)
		for _, coin := range out.Coins {
			// amount(int -> dec) * GetMaxParentDepth = newAmount(dec -> int)
			// func GetShareValue(amount sdk.Int, ratio sdk.Dec) sdk.Int {
			// 	return amount.ToDec().MulTruncate(ratio).TruncateInt()
			// }
			newAmtForParent := coin.Amount.ToDec().MulTruncate(k.GetSwapFeeReferralRate(ctx)).TruncateInt()
			newAmtForOwner := coin.Amount.Sub(newAmtForParent)
			newCoinsForOwner = append(newCoinsForOwner, sdk.NewCoin(coin.Denom, newAmtForOwner))
			if newAmtForParent.GT(sdk.NewInt(0)) {
				newCoinsForParent = append(newCoinsForParent, sdk.NewCoin(coin.Denom, newAmtForParent))
			}
		}
		outputForOwner := banktypes.Output{
			Address: out.Address,
			Coins:   newCoinsForOwner,
		}
		outputForParent := banktypes.Output{
			Address: parentReferral.Addr,
			Coins:   newCoinsForParent,
		}
		newOutputs = append(newOutputs, outputForOwner, outputForParent)
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
	fmt.Println(newOutputs)
	return k.bankKeeper.InputOutputCoins(ctx, inputs, newOutputs)
}
