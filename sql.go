package xlog

import (
	"context"
	"fmt"
	"gorm.io/gorm/logger"
	"time"
)

var GormLogger = gormLogger{}

type gormLogger struct{}

func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *gormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	NewCtxLogger(ctx).Info(s)
}

func (l *gormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
}

func (l *gormLogger) Error(ctx context.Context, s string, i ...interface{}) {
}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	l.Info(ctx, fmt.Sprintf("[Gorm] %dms %s rows: %d\n", elapsed.Milliseconds(), sql, rows))
}
