package database

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Q interface {
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
}

type Queryable struct {
	q Q
	db *sqlx.DB
	tx *sqlx.Tx
}

func NewQueryable(db interface{}) Queryable {
	var newQ Q
	var newDb *sqlx.DB
	var newTx *sqlx.Tx

	qTx, ok := db.(*sqlx.Tx)
	if !ok {
		qDb := db.(*sqlx.DB)
		newQ = Q(qDb)
	} else {
		newQ = Q(qTx)
	}

	return Queryable{
		q:  newQ,
		db: newDb,
		tx: newTx,
	}
}

func (q *Queryable) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return q.q.BindNamed(query, arg)
}

func (q *Queryable) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return q.q.ExecContext(ctx, query, args)
}

func (q *Queryable) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return q.q.NamedExecContext(ctx, query, arg)
}

func (q *Queryable) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	return q.q.NamedQuery(query, arg)
}

func (q *Queryable) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return q.q.QueryRowxContext(ctx, query, args)
}

func (q *Queryable) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return q.q.QueryxContext(ctx, query, args)
}

func (q *Queryable) NamedQueryRowxContext(ctx context.Context, query string, arg interface{}) (*sqlx.Row, error) {
	query, args, err := q.q.BindNamed(query, arg)
	if nil != err {
		return nil, err
	}

	return q.q.QueryRowxContext(ctx, query, args...), nil
}

func (q *Queryable) NamedQueryxContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	query, args, err := q.q.BindNamed(query, arg)
	if nil != err {
		return nil, err
	}

	return q.q.QueryxContext(ctx, query, args...)
}