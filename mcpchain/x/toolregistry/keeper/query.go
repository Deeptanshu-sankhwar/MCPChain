package keeper

import (
	"mcpchain/x/toolregistry/types"
)

var _ types.QueryServer = Keeper{}
