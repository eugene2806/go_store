package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"store/internal/utils"
)

const apiPrefix string = "/api"

var (
	port          string = "8080"
	storeDBPrefix        = apiPrefix + "/storeDB"
)

func main() {
	log.Println("Start store server")
	router := mux.NewRouter()

	utils.BuildStoreDB(router, storeDBPrefix)

	log.Println("Router configured successfully")

	log.Fatal(http.ListenAndServe(":"+port, router))

	//for _, item := range itemsDB.DB {
	//	fmt.Println(item)
	//}
	//key, err := uuid.Parse("4424a5c9-b0e9-47d9-b250-d685158a7c60")
	//if err != nil {
	//	fmt.Println("Invalid UUID string:", err)
	//
	//	return
	//}
	//
	//if value, ok := itemsDB.DB[key]; ok {
	//	cart.AddItem(value)
	//} else {
	//	fmt.Println("Item not found in DB")
	//}

}
