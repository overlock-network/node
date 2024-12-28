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

func (k Keeper) ListXrd(goCtx context.Context, req *types.QueryListXrdRequest) (*types.QueryListXrdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.XRDKey))

	var xrds []types.CompositeResourceDefinition
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var compositeResourceDefinition types.CompositeResourceDefinition
		if err := k.cdc.Unmarshal(value, &compositeResourceDefinition); err != nil {
			return err
		}

		xrds = append(xrds, compositeResourceDefinition)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListXrdResponse{Xrd: xrds, Pagination: pageRes}, nil
}
