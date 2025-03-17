package keeper

import (
	"context"

	"github.com/web-seven/overlock-api/go/node/overlock/storage/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Params(goCtx context.Context, req *v1beta1.QueryParamsRequest) (*v1beta1.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &v1beta1.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
