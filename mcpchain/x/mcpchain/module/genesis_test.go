package mcpchain_test

import (
	"testing"

	keepertest "mcpchain/testutil/keeper"
	"mcpchain/testutil/nullify"
	mcpchain "mcpchain/x/mcpchain/module"
	"mcpchain/x/mcpchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.McpchainKeeper(t)
	mcpchain.InitGenesis(ctx, k, genesisState)
	got := mcpchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
