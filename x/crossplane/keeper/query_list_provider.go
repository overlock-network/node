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

func (k Keeper) ListProvider(goCtx context.Context, req *v1beta1.QueryListProviderRequest) (*v1beta1.QueryListProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.ProviderKey))

	var providers []v1beta1.Provider
	pageRes, err := query.FilteredPaginate(store, req.Pagination, func(key, value []byte, accumulate bool) (bool, error) {
		var env v1beta1.Provider
		if err := k.cdc.Unmarshal(value, &env); err != nil {
			return false, err
		}

		if req.Creator.Value == "" || env.Creator == req.Creator.Value {
			if accumulate {
				providers = append(providers, env)
			}
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1beta1.QueryListProviderResponse{Providers: providers, Pagination: pageRes}, nil
}
