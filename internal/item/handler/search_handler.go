//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package handler

import (
	"char5742/ecsite-sample/internal/app/handler"
	app "char5742/ecsite-sample/internal/app/infra"
	"char5742/ecsite-sample/internal/item/domain"
	"char5742/ecsite-sample/internal/item/infra"
	"char5742/ecsite-sample/pkg/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ItemListSearchHandler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type ItemListSearchHandlerImpl struct {
	query infra.ItemListQuery
	db    db.DB
}

func NewItemListSearchHandler(query infra.ItemListQuery, db db.DB) ItemListSearchHandler {
	return &ItemListSearchHandlerImpl{query: query, db: db}
}

func (h *ItemListSearchHandlerImpl) Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("failed to begin transaction: %v", err)
		return
	}

	query := SearchQueryRequest{}
	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		http.Error(w, "Argument Error", http.StatusBadRequest)
		return
	}
	fmt.Printf("query: %v", query)

	condition := infra.ItemListCondition{
		PriceCond: infra.PriceCondition{
			Min:   query.Search.MinPrice,
			Max:   query.Search.MaxPrice,
			Valid: true,
		},
		ColorCond: infra.ColorCondition{
			ColorIDList: query.Search.ColorList,
		},
		BreedCond: infra.BreedCondition{
			BreedIDList: []string{query.Search.Breed},
		},
		Pagination: app.Pagination{
			CurrentPage: query.Page.CurrentPage,
			PerPage:     query.Page.PerPage,
		},
	}

	items, err := h.query.ItemListByCondition(ctx, tx, condition)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("failed to get item list: %v", err)
		return
	}

	t := buildItemListSearchResponse(items)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Printf("failed to encode response: %v", err)
		return
	}
}

func buildItemListSearchResponse(items []domain.Item) handler.Page[ItemResponse] {
	records := []ItemResponse{}
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
		records = append(records, itemResponse)
	}
	var m handler.Metadata

	m.CurrentPage = 1
	m.PerPage = 10
	m.Total = len(records)
	m.LastPage = m.Total / m.PerPage
	if m.Total%m.PerPage != 0 {
		m.LastPage++
	}

	return handler.Page[ItemResponse]{
		Metadata: m,
		Records:  records,
	}
}

type SearchQueryRequest struct {
	Search struct {
		MaxPrice  int      `json:"maxPrice"`
		MinPrice  int      `json:"minPrice"`
		ColorList []string `json:"colorList"`
		Breed     string   `json:"breed"`
	} `json:"search"`
	Page struct {
		CurrentPage int `json:"currentPage"`
		PerPage     int `json:"perPage"`
	} `json:"page"`
}
