package keeper

import (
	"context"
	"strconv"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateXrd(goCtx context.Context, msg *v1beta1.MsgCreateXrd) (*v1beta1.MsgCreateXrdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var xrd = v1beta1.CompositeResourceDefinition{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
	}

	id := k.AppendCompositeResourceDefinition(
		ctx,
		xrd,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.XRDCreatedEvent,
			sdk.NewAttribute(v1beta1.XRDIndex, strconv.FormatUint(id, 10)),
		),
	)

	return &v1beta1.MsgCreateXrdResponse{Id: id}, nil
}
