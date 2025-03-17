package keeper

import (
	"context"
	"strconv"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateComposition(goCtx context.Context, msg *v1beta1.MsgCreateComposition) (*v1beta1.MsgCreateCompositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var composition = v1beta1.Composition{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
	}
	id := k.AppendComposition(
		ctx,
		composition,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.CompositionCreatedEvent,
			sdk.NewAttribute(v1beta1.CompositionIndex, strconv.FormatUint(id, 10)),
		),
	)

	return &v1beta1.MsgCreateCompositionResponse{Id: id}, nil
}
