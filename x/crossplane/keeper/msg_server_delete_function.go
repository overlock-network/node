package keeper

import (
	"context"
	"fmt"

	"overlock/x/crossplane/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteFunction(goCtx context.Context, msg *types.MsgDeleteFunction) (*types.MsgDeleteFunctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetFunction(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	k.RemoveFunction(ctx, msg.Id)

	return &types.MsgDeleteFunctionResponse{Id: msg.Id}, nil
}
