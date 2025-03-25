package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/web-seven/overlock-api/go/node/overlock/storage/v1beta1"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteRegistry(goCtx context.Context, msg *v1beta1.MsgDeleteRegistry) (*v1beta1.MsgDeleteRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetRegistry(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	k.RemoveRegistry(ctx, msg.Id)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.RegistryDeletedEvent,
			sdk.NewAttribute(v1beta1.RegistryIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)
	return &v1beta1.MsgDeleteRegistryResponse{}, nil
}
