package keeper

import (
	"context"

	"github.com/web-seven/overlock-api/go/node/overlock/storage/v1beta1"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListRegistry(goCtx context.Context, req *v1beta1.QueryListRegistryRequest) (*v1beta1.QueryListRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.RegistryKey))

	var registries []v1beta1.Registry
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var registry v1beta1.Registry
		if err := k.cdc.Unmarshal(value, &registry); err != nil {
			return err
		}

		registries = append(registries, registry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1beta1.QueryListRegistryResponse{Registry: registries, Pagination: pageRes}, nil
}
