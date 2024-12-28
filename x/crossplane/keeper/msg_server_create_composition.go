package keeper

import (
	"context"
	"strconv"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateComposition(goCtx context.Context, msg *types.MsgCreateComposition) (*types.MsgCreateCompositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var composition = types.Composition{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
	}
	id := k.AppendComposition(
		ctx,
		composition,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.CompositionCreatedEvent,
			sdk.NewAttribute(types.CompositionIndex, strconv.FormatUint(id, 10)),
		),
	)

	return &types.MsgCreateCompositionResponse{Id: id}, nil
}
