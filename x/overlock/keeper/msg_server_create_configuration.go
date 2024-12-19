package keeper

import (
	"context"

	"overlock/x/overlock/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateConfiguration(goCtx context.Context, msg *types.MsgCreateConfiguration) (*types.MsgCreateConfigurationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var configuration = types.Configuration{
		ApiVersion: msg.ApiVersion,
		Kind:       msg.Kind,
		Metadata:   msg.Metadata,
		Spec:       &types.ConfigurationSpec{},
	}
	id := k.AppendConfiguration(
		ctx,
		configuration,
	)

	return &types.MsgCreateConfigurationResponse{Id: id}, nil
}
