package keeper

import (
	"github.com/playerfury/xfury/x/orderbook/types"
)

var _ types.QueryServer = Keeper{}
