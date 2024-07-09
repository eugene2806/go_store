package itemsDB

import (
	"github.com/google/uuid"
	"store/internal/model/basket"
)

var DB map[uuid.UUID]*basket.Item

func init() {
	DB = make(map[uuid.UUID]*basket.Item)

	item1 := basket.NewItem("Bread", 200, 70)
	item2 := basket.NewItem("Milk", 900, 80)

	DB[item1.ID] = item1
	DB[item2.ID] = item2

}
