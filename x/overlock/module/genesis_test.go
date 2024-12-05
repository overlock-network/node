package overlock_test

import (
	"testing"

	keepertest "overlock/testutil/keeper"
	"overlock/testutil/nullify"
	overlock "overlock/x/overlock/module"
	"overlock/x/overlock/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OverlockKeeper(t)
	overlock.InitGenesis(ctx, k, genesisState)
	got := overlock.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
