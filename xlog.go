package xlog

import (
	"context"
	"os"
)

var (
	RequestIdKey = "X-request-Id"
)

type XLogger struct {
	*Logger
	context.Context
	reqId string
}

func (x *XLogger) Value(arg any) any {
	return x.reqId
}

// NewXLogger 创建一个具有reqId前缀的Logger
func NewXLogger(reqId string) *XLogger {
	prefix := "[" + reqId + "] "
	return &XLogger{
		Logger:  New(os.Stdout, prefix, Ldefault),
		Context: context.Background(),
		reqId:   reqId,
	}
}

// NewLogger 创建无需追踪功能的logger
func NewLogger() *XLogger {
	return &XLogger{
		Logger:  New(os.Stdout, "", Ldefault),
		Context: context.Background(),
	}
}

func NewCtxLogger(ctx context.Context) *XLogger {
	reqId := ctx.Value(RequestIdKey).(string)
	prefix := "[" + reqId + "]"
	return &XLogger{
		Logger:  New(os.Stdout, prefix, Ldefault),
		Context: context.Background(),
		reqId:   reqId,
	}
}
