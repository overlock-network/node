package keeper

import (
	"context"

	"overlock/x/overlock/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateXrd(goCtx context.Context, msg *types.MsgCreateXrd) (*types.MsgCreateXrdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var xrd = types.CompositeResourceDefinition{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
	}
	id := k.AppendCompositeResourceDefinition(
		ctx,
		xrd,
	)

	return &types.MsgCreateXrdResponse{Id: id}, nil
}
