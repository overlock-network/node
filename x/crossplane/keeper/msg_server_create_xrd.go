package keeper

import (
	"context"
	"strconv"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateXrd(goCtx context.Context, msg *types.MsgCreateXrd) (*types.MsgCreateXrdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var xrd = types.CompositeResourceDefinition{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
	}

	id := k.AppendCompositeResourceDefinition(
		ctx,
		xrd,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.XRDCreatedEvent,
			sdk.NewAttribute(types.XRDIndex, strconv.FormatUint(id, 10)),
		),
	)

	return &types.MsgCreateXrdResponse{Id: id}, nil
}
