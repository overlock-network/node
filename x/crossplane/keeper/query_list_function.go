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

func (k Keeper) ListFunction(goCtx context.Context, req *types.QueryListFunctionRequest) (*types.QueryListFunctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FunctionKey))

	var functions []types.Function
	pageRes, err := query.FilteredPaginate(store, req.Pagination, func(key, value []byte, accumulate bool) (bool, error) {
		var env types.Function
		if err := k.cdc.Unmarshal(value, &env); err != nil {
			return false, err
		}

		if req.Creator == "" || env.Creator == req.Creator {
			if accumulate {
				functions = append(functions, env)
			}
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListFunctionResponse{Functions: functions, Pagination: pageRes}, nil
}
