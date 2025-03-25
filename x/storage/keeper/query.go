package keeper

import "github.com/web-seven/overlock-api/go/node/overlock/storage/v1beta1"

var _ v1beta1.QueryServer = Keeper{}
