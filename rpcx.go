package xlog

import (
	"context"
	"github.com/smallnest/rpcx/protocol"
)

var (
	RPCXServerPreHandle = &serverPreHandle{}
)

type serverPreHandle struct{}

func (handler *serverPreHandle) PostReadRequest(ctx context.Context, r *protocol.Message, e error) error {
	reqId, ok := r.Metadata[ReqHeader]
	if !r.IsOneway() && !r.IsHeartbeat() {
		xl := NewLogger()
		if ok {
			xl.reqId = reqId
		}
		xl.Infof("[RPC Handle] %s, Args: %s \n", string(r.Payload))
	}
	return nil
}
