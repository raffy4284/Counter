package controllers

import (
	"encoding/json"
	"net/http"

	mux "github.com/gorilla/mux"
	models "github.com/raffy4284/counter/models"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.GetPeople())
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range models.GetPeople() {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Person{})
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	var people = models.GetPeople()
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func DeletePerson() {}
