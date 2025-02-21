package keeper

import (
	"context"
	"fmt"

	"overlock/x/crossplane/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteXrd(goCtx context.Context, msg *types.MsgDeleteXrd) (*types.MsgDeleteXrdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetCompositeResourceDefinition(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.RemoveCompositeResourceDefinition(ctx, msg.Id)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.XRDDeletedEvent,
			sdk.NewAttribute(types.XRDIndex, strconv.FormatUint(id, 10)),
		),
	)
	return &types.MsgDeleteXrdResponse{}, nil
}
