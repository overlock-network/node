package keeper

import (
	"context"

	"overlock/x/crossplane/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListEnvironment(goCtx context.Context, req *types.QueryListEnvironmentRequest) (*types.QueryListEnvironmentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EnvironmentKey))

	var environments []types.Environment
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var env types.Environment
		if err := k.cdc.Unmarshal(value, &env); err != nil {
			return err
		}

		environments = append(environments, env)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListEnvironmentResponse{Environments: environments, Pagination: pageRes}, nil
}
