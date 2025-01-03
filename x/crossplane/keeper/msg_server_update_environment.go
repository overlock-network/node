package keeper

import (
	"context"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateEnvironment(goCtx context.Context, msg *types.MsgUpdateEnvironment) (*types.MsgUpdateEnvironmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateEnvironmentResponse{}, nil
}
