package basket

import (
	"github.com/google/uuid"
	"store/internal/model/items"
)

type Basket struct {
	Items []*ItemBasket
}

type ItemBasket struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	TotalWeight   int       `json:"total_weight"`
	TotalPriceRUB int       `json:"total_price_rub"`
	CountItem     int       `json:"count_item"`
}

func NewItemBasket(newID uuid.UUID, productName string, newWeight, newPriceRUB int) *ItemBasket {
	return &ItemBasket{
		ID:            newID,
		Name:          productName,
		TotalWeight:   newWeight,
		TotalPriceRUB: newPriceRUB,
		CountItem:     1,
	}
}

func NewCart() *Basket {
	return &Basket{Items: make([]*ItemBasket, 0, 10)}
}

func (c *Basket) AddItem(item *ItemBasket, product *items.ItemDB) {
	for index, i := range c.Items {
		if i.ID == item.ID {

			c.Items[index].CountItem++
			c.Items[index].TotalWeight += product.Weight
			c.Items[index].TotalPriceRUB += product.PriceRUB

			return
		}
	}

	c.Items = append(c.Items, item)
}

func (c *Basket) RemoveItem(item *ItemBasket, product *items.ItemDB) {
	for index, i := range c.Items {
		if i.ID == item.ID {

			if c.Items[index].CountItem > 1 {

				c.Items[index].CountItem--
				c.Items[index].TotalWeight -= product.Weight
				c.Items[index].TotalPriceRUB -= product.PriceRUB

				return
			}

			c.Items = append(c.Items[:index], c.Items[index+1:]...)
		}
	}
}

var NewBasket *Basket

func init() {
	NewBasket = NewCart()
}
