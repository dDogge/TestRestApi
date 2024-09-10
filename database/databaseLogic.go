package database

import (
	"database/sql"
	"fmt"
	"log"
)

type Person struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func CreateTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS persons (
				id INTEGER PRIMARY KEY AUTOINCREMENT, 
				firstname VARCHAR(255) NOT NULL, 
				lastname VARCHAR(255) NOT NULL
			 );`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}
}

func AddPerson(db *sql.DB, firstname, lastname string) {
	stmt, err := db.Prepare("INSERT INTO persons(firstName, lastName) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(firstname, lastname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Person added!")
}

func GetPersons(db *sql.DB) []Person {
	rows, err := db.Query("SELECT firstName, lastName FROM persons")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var persons []Person

	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Firstname, &p.Lastname)
		if err != nil {
			log.Fatal(err)
		}
		persons = append(persons, p)
	}
	return persons
}
