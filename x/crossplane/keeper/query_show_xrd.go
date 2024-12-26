package keeper

import (
	"context"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowXrd(goCtx context.Context, req *types.QueryShowXrdRequest) (*types.QueryShowXrdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	xrd, found := k.GetCompositeResourceDefinition(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryShowXrdResponse{Xrd: xrd}, nil
}
