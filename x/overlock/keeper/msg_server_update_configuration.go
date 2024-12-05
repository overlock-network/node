package keeper

import (
	"context"
	"fmt"

	"overlock/x/overlock/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateConfiguration(goCtx context.Context, msg *types.MsgUpdateConfiguration) (*types.MsgUpdateConfigurationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var configuration = types.Configuration{
		Id:   msg.Id,
		Name: msg.Name,
		Spec: msg.Spec,
	}

	_, found := k.GetConfiguration(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	k.SetConfiguration(ctx, configuration)
	return &types.MsgUpdateConfigurationResponse{}, nil
}
