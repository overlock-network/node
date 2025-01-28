package keeper

import (
	"context"
	"strconv"

	"overlock/x/storage/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateRegistry(goCtx context.Context, msg *types.MsgCreateRegistry) (*types.MsgCreateRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var registry = types.Registry{
		Name:     msg.Name,
		Provider: msg.Provider,
	}
	id := k.AppendRegistry(
		ctx,
		registry,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.RegistryCreatedEvent,
			sdk.NewAttribute(types.RegistryIndex, strconv.FormatUint(id, 10)),
		),
	)

	return &types.MsgCreateRegistryResponse{Id: id}, nil
}
