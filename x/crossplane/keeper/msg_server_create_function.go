package keeper

import (
	"context"
	"strconv"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateFunction(goCtx context.Context, msg *types.MsgCreateFunction) (*types.MsgCreateFunctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var function = types.Function{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
	}
	id := k.AppendFunction(
		ctx,
		function,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.FunctionCreatedEvent,
			sdk.NewAttribute(types.FunctionIndex, strconv.FormatUint(id, 10)),
		),
	)

	return &types.MsgCreateFunctionResponse{}, nil
}
