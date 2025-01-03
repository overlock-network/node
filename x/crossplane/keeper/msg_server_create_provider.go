package keeper

import (
	"context"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateProvider(goCtx context.Context, msg *types.MsgCreateProvider) (*types.MsgCreateProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateProviderResponse{}, nil
}
