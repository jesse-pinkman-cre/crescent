package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/jesse-pinkman-cre/crescent/x/marker/types"
)

// Keeper of the module's store.
type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	paramSpace paramstypes.Subspace
}

// NewKeeper creates a new Keeper instance.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	paramSpace paramstypes.Subspace,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
