package keeper

import (
	"context"

	"github.com/overlock-network/api/go/node/overlock/storage/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowRegistry(goCtx context.Context, req *v1beta1.QueryShowRegistryRequest) (*v1beta1.QueryShowRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	registry, found := k.GetRegistry(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &v1beta1.QueryShowRegistryResponse{Registry: registry}, nil
}
