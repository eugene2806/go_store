package utils

import (
	"github.com/gorilla/mux"
	"store/internal/model/handlers"
)

func BuildStoreDB(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllItemsDB).Methods("GET")
}
