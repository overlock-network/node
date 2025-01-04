package keeper

import (
	"context"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteProvider(goCtx context.Context, msg *types.MsgDeleteProvider) (*types.MsgDeleteProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeleteProviderResponse{}, nil
}
