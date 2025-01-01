package item

import (
	"net/http"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/item/", http.StripPrefix("/item", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// GETリクエストの処理
		case http.MethodPost:
			// POSTリクエストの処理
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	},
	)))
	return mux
}
