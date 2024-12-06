package keeper

import (
	"overlock/x/blocks/types"
)

var _ types.QueryServer = Keeper{}
