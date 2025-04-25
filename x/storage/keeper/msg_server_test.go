package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "overlock/testutil/keeper"
	"overlock/x/storage/keeper"

	"github.com/overlock-network/api/go/node/overlock/storage/v1beta1"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, v1beta1.MsgServer, context.Context) {
	k, ctx := keepertest.StorageKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
