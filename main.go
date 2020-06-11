package main

//go:generate sqlboiler --wipe psql

import (
	"boilerplate/app"
	"boilerplate/database"
	"fmt"
	"net/http"

	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	// m, err := migrate.New("file://database/migrations",  fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", dbType, dbUser, dbPass, dbHost, dbPort, dbName, dbSSLMode))

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := m.Up(); err != nil {
	// 	log.Fatal(err)
	// }

	router := app.ChargeRoutes()
	database.OpenConnection()

	port := "8000"

	// log := logger.GetLogger()
	// log.Debugf("Starting API server at %s", port)
	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
