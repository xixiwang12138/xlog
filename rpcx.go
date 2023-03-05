package xlog

import (
	"context"
	"github.com/smallnest/rpcx/protocol"
)

var (
	RPCXServerPreHandle   = &serverPreHandle{}
	RPCXServerAfterHandle = &serverAfterHandle{}
)

type serverPreHandle struct{}

func (handler *serverPreHandle) PostReadRequest(ctx context.Context, r *protocol.Message, e error) error {
	reqId, ok := r.Metadata[ReqHeader]
	if !r.IsOneway() && !r.IsHeartbeat() {
		xl := NewLogger()
		xl.SetFlags(Ldate | Ltime | Llevel)
		if ok {
			xl.SetPrefix("[" + reqId + "]")
		}
		xl.Infof("[RPC Handle] %s, Args: %s \n", r.ServiceMethod, string(r.Payload))
	}
	return nil
}

type serverAfterHandle struct{}

func (handler *serverAfterHandle) PreWriteResponse(ctx context.Context, r *protocol.Message, resp *protocol.Message, err error) error {
	reqId, ok := r.Metadata[ReqHeader]
	if !r.IsOneway() && !r.IsHeartbeat() {
		xl := NewLogger()
		xl.SetFlags(Ldate | Ltime | Llevel)
		if ok {
			xl.SetPrefix("[" + reqId + "]")
		}
		if err != nil {
			xl.Error("[Rpc Server Internal] Handle: %s, error: %s", r.ServiceMethod, err.Error())
			return nil
		}
		xl.Infof("[RPC Reply] %s, Response: %s \n", r.ServiceMethod, string(resp.Payload))
	}
	return nil
}
