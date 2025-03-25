package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "overlock/testutil/keeper"
	"overlock/x/crossplane/keeper"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, v1beta1.MsgServer, context.Context) {
	k, ctx := keepertest.OverlockKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
