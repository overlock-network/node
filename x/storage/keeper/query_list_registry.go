package keeper

import (
	"context"

	"overlock/x/storage/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListRegistry(goCtx context.Context, req *types.QueryListRegistryRequest) (*types.QueryListRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RegistryKey))

	var registries []types.Registry
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var registry types.Registry
		if err := k.cdc.Unmarshal(value, &registry); err != nil {
			return err
		}

		registries = append(registries, registry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListRegistryResponse{Registry: registries, Pagination: pageRes}, nil
}
