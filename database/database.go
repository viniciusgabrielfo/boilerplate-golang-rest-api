package database

import (
	"database/sql"
	"fmt"
)

var InstanceDB *sql.DB

func OpenConnection() {

	dbDriver := "postgres"
	dbType := "postgres"
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPass := "root"
	dbName := "api-boilerplate"
	dbSSLMode := "disable"

	db, err := sql.Open(dbDriver, fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", dbType, dbUser, dbPass, dbHost, dbPort, dbName, dbSSLMode))

	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	InstanceDB = db
	fmt.Println("connected")
}
