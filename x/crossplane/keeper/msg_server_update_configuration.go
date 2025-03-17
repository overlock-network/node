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

func (k msgServer) UpdateConfiguration(goCtx context.Context, msg *v1beta1.MsgUpdateConfiguration) (*v1beta1.MsgUpdateConfigurationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var configuration = v1beta1.Configuration{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
		Creator:  msg.Creator,
		Id:       msg.Id,
	}
	_, found := k.GetConfiguration(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	k.SetConfiguration(ctx, configuration)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.ConfigurationUpdatedEvent,
			sdk.NewAttribute(v1beta1.ConfigurationIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)

	return &v1beta1.MsgUpdateConfigurationResponse{}, nil
}
