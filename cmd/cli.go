package main

import (
	"fmt"
	"store/internal/model"
)

func main() {
	item1 := model.NewItem("Bread", 200, 70)
	cart := model.NewCart()

	cart.AddItem(item1)
	cart.AddItem(item1)
	cart.AddItem(item1)
	for key, val := range cart.Items {
		fmt.Println("Key", key, "Val", val)
	}
	cart.RemoveItem(item1)

}
