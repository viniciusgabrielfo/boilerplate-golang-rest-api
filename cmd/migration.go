package main

import (
	"boilerplate/database"
	"log"

	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joho/godotenv"
)

func main() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	database.GenerateDatabaseURL()
	database.Migrate()

	log.Println("Migration was succesfully executed.")
}
