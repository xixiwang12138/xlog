package xlog

import (
	"context"
	"os"
)

type XLogger struct {
	*Logger
	context.Context
	reqId string
}

func (x *XLogger) PutValue(key, val any) {
	x.Context = context.WithValue(x.Context, key, val)
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
