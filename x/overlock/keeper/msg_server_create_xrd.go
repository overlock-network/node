package keeper

import (
	"context"

	"overlock/x/overlock/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateXrd(goCtx context.Context, msg *types.MsgCreateXrd) (*types.MsgCreateXrdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var xrd = types.CompositeResourceDefinition{
		ApiVersion: msg.ApiVersion,
		Kind:       msg.Kind,
		Metadata:   msg.Metadata,
		Spec:       &types.XrdSpec{},
	}
	id := k.AppendCompositeResourceDefinition(
		ctx,
		xrd,
	)

	return &types.MsgCreateXrdResponse{Id: id}, nil
}
