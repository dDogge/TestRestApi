package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dDogge/TestRestApi/api"
	"github.com/dDogge/TestRestApi/database"
	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

func main() {
	var err error
	db, err := sql.Open("sqlite", "./temp_db.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.CreateTable(db)
	router := mux.NewRouter()
	api.SetupRoutes(router, db)

	fs := http.FileServer(http.Dir("./frontend/build"))
	router.PathPrefix("/").Handler(fs)

	log.Println("Server started on :3000")
	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
