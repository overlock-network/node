package keeper

import (
	"context"

	"github.com/overlock-network/api/go/node/overlock/crossplane/v1beta1"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListEnvironment(goCtx context.Context, req *v1beta1.QueryListEnvironmentRequest) (*v1beta1.QueryListEnvironmentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.EnvironmentKey))

	var environments []v1beta1.Environment
	pageRes, err := query.FilteredPaginate(store, req.Pagination, func(key, value []byte, accumulate bool) (bool, error) {
		var env v1beta1.Environment
		if err := k.cdc.Unmarshal(value, &env); err != nil {
			return false, err
		}

		if req.Creator == "" || env.Creator == req.Creator {
			if accumulate {
				environments = append(environments, env)
			}
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1beta1.QueryListEnvironmentResponse{Environments: environments, Pagination: pageRes}, nil
}
