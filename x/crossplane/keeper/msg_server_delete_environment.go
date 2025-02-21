package keeper

import (
	"context"
	"fmt"

	"overlock/x/crossplane/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteEnvironment(goCtx context.Context, msg *types.MsgDeleteEnvironment) (*types.MsgDeleteEnvironmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetEnvironment(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	k.RemoveEnvironment(ctx, msg.Id)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EnvironmentDeletedEvent,
			sdk.NewAttribute(types.EnvironmentIndex, strconv.FormatUint(id, 10)),
		),
	)
	return &types.MsgDeleteEnvironmentResponse{}, nil
}
