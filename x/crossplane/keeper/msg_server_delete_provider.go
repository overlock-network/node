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

func (k msgServer) DeleteProvider(goCtx context.Context, msg *types.MsgDeleteProvider) (*types.MsgDeleteProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetProvider(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	k.RemoveProvider(ctx, msg.Id)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ProviderDeletedEvent,
			sdk.NewAttribute(types.ProviderIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)
	return &types.MsgDeleteProviderResponse{Id: msg.Id}, nil
}
