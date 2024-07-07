package main

import (
	"database/sql"

	"log"
	"my-go-api/pkg/db"
	"my-go-api/pkg/routes"

	_ "github.com/lib/pq"
)

func main() {

	// Connect to mongodb
	client := db.ConnectDB()
	defer db.DisconnectDB()

	// Connect to postgres
	db, err := db.ConnectToPSQL()
	if err != nil {
		log.Fatal(err)
	}
	defer ClosePSQL(db)

	// Set up the Gin router
	r := routes.SetupRouter(client)

	// Start http server
	r.Run(":8081")
}

func ClosePSQL(db *sql.DB) {
	panic("unimplemented")
}
