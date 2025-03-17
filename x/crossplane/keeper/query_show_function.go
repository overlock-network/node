package keeper

import (
	"context"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowFunction(goCtx context.Context, req *v1beta1.QueryShowFunctionRequest) (*v1beta1.QueryShowFunctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	function, found := k.GetFunction(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &v1beta1.QueryShowFunctionResponse{Function: &function}, nil
}
