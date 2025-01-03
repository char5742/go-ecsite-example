package app

import (
	"char5742/ecsite-sample/internal/item"
	"net/http"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	itemMux := item.NewMux()
	mux.Handle("/", itemMux)

	return mux
}
