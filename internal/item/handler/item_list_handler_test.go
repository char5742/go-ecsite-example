package handler

import (
	"char5742/ecsite-sample/internal/item/domain"
	"char5742/ecsite-sample/internal/item/infra"
	"char5742/ecsite-sample/pkg/db"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestItemListHandler_Handler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQuery := infra.NewMockItemListQuery(ctrl)
	mockDB := db.NewMockDB(ctrl)
	mockTX := db.NewMockTX(ctrl)
	handler := NewItemListHandler(mockQuery, mockDB)
	t.Log("TestItemListHandler_Handler")
	// Mock data
	items := []domain.Item{
		{
			ID:          domain.NewItemID("44444444-4444-4444-4444-444444444444"),
			Description: "Item 1",
			Price:       100,
			Birthday:    time.Now(),
			Image:       "image1.png",
			IsDeleted:   false,
			Gender:      domain.Gender{Name: "Male"},
			Breed: domain.Breed{ID: domain.NewBreedID("44444444-4444-4444-4444-444444444444"),
				Name: "Breed 1"},
			Color: domain.Color{ID: domain.NewColorID("44444444-4444-4444-4444-444444444444"),
				Name: "Color 1"},
		},
	}

	mockQuery.EXPECT().ItemList(gomock.Any(), gomock.Any()).Return(items, nil)
	mockDB.EXPECT().BeginTx(gomock.Any(), gomock.Any()).Return(mockTX, nil)
	req, err := http.NewRequest("GET", "/items", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.Handler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json; charset=utf-8", rr.Header().Get("Content-Type"))

	var response []ItemResponse
	t.Logf("%s", rr.Body.String())
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Len(t, response, 1)
	assert.Equal(t, "44444444-4444-4444-4444-444444444444", response[0].ID)
	assert.Equal(t, "Item 1", response[0].Description)
	assert.Equal(t, 100, response[0].Price)
	assert.Equal(t, "image1.png", response[0].Image)
	assert.Equal(t, "Male", response[0].Gender)
	assert.Equal(t, "44444444-4444-4444-4444-444444444444", response[0].Breed.ID)
	assert.Equal(t, "Breed 1", response[0].Breed.Name)
	assert.Equal(t, "44444444-4444-4444-4444-444444444444", response[0].Color.ID)
	assert.Equal(t, "Color 1", response[0].Color.Name)
}
