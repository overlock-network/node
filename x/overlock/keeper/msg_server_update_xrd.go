package keeper

import (
	"context"
	"fmt"

	"overlock/x/overlock/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateXrd(goCtx context.Context, msg *types.MsgUpdateXrd) (*types.MsgUpdateXrdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var xrd = types.CompositeResourceDefinition{
		ApiVersion: msg.ApiVersion,
		Kind:       msg.Kind,
		Metadata:   msg.Metadata,
		Spec:       &types.XrdSpec{},
	}
	_, found := k.GetCompositeResourceDefinition(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.SetCompositeResourceDefinition(ctx, xrd)

	return &types.MsgUpdateXrdResponse{}, nil
}
