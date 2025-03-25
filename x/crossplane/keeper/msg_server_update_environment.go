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

func (k msgServer) UpdateEnvironment(goCtx context.Context, msg *v1beta1.MsgUpdateEnvironment) (*v1beta1.MsgUpdateEnvironmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var env = v1beta1.Environment{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
		Creator:  msg.Creator,
		Id:       msg.Id,
	}
	_, found := k.GetEnvironment(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.SetEnvironment(ctx, env)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.EnvironmentUpdatedEvent,
			sdk.NewAttribute(v1beta1.EnvironmentIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)

	return &v1beta1.MsgUpdateEnvironmentResponse{Id: env.Id}, nil
}
