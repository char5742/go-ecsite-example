package domain

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ItemID uuid.NullUUID

func NewItemID(id string) ItemID {
	return ItemID{UUID: uuid.MustParse(id)}
}

func (id ItemID) Value() (driver.Value, error) {
	return id.UUID.String(), nil
}

// sql.Scanner を実装
func (id *ItemID) Scan(value interface{}) error {
	if v, ok := value.([]uint8); ok {
		*id = NewItemID(string(v))
		return nil
	}
	return fmt.Errorf("cannot scan type %T into ItemID", value)
}

type Item struct {
	ID ItemID
	// 詳細
	Description string
	// 価格
	Price int
	// 誕生日
	Birthday time.Time
	// 画像パス
	Image string
	// 性別
	Gender Gender
	// 種別
	Breed Breed
	// 色
	Color Color
}

// 性別
type GenderID uuid.NullUUID

func NewGenderID(id string) GenderID {
	return GenderID{UUID: uuid.MustParse(id)}
}

func (id GenderID) Value() (driver.Value, error) {
	return id.UUID.String(), nil
}

func (id *GenderID) Scan(value interface{}) error {
	if v, ok := value.([]uint8); ok {
		*id = NewGenderID(string(v))
		return nil
	}
	return fmt.Errorf("cannot scan type %T into GenderID", value)
}

type Gender struct {
	ID   GenderID
	Name string
}

// 種別
type BreedID uuid.NullUUID

func NewBreedID(id string) BreedID {
	return BreedID{UUID: uuid.MustParse(id)}
}

func (id BreedID) Value() (driver.Value, error) {
	return id.UUID.String(), nil
}

func (id *BreedID) Scan(value interface{}) error {
	if v, ok := value.([]uint8); ok {
		*id = NewBreedID(string(v))
		return nil
	}
	return fmt.Errorf("cannot scan type %T into BreedID", value)
}

type Breed struct {
	ID   BreedID
	Name string
}

// 色
type ColorID uuid.NullUUID

func NewColorID(id string) ColorID {
	return ColorID{UUID: uuid.MustParse(id)}
}

func (id ColorID) Value() (driver.Value, error) {
	return id.UUID.String(), nil
}

func (id *ColorID) Scan(value interface{}) error {
	if v, ok := value.([]uint8); ok {
		*id = NewColorID(string(v))
		return nil
	}
	return fmt.Errorf("cannot scan type %T into ColorID", value)
}

type Color struct {
	ID   ColorID
	Name string
}

func NewItem(id ItemID, description string, price int, birthday time.Time, image string, gender Gender, breed Breed, color Color) *Item {
	return &Item{
		ID:          id,
		Description: description,
		Price:       price,
		Birthday:    birthday,
		Image:       image,
		Gender:      gender,
		Breed:       breed,
		Color:       color,
	}
}
