package keeper

import (
	"context"
	"fmt"

	"github.com/web-seven/overlock-api/go/node/overlock/storage/v1beta1"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateRegistry(goCtx context.Context, msg *v1beta1.MsgUpdateRegistry) (*v1beta1.MsgUpdateRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var registry = v1beta1.Registry{
		Name:     msg.Name,
		Provider: msg.Provider,
		Id:       msg.Id,
	}
	_, found := k.GetRegistry(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.SetRegistry(ctx, registry)

	return &v1beta1.MsgUpdateRegistryResponse{}, nil
}
