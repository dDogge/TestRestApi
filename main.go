package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dDogge/TestRestApi/database"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./temp_db.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.CreateTable(db)

	database.AddPerson(db, "Mister", "Anderson")
	database.AddPerson(db, "Neo", "Fetch")
	database.AddPerson(db, "Cave", "Goblin")

	database.GetPersons(db)

	fs := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", fs)

	log.Println("Server started on :3000")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
