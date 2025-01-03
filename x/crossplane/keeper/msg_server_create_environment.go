package keeper

import (
	"context"
	"strconv"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateEnvironment(goCtx context.Context, msg *types.MsgCreateEnvironment) (*types.MsgCreateEnvironmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var env = types.Environment{
		Name: msg.Name,
	}
	id := k.AppendEnvironment(
		ctx,
		env,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EnvironmentCreatedEvent,
			sdk.NewAttribute(types.EnvironmentIndex, strconv.FormatUint(id, 10)),
		),
	)
	return &types.MsgCreateEnvironmentResponse{Id: id}, nil
}
