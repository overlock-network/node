package keeper

import (
	"context"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListComposition(goCtx context.Context, req *v1beta1.QueryListCompositionRequest) (*v1beta1.QueryListCompositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.CompositionKey))

	var compositions []v1beta1.Composition
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var composition v1beta1.Composition
		if err := k.cdc.Unmarshal(value, &composition); err != nil {
			return err
		}

		compositions = append(compositions, composition)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1beta1.QueryListCompositionResponse{Compositions: compositions, Pagination: pageRes}, nil
}
