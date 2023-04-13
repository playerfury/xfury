package keeper

import (
	"github.com/playerfury/xfury/x/bet/types"
	sporteventtypes "github.com/playerfury/xfury/x/sportevent/types"
)

// KeeperTest is a wrapper object for the keeper, It is being used
// to export unexported methods of the keeper
type KeeperTest = Keeper

func (k KeeperTest) ProcessBetResultAndStatus(bet *types.Bet, sportEvent sporteventtypes.SportEvent) error {
	return processBetResultAndStatus(bet, sportEvent)
}

func (k KeeperTest) CheckBetStatus(bet *types.Bet) error {
	return checkBetStatus(bet.Status)
}
