package keeper

import (
	"overlock/x/storage/types"
)

var _ types.QueryServer = Keeper{}
