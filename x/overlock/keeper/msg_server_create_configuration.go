package keeper

import (
	"context"

	"overlock/x/overlock/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateConfiguration(goCtx context.Context, msg *types.MsgCreateConfiguration) (*types.MsgCreateConfigurationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateConfigurationResponse{}, nil
}
