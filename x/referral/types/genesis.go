package types

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return NewGenesisState(DefaultParams(), 0, nil)
}

func NewGenesisState(params Params, lastReferralId uint64, referrals []Referral) *GenesisState {
	return &GenesisState{
		Params:         params,
		LastReferralId: lastReferralId,
		Referrals:      referrals,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	return nil
}
