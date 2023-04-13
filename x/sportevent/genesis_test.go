package sportevent_test

import (
	"testing"

	"github.com/playerfury/xfury/testutil/nullify"
	simappUtil "github.com/playerfury/xfury/testutil/simapp"
	"github.com/playerfury/xfury/x/sportevent"
	"github.com/playerfury/xfury/x/sportevent/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SportEventList: []types.SportEvent{
			{
				UID: "0",
			},
			{
				UID: "1",
			},
		},
	}

	tApp, ctx, err := simappUtil.GetTestObjects()
	require.NoError(t, err)

	sportevent.InitGenesis(ctx, tApp.SportEventKeeper, genesisState)
	got := sportevent.ExportGenesis(ctx, tApp.SportEventKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.SportEventList, got.SportEventList)
}
