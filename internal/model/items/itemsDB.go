package items

import (
	"github.com/google/uuid"
)

type ItemDB struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Weight   int       `json:"weight"`
	PriceRUB int       `json:"price_rub"`
	InStock  int       `json:"in_stock"`
}

var DB map[uuid.UUID]ItemDB

func NewItemDB(productName string, newWeight, newPriceRUB, newInStock int) *ItemDB {
	return &ItemDB{
		ID:       uuid.New(),
		Name:     productName,
		Weight:   newWeight,
		PriceRUB: newPriceRUB,
		InStock:  newInStock,
	}
}

func FindItemDBByID(id uuid.UUID) (*ItemDB, bool) {
	var find bool

	if item, ok := DB[id]; ok {
		find = true

		return &item, find
	}

	return nil, find
}

func init() {
	DB = make(map[uuid.UUID]ItemDB)

	item1 := NewItemDB("Bread", 200, 80, 10)
	item2 := NewItemDB("Milk", 900, 80, 5)

	DB[item1.ID] = *item1
	DB[item2.ID] = *item2

}
