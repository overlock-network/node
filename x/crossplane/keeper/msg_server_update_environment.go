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

func (k msgServer) UpdateEnvironment(goCtx context.Context, msg *types.MsgUpdateEnvironment) (*types.MsgUpdateEnvironmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var env = types.Environment{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
		Creator:  msg.Creator,
		Id:       msg.Id,
	}
	_, found := k.GetEnvironment(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.SetEnvironment(ctx, env)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EnvironmentUpdatedEvent,
			sdk.NewAttribute(types.EnvironmentIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)

	return &types.MsgUpdateEnvironmentResponse{Id: env.Id}, nil
}
