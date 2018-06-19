package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	database "github.com/raffy4284/counter/models/database"
)

func main() {
	db := database.GetDatabase()
	defer db.DB.Close()
	database.CreateSchemas()
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
