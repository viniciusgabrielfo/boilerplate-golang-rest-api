package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
)

// InstanceDB is a variable that contains a instance of database connection.
var InstanceDB *sql.DB

var databaseURL string

// GenerateDatabaseURL is a function to generate database connection URL string for other functions in this file
func GenerateDatabaseURL() {
	databaseURL = fmt.Sprintf("%v://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("db_type"),
		os.Getenv("db_user"),
		os.Getenv("db_pass"),
		os.Getenv("db_host"),
		os.Getenv("db_port"),
		os.Getenv("db_name"),
		os.Getenv("db_ssl_mode"))
}

// OpenConnectionDatabase is a function to connect and create instance for database (InstanceDB).
func OpenConnectionDatabase() {
	// Validates if databaseURL was defined
	if databaseURL == "" {
		log.Fatal("Call GenerateDatabaseURL() before call this function.")
	}

	db, err := sql.Open(os.Getenv("db_driver"), databaseURL)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}

	InstanceDB = db
	log.Println("Database connected and InstanceDB was generated.")
}

// Migrate is a function to execute migrations.
func Migrate() {
	// Validates if databaseURL was defined
	if databaseURL == "" {
		log.Fatal("Call GenerateDatabaseURL() before call this function.")
	}

	m, err := migrate.New("file://database/migrations", databaseURL)

	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
