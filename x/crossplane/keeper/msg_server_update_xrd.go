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

func (k msgServer) UpdateXrd(goCtx context.Context, msg *types.MsgUpdateXrd) (*types.MsgUpdateXrdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var xrd = types.CompositeResourceDefinition{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
		Id:       msg.Id,
	}
	_, found := k.GetCompositeResourceDefinition(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.SetCompositeResourceDefinition(ctx, xrd)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.XRDUpdateEvent,
			sdk.NewAttribute(types.XRDIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)

	return &types.MsgUpdateXrdResponse{}, nil
}
