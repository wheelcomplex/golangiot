package remoting

import (
	"github.com/ttstringiot/golangiot/stgnet/netm"
	"github.com/ttstringiot/golangiot/stgnet/protocol"
)

// RPCHook rpc hook, use send msg
type RPCHook interface {
	DoBeforeRequest(ctx netm.Context, request *protocol.RemotingCommand)
	DoAfterResponse(ctx netm.Context, request *protocol.RemotingCommand, response *protocol.RemotingCommand)
}
