package main

//go:generate sqlboiler --wipe psql

import (
	"boilerplate/app"
	"boilerplate/database"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	router := app.LoadRoutes()
	database.GenerateDatabaseURL()
	database.OpenConnectionDatabase()
	database.Migrate()

	port := "8000"
	log.Printf("Starting API server at %s port.\n", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
