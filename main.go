package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jdearly/go-microservice/app"
	"github.com/jdearly/go-microservice/db"
)

func main() {
	// TODO: exponential retry algorithm instead? Just have to wait for MySQL
	// service to be ready for connections
	time.Sleep(5 * time.Second)
	database, err := db.CreateDB()

	if err != nil {
		log.Fatalf("Connection failed: %s", err.Error())
	}

	fmt.Println("Connection successful, listening")

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.Setup()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
