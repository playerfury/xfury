package keeper

import (
	"github.com/playerfury/xfury/x/dvm/types"
)

var _ types.QueryServer = Keeper{}
