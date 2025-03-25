package keeper

import (
	"context"
	"strconv"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateProvider(goCtx context.Context, msg *v1beta1.MsgCreateProvider) (*v1beta1.MsgCreateProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var provider = v1beta1.Provider{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
		Creator:  msg.Creator,
	}
	id := k.AppendProvider(
		ctx,
		provider,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.ProviderCreatedEvent,
			sdk.NewAttribute(v1beta1.ProviderIndex, strconv.FormatUint(id, 10)),
		),
	)
	return &v1beta1.MsgCreateProviderResponse{}, nil
}
