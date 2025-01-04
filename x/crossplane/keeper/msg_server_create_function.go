package keeper

import (
	"context"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateFunction(goCtx context.Context, msg *types.MsgCreateFunction) (*types.MsgCreateFunctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateFunctionResponse{}, nil
}
