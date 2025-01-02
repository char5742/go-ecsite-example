//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package db

import (
	"context"
	"database/sql"
)

// DB interface to abstract the database operations
type TX interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Rollback() error
	Commit() error
}

type DB interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (TX, error)
}
