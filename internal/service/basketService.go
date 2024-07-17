package service

import (
	"store/internal/model/basket"
	"store/internal/model/items"
)

type BasketService struct {
}

func (bs *BasketService) AddItemToBasket(product *items.ItemDB) *basket.ItemBasket {
	item := basket.NewItemBasket(product.ID, product.Name, product.Weight, product.PriceRUB)

	basket.NewBasket.AddItem(item, product)

	return item
}
