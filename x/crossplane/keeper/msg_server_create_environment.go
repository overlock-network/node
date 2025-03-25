package keeper

import (
	"context"
	"strconv"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateEnvironment(goCtx context.Context, msg *v1beta1.MsgCreateEnvironment) (*v1beta1.MsgCreateEnvironmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var env = v1beta1.Environment{
		Metadata: msg.Metadata,
		Spec:     msg.Spec,
		Creator:  msg.Creator,
	}
	id := k.AppendEnvironment(
		ctx,
		env,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(v1beta1.EnvironmentCreatedEvent,
			sdk.NewAttribute(v1beta1.EnvironmentIndex, strconv.FormatUint(id, 10)),
		),
	)
	return &v1beta1.MsgCreateEnvironmentResponse{Id: id}, nil
}
