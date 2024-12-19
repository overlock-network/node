package keeper

import (
	"context"

	"overlock/x/overlock/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateComposition(goCtx context.Context, msg *types.MsgCreateComposition) (*types.MsgCreateCompositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var composition = types.Composition{
		ApiVersion: msg.ApiVersion,
		Kind:       msg.Kind,
		Metadata:   msg.Metadata,
		Spec:       &types.CompositionSpec{},
	}
	id := k.AppendComposition(
		ctx,
		composition,
	)

	return &types.MsgCreateCompositionResponse{Id: id}, nil
}
