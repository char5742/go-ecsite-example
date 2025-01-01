package infra

import (
	"context"
	"fmt"
	"time"

	"char5742/ecsite-sample/internal/item/domain"
	"char5742/ecsite-sample/internal/share/infra/db"
)

type ItemRepositoryImpl struct {
	conn *db.DatabaseConnection
}

// インターフェース実装用のコンストラクタ
func NewItemRepositoryImpl(conn *db.DatabaseConnection) domain.ItemRepository {
	return &ItemRepositoryImpl{conn: conn}
}

// Read は指定された ID の Item を取得します。
// 指定された ID の Item が存在しない場合は `sql.ErrNoRows` を返却します。
func (t *ItemRepositoryImpl) Read(ctx context.Context, id domain.ItemID) (*domain.Item, error) {
	query := `
		SELECT
			items.id,
			items.description,
			items.price,
			items.birthday,
			items.image,
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
	row := t.conn.QueryRowContext(ctx, query, args...)
	var item domain.Item
	if err := row.Scan(
		&item.ID,
		&item.Description,
		&item.Price,
		&item.Birthday,
		&item.Image,
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

func (t *ItemRepositoryImpl) Save(ctx context.Context, item *domain.Item) error {
	query := `
		INSERT INTO items (
			id,
			description,
			price,
			birthday,
			image,
			gender_id,
			breed_id,
			color_id,
			created_by,
			update_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (id) DO UPDATE SET
			description = $2,
			price = $3,
			birthday = $4,
			image = $5,
			gender_id = $6,
			breed_id = $7,
			color_id = $8,
			update_by = $10
	`
	args := []interface{}{
		item.ID,
		item.Description,
		item.Price,
		item.Birthday,
		item.Image,
		item.Gender.ID,
		item.Breed.ID,
		item.Color.ID,
		ctx.Value("account"),
	}

	if _, err := t.conn.ExecContext(ctx, query, args...); err != nil {
		err = fmt.Errorf("failed to save item: %w", err)
		return err
	}
	return nil
}

func insertGender(ctx context.Context, conn *db.DatabaseConnection, g *domain.Gender) error {
	account := ctx.Value("account").(string)
	query := `
		INSERT INTO
			id,
			name,
			created_at,
			created_by,
			update_at,
			update_by
		VALUES (?, ?, ?, ?, ?, ?)
	`
	args := []interface{}{
		g.ID,
		g.Name,
		time.Now(),
		account,
		time.Now(),
		account,
	}

	if _, err := conn.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("failed to save Gender: %w", err)
	}
	return nil
}

func insertBreed(ctx context.Context, conn *db.DatabaseConnection, b *domain.Breed) error {
	account := ctx.Value("account").(string)
	query := `
		INSERT INTO
			id,
			name,
			created_at,
			created_by,
			update_at,
			update_by
		VALUES (?, ?, ?, ?, ?, ?)
	`
	args := []interface{}{
		b.ID,
		b.Name,
		time.Now(),
		account,
		time.Now(),
		account,
	}

	if _, err := conn.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("failed to save Breed: %w", err)
	}
	return nil
}

func insertColor(ctx context.Context, conn *db.DatabaseConnection, c *domain.Color) error {
	account := ctx.Value("account").(string)
	query := `
		INSERT INTO
			id,
			name,
			created_at,
			created_by,
			update_at,
			update_by
		VALUES (?, ?, ?, ?, ?, ?)
	`
	args := []interface{}{
		c.ID,
		c.Name,
		time.Now(),
		account,
		time.Now(),
		account,
	}

	if _, err := conn.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("failed to save Color: %w", err)
	}
	return nil
}
