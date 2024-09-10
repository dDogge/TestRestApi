package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/dDogge/TestRestApi/database"
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/api/person", func(w http.ResponseWriter, r *http.Request) {
		addPerson(w, r, db)
	}).Methods("POST")

	router.HandleFunc("/api/persons", func(w http.ResponseWriter, r *http.Request) {
		getPersons(w, r, db)
	}).Methods("GET")
}

func addPerson(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var person database.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	database.AddPerson(db, person.Firstname, person.Lastname)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Person added!"))
}

func getPersons(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	persons := database.GetPersons(db)
	json.NewEncoder(w).Encode(persons)
}
