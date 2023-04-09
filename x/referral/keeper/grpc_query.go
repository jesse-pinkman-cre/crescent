package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jesse-pinkman-cre/crescent/x/referral/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(k Keeper) Querier {
	return Querier{k}
}

func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Querier) Referrals(c context.Context, req *types.QueryReferralsRequest) (*types.QueryReferralsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	planStore := prefix.NewStore(store, types.ReferralKeyPrefix)
	var referrals []types.Referral
	pageRes, err := query.Paginate(planStore, req.Pagination, func(key, value []byte) error {
		var referral types.Referral
		if err := k.cdc.Unmarshal(value, &referral); err != nil {
			return err
		}
		referrals = append(referrals, referral)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryReferralsResponse{
		Referrals:  referrals,
		Pagination: pageRes,
	}, nil
}

func (k Querier) Referral(c context.Context, req *types.QueryReferralRequest) (*types.QueryReferralResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	referral, found := k.GetReferral(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "referral %d not found", req.Id)
	}
	return &types.QueryReferralResponse{
		Referral: referral,
	}, nil
}

func (k Querier) ReferralByAddr(c context.Context, req *types.QueryReferralByAddrRequest) (*types.QueryReferralResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	referral, found := k.GetReferralByAddr(ctx, sdk.MustAccAddressFromBech32(req.Addr))
	if !found {
		return nil, status.Errorf(codes.NotFound, "referral by addr %s not found", req.Addr)
	}
	return &types.QueryReferralResponse{
		Referral: referral,
	}, nil
}

func (k Querier) ReferralByCode(c context.Context, req *types.QueryReferralByCodeRequest) (*types.QueryReferralResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	referral, found := k.GetReferralByCode(ctx, req.Code)
	if !found {
		return nil, status.Errorf(codes.NotFound, "referral by code %s not found", req.Code)
	}
	return &types.QueryReferralResponse{
		Referral: referral,
	}, nil
}
