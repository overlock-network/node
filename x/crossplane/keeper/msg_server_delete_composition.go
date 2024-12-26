package keeper

import (
	"context"
	"fmt"

	"overlock/x/crossplane/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteComposition(goCtx context.Context, msg *types.MsgDeleteComposition) (*types.MsgDeleteCompositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetComposition(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	k.RemoveComposition(ctx, msg.Id)

	return &types.MsgDeleteCompositionResponse{}, nil
}
