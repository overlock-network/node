package keeper

import (
	"context"
	"fmt"

	"overlock/x/crossplane/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteConfiguration(goCtx context.Context, msg *types.MsgDeleteConfiguration) (*types.MsgDeleteConfigurationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetConfiguration(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	k.RemoveComposition(ctx, msg.Id)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ConfigurationDeletedEvent,
			sdk.NewAttribute(types.ConfigurationIndex, strconv.FormatUint(id, 10)),
		),
	)
	return &types.MsgDeleteConfigurationResponse{}, nil
}
