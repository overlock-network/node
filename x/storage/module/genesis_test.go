package storage_test

import (
	"testing"

	keepertest "overlock/testutil/keeper"
	"overlock/testutil/nullify"
	storage "overlock/x/storage/module"

	"github.com/web-seven/overlock-api/go/node/overlock/storage/v1beta1"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := v1beta1.GenesisState{
		Params: v1beta1.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StorageKeeper(t)
	storage.InitGenesis(ctx, k, genesisState)
	got := storage.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
