package keeper

import (
	"context"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateEnvironment(goCtx context.Context, msg *types.MsgCreateEnvironment) (*types.MsgCreateEnvironmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateEnvironmentResponse{}, nil
}
