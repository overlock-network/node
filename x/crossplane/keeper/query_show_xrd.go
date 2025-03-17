package keeper

import (
	"context"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowXrd(goCtx context.Context, req *v1beta1.QueryShowXrdRequest) (*v1beta1.QueryShowXrdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	xrd, found := k.GetCompositeResourceDefinition(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &v1beta1.QueryShowXrdResponse{Xrd: xrd}, nil
}
