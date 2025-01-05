//go:build medium

package medium_test

import (
	"context"
	"testing"

	app "char5742/ecsite-sample/internal/app/infra"
	"char5742/ecsite-sample/internal/item/infra"
	"char5742/ecsite-sample/test/share"
)

func TestItemListQueryImpl_ItemList(t *testing.T) {
	ctx := context.Background()

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
	share.ExecSQLFile(t, ctx, "sql/init-list-query.sql", tx)

	// ここから先は元のテストと同じ
	itemListQuery := infra.NewItemListQuery()

	items, err := itemListQuery.ItemList(ctx, tx)
	if err != nil {
		t.Fatalf("unexpected error calling ItemList: %v", err)
	}

	if len(items) != 8 {
		t.Errorf("expected 8 items, got %d", len(items))
	}
}

func TestItemListQueryImpl_ItemListByCondition(t *testing.T) {
	ctx := context.Background()
	tx, err := tdb.BeginTx(ctx, nil)
	defer func() {
		if err := tx.Rollback(); err != nil {
			t.Fatalf("rollback failed: %v", err)
		}
	}()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	share.ExecSQLFile(t, ctx, "sql/init-list-query.sql", tx)
	itemListQuery := infra.NewItemListQuery()

	t.Run("Case1: Gender=Male, Breed=Bulldog, Color=Brown, Price=0..15000", func(t *testing.T) {
		cond := infra.ItemListCondition{
			GenderCond: infra.GenderCondition{
				GenderIDList: []string{"11111111-1111-1111-1111-111111111111"},
			},
			BreedCond: infra.BreedCondition{
				BreedIDList: []string{"33333333-3333-3333-3333-333333333333"},
			},
			ColorCond: infra.ColorCondition{
				ColorIDList: []string{"55555555-5555-5555-5555-555555555555"},
			},
			PriceCond: infra.PriceCondition{
				Min:   0,
				Max:   15000,
				Valid: true,
			},
			Pagination: app.Pagination{
				CurrentPage: 1,
				PerPage:     10,
			},
		}
		items, err := itemListQuery.ItemListByCondition(ctx, tx, cond)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(items) != 1 {
			t.Errorf("unexpected items count: got=%d, want=2", len(items))
		}
	})

	t.Run("Case2: Gender=Male, Breed=Golden Retriever, Color=White, Price=0..15000", func(t *testing.T) {
		cond := infra.ItemListCondition{
			GenderCond: infra.GenderCondition{
				GenderIDList: []string{"11111111-1111-1111-1111-111111111111"},
			},
			BreedCond: infra.BreedCondition{
				BreedIDList: []string{"BBBBBBBB-BBBB-BBBB-BBBB-BBBBBBBBBBBB"},
			},
			ColorCond: infra.ColorCondition{
				ColorIDList: []string{"66666666-6666-6666-6666-666666666666"},
			},
			PriceCond: infra.PriceCondition{
				Min:   0,
				Max:   15000,
				Valid: true,
			},
			Pagination: app.Pagination{
				CurrentPage: 1,
				PerPage:     10,
			},
		}
		items, err := itemListQuery.ItemListByCondition(ctx, tx, cond)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// 12000円のGolden Retrieverが該当するはず
		if len(items) != 1 {
			t.Errorf("unexpected items count: got=%d, want=1", len(items))
		}
	})

	t.Run("Case3: Gender=Female, Breed=Shiba Inu, Color=Black, Price=0..40000", func(t *testing.T) {
		cond := infra.ItemListCondition{
			GenderCond: infra.GenderCondition{
				GenderIDList: []string{"22222222-2222-2222-2222-222222222222"},
			},
			BreedCond: infra.BreedCondition{
				BreedIDList: []string{"CCCCCCCC-CCCC-CCCC-CCCC-CCCCCCCCCCCC"},
			},
			ColorCond: infra.ColorCondition{
				ColorIDList: []string{"DDDDDDDD-DDDD-DDDD-DDDD-DDDDDDDDDDDD"},
			},
			PriceCond: infra.PriceCondition{
				Min:   0,
				Max:   40000,
				Valid: true,
			},
			Pagination: app.Pagination{
				CurrentPage: 1,
				PerPage:     10,
			},
		}
		items, err := itemListQuery.ItemListByCondition(ctx, tx, cond)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// 25000円のBlack Shiba Inuが該当するはず
		if len(items) != 1 {
			t.Errorf("unexpected items count: got=%d, want=1", len(items))
		}
	})

	t.Run("Case4: Gender=Undefined, Breed=Poodle, Color=White, Price=0..15000", func(t *testing.T) {
		cond := infra.ItemListCondition{
			GenderCond: infra.GenderCondition{
				GenderIDList: []string{"AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA"},
			},
			BreedCond: infra.BreedCondition{
				BreedIDList: []string{"44444444-4444-4444-4444-444444444444"},
			},
			ColorCond: infra.ColorCondition{
				ColorIDList: []string{"66666666-6666-6666-6666-666666666666"},
			},
			PriceCond: infra.PriceCondition{
				Min:   0,
				Max:   15000,
				Valid: true,
			},
			Pagination: app.Pagination{
				CurrentPage: 1,
				PerPage:     10,
			},
		}
		items, err := itemListQuery.ItemListByCondition(ctx, tx, cond)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// 15000円のUndefined Poodleが該当するはず
		if len(items) != 1 {
			t.Errorf("unexpected items count: got=%d, want=1", len(items))
		}
	})

	t.Run("Case5: Gender=Male, Breed=Shiba Inu, Color=Golden, Price=0..10000 (範囲外データのみ存在)", func(t *testing.T) {
		cond := infra.ItemListCondition{
			GenderCond: infra.GenderCondition{
				GenderIDList: []string{"11111111-1111-1111-1111-111111111111"},
			},
			BreedCond: infra.BreedCondition{
				BreedIDList: []string{"CCCCCCCC-CCCC-CCCC-CCCC-CCCCCCCCCCCC"},
			},
			ColorCond: infra.ColorCondition{
				ColorIDList: []string{"EEEEEEEE-EEEE-EEEE-EEEE-EEEEEEEEEEEE"},
			},
			PriceCond: infra.PriceCondition{
				Min:   0,
				Max:   10000,
				Valid: true,
			},
			Pagination: app.Pagination{
				CurrentPage: 1,
				PerPage:     10,
			},
		}
		items, err := itemListQuery.ItemListByCondition(ctx, tx, cond)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// 30000円のGolden ShibaがDBにあるが、Max=10000なので該当なし
		if len(items) != 0 {
			t.Errorf("unexpected items count: got=%d, want=0", len(items))
		}
	})

	t.Run("Case6: Gender=Unknown, Breed=Bulldog, Color=Brown, Price=0..10000", func(t *testing.T) {
		cond := infra.ItemListCondition{
			GenderCond: infra.GenderCondition{
				GenderIDList: []string{"99999999-9999-9999-9999-999999999999"},
			},
			BreedCond: infra.BreedCondition{
				BreedIDList: []string{"33333333-3333-3333-3333-333333333333"},
			},
			ColorCond: infra.ColorCondition{
				ColorIDList: []string{"55555555-5555-5555-5555-555555555555"},
			},
			PriceCond: infra.PriceCondition{
				Min:   0,
				Max:   10000,
				Valid: true,
			},
			Pagination: app.Pagination{
				CurrentPage: 1,
				PerPage:     10,
			},
		}
		items, err := itemListQuery.ItemListByCondition(ctx, tx, cond)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// 5000円のUnknown Bulldogが該当するはず
		if len(items) != 1 {
			t.Errorf("unexpected items count: got=%d, want=1", len(items))
		}
	})

	t.Run("Case7: Gender=Female, Breed=Golden Retriever, Color=Golden, Price=0..99999", func(t *testing.T) {
		cond := infra.ItemListCondition{
			GenderCond: infra.GenderCondition{
				GenderIDList: []string{"22222222-2222-2222-2222-222222222222"},
			},
			BreedCond: infra.BreedCondition{
				BreedIDList: []string{"BBBBBBBB-BBBB-BBBB-BBBB-BBBBBBBBBBBB"},
			},
			ColorCond: infra.ColorCondition{
				ColorIDList: []string{"EEEEEEEE-EEEE-EEEE-EEEE-EEEEEEEEEEEE"},
			},
			PriceCond: infra.PriceCondition{
				Min:   0,
				Max:   99999,
				Valid: true,
			},
			Pagination: app.Pagination{
				CurrentPage: 1,
				PerPage:     10,
			},
		}
		items, err := itemListQuery.ItemListByCondition(ctx, tx, cond)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// 35000円のFemale Golden Retrieverが該当するはず
		if len(items) != 1 {
			t.Errorf("unexpected items count: got=%d, want=1", len(items))
		}
	})
}
