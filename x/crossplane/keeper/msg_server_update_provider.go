package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/overlock-network/api/go/node/overlock/crossplane/v1beta1"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateProvider(goCtx context.Context, msg *v1beta1.MsgUpdateProvider) (*v1beta1.MsgUpdateProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var provider = v1beta1.Provider{
		Metadata: msg.Metadata,
		Creator:  msg.Creator,
		Id:       msg.Id,
	}
	_, found := k.GetProvider(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.SetProvider(ctx, provider)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.ProviderUpdatedEvent,
			sdk.NewAttribute(v1beta1.ProviderIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)

	return &v1beta1.MsgUpdateProviderResponse{Id: msg.Id}, nil
}
