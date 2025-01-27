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

func (k Keeper) ListConfiguration(goCtx context.Context, req *types.QueryListConfigurationRequest) (*types.QueryListConfigurationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ConfigurationKey))

	var configurations []types.Configuration

	pageRes, err := query.FilteredPaginate(store, req.Pagination, func(key, value []byte, accumulate bool) (bool, error) {
		var configuration types.Configuration
		if err := k.cdc.Unmarshal(value, &configuration); err != nil {
			return false, err
		}

		if req.Creator == "" || configuration.Creator == req.Creator {
			if accumulate {
				configurations = append(configurations, configuration)
			}
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListConfigurationResponse{Configurations: configurations, Pagination: pageRes}, nil
}
