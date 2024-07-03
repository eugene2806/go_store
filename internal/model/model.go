package model

type Item struct {
	name      string
	weight    int
	priceRUB  int
	totalItem TotalItem
}

type Cart struct {
	Items []*Item
}

type TotalItem struct {
	weightSum int
	priceSum  int
	countItem int
}

func NewTotalItem(newWeight, newPriceRUB int) *TotalItem {
	return &TotalItem{weightSum: newWeight, priceSum: newPriceRUB, countItem: 1}
}

func NewItem(productName string, newWeight, newPriceRUB int) *Item {
	return &Item{
		name:      productName,
		weight:    newWeight,
		priceRUB:  newPriceRUB,
		totalItem: *NewTotalItem(newWeight, newPriceRUB),
	}
}

func NewCart() *Cart {
	return &Cart{Items: make([]*Item, 0, 10)}
}

func (c *Cart) AddItem(item *Item) {
	for _, i := range c.Items {
		if i.name == item.name {

			item.totalItem.weightSum += item.weight
			item.totalItem.priceSum += item.priceRUB
			item.totalItem.countItem += 1

			return
		}
	}

	c.Items = append(c.Items, item)
}

func (c *Cart) RemoveItem(item *Item) {
	for index, i := range c.Items {
		if i.name == item.name {

			if i.totalItem.countItem > 1 {

				item.totalItem.weightSum -= item.weight
				item.totalItem.priceSum -= item.priceRUB
				item.totalItem.countItem -= 1

				return
			}

			c.Items = append(c.Items[:index], c.Items[index+1:]...)

			return
		}
	}
}
