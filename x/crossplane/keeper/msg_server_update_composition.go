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

func (k msgServer) UpdateComposition(goCtx context.Context, msg *v1beta1.MsgUpdateComposition) (*v1beta1.MsgUpdateCompositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var composition = v1beta1.Composition{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
		Id:       msg.Id,
	}
	_, found := k.GetComposition(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.SetComposition(ctx, composition)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.CompositionUpdatedEvent,
			sdk.NewAttribute(v1beta1.CompositionIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)

	return &v1beta1.MsgUpdateCompositionResponse{}, nil
}
