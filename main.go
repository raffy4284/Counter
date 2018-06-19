package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/raffy4284/counter/models"
	database "github.com/raffy4284/counter/models/database"
)

func main() {
	db := database.GetDatabase()
	defer db.DB.Close()
	createSchemas(db)
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createSchemas(db *database.DBHandler) {
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Counter{})
}
