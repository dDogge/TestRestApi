package main

import (
	"database/sql"
	"log"

	"github.com/dDogge/TestRestApi/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.01:3306)/temp_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.AddPerson(db, "Mister", "Anderson")
	database.AddPerson(db, "Neo", "Fetch")
	database.AddPerson(db, "Cave", "Goblin")

	database.GetPersons(db)
}
