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

func (k msgServer) DeleteConfiguration(goCtx context.Context, msg *v1beta1.MsgDeleteConfiguration) (*v1beta1.MsgDeleteConfigurationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetConfiguration(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	k.RemoveConfiguration(ctx, msg.Id)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.ConfigurationDeletedEvent,
			sdk.NewAttribute(v1beta1.ConfigurationIndex, strconv.FormatUint(msg.Id, 10)),
		),
	)
	return &v1beta1.MsgDeleteConfigurationResponse{}, nil
}
