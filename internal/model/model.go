package model

type Item struct {
	name      string
	weight    int
	priceRUB  int
	totalItem TotalItem
}

type Cart struct {
	Items []Item
}

type TotalItem struct {
	weightSum int
	priceSum  int
	countItem int
}

func NewItem(productName string, newWeight, newPriceRUB int) *Item {
	return &Item{
		name:     productName,
		weight:   newWeight,
		priceRUB: newPriceRUB,
		totalItem: TotalItem{
			weightSum: newWeight,
			priceSum:  newPriceRUB,
			countItem: 1,
		},
	}
}

func NewCart() *Cart {
	return &Cart{Items: make([]Item, 0, 10)}
}

func (c *Cart) AddItem(item *Item) {
	for index, i := range c.Items {
		if i.name == item.name {

			newItem := Item{
				name:     i.name,
				weight:   i.weight,
				priceRUB: i.priceRUB,
				totalItem: TotalItem{
					weightSum: i.totalItem.weightSum + i.weight,
					priceSum:  i.totalItem.priceSum + i.priceRUB,
					countItem: i.totalItem.countItem + 1,
				},
			}

			c.Items[index] = newItem

			return
		}
	}

	c.Items = append(c.Items, *item)
}

func (c *Cart) RemoveItem(item *Item) {
	for index, i := range c.Items {
		if i.name == item.name {
			if i.totalItem.countItem > 1 {
				newItem := Item{
					name:     i.name,
					weight:   i.weight,
					priceRUB: i.priceRUB,
					totalItem: TotalItem{
						weightSum: i.totalItem.weightSum - i.weight,
						priceSum:  i.totalItem.priceSum - i.priceRUB,
						countItem: i.totalItem.countItem - 1,
					},
				}
				c.Items[index] = newItem

				return
			}
			c.Items = append(c.Items[:index], c.Items[index+1:]...)
			return

		}
	}
}
