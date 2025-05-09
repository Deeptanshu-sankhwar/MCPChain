package toolregistry_test

import (
	"testing"

	keepertest "mcpchain/testutil/keeper"
	"mcpchain/testutil/nullify"
	toolregistry "mcpchain/x/toolregistry/module"
	"mcpchain/x/toolregistry/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ToolregistryKeeper(t)
	toolregistry.InitGenesis(ctx, k, genesisState)
	got := toolregistry.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
