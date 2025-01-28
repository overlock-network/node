package keeper

import (
	"context"
	"fmt"

	"overlock/x/storage/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateRegistry(goCtx context.Context, msg *types.MsgUpdateRegistry) (*types.MsgUpdateRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var registry = types.Registry{
		Name:     msg.Name,
		Provider: msg.Provider,
		Id:       msg.Id,
	}
	_, found := k.GetRegistry(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.SetRegistry(ctx, registry)

	return &types.MsgUpdateRegistryResponse{}, nil
}
