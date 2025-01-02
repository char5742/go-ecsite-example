package infra

import (
	"context"
	"fmt"

	"char5742/ecsite-sample/internal/item/domain"
	"char5742/ecsite-sample/pkg/db"
)

type ItemRepository struct {
}

// インターフェース実装用のコンストラクタ
func NewItemRepository() *ItemRepository {
	return &ItemRepository{}
}

// Read は指定された ID の Item を取得します。
// 指定された ID の Item が存在しない場合は `sql.ErrNoRows` を返却します。
func (t *ItemRepository) Read(ctx context.Context, tx db.TX, id domain.ItemID) (*domain.Item, error) {
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
		WHERE
			items.id = $1
	`
	args := []interface{}{id}
	row := tx.QueryRowContext(ctx, query, args...)
	var item domain.Item
	if err := row.Scan(
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
	return &item, nil
}

func (t *ItemRepository) Save(ctx context.Context, tx db.TX, item *domain.Item) error {
	query := `
		INSERT INTO items (
			id,
			description,
			price,
			birthday,
			image,
			is_deleted,
			gender_id,
			breed_id,
			color_id,
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (id) DO UPDATE SET
			description = $2,
			price = $3,
			birthday = $4,
			image = $5,
			gender_id = $6,
			breed_id = $7,
			color_id = $8,
	`
	args := []interface{}{
		item.ID,
		item.Description,
		item.Price,
		item.Birthday,
		item.Image,
		item.IsDeleted,
		item.Gender.ID,
		item.Breed.ID,
		item.Color.ID,
	}

	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		err = fmt.Errorf("failed to save item: %w", err)
		return err
	}
	return nil
}
