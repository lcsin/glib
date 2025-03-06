package igrom

import (
	"context"
	"database/sql"
	"errors"
	"sync/atomic"

	"gorm.io/gorm"
)

const (
	PatternSrcOnly  = "SrcOnly"
	PatternSrcFirst = "SrcFirst"
	PatternDstOnly  = "DstOnly"
	PatternDstFirst = "DstFirst"
)

var errUnknownPattern = errors.New("未知的双写模式")

type DoubleWritePool struct {
	src gorm.ConnPool // 源数据库
	dst gorm.ConnPool // 目标数据库

	pattern atomic.Value // 双写模式
}

func NewDoubleWritePool(srcDB *gorm.DB, dst *gorm.DB) *DoubleWritePool {
	var pattern atomic.Value
	pattern.Store(PatternSrcOnly)

	return &DoubleWritePool{
		src:     srcDB.ConnPool,
		dst:     dst.ConnPool,
		pattern: pattern,
	}
}

func (d *DoubleWritePool) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return nil, errors.New("双写模式下不支持")
}

// ExecContext 增、删、改、表结构的修改语句都会走这里
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
		return nil, errUnknownPattern
	}
}

// QueryContext 查询语句走这里
func (d *DoubleWritePool) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	switch d.pattern.Load() {
	case PatternSrcOnly, PatternSrcFirst:
		return d.src.QueryContext(ctx, query, args...)
	case PatternDstOnly, PatternDstFirst:
		return d.dst.QueryContext(ctx, query, args...)
	default:
		return nil, errUnknownPattern
	}
}

// QueryRowContext 查询语句走这里
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

// BeginTx 开启事物
func (d *DoubleWritePool) BeginTx(ctx context.Context, opts *sql.TxOptions) (gorm.ConnPool, error) {
	pattern := d.pattern.Load()
	switch pattern {
	case PatternSrcOnly:
		tx, err := d.src.(gorm.TxBeginner).BeginTx(ctx, opts)
		return &DoubleWritePoolTx{src: tx, pattern: PatternSrcOnly}, err
	case PatternSrcFirst:
		return d.startDoubleWriteTx(ctx, opts, d.src, d.dst, PatternSrcFirst)
	case PatternDstOnly:
		tx, err := d.dst.(gorm.TxBeginner).BeginTx(ctx, opts)
		return &DoubleWritePoolTx{src: tx, pattern: PatternDstOnly}, err
	case PatternDstFirst:
		return d.startDoubleWriteTx(ctx, opts, d.dst, d.src, PatternDstFirst)
	default:
		return nil, errUnknownPattern
	}
}

func (d *DoubleWritePool) startDoubleWriteTx(ctx context.Context, opts *sql.TxOptions,
	first, second gorm.ConnPool, pattern string) (*DoubleWritePoolTx, error) {
	src, err := first.(gorm.TxBeginner).BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	dst, err := second.(gorm.TxBeginner).BeginTx(ctx, opts)
	if err != nil {
		// 记录日志
		_ = src.Rollback()
	}

	return &DoubleWritePoolTx{src: src, dst: dst, pattern: pattern}, nil
}

func (d *DoubleWritePool) ChangePattern(pattern string) {
	d.pattern.Store(pattern)
}

type DoubleWritePoolTx struct {
	src     *sql.Tx // 源库事物
	dst     *sql.Tx // 目标库事物
	pattern string  // 双写模式
}

func (d *DoubleWritePoolTx) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return nil, errors.New("双写模式下不支持")
}

func (d *DoubleWritePoolTx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	switch d.pattern {
	case PatternSrcOnly: // 只写源库
		return d.src.ExecContext(ctx, query, args...)
	case PatternSrcFirst: // 双写，先写源库再写目标库
		res, err := d.src.ExecContext(ctx, query, args...)
		if err != nil {
			return res, err
		}

		// 事物可能开启失败
		if d.dst == nil {
			return res, err
		}
		_, err1 := d.dst.ExecContext(ctx, query, args...)
		if err1 != nil { // 写目标库失败，记录日志
			// 记录日志
		}
		return res, err
	case PatternDstOnly: // 只写目标库
		return d.dst.ExecContext(ctx, query, args...)
	case PatternDstFirst: // 双写，先写目标库再写源库
		res, err := d.dst.ExecContext(ctx, query, args...)
		if err != nil {
			return res, err
		}

		// 事物可能开启失败
		if d.src == nil {
			return res, err
		}
		_, err1 := d.src.ExecContext(ctx, query, args...)
		if err1 != nil { // 写源库失败，记录日志
			// 记录日志
		}
		return res, err
	default:
		return nil, errUnknownPattern
	}
}

func (d *DoubleWritePoolTx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	switch d.pattern {
	case PatternSrcOnly, PatternSrcFirst:
		return d.src.QueryContext(ctx, query, args...)
	case PatternDstOnly, PatternDstFirst:
		return d.dst.QueryContext(ctx, query, args...)
	default:
		return nil, errUnknownPattern
	}
}

func (d *DoubleWritePoolTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	switch d.pattern {
	case PatternSrcOnly, PatternSrcFirst:
		return d.src.QueryRowContext(ctx, query, args...)
	case PatternDstOnly, PatternDstFirst:
		return d.dst.QueryRowContext(ctx, query, args...)
	default:
		return nil
	}
}

func (d *DoubleWritePoolTx) Commit() error {
	switch d.pattern {
	case PatternSrcOnly:
		return d.src.Commit()
	case PatternSrcFirst:
		err := d.src.Commit()
		err1 := d.dst.Commit()
		if err1 != nil {
			// 记录日志
		}
		return err
	case PatternDstOnly:
		return d.dst.Commit()
	case PatternDstFirst:
		err := d.dst.Commit()
		err1 := d.src.Commit()
		if err1 != nil {
			// 记录日志
		}
		return err
	default:
		return errUnknownPattern
	}
}

func (d *DoubleWritePoolTx) Rollback() error {
	switch d.pattern {
	case PatternSrcFirst:
		err := d.src.Rollback()
		err1 := d.dst.Rollback()
		if err1 != nil {
			// 记录日志
		}
		return err
	case PatternSrcOnly:
		return d.src.Rollback()
	case PatternDstOnly:
		return d.dst.Rollback()
	case PatternDstFirst:
		err := d.dst.Rollback()
		err1 := d.src.Rollback()
		if err1 != nil {
			// 记录日志
		}
		return err
	default:
		return errUnknownPattern
	}
}
