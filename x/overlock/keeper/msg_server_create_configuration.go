package keeper

import (
	"context"

	"overlock/x/overlock/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateConfiguration(goCtx context.Context, msg *types.MsgCreateConfiguration) (*types.MsgCreateConfigurationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var post = types.Configuration{
		Name: msg.Name,
	}
	id := k.AppendPost(
		ctx,
		post,
	)
	return &types.MsgCreateConfigurationResponse{
		Id: id,
	}, nil

	return &types.MsgCreateConfigurationResponse{}, nil
}
