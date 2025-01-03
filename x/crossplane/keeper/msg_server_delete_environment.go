package keeper

import (
	"context"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteEnvironment(goCtx context.Context, msg *types.MsgDeleteEnvironment) (*types.MsgDeleteEnvironmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeleteEnvironmentResponse{}, nil
}
