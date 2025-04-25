package keeper

import (
	"context"
	"strconv"

	"github.com/overlock-network/api/go/node/overlock/crossplane/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateConfiguration(goCtx context.Context, msg *v1beta1.MsgCreateConfiguration) (*v1beta1.MsgCreateConfigurationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var configuration = v1beta1.Configuration{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
		Creator:  msg.Creator,
	}
	id := k.AppendConfiguration(
		ctx,
		configuration,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.ConfigurationCreatedEvent,
			sdk.NewAttribute(v1beta1.ConfigurationIndex, strconv.FormatUint(id, 10)),
		),
	)

	return &v1beta1.MsgCreateConfigurationResponse{Id: id}, nil
}
