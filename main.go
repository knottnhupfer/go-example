package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"service/handlers"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {

	configureDatabase()

	handlers.RegisterHandlers()

	log.Println("Starting server...")
	http.ListenAndServe("0.0.0.0:7080", nil)
}

// https://golang.org/doc/tutorial/database-access
// https://data-nerd.blog/2020/04/11/connecting-to-postgresql-from-go-lang-project/
func configureDatabase() {

	connStr := "user=postgres dbname=demo password=secure host=0.0.0.0 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")
}
