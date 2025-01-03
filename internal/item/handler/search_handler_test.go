package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"char5742/ecsite-sample/internal/item/domain"
	"char5742/ecsite-sample/internal/item/infra"
	"char5742/ecsite-sample/pkg/db"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestItemListSearchHandler_Handler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuery := infra.NewMockItemListQuery(ctrl)
	mockDB := db.NewMockDB(ctrl)
	mockTX := db.NewMockTX(ctrl)
	handler := NewItemListSearchHandler(mockQuery, mockDB)

	t.Run("successful search", func(t *testing.T) {
		reqBody := SearchQueryRequest{
			Search: struct {
				MaxPrice  int      `json:"maxPrice"`
				MinPrice  int      `json:"minPrice"`
				ColorList []string `json:"colorList"`
				Breed     string   `json:"breed"`
			}{
				MaxPrice:  1000,
				MinPrice:  100,
				ColorList: []string{"red", "blue"},
				Breed:     "bulldog",
			},
			Page: struct {
				CurrentPage int `json:"currentPage"`
				PerPage     int `json:"perPage"`
			}{
				CurrentPage: 1,
				PerPage:     10,
			},
		}

		reqBodyBytes, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/search", bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()

		ctx := context.Background()

		mockDB.EXPECT().BeginTx(ctx, nil).Return(mockTX, nil)

		mockQuery.EXPECT().ItemListByCondition(ctx, mockTX, gomock.Any()).Return([]domain.Item{
			{
				ID:          domain.NewItemID("44444444-4444-4444-4444-444444444444"),
				Description: "Sample Item",
				Price:       500,
				Birthday:    time.Now(),
				Image:       "image.jpg",
				IsDeleted:   false,
				Gender:      domain.Gender{Name: "Male"},
				Breed:       domain.Breed{ID: domain.NewBreedID("44444444-4444-4444-4444-444444444444"), Name: "Bulldog"},
				Color:       domain.Color{ID: domain.NewColorID("44444444-4444-4444-4444-444444444444"), Name: "Red"},
			},
		}, nil)

		handler.Handler(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	})

	t.Run("failed to begin transaction", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/search", nil)
		w := httptest.NewRecorder()

		ctx := context.Background()

		mockDB.EXPECT().BeginTx(ctx, nil).Return(nil, errors.New("failed to begin transaction"))

		handler.Handler(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("failed to decode request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/search", bytes.NewReader([]byte("invalid json")))
		w := httptest.NewRecorder()

		mockDB.EXPECT().BeginTx(gomock.Any(), nil).Return(mockTX, nil)

		handler.Handler(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("failed to get item list", func(t *testing.T) {
		reqBody := SearchQueryRequest{
			Search: struct {
				MaxPrice  int      `json:"maxPrice"`
				MinPrice  int      `json:"minPrice"`
				ColorList []string `json:"colorList"`
				Breed     string   `json:"breed"`
			}{
				MaxPrice:  1000,
				MinPrice:  100,
				ColorList: []string{"red", "blue"},
				Breed:     "bulldog",
			},
			Page: struct {
				CurrentPage int `json:"currentPage"`
				PerPage     int `json:"perPage"`
			}{
				CurrentPage: 1,
				PerPage:     10,
			},
		}

		reqBodyBytes, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/search", bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()

		ctx := context.Background()
		mockDB.EXPECT().BeginTx(ctx, nil).Return(mockTX, nil)
		mockQuery.EXPECT().ItemListByCondition(ctx, mockTX, gomock.Any()).Return(nil, errors.New("failed to get item list"))

		handler.Handler(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
