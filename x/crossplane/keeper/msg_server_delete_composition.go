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

func (k msgServer) DeleteComposition(goCtx context.Context, msg *v1beta1.MsgDeleteComposition) (*v1beta1.MsgDeleteCompositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetComposition(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	k.RemoveComposition(ctx, msg.Id)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.CompositionDeletedEvent,
			sdk.NewAttribute(v1beta1.CompositionIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)
	return &v1beta1.MsgDeleteCompositionResponse{}, nil
}
