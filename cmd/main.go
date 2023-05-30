package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jserva90/Erply-test-task/database"
)

type application struct {
	DB database.SqliteDB
}

const (
	PORT = 8080
)

func main() {
	var app application

	// sessionKey := "ce8778a3fb64044ef21c7ef65838be765ee61f276ddf"

	// res, _ := app.verifySessionValidity("531748", sessionKey)

	// var response models.GetSessionKeyInfoResponse
	// err := json.Unmarshal([]byte(res), &response)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Response:")
	// fmt.Printf("Status:\n")
	// fmt.Printf("  Request: %s\n", response.Status.Request)
	// fmt.Printf("  RequestUnixTime: %d\n", response.Status.RequestUnixTime)
	// fmt.Printf("  ResponseStatus: %s\n", response.Status.ResponseStatus)
	// fmt.Printf("  ErrorCode: %d\n", response.Status.ErrorCode)
	// fmt.Printf("  GenerationTime: %.6f\n", response.Status.GenerationTime)
	// fmt.Printf("  RecordsTotal: %d\n", response.Status.RecordsTotal)
	// fmt.Printf("  RecordsInResponse: %d\n", response.Status.RecordsInResponse)

	// fmt.Println("Records:")
	// for i, record := range response.Records {
	// 	fmt.Printf("  Record %d:\n", i+1)
	// 	fmt.Printf("    CreationUnixTime: %s\n", record.CreationUnixTime)
	// 	fmt.Printf("    ExpireUnixTime: %s\n", record.ExpireUnixTime)
	// }

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
