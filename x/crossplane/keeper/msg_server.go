package keeper

import "github.com/overlock-network/api/go/node/overlock/crossplane/v1beta1"

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) v1beta1.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ v1beta1.MsgServer = msgServer{}
