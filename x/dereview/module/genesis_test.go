package dereview_test

import (
	"testing"

	keepertest "dereview/testutil/keeper"
	"dereview/testutil/nullify"
	dereview "dereview/x/dereview/module"
	"dereview/x/dereview/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.DereviewKeeper(t)
	dereview.InitGenesis(ctx, k, genesisState)
	got := dereview.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
