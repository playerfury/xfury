package keeper

import (
	"github.com/playerfury/xfury/x/house/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the house MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}
