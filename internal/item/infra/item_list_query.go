//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package infra

import (
	"context"
	"fmt"
	"strings"

	"char5742/ecsite-sample/internal/item/domain"
	"char5742/ecsite-sample/pkg/db"
)

type ItemListQuery interface {
	ItemList(ctx context.Context, tx db.TX) ([]domain.Item, error)
	ItemListByCondition(ctx context.Context, tx db.TX, condition ItemListCondition) ([]domain.Item, error)
}

type ItemListQueryImpl struct {
}

// インターフェース実装用のコンストラクタ
func NewItemListQuery() ItemListQuery {
	return &ItemListQueryImpl{}
}

func (t *ItemListQueryImpl) ItemList(ctx context.Context, tx db.TX) ([]domain.Item, error) {
	query := `
		SELECT
			items.id,
			items.description,
			items.price,
			items.birthday,
			items.image,
			items.is_deleted,
			genders.id AS gender_id,
			genders.name AS gender_name,
			breeds.id AS breed_id,
			breeds.name AS breed_name,
			colors.id AS color_id,
			colors.name AS color_name
		FROM
			items
		JOIN 
			genders ON items.gender_id = genders.id
		JOIN
			breeds ON items.breed_id = breeds.id
		JOIN
			colors ON items.color_id = colors.id
	`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.Item
	for rows.Next() {
		var item domain.Item
		if err := rows.Scan(
			&item.ID,
			&item.Description,
			&item.Price,
			&item.Birthday,
			&item.Image,
			&item.IsDeleted,
			&item.Gender.ID,
			&item.Gender.Name,
			&item.Breed.ID,
			&item.Breed.Name,
			&item.Color.ID,
			&item.Color.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
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

func (t *ItemListQueryImpl) ItemListByCondition(ctx context.Context, tx db.TX, condition ItemListCondition) ([]domain.Item, error) {
	baseQuery := `
		SELECT
			i.id,
			i.description,
			i.price,
			i.birthday,
			i.image,
			i.is_deleted,
			g.id AS gender_id,
			g.name AS gender_name,
			b.id AS breed_id,
			b.name AS breed_name,
			c.id AS color_id,
			c.name AS color_name
		FROM
			items i
			JOIN genders g   ON i.gender_id = g.id
			JOIN breeds b   ON i.breed_id = b.id
			JOIN colors c   ON i.color_id = c.id
		WHERE 1=1
	`

	var (
		conds []string
		args  []interface{}
	)

	// プレースホルダに使う$番号を管理する
	placeHolderIndex := 1

	// Gender
	if condition.GenderCond != nil && len(condition.GenderCond.GenderIDList) > 0 {
		q, a, newIndex := buildInCondition("i.gender_id", condition.GenderCond.GenderIDList, placeHolderIndex)
		conds = append(conds, q)
		args = append(args, a...)
		placeHolderIndex = newIndex
	}

	// Breed
	if condition.BreedCond != nil && len(condition.BreedCond.BreedIDList) > 0 {
		q, a, newIndex := buildInCondition("i.breed_id", condition.BreedCond.BreedIDList, placeHolderIndex)
		conds = append(conds, q)
		args = append(args, a...)
		placeHolderIndex = newIndex
	}

	// Color
	if condition.ColorCond != nil && len(condition.ColorCond.ColorIDList) > 0 {
		q, a, newIndex := buildInCondition("i.color_id", condition.ColorCond.ColorIDList, placeHolderIndex)
		conds = append(conds, q)
		args = append(args, a...)
		placeHolderIndex = newIndex
	}

	// Price
	if condition.PriceCond != nil {
		q, a, _ := buildPriceCondition(*condition.PriceCond, placeHolderIndex)
		if q != "" {
			conds = append(conds, q)
			args = append(args, a...)
		}
	}

	// WHERE句追加
	if len(conds) > 0 {
		baseQuery += " AND " + strings.Join(conds, " AND ")
	}

	rows, err := tx.QueryContext(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.Item
	for rows.Next() {
		var item domain.Item
		if err := rows.Scan(
			&item.ID,
			&item.Description,
			&item.Price,
			&item.Birthday,
			&item.Image,
			&item.IsDeleted,
			&item.Gender.ID,
			&item.Gender.Name,
			&item.Breed.ID,
			&item.Breed.Name,
			&item.Color.ID,
			&item.Color.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// buildInCondition
// field: カラム名 ("i.gender_id"など)
// idList: IN (...) でバインドするIDリスト
// startIndex: $1, $2 などの開始番号
func buildInCondition(field string, idList []string, startIndex int) (string, []interface{}, int) {
	if len(idList) == 0 {
		return "", nil, startIndex
	}
	placeholders := make([]string, len(idList))
	args := make([]interface{}, 0, len(idList))

	for i, id := range idList {
		placeholders[i] = fmt.Sprintf("$%d", startIndex)
		args = append(args, id)
		startIndex++
	}

	query := fmt.Sprintf("%s IN (%s)", field, strings.Join(placeholders, ","))
	return query, args, startIndex
}

// buildPriceCondition
// startIndex: $1, $2 などの開始番号
func buildPriceCondition(price PriceCondition, startIndex int) (string, []interface{}, int) {
	if price.Min == 0 && price.Max == 0 {
		return "", nil, startIndex
	}
	// Maxのみ
	if price.Min == 0 {
		query := fmt.Sprintf("i.price <= $%d", startIndex)
		args := []interface{}{price.Max}
		startIndex++
		return query, args, startIndex
	}
	// Minのみ
	if price.Max == 0 {
		query := fmt.Sprintf("i.price >= $%d", startIndex)
		args := []interface{}{price.Min}
		startIndex++
		return query, args, startIndex
	}
	// MinとMax両方
	query := fmt.Sprintf("i.price BETWEEN $%d AND $%d", startIndex, startIndex+1)
	args := []interface{}{price.Min, price.Max}
	startIndex += 2
	return query, args, startIndex
}
