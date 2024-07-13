package main

import (
	"database/sql"

	"log"
	"my-go-api/pkg/db"
	"my-go-api/pkg/routes"
	"my-go-api/pkg/services"

	_ "github.com/lib/pq"
)

func main() {

	// Connect to mongodb
	mongoClient := db.ConnectDB()
	defer db.DisconnectDB()

	// Connect to postgres
	postgresClient, err := db.ConnectToPSQL()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer postgresClient.Close()

	// Create table if it doesn't exist
	err = services.CreateTableIfNotExists(postgresClient)
	if err != nil {
		log.Fatal(err)
	}

	// Set up the Gin router for both MongoDB and PostgreSQL
	r := routes.SetupRoutes(mongoClient, postgresClient)

	// Start http server
	r.Run(":8081")
}

func ClosePSQL(db *sql.DB) {
	panic("unimplemented")
}
