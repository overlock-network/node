package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateFunction(goCtx context.Context, msg *v1beta1.MsgUpdateFunction) (*v1beta1.MsgUpdateFunctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var function = v1beta1.Function{
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
		sdk.NewEvent(v1beta1.FunctionUpdatedEvent,
			sdk.NewAttribute(v1beta1.FunctionIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)

	return &v1beta1.MsgUpdateFunctionResponse{Id: msg.Id}, nil
}
