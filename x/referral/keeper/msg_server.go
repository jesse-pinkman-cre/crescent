package keeper

import (
	"github.com/cosmos-builders/chaos/x/amm/types"
	//"github.com/jesse-pinkman-cre/crescent/x/referral/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}
