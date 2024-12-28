package keeper

import (
	"context"
	"strconv"

	"overlock/x/crossplane/types"

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

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ConfigurationCreatedEvent,
			sdk.NewAttribute(types.ConfigurationIndex, strconv.FormatUint(id, 10)),
		),
	)

	return &types.MsgCreateConfigurationResponse{Id: id}, nil
}
