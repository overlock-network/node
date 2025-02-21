package keeper

import (
	"context"
	"fmt"
	"strconv"

	"overlock/x/storage/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteRegistry(goCtx context.Context, msg *types.MsgDeleteRegistry) (*types.MsgDeleteRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetRegistry(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	k.RemoveRegistry(ctx, msg.Id)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.RegistryDeletedEvent,
			sdk.NewAttribute(types.RegistryIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)
	return &types.MsgDeleteRegistryResponse{}, nil
}
