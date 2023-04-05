package liquidstaking

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jesse-pinkman-cre/crescent/x/liquidstaking/keeper"
	"github.com/jesse-pinkman-cre/crescent/x/liquidstaking/types"
)

// BeginBlocker updates liquid validator set changes for the current block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	k.UpdateLiquidValidatorSet(ctx)
}
