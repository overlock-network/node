package storage

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"overlock/x/storage/keeper"

	"github.com/web-seven/overlock-api/go/node/overlock/storage/v1beta1"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState v1beta1.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *v1beta1.GenesisState {
	genesis := v1beta1.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
