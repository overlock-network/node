package keeper

import (
	"context"
	"fmt"

	"overlock/x/overlock/types"

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

	return &types.MsgDeleteConfigurationResponse{}, nil
}
