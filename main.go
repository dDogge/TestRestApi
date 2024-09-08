package main

import (
	"database/sql"
	"log"

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
}
