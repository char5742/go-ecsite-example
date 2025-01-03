package handler

type Page[T any] struct {
	Metadata Metadata `json:"metadata"`
	Records  []T      `json:"records"`
}

type Metadata struct {
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"perPage"`
	LastPage    int `json:"lastPage"`
	Total       int `json:"total"`
}
