package remoting

import (
	"github.com/ttstringiot/golangiot/stgnet/netm"
	"github.com/ttstringiot/golangiot/stgnet/protocol"
)

// RemotingClient remoting client define
type RemotingClient interface {
	InvokeSync(addr string, request *protocol.RemotingCommand, timeoutMillis int64) (*protocol.RemotingCommand, error)
	InvokeAsync(addr string, request *protocol.RemotingCommand, timeoutMillis int64, invokeCallback InvokeCallback) error
	InvokeOneway(addr string, request *protocol.RemotingCommand, timeoutMillis int64) error
	RegisterProcessor(requestCode int32, processor RequestProcessor)
	RegisterRPCHook(rpcHook RPCHook)
	GetNameServerAddressList() []string
	UpdateNameServerAddressList(addrs []string)
	RegisterContextListener(contextListener netm.ContextListener)
	Start()
	Shutdown()
}
