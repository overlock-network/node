package keeper

import (
	"context"

	"overlock/x/overlock/types"

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
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var configuration types.Configuration
		if err := k.cdc.Unmarshal(value, &configuration); err != nil {
			return err
		}

		configurations = append(configurations, configuration)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListConfigurationResponse{Configurations: configurations, Pagination: pageRes}, nil
}
