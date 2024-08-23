package service

import "store/internal/model/items"

type ItemsService struct {
}

func (is *ItemsService) AddItemToDB(item []items.ItemDB) {
	for index, i := range item {

		product := items.NewItemDB(i.Name, i.Weight, i.PriceRUB, i.InStock)
		item[index].ID = product.ID
		items.DB[product.ID] = *product
	}
}
