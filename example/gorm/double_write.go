package gorm

import (
	"context"
	"database/sql"
	"errors"
	"sync/atomic"

	"gorm.io/gorm"
)

type DoubleWritePool struct {
	src gorm.ConnPool // 源数据库
	dst gorm.ConnPool // 目标数据库

	pattern atomic.Value // 双写模式
}

const (
	PatternSrcOnly  = "SrcOnly"
	PatternSrcFirst = "SrcFirst"
	PatternDstOnly  = "DstOnly"
	PatternDstFirst = "DstFirst"
)

func (d *DoubleWritePool) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DoubleWritePool) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	switch d.pattern.Load() {
	case PatternSrcOnly: // 只写源库
		return d.src.ExecContext(ctx, query, args...)
	case PatternSrcFirst: // 双写，先写源库再写目标库
		res, err := d.src.ExecContext(ctx, query, args...)
		if err == nil { // 源库写入成功后再写目标库
			_, err1 := d.dst.ExecContext(ctx, query, args...)
			if err1 != nil { // 写目标库失败，记录日志
				// 记录日志
			}
		}
		return res, err
	case PatternDstOnly: // 只写目标库
		return d.dst.ExecContext(ctx, query, args...)
	case PatternDstFirst: // 双写，先写目标库再写源库
		res, err := d.dst.ExecContext(ctx, query, args...)
		if err == nil { // 目标库写入成功后再写源库
			_, err1 := d.src.ExecContext(ctx, query, args...)
			if err1 != nil { // 写源库失败，记录日志
				// 记录日志
			}
		}
		return res, err
	default:
		return nil, errors.New("未知的双写模式")
	}
}

func (d *DoubleWritePool) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	switch d.pattern.Load() {
	case PatternSrcOnly, PatternSrcFirst:
		return d.src.QueryContext(ctx, query, args...)
	case PatternDstOnly, PatternDstFirst:
		return d.dst.QueryContext(ctx, query, args...)
	default:
		return nil, errors.New("未知的双写模式")
	}
}

func (d *DoubleWritePool) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	switch d.pattern.Load() {
	case PatternSrcOnly, PatternSrcFirst:
		return d.src.QueryRowContext(ctx, query, args...)
	case PatternDstOnly, PatternDstFirst:
		return d.dst.QueryRowContext(ctx, query, args...)
	default:
		return nil
	}
}
