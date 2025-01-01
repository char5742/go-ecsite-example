package domain

import "context"

type ItemRepository interface {
	Read(ctx context.Context, id ItemID) (*Item, error)
	Save(ctx context.Context, item *Item) error
}
