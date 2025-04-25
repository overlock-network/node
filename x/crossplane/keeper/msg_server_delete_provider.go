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

func (k msgServer) DeleteProvider(goCtx context.Context, msg *v1beta1.MsgDeleteProvider) (*v1beta1.MsgDeleteProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetProvider(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	k.RemoveProvider(ctx, msg.Id)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.ProviderDeletedEvent,
			sdk.NewAttribute(v1beta1.ProviderIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)
	return &v1beta1.MsgDeleteProviderResponse{}, nil
}
