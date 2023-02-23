package xlog

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
)

var NotRegisterMiddleware = errors.New("not register middleware tracing logger")

const (
	hLoggerGinContext = "xlog_gin_ctx"
)

func TracingLogger(ctxGenerator func() context.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := ctx.GetHeader(ReqHeader)
		if traceId == "" {
			traceId = RandString(DefaultReqIdLen)
			ctx.Header(ReqHeader, traceId)
		}
		xl := NewXLogger(traceId)
		if ctxGenerator != nil {
			xl.Context = ctxGenerator()
		}
		ctx.Set(hLoggerGinContext, xl)
		ctx.Next()
		ctx.Header(ReqHeader, traceId)
	}
}

func FromGin(ctx *gin.Context) *XLogger {
	logger, ok := ctx.Get(hLoggerGinContext)
	if !ok {
		panic(NotRegisterMiddleware.Error())
	}
	return logger.(*XLogger)
}
