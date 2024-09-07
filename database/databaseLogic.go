package database

import (
	"database/sql"
	"fmt"
	"log"
)

type Person struct {
	ID        int
	Firstname string
	Lastname  string
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

func GetPersons(db *sql.DB) {
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

	for _, p := range persons {
		fmt.Printf("Firstname: %s, Lastname: %s\n", p.Firstname, p.Lastname)
	}
}
