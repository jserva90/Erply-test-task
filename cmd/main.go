package main

import (
	"fmt"
	"log"
	"net/http"

	database "github.com/jserva90/Erply-test-task/persistance"
)

type application struct {
	DB database.SqliteDB
}

const (
	PORT = 8080
)

func main() {
	var app application

	conn, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = database.SqliteDB{DB: conn}

	defer app.DB.Connection().Close()

	log.Printf("Starting server at http://localhost:%d/\n", PORT)
	err = http.ListenAndServe(fmt.Sprintf(":%d", PORT), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
