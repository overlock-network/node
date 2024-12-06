package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "overlock/testutil/keeper"
	"overlock/x/blocks/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BlocksKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
