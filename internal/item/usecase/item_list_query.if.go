package usecase

import (
	"context"
	"database/sql"

	"char5742/ecsite-sample/internal/item/domain"
)

type ItemListQuery interface {
	// 商品一覧を取得する
	ItemList(ctx context.Context, tx *sql.Tx) ([]*domain.Item, error)
	// 条件に合致する商品一覧を取得する
	ItemListByCondition(ctx context.Context, tx *sql.Tx, condition ItemListCondition) ([]*domain.Item, error)
}

// 商品一覧取得条件
type ItemListCondition struct {
	// 性別
	GenderCond *GenderCondition
	// 種別
	BreedCond *BreedCondition
	// 色
	ColorCond *ColorCondition
	// 価格
	PriceCond *PriceCondition
}

type GenderCondition struct {
	GenderIDList []string
}

type BreedCondition struct {
	BreedIDList []string
}

type ColorCondition struct {
	ColorIDList []string
}

type PriceCondition struct {
	// 最小価格
	Min int
	// 最大価格
	Max int
}
