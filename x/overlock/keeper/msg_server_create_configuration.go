package keeper

import (
	"context"

	"overlock/x/overlock/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateConfiguration(goCtx context.Context, msg *types.MsgCreateConfiguration) (*types.MsgCreateConfigurationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var configuration = types.Configuration{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
	}
	id := k.AppendConfiguration(
		ctx,
		configuration,
	)

	return &types.MsgCreateConfigurationResponse{Id: id}, nil
}
