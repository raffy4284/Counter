package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controllers "github.com/raffy4284/counter/controllers"
	database "github.com/raffy4284/counter/models/database"
)

func main() {
	db := database.GetDatabase()
	defer db.DB.Close()
	database.CreateSchemas()

	router := mux.NewRouter()

	// ************************************
	// Sample endpoints; will delete later
	router.HandleFunc("/people", controllers.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", controllers.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", controllers.CreatePerson).Methods("POST")
	// ************************************

	log.Fatal(http.ListenAndServe(":8080", router))
}
