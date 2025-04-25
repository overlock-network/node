package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "overlock/testutil/keeper"

	"github.com/overlock-network/api/go/node/overlock/storage/v1beta1"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.StorageKeeper(t)
	params := v1beta1.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
