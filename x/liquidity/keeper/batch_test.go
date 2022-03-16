package keeper_test

import (
	utils "github.com/crescent-network/crescent/types"

	_ "github.com/stretchr/testify/suite"
)

func (s *KeeperTestSuite) TestDepositWithdraw() {
	k, ctx := s.keeper, s.ctx

	params := k.GetParams(ctx)

	// Create a normal pool
	creator := s.addr(0)
	s.createPair(creator, "denom1", "denom2", true)
	s.createPool(creator, 1, utils.ParseCoins("1000000denom1,1000000denom2"), true)

	pool, found := k.GetPool(ctx, 1)
	s.Require().True(found)
	s.Require().Equal(params.MinInitialPoolCoinSupply, s.getBalance(creator, pool.PoolCoinDenom).Amount)

	// A depositor makes a deposit
	depositor := s.addr(1)
	s.deposit(depositor, pool.Id, utils.ParseCoins("500000denom1,500000denom2"), true)
	s.nextBlock()

	// The depositor withdraws pool coin
	poolCoin := s.getBalance(depositor, pool.PoolCoinDenom)
	s.withdraw(depositor, pool.Id, poolCoin)
	s.nextBlock()
}
