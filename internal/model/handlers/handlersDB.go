package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"store/internal/model/items"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GetAllItemsDB(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	log.Println("Get all items DB")
	writer.WriteHeader(200)

	json.NewEncoder(writer).Encode(items.DB)

}

func AddItemDB(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)

	log.Println("Add item DB")

	var item []items.ItemDB

	err := json.NewDecoder(request.Body).Decode(&item)

	if err != nil {
		msg := Message{Message: "Provided JSON file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	for index, i := range item {
		product := items.NewItemDB(i.Name, i.Weight, i.PriceRUB, i.InStock)
		item[index].ID = product.ID
		items.DB[product.ID] = *product
	}

	writer.WriteHeader(201)

	json.NewEncoder(writer).Encode(item)
}
