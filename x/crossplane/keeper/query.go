package keeper

import "github.com/overlock-network/api/go/node/overlock/crossplane/v1beta1"

var _ v1beta1.QueryServer = Keeper{}
