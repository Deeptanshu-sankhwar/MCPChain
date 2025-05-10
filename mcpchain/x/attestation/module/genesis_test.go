package attestation_test

import (
	"testing"

	keepertest "mcpchain/testutil/keeper"
	"mcpchain/testutil/nullify"
	attestation "mcpchain/x/attestation/module"
	"mcpchain/x/attestation/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AttestationKeeper(t)
	attestation.InitGenesis(ctx, k, genesisState)
	got := attestation.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
