package xlog

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var GormLogger = gormLogger{}

type gormLogger struct{}

func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *gormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	xl, ok := ctx.(*XLogger)
	if !ok {
		panic("invalid parameter in gorm, first must pass xlog pointer")
	}
	xl.SetFlags(Ldate | Ltime)
	xl.Info(s)
	xl.SetFlags(Ldefault)
}

func (l *gormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
}

func (l *gormLogger) Error(ctx context.Context, s string, i ...interface{}) {
}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	l.Info(ctx, fmt.Sprintf("[SQL] %s ==> (%dms) | rows: %d", sql, elapsed.Milliseconds(), rows))
}

func (x *XLogger) SqlTxBegin(db *gorm.DB, opts ...*sql.TxOptions) {
	x.PutValue(GormTransactionHeader, db.Begin(opts...))
}

func (x *XLogger) Tx() (tx *gorm.DB) {
	tx, ok := x.Value(GormTransactionHeader).(*gorm.DB)
	if ok {
		return nil
	}
	return tx
}
