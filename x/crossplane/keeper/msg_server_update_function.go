package keeper

import (
	"context"
	"fmt"
	"strconv"

	"overlock/x/crossplane/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateFunction(goCtx context.Context, msg *types.MsgUpdateFunction) (*types.MsgUpdateFunctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var function = types.Function{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
		Creator:  msg.Creator,
		Id:       msg.Id,
	}
	_, found := k.GetFunction(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.SetFunction(ctx, function)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.FunctionUpdatedEvent,
			sdk.NewAttribute(types.FunctionIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)

	return &types.MsgUpdateFunctionResponse{Id: msg.Id}, nil
}
