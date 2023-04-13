package bet_test

import (
	"testing"

	"github.com/playerfury/xfury/testutil/nullify"
	simappUtil "github.com/playerfury/xfury/testutil/simapp"
	"github.com/playerfury/xfury/x/bet"
	"github.com/playerfury/xfury/x/bet/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	tApp, ctx, err := simappUtil.GetTestObjects()
	require.NoError(t, err)

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		BetList: []types.Bet{
			{
				UID:     "0",
				Creator: simappUtil.TestParamUsers["user1"].Address.String(),
			},
			{
				UID:     "1",
				Creator: simappUtil.TestParamUsers["user2"].Address.String(),
			},
		},
		Uid2IdList: []types.UID2ID{
			{
				UID: "0",
				ID:  1,
			},
			{
				UID: "1",
				ID:  2,
			},
		},
		ActiveBetList: []types.ActiveBet{
			{
				UID:     "1",
				Creator: simappUtil.TestParamUsers["user1"].Address.String(),
			},
		},
		SettledBetList: []types.SettledBet{
			{
				UID:           "1",
				BettorAddress: simappUtil.TestParamUsers["user1"].Address.String(),
			},
		},
		Stats: types.BetStats{
			Count: 2,
		},
	}

	bet.InitGenesis(ctx, tApp.BetKeeper, genesisState)
	got := bet.ExportGenesis(ctx, tApp.BetKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.BetList, got.BetList)
}
