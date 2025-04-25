package keeper

import (
	"context"
	"strconv"

	"github.com/overlock-network/api/go/node/overlock/storage/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateRegistry(goCtx context.Context, msg *v1beta1.MsgCreateRegistry) (*v1beta1.MsgCreateRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var registry = v1beta1.Registry{
		Name:     msg.Name,
		Provider: msg.Provider,
	}
	id := k.AppendRegistry(
		ctx,
		registry,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.RegistryCreatedEvent,
			sdk.NewAttribute(v1beta1.RegistryIndex, strconv.FormatUint(id, 10)),
		),
	)

	return &v1beta1.MsgCreateRegistryResponse{Id: id}, nil
}
