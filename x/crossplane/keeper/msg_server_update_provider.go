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

func (k msgServer) UpdateProvider(goCtx context.Context, msg *types.MsgUpdateProvider) (*types.MsgUpdateProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var provider = types.Provider{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
		Creator:  msg.Creator,
		Id:       msg.Id,
	}
	_, found := k.GetProvider(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.SetProvider(ctx, provider)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ProviderUpdatedEvent,
			sdk.NewAttribute(types.ProviderIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)

	return &types.MsgUpdateProviderResponse{Id: msg.Id}, nil
}
