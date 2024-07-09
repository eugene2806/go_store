package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"store/internal/model/itemsDB"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GetAllItemsDB(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Get all items DB")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(itemsDB.DB)
}
