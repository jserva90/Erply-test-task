package main

import (
	"fmt"
	"log"
	"net/http"

	database "github.com/jserva90/Erply-test-task/database"
	_ "github.com/jserva90/Erply-test-task/docs"
)

type application struct {
	DB database.DatabaseRepo
}

const (
	port = 8080
)

// @title Erply Test Task
// @version 1.0
// @description This is a test task for Erply using Erply API.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey SessionKeyAuth
// @in header
// @name sessionKey
func main() {
	var app application

	conn, err := database.OpenDB("./database/database.db")
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &database.SqliteDB{DB: conn}
	defer app.DB.Connection().Close()

	log.Printf("Starting server at http://localhost:%d/\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
