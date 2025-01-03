package keeper

import (
	"context"

	"overlock/x/crossplane/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListEnvironment(goCtx context.Context, req *types.QueryListEnvironmentRequest) (*types.QueryListEnvironmentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryListEnvironmentResponse{}, nil
}
