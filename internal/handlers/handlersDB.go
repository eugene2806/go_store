package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"store/internal/model/items"
	"store/internal/service"
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
		msg := Message{Message: "Provided JSON is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	itemsService := service.ItemsService{}
	itemsService.AddItemToDB(item)

	writer.WriteHeader(201)

	json.NewEncoder(writer).Encode(item)
}
