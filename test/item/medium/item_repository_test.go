//go:build medium

package medium_test

import (
	"char5742/ecsite-sample/internal/item/domain"
	"char5742/ecsite-sample/internal/item/infra"
	"char5742/ecsite-sample/test/share"
	"context"
	"testing"
	"time"
)

func TestItemRepository_Read(t *testing.T) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "account", "00000000-0000-0000-0000-000000000000")

	tx, err := tdb.BeginTx(ctx, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			t.Fatalf("rollback failed: %v", err)
		}
	}()

	// 切り出したSQLファイルを実行
	share.ExecSQLFile(t, ctx, "sql/init-read.sql", tx)

	// ここから先は元のテストと同じ
	repo := infra.NewItemRepositoryImpl()

	item, err := repo.Read(ctx, tx, domain.NewItemID("44444444-4444-4444-4444-444444444444"))
	if err != nil {
		t.Fatalf("unexpected error calling Read: %v", err)
	}

	if item.ID.UUID.String() != "44444444-4444-4444-4444-444444444444" {
		t.Errorf("want 44444444-4444-4444-4444-444444444444, have %s", item.ID.UUID.String())
	}

	wantDescription := "Cute dog"
	if item.Description != wantDescription {
		t.Errorf("want %s, have %s", wantDescription, item.Description)
	}

	wantPrice := 10000
	if item.Price != wantPrice {
		t.Errorf("want %d, have %d", wantPrice, item.Price)
	}

	wantBirthday, _ := time.Parse("2006-01-02", "2020-01-01")
	if !item.Birthday.Equal(wantBirthday) {
		t.Errorf("want %s, have %s", wantBirthday, item.Birthday)
	}

	wantImage := "dog.jpg"
	if item.Image != wantImage {
		t.Errorf("want image '%s', have %s", wantImage, item.Image)
	}

	wantGenderName := "Male"
	if item.Gender.Name != wantGenderName {
		t.Errorf("want Gender.Name '%s', have %s", wantGenderName, item.Gender.Name)
	}

	wantBreedName := "Bulldog"
	if item.Breed.Name != wantBreedName {
		t.Errorf("want Breed.Name '%s', have %s", wantBreedName, item.Breed.Name)
	}

	wantColorName := "Brown"
	if item.Color.Name != wantColorName {
		t.Errorf("want Color.Name '%s', have %s", wantColorName, item.Color.Name)
	}

}
