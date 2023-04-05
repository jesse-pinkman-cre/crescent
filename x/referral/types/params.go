package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	//"cosmossdk.io/math"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	farmingtypes "github.com/jesse-pinkman-cre/crescent/x/farming/types"
	"gopkg.in/yaml.v2"
)

var _ paramstypes.ParamSet = (*Params)(nil)

// General constants
const (
	//PoolReserveAddressPrefix  = "PoolReserveAddress"
	//PairEscrowAddressPrefix   = "PairEscrowAddress"
	ModuleAddressNameSplitter = "|"
	AddressType               = farmingtypes.AddressType32Bytes
)

var (
	DefaultMaxParentDepth           = sdk.NewInt(3)
	DefaultSwapFeeRate              = sdk.ZeroDec()
	DefaultSwapFeeReferralRate      = sdk.ZeroDec()
	DefaultSwapFeeParentRate        = sdk.NewDecWithPrec(1, 1) // 10%
	DefaultMinNumChildForTierSilver = sdk.NewInt(15)
	DefaultFeeCollectorAddress      = farmingtypes.DeriveAddress(AddressType, ModuleName, "FeeCollector")
)

var (
	KeyMaxParentDepth           = []byte("MaxParentDepth")
	KeySwapFeeRate              = []byte("SwapFeeRate")
	KeySwapFeeReferralRate      = []byte("SwapFeeReferralRate")
	KeySwapFeeParentRate        = []byte("SwapFeeParentRate")
	KeyMinNumChildForTierSilver = []byte("MinNumChildForTierSilver")
	KeyFeeCollectorAddress      = []byte("FeeCollectorAddress")
)

func DefaultParams() Params {
	return Params{
		MaxParentDepth:           DefaultMaxParentDepth,
		SwapFeeRate:              DefaultSwapFeeRate,
		SwapFeeReferralRate:      DefaultSwapFeeReferralRate,
		SwapFeeParentRate:        DefaultSwapFeeParentRate,
		MinNumChildForTierSilver: DefaultMinNumChildForTierSilver,
		FeeCollectorAddress:      DefaultFeeCollectorAddress.String(),
	}
}
func (params *Params) Validate() error {
	for _, field := range []struct {
		val          interface{}
		validateFunc func(i interface{}) error
	}{
		{params.MaxParentDepth, validateMaxParentDepth},
		{params.SwapFeeRate, validateSwapFeeRate},
		{params.SwapFeeReferralRate, validateSwapFeeReferralRate},
		{params.SwapFeeParentRate, validateSwapFeeParentRate},
		{params.MinNumChildForTierSilver, validateMinNumChildForTierSilver},
		{params.FeeCollectorAddress, validateFeeCollectorAddress},
	} {
		if err := field.validateFunc(field.val); err != nil {
			return err
		}
	}
	return nil
}
func (params *Params) String() string {
	out, _ := yaml.Marshal(params)
	return string(out)
}

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyMaxParentDepth, &params.MaxParentDepth, validateMaxParentDepth),
		paramstypes.NewParamSetPair(KeySwapFeeRate, &params.SwapFeeRate, validateSwapFeeRate),
		paramstypes.NewParamSetPair(KeySwapFeeReferralRate, &params.SwapFeeReferralRate, validateSwapFeeReferralRate),
		paramstypes.NewParamSetPair(KeySwapFeeParentRate, &params.SwapFeeParentRate, validateSwapFeeParentRate),
		paramstypes.NewParamSetPair(KeyMinNumChildForTierSilver, &params.MinNumChildForTierSilver, validateMinNumChildForTierSilver),
		//paramstypes.NewParamSetPair(KeyFeeCollectorAddress, &params.FeeCollectorAddress, validateFeeCollectorAddress),
	}
}

func validateMaxParentDepth(i interface{}) error {
	return nil
}

func validateSwapFeeRate(i interface{}) error {
	feeRate, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type for referral DateSwapFeeRate: %T", i)
	}
	if feeRate.IsNegative() {
		return fmt.Errorf("invalid parameter value: %v", feeRate)
	}
	return nil
}

func validateSwapFeeReferralRate(i interface{}) error {
	return nil
}

func validateSwapFeeParentRate(i interface{}) error {
	return nil
}

func validateMinNumChildForTierSilver(i interface{}) error {
	return nil
}

func validateFeeCollectorAddress(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if _, err := sdk.AccAddressFromBech32(v); err != nil {
		return fmt.Errorf("invalid fee collector address: %w", err)
	}

	return nil
}
