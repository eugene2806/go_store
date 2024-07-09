package basket

import "github.com/google/uuid"

type Basket struct {
	Items []*Item
}

type Item struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Weight    int       `json:"weight"`
	PriceRUB  int       `json:"price_rub"`
	TotalItem TotalItem
}

type TotalItem struct {
	WeightSum int `json:"weight_sum"`
	PriceSum  int `json:"price_sum"`
	CountItem int `json:"count_item"`
}

func NewItem(productName string, newWeight, newPriceRUB int) *Item {
	return &Item{
		ID:       uuid.New(),
		Name:     productName,
		Weight:   newWeight,
		PriceRUB: newPriceRUB,
		TotalItem: TotalItem{
			PriceSum:  newPriceRUB,
			WeightSum: newWeight,
			CountItem: 1,
		},
	}
}

func NewCart() *Basket {
	return &Basket{Items: make([]*Item, 0, 10)}
}

func (c *Basket) AddItem(item *Item) {
	for _, i := range c.Items {
		if i.Name == item.Name {

			i.TotalItem.WeightSum += item.Weight
			i.TotalItem.PriceSum += item.PriceRUB
			i.TotalItem.CountItem++

			return
		}
	}

	c.Items = append(c.Items, item)
}

func (c *Basket) RemoveItem(item *Item) {
	for index, i := range c.Items {
		if i.Name == item.Name {

			if i.TotalItem.CountItem > 1 {
				i.TotalItem.WeightSum -= item.Weight
				i.TotalItem.PriceSum -= item.PriceRUB
				i.TotalItem.CountItem--

				return
			}

			c.Items = append(c.Items[:index], c.Items[index+1:]...)

			return
		}
	}
}
