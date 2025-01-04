package keeper

import (
	"context"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteFunction(goCtx context.Context, msg *types.MsgDeleteFunction) (*types.MsgDeleteFunctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeleteFunctionResponse{}, nil
}
