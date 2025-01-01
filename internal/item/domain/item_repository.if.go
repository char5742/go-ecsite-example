package domain

import (
	"context"
	"database/sql"
)

type ItemRepository interface {
	Read(ctx context.Context, tx *sql.Tx, id ItemID) (*Item, error)
	Save(ctx context.Context, tx *sql.Tx, item *Item) error
}
