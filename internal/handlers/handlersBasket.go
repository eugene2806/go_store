package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"store/internal/model/basket"
	"store/internal/model/items"
	"store/internal/service"
)

func GetAllItemsBasket(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	log.Println("Get all items Basket")
	writer.WriteHeader(200)

	json.NewEncoder(writer).Encode(basket.NewBasket)
}

func AddItemBasket(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Add item Basket")

	var itemInput struct {
		ID uuid.UUID `json:"id"`
	}

	err := json.NewDecoder(request.Body).Decode(&itemInput)

	if err != nil {
		msg := Message{Message: "Provided JSON is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
	}

	product, ok := items.FindItemDBByID(itemInput.ID)

	if !ok {
		msg := Message{Message: "Product not found"}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	basketService := service.BasketService{}
	item := basketService.AddItemToBasket(product)

	writer.WriteHeader(201)

	json.NewEncoder(writer).Encode(item)
}
