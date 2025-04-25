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

func (k msgServer) DeleteEnvironment(goCtx context.Context, msg *v1beta1.MsgDeleteEnvironment) (*v1beta1.MsgDeleteEnvironmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetEnvironment(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	k.RemoveEnvironment(ctx, msg.Id)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.EnvironmentDeletedEvent,
			sdk.NewAttribute(v1beta1.EnvironmentIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)
	return &v1beta1.MsgDeleteEnvironmentResponse{}, nil
}
