package overlock_test

import (
	"testing"

	keepertest "overlock/testutil/keeper"
	"overlock/testutil/nullify"
	overlock "overlock/x/crossplane/module"

	"github.com/overlock-network/api/go/node/overlock/crossplane/v1beta1"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := v1beta1.GenesisState{
		Params: v1beta1.DefaultParams(),

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
