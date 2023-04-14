package marker

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/jesse-pinkman-cre/crescent/x/marker/keeper"
	"github.com/jesse-pinkman-cre/crescent/x/marker/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
	k.SetLastBlockTime(ctx, ctx.BlockTime())
}
