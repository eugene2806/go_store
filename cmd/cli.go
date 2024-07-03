package main

import (
	"fmt"
	"store/internal/model"
)

func main() {
	item1 := model.NewItem("Bread", 200, 70)
	item2 := model.NewItem("Milk", 900, 80)
	cart := model.NewCart()

	cart.AddItem(item1)
	cart.AddItem(item1)
	cart.AddItem(item2)

	cart.RemoveItem(item1)
	cart.RemoveItem(item2)
	cart.RemoveItem(item1)
	for _, item := range cart.Items {
		fmt.Println("item:", item)
	}

}
