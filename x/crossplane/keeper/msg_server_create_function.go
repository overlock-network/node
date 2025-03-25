package keeper

import (
	"context"
	"strconv"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateFunction(goCtx context.Context, msg *v1beta1.MsgCreateFunction) (*v1beta1.MsgCreateFunctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var function = v1beta1.Function{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
		Creator:  msg.Creator,
	}
	id := k.AppendFunction(
		ctx,
		function,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.FunctionCreatedEvent,
			sdk.NewAttribute(v1beta1.FunctionIndex, strconv.FormatUint(id, 10)),
		),
	)

	return &v1beta1.MsgCreateFunctionResponse{}, nil
}
