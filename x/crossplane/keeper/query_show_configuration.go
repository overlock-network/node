package keeper

import (
	"context"

	"github.com/overlock-network/api/go/node/overlock/crossplane/v1beta1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowConfiguration(goCtx context.Context, req *v1beta1.QueryShowConfigurationRequest) (*v1beta1.QueryShowConfigurationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	configuration, found := k.GetConfiguration(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &v1beta1.QueryShowConfigurationResponse{Configuration: configuration}, nil
}
