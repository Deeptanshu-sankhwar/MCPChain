package keeper

import (
	"mcpchain/x/attestation/types"
)

var _ types.QueryServer = Keeper{}
