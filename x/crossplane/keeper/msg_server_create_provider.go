package keeper

import (
	"context"
	"strconv"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateProvider(goCtx context.Context, msg *types.MsgCreateProvider) (*types.MsgCreateProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var provider = types.Provider{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
	}
	id := k.AppendProvider(
		ctx,
		provider,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ProviderCreatedEvent,
			sdk.NewAttribute(types.ProviderIndex, strconv.FormatUint(id, 10)),
		),
	)

	return &types.MsgCreateProviderResponse{}, nil
}
