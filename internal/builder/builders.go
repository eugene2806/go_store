package builder

import (
	"github.com/gorilla/mux"
	"store/internal/handlers"
)

func BuildStoreDB(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllItemsDB).Methods("GET")
	router.HandleFunc(prefix, handlers.AddItemDB).Methods("POST")
}

func BuildBasket(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllItemsBasket).Methods("GET")
	router.HandleFunc(prefix, handlers.AddItemBasket).Methods("POST")
}
