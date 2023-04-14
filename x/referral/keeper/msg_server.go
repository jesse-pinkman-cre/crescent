package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jesse-pinkman-cre/crescent/x/referral/types"
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

func (m msgServer) AddReferral(c context.Context, msg *types.MsgAddReferral) (*types.MsgAddReferralResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	referral, err := m.Keeper.AddReferral(
		ctx, sdk.MustAccAddressFromBech32(msg.Sender), msg.Code,
	)
	if err != nil {
		return nil, err
	}
	return &types.MsgAddReferralResponse{
		Referral: referral,
	}, nil
}

func (m msgServer) SetReferral(c context.Context, msg *types.MsgSetReferral) (*types.MsgSetReferralResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	referral, err := m.Keeper.SetParent(
		ctx, sdk.MustAccAddressFromBech32(msg.Sender), msg.Parent)
	if err != nil {
		return nil, err
	}
	return &types.MsgSetReferralResponse{
		Referral: referral,
	}, nil
}
