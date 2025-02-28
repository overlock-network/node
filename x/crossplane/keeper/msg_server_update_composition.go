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

func (k msgServer) UpdateComposition(goCtx context.Context, msg *types.MsgUpdateComposition) (*types.MsgUpdateCompositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var composition = types.Composition{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
		Id:       msg.Id,
	}
	_, found := k.GetComposition(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.SetComposition(ctx, composition)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.CompositionUpdatedEvent,
			sdk.NewAttribute(types.CompositionIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)

	return &types.MsgUpdateCompositionResponse{}, nil
}
