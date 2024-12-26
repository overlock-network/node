package keeper

import (
	"overlock/x/crossplane/types"
)

var _ types.QueryServer = Keeper{}
