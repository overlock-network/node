package keeper

import (
	"context"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateFunction(goCtx context.Context, msg *types.MsgUpdateFunction) (*types.MsgUpdateFunctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateFunctionResponse{}, nil
}
