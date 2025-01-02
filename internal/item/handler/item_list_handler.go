//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package handler

import (
	"char5742/ecsite-sample/internal/item/infra"
	"char5742/ecsite-sample/pkg/db"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ItemListHandler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type ItemListHandlerImpl struct {
	query infra.ItemListQuery
	db    db.DB
}

func NewItemListHandler(query infra.ItemListQuery, db db.DB) ItemListHandler {
	return &ItemListHandlerImpl{query: query, db: db}
}

func (h *ItemListHandlerImpl) Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	items, err := h.query.ItemList(ctx, tx)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var t []ItemResponse
	for _, item := range items {
		var itemResponse ItemResponse
		itemResponse.ID = item.ID.UUID.String()
		itemResponse.Description = item.Description
		itemResponse.Price = item.Price
		itemResponse.BirthDay = item.Birthday
		itemResponse.Image = item.Image
		itemResponse.IsDeleted = item.IsDeleted
		itemResponse.Gender = item.Gender.Name
		itemResponse.Breed.ID = item.Breed.ID.UUID.String()
		itemResponse.Breed.Name = item.Breed.Name
		itemResponse.Color.ID = item.Color.ID.UUID.String()
		itemResponse.Color.Name = item.Color.Name
		t = append(t, itemResponse)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		log.Println(err)
	}
}

type ItemResponse struct {
	ID string `json:"id"`
	// 詳細
	Description string `json:"description"`
	// 価格
	Price int `json:"price"`
	// 誕生日
	BirthDay time.Time `json:"birthDay"`
	// 画像パス
	Image string `json:"image"`
	// 削除フラグ
	IsDeleted bool `json:"deleted"`
	// 性別
	Gender string `json:"gender"`
	// 種別
	Breed struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"breed"`
	// 色
	Color struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"color"`
}
