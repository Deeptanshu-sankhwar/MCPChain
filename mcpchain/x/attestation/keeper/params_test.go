package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "mcpchain/testutil/keeper"
	"mcpchain/x/attestation/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AttestationKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
