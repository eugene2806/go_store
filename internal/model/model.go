package model

type Item struct {
	name     string
	weight   int
	priceRUB int
}

type Cart struct {
	Items map[string]CardItem
}

type CardItem struct {
	item      *Item
	weightSum int
	priceSum  int
	countItem int
}

func NewItem(productName string, newWeight, newPriceRUB int) *Item {
	return &Item{name: productName, weight: newWeight, priceRUB: newPriceRUB}
}

func NewCartItem(item *Item) *CardItem {
	return &CardItem{
		item:      item,
		weightSum: item.weight,
		priceSum:  item.priceRUB,
		countItem: 1,
	}
}

func NewCart() *Cart {
	return &Cart{Items: make(map[string]CardItem)}
}

func (c *Cart) AddItem(item *Item) {
	nameItem, ok := c.Items[item.name]

	if !ok {
		c.Items[item.name] = *NewCartItem(item)

		return
	}

	nameItem.countItem++
	nameItem.priceSum += nameItem.item.priceRUB
	nameItem.weightSum += nameItem.item.weight
	c.Items[item.name] = nameItem
}

func (c *Cart) RemoveItem(item *Item) {
	nameItem, ok := c.Items[item.name]

	if !ok {

		return
	}

	if c.Items[item.name].countItem > 1 {
		nameItem.countItem--
		nameItem.weightSum -= nameItem.item.weight
		nameItem.priceSum -= nameItem.item.priceRUB
		c.Items[item.name] = nameItem
		return
	}

	delete(c.Items, item.name)
}
