package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jesse-pinkman-cre/crescent/x/liquidity/types"
)

// RegisterInvariants registers all liquidity module invariants.
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "deposit-coins-escrow", DepositCoinsEscrowInvariant(k))
	ir.RegisterRoute(types.ModuleName, "pool-coin-escrow", PoolCoinEscrowInvariant(k))
	ir.RegisterRoute(types.ModuleName, "remaining-offer-coin-escrow", RemainingOfferCoinEscrowInvariant(k))
	ir.RegisterRoute(types.ModuleName, "pool-status", PoolStatusInvariant(k))
	ir.RegisterRoute(types.ModuleName, "num-mm-orders", NumMMOrdersInvariant(k))
}

// AllInvariants returns a combined invariant of the liquidity module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		for _, inv := range []func(Keeper) sdk.Invariant{
			DepositCoinsEscrowInvariant,
			PoolCoinEscrowInvariant,
			RemainingOfferCoinEscrowInvariant,
			PoolStatusInvariant,
			NumMMOrdersInvariant,
		} {
			res, stop := inv(k)(ctx)
			if stop {
				return res, stop
			}
		}
		return "", false
	}
}

// DepositCoinsEscrowInvariant checks that the amount of coins in the global
// escrow address is greater or equal than remaining deposit coins in all
// deposit requests.
func DepositCoinsEscrowInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		escrowDepositCoins := sdk.Coins{}
		_ = k.IterateAllDepositRequests(ctx, func(req types.DepositRequest) (stop bool, err error) {
			if req.Status == types.RequestStatusNotExecuted {
				escrowDepositCoins = escrowDepositCoins.Add(req.DepositCoins...)
			}
			return false, nil
		})
		balances := k.bankKeeper.SpendableCoins(ctx, types.GlobalEscrowAddress)
		broken := !balances.IsAllGTE(escrowDepositCoins)
		return sdk.FormatInvariant(
			types.ModuleName, "deposit-coins-escrow",
			fmt.Sprintf("escrow amount %s is smaller than expected %s", balances, escrowDepositCoins),
		), broken
	}
}

// PoolCoinEscrowInvariant checks that the amount of coins in the global
// escrow address is greater or equal than remaining withdrawing pool
// coins in all withdrawal requests.
func PoolCoinEscrowInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		escrowPoolCoins := sdk.Coins{}
		_ = k.IterateAllWithdrawRequests(ctx, func(req types.WithdrawRequest) (stop bool, err error) {
			if req.Status == types.RequestStatusNotExecuted {
				escrowPoolCoins = escrowPoolCoins.Add(req.PoolCoin)
			}
			return false, nil
		})
		balances := k.bankKeeper.SpendableCoins(ctx, types.GlobalEscrowAddress)
		broken := !balances.IsAllGTE(escrowPoolCoins)
		return sdk.FormatInvariant(
			types.ModuleName, "pool-coin-escrow",
			fmt.Sprintf("escrow amount %s is smaller than expected %s", balances, escrowPoolCoins),
		), broken
	}
}

// RemainingOfferCoinEscrowInvariant checks that the amount of coins in each pair's
// escrow address is greater or equal than remaining offer coins in the pair's
// orders.
func RemainingOfferCoinEscrowInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var (
			count int
			msg   string
		)
		_ = k.IterateAllPairs(ctx, func(pair types.Pair) (stop bool, err error) {
			remainingOfferCoins := sdk.Coins{}
			_ = k.IterateOrdersByPair(ctx, pair.Id, func(req types.Order) (stop bool, err error) {
				if !req.Status.ShouldBeDeleted() {
					remainingOfferCoins = remainingOfferCoins.Add(req.RemainingOfferCoin)
				}
				return false, nil
			})
			balances := k.bankKeeper.SpendableCoins(ctx, pair.GetEscrowAddress())
			if !balances.IsAllGTE(remainingOfferCoins) {
				count++
				msg += fmt.Sprintf("\tpair %d has %s, which is smaller than %s\n", pair.Id, balances, remainingOfferCoins)
			}
			return false, nil
		})
		broken := count != 0
		return sdk.FormatInvariant(
			types.ModuleName, "remaining-offer-coin-escrow",
			fmt.Sprintf("%d pair(s) with insufficient escrow amount found\n%s", count, msg),
		), broken
	}
}

// PoolStatusInvariant checks that the pools with zero pool coin supply have
// been marked as disabled.
func PoolStatusInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var (
			count int
			msg   string
		)
		_ = k.IterateAllPools(ctx, func(pool types.Pool) (stop bool, err error) {
			if !pool.Disabled {
				ps := k.GetPoolCoinSupply(ctx, pool)
				if ps.IsZero() {
					count++
					msg += fmt.Sprintf("\tpool %d should be disabled, but not\n", pool.Id)
				}
			}
			return false, nil
		})
		broken := count != 0
		return sdk.FormatInvariant(
			types.ModuleName, "pool-status",
			fmt.Sprintf("%d pool(s) with wrong status found\n%s", count, msg),
		), broken
	}
}

// NumMMOrdersInvariant checks the actual number of mm orders with NumMMOrders
// records.
func NumMMOrdersInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		// orderer address => (pair id => num mm orders)
		numMMOrdersMap := map[string]map[uint64]uint32{}
		_ = k.IterateAllOrders(ctx, func(order types.Order) (stop bool, err error) {
			if order.Type == types.OrderTypeMM && !order.Status.ShouldBeDeleted() {
				numMMOrdersByPairId, ok := numMMOrdersMap[order.Orderer]
				if !ok {
					numMMOrdersByPairId = map[uint64]uint32{}
					numMMOrdersMap[order.Orderer] = numMMOrdersByPairId
				}
				numMMOrdersByPairId[order.PairId]++
			}
			return false, nil
		})
		var (
			count int
			msg   string
		)
		k.IterateAllNumMMOrders(ctx, func(ordererAddr sdk.AccAddress, pairId uint64, numMMOrders uint32) (stop bool) {
			orderer := ordererAddr.String()
			numMMOrdersByPairId, ok := numMMOrdersMap[orderer]
			if !ok {
				count++
				msg += fmt.Sprintf("\torderer %s has no mm orders\n", orderer)
				return false
			}
			correctNumMMOrders, ok := numMMOrdersByPairId[pairId]
			if !ok {
				count++
				msg += fmt.Sprintf("\torderer %s has no mm orders in pair %d\n", orderer, pairId)
				return false
			}
			if numMMOrders != correctNumMMOrders {
				count++
				msg += fmt.Sprintf("\torderer %s has wrong number of mm orders in pair %d; got %d, expected %d\n", orderer, pairId, numMMOrders, correctNumMMOrders)
			}
			return false
		})
		broken := count != 0
		return sdk.FormatInvariant(
			types.ModuleName, "num-mm-orders",
			fmt.Sprintf("%d orderer-pair pairs with wrong status\n%s", count, msg),
		), broken
	}
}
