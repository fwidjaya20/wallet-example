package database

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var QUERY_KEY = "query"

func NewQueryContext(ctx context.Context, queryable Queryable) context.Context {
	return context.WithValue(ctx, QUERY_KEY, queryable)
}

func QueryFromContext(ctx context.Context) (Queryable, bool) {
	q, ok := ctx.Value(QUERY_KEY).(Queryable)
	return q, ok
}

type TransactionCallback func(ctx context.Context) error

func RunInTransaction(ctx context.Context, db *sqlx.DB, fn TransactionCallback) error {
	tx, err := db.Beginx()
	if nil != err {
		return err
	}

	ctx = NewQueryContext(ctx, NewQueryable(tx))
	err = fn(ctx)

	if nil != err {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if nil != err {
		return fmt.Errorf("error when committing transaction: %v", err)
	}

	return nil
}