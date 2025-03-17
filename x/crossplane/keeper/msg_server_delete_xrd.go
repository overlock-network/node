package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteXrd(goCtx context.Context, msg *v1beta1.MsgDeleteXrd) (*v1beta1.MsgDeleteXrdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetCompositeResourceDefinition(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.RemoveCompositeResourceDefinition(ctx, msg.Id)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.XRDDeletedEvent,
			sdk.NewAttribute(v1beta1.XRDIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)
	return &v1beta1.MsgDeleteXrdResponse{}, nil
}
