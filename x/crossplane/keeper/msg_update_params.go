package keeper

import (
	"context"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateParams(goCtx context.Context, req *v1beta1.MsgUpdateParams) (*v1beta1.MsgUpdateParamsResponse, error) {
	if k.GetAuthority() != req.Authority {
		return nil, errorsmod.Wrapf(v1beta1.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetAuthority(), req.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.SetParams(ctx, req.Params); err != nil {
		return nil, err
	}

	return &v1beta1.MsgUpdateParamsResponse{}, nil
}
