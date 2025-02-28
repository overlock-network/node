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

func (k msgServer) UpdateConfiguration(goCtx context.Context, msg *types.MsgUpdateConfiguration) (*types.MsgUpdateConfigurationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var configuration = types.Configuration{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
		Creator:  msg.Creator,
		Id:       msg.Id,
	}
	_, found := k.GetConfiguration(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.SetConfiguration(ctx, configuration)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ConfigurationUpdatedEvent,
			sdk.NewAttribute(types.ConfigurationIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)

	return &types.MsgUpdateConfigurationResponse{}, nil
}
