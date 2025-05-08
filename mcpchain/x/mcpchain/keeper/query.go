package keeper

import (
	"mcpchain/x/mcpchain/types"
)

var _ types.QueryServer = Keeper{}
