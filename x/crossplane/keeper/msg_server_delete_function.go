package keeper

import (
	"context"
	"fmt"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteFunction(goCtx context.Context, msg *v1beta1.MsgDeleteFunction) (*v1beta1.MsgDeleteFunctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetFunction(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	k.RemoveFunction(ctx, msg.Id)

	return &v1beta1.MsgDeleteFunctionResponse{Id: msg.Id}, nil
}
