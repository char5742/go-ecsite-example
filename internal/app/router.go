package app

import (
	"net/http"

	"github.com/char5742/ecsite-sample/internal/item"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	itemMux := item.NewMux()
	mux.Handle("/item/", http.StripPrefix("/item", itemMux))
	return mux
}
