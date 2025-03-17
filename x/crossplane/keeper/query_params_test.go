package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "overlock/testutil/keeper"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.OverlockKeeper(t)
	params := v1beta1.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &v1beta1.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &v1beta1.QueryParamsResponse{Params: params}, response)
}
