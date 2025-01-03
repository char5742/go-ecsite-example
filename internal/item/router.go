package item

import (
	"char5742/ecsite-sample/internal/item/handler"
	"char5742/ecsite-sample/internal/item/infra"
	"char5742/ecsite-sample/pkg/db"
	"net/http"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	db, err := db.OpenDB()
	if err != nil {
		panic(err)
	}
	query := infra.NewItemListQuery()
	itemListHandler := handler.NewItemListHandler(query, db)
	mux.HandleFunc("/api/getItemList", itemListHandler.Handler)
	itemListSearchHandler := handler.NewItemListSearchHandler(query, db)
	mux.HandleFunc("/api/getItemList/page", itemListSearchHandler.Handler)

	return mux
}
